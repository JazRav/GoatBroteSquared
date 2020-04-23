/*
+++Plan+++
Twitter manager, coming soon, smoked too much weed to do this now, keep getting distracted, will do later
-Per Channel Twitter accounts
-Stored in server INI, so rework it from global INI
-Automatic replies, retweets and likes posting in discord, going to rate limit it and combin it

*/
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
	makeCmd("twitfollow", cmdTwitFollow).helpText("Follows account on global twitter").owner().add()
  makeCmd("twitmassfollow", cmdTwitMassFollow).helpText("Follows accounts on global twitter via uploaded .txt file").owner().add()
  makeCmd("tweet", cmdTweet).helpText("Tweets, can upload a single image, 5mb limit").add()
  makeCmd("twit", cmdTwitSwitch).helpText("Manages the twitter config file\n`SET` to set twitter user for global\n`LIST` to list twitter accounts").owner().add()
  makeCmd("twitowner", cmdTwitForAll).helpText("Toggles twitter for everyone (global and local)").owner().add()
  makeCmd("chantwitowner", cmdTwitForAll).helpText("Toggles twitter for everyone (global and local)").owner().add()
  makeCmd("twitlock", cmdTwitLock).helpText("Locks global twitter to non-admins").owner().add()
  makeCmd("chantwitlist", cmdTwitListChans).helpText("List whatever twitter account is tied channel").owner().add()
  makeCmd("chantwitset", cmdTwitListChans).helpText("Set twitter config to this channel").owner().disableDM().add()
  makeCmd("chantwitremove", cmdTwitRemoveChan).helpText("Removes whatever twitter account is tied channel").owner().disableDM().add()
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
func cmdTwitLock(message []string, s *discordgo.Session, m *discordgo.MessageCreate) {
  if twit.Lock {
    twit.Lock = false
    cfg.Section("bot").Key("twitterLock").SetValue("false")
    cfg.SaveTo(cfgFile)
    s.ChannelMessageSend(m.ChannelID, "Twitter unlocked for all channels (if owner only mode isn't on)")
  } else if !twit.Lock {
    twit.Lock = true
    cfg.Section("bot").Key("twitterLock").SetValue("true")
    cfg.SaveTo(cfgFile)
    s.ChannelMessageSend(m.ChannelID, "Twitter locked to channel")
    log.Println("ALL CAN TWEET DISABLED")
  }
}
func cmdTwitListChans(message []string, s *discordgo.Session, m *discordgo.MessageCreate) {

}
func cmdTwitSetChan(message []string, s *discordgo.Session, m *discordgo.MessageCreate) {
  //Set
}
func cmdTwitRemoveChan(message []string, s *discordgo.Session, m *discordgo.MessageCreate) {

}

func cmdTwitForAll(message []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	if twit.All == false {
		twit.All = true
		cfg.Section("bot").Key("twitterForAll").SetValue("true")
		cfg.SaveTo(cfgFile)
		s.ChannelMessageSend(m.ChannelID, "TWITTER FOR EVERYONE ENABLED")
		log.Println("ALL CAN TWEET ENABLED")
	} else {
		twit.All = false
		cfg.Section("bot").Key("twitterForAll").SetValue("false")
		cfg.SaveTo(cfgFile)
		s.ChannelMessageSend(m.ChannelID, "TWITTER FOR EVERYONE DISABLED")
		log.Println("ALL CAN TWEET DISABLED")
	}
}

func cmdTweet(message []string, s *discordgo.Session, m *discordgo.MessageCreate) {
  if (m.Author.ID != ownerID) && !twit.All{
    s.ChannelMessageSend(m.ChannelID, "No tweets for you")
    return
  }
  if m.Author.ID != ownerID && twit.Lock {
    s.ChannelMessageSend(m.ChannelID, "Tweets locked too <#" + m.ChannelID + ">")
    return
  }
  status := strings.TrimPrefix(m.Content, message[0])
  status = strings.Replace(status, "`", "", -1)

  if twit.All && (m.Author.ID != ownerID) {
    status = status + "\nby " + m.Author.Username + "#" + m.Author.Discriminator
  }
  var urlink string
  var err error

  //msgContainsFileLink := false
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
      twitMedia, twitFileErr := twitter(true, "").UploadMedia(mediaFile64)
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
    } else {
      //yes, i know its a bad idea to fucking do this like this, sue me
        urlink = "You fool, that file is bigger than `5mb`, which is the limit for twitter for some reason"
    }
  } else {
    urlink, err = twitTweet(status, nil)
  }


  if err != nil {
    s.ChannelMessageSend(m.ChannelID, "It borked with: " + err.Error())
  } else {
    s.ChannelMessageSend(m.ChannelID, urlink)
  }
}

func twitter(global bool, chanID string) *anaconda.TwitterApi {
  if global {
    api := anaconda.NewTwitterApiWithCredentials(twit.AccessToken, twit.AccessTokenSecret, twit.ConsumerKey, twit.ConsumerSecret)
    return api
  }

  api := anaconda.NewTwitterApiWithCredentials(twit.AccessToken, twit.AccessTokenSecret, twit.ConsumerKey, twit.ConsumerSecret)
  return api
}

func twitTweet(status string, v url.Values) (url string, err error) {
  tweet, err := twitter(true, "").PostTweet(status, v)
  url = "https://twitter.com/"+tweet.User.ScreenName+"/status/"+tweet.IdStr
  twitter(true, "").Close()
  return url, err
}


func twitFollow(twitterUser string) (string, error){
  user, err := twitter(true, "").FollowUser(twitterUser)
  log.Println(err)
  twitter(true, "").Close()
  return user.Name, err
}

func twitChanPaser(chanID string) (){

}

type twitAPIKeys struct {
	DefaultConfig string
	CurrentConfg string
	ConsumerKey string
	ConsumerSecret string
	AccessToken string
	AccessTokenSecret string
	Delay time.Duration
	All bool
	Lock bool
	PerChanInfo string
	ChanInfo []twitPerChanInfo
}

type twitPerChanInfo struct {
	ChanID string
	TwitterAccount string
	GetReplies bool //When I work in a checks
	GetLikes bool // ditto
	GetRetweets bool //ditto
	OwnerOnly bool
}
