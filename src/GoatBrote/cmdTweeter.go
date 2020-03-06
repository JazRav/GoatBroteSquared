
package main

import (
  "github.com/bwmarrin/discordgo"
  "github.com/ChimeraCoder/anaconda"
  log "github.com/Sirupsen/logrus"
  "net/url"
  "strings"
  "io/ioutil"
  "bufio"
  "os"
  "github.com/go-ini/ini"
  "strconv"
  "time"
)

func init() {
	makeCmd("twitfollow", cmdTwitFollow).helpText("Follows account").owner().add()
  makeCmd("twitmassfollow", cmdTwitMassFollow).helpText("Follows accounts via uploaded .txt file").owner().add()
  makeCmd("tweet", cmdTweet).helpText("Tweets, can upload a single image, 5mb limit").owner().add()
  makeCmd("twit", cmdTwitSwitch).helpText("Manages the twitter config file\n`SET` to set config files\n`LIST` to list config files").owner().add()
}
//lol, ngl, i made this to mass follow porn accounts on my AD account, since i don't want to be horny on main anymore
func cmdTwitMassFollow(message []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	if len(m.Message.Attachments) == 0 {
    s.ChannelMessageSend(m.ChannelID, "Attach a file, idiot")
  } else {
    file := m.Message.Attachments
    fileName := file[0].Filename
    if strings.HasSuffix(fileName, ".txt") {
          downloadErr := fileGetter(file[0].URL, "temp/"+fileName)
          if downloadErr == nil {
            locFile, err := os.Open("temp/"+fileName)
            if err == nil {
                defer locFile.Close()
                i := 1
                var twitterList []string
                scanner := bufio.NewScanner(locFile)
                for scanner.Scan() {
                  twitterList = append(twitterList, scanner.Text())
                  i++
                }
                //s.ChannelMessageSend(m.ChannelID, twitterList[0])
                twitAccNum := len(twitterList)
                acctMsg, acctMsgErr := s.ChannelMessageSend(m.ChannelID, "Followed 0 of " + strconv.Itoa(twitAccNum) + " accounts in this file")
                t := 1
                if acctMsgErr == nil {
                  for n := 0; n < twitAccNum; n++ {
                    twitterList[n] = strings.TrimPrefix(twitterList[n], "https://twitter.com/")
                    twitterList[n] = strings.TrimSuffix(twitterList[n], "/")
                    _, followErr := twitFollow(twitterList[n])
                    if followErr != nil {
                      s.ChannelMessageSend(m.ChannelID, "Error: " + followErr.Error())
                    }
                    s.ChannelMessageEdit(acctMsg.ChannelID, acctMsg.ID, "Followed " + strconv.Itoa(n+1) + " of " +strconv.Itoa(twitAccNum) + " accounts in this file\nLast follow: " + twitterList[n] )
                    t++
                    if t > 5 && n+1 != twitAccNum{
                      t = 1
                      time.Sleep(60 * time.Second)
                    }
                  }
                } else {
                  log.Println(acctMsgErr.Error())
                }
                s.ChannelMessageEdit(acctMsg.ChannelID, acctMsg.ID, "Followed all " +strconv.Itoa(twitAccNum) + " accounts in this file" )
                s.ChannelMessageSend(m.ChannelID, "Done following all the accounts")
            } else {
              s.ChannelMessageSend(m.ChannelID, "file failed to open, thanks bot")
            }
          } else {
            s.ChannelMessageSend(m.ChannelID, "file failed to download, thanks discord\nFile: " + file[0].URL + "\nDownload err: `" + downloadErr.Error() +"`" )
          }

    } else {
        s.ChannelMessageSend(m.ChannelID, "file needs to be .txt")
    }

  }
}

func cmdTwitFollow(message []string, s *discordgo.Session, m *discordgo.MessageCreate) {
  log.Println("aaaaaaaaaaaaaaaaaaaaaaa")
  if len(message) > 1 {
    msg, _ := s.ChannelMessageSend(m.ChannelID, "Following " + message[1]+"...")
    twitterName, errFollow := twitFollow(message[1])
    if errFollow != nil {
      s.ChannelMessageEdit(msg.ChannelID, msg.ID, "Failled to follow " + message[1] + "\nError: `" + errFollow.Error() + "`")
    } else {
      s.ChannelMessageEdit(msg.ChannelID, msg.ID, "Now following " + twitterName)
    }
  } else {
    s.ChannelMessageSend(m.ChannelID, "You need to include a user")
  }
}

func cmdTwitSwitch(message []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	if len(message) > 1 {
    if message[1] == "list" || message[1] == "cfgs" {
      cfg := ""
      files, err := ioutil.ReadDir("config/twitter")
      if err != nil {
        s.ChannelMessageSend(m.ChannelID, "Error: " + err.Error())
      } else {
        for _, file := range files {
          if strings.HasSuffix(file.Name(), ".ini") {
            log.Println("var file.Name() = " + file.Name())
            fileTemp := strings.TrimSuffix(file.Name(), ".ini")
            cfg = cfg + fileTemp +"\n"
          }
        }
        s.ChannelMessageSend(m.ChannelID, "Configs \n```\n"+cfg+"```")
      }
    } else if message[1] == "set" {
      if len(message) > 2 {
        config := strings.TrimPrefix(m.Content, message[0]+" "+message[1] + " ")
        log.Println(config)
        configFile, err := ini.Load("config/twitter/"+config+".ini")
        if err != nil {
          s.ChannelMessageSend(m.ChannelID, "Failed to set config to `"+config+"` " + err.Error())
        } else {
          twit.AccessToken = configFile.Section("").Key("token").String()
          twit.AccessTokenSecret = configFile.Section("").Key("tokenSecret").String()
          twit.ConsumerKey = configFile.Section("").Key("consumer").String()
          twit.ConsumerSecret = configFile.Section("").Key("consumerSecret").String()
          twit.CurrentConfg = config
          s.ChannelMessageSend(m.ChannelID, "Set config to : `"+ config + "`")
        }
      } else {
        s.ChannelMessageSend(m.ChannelID, "PUT SOMETHING THERE IDIOT")
      }
    }

  } else {
    s.ChannelMessageSend(m.ChannelID, "Current twitter config: `" + twit.CurrentConfg + "`, default: `" + twit.DefaultConfig + "`")
  }
}

func cmdTweet(message []string, s *discordgo.Session, m *discordgo.MessageCreate) {
  status := strings.TrimPrefix(m.Content, message[0])
  status = strings.Replace(status, "`", "", -1)

  var urlink string
  var err error

  shouldDelete := false

  tweetMediaFile := "images/imagebork.png"
  if len(m.Attachments) == 1 {
    log.Println("discord tweet file size is: " + strconv.Itoa(m.Attachments[0].Size) + "b")
    if m.Attachments[0].Size < 5242880 {
      fileErr := fileGetter(m.Attachments[0].URL, "temp/" + m.Attachments[0].Filename)
      if fileErr == nil {
        tweetMediaFile = "temp/" + m.Attachments[0].Filename
        shouldDelete = true
      }
      mediaFile64 := fileToBase64(tweetMediaFile)
      //log.Println(mediaFile64)
      twitMedia, twitFileErr := twitter().UploadMedia(mediaFile64)
     if twitFileErr != nil {
        log.Println("twitFileErr: " + twitFileErr.Error())
      }
      log.Println("Media ID: " + twitMedia.MediaIDString +" Size: "+ strconv.Itoa(twitMedia.Size))
      if shouldDelete{
        os.Remove(tweetMediaFile)
      }
      v := url.Values{}
      if twitMedia.MediaIDString != "" {
        v.Set("media_ids", strconv.FormatInt(twitMedia.MediaID, 10) )
        urlink, err = twitTweet(status, v)
      }
  }
    //yes, i know its a bad idea to fucking do this like this, sue me
    urlink = "You fool, that file is bigger than `5mb`, which is the limit for twitter for some reason"
  } else {
    urlink, err = twitTweet(status, nil)
  }


  if err != nil {
    s.ChannelMessageSend(m.ChannelID, "It borked with: " + err.Error())
  } else {
    s.ChannelMessageSend(m.ChannelID, urlink)
  }
}

func twitter() *anaconda.TwitterApi {
  api := anaconda.NewTwitterApiWithCredentials(twit.AccessToken, twit.AccessTokenSecret, twit.ConsumerKey, twit.ConsumerSecret)
  return api
}

func twitTweet(status string, v url.Values) (url string, err error) {
  tweet, err := twitter().PostTweet(status, v)
  url = "https://twitter.com/"+tweet.User.ScreenName+"/status/"+tweet.IdStr
  twitter().Close()
  return url, err
}


func twitFollow(twitterUser string) (string, error){
  user, err := twitter().FollowUser(twitterUser)
  log.Println(err)
  twitter().Close()
  return user.Name, err
}
