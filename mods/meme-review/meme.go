package memereview

import (
  "math/rand"
  "strconv"
  "time"

  log "github.com/Sirupsen/logrus"
  "github.com/bwmarrin/discordgo"

  "github.com/dokvis/goatbrotesquared/cmd"
)

//Load - Loads meme review
func Load() {
  log.Println("Loading Meme Review plugin")
	cmd.Make("meme","Fun", cmdMemeReview).HelpText("Reviews meme").Hidden().Add()
  cmd.Make("mr","Fun", cmdMemeReview).HelpText("Reviews meme").Hidden().Add()
  cmd.Make("memereview","Fun", cmdMemeReview).HelpText("Reviews meme").Add()
}


func cmdMemeReview(message []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	time, err := m.Timestamp.Parse()
	var meme memeReview
  var me editMemeSt
	if err != nil {
		meme.Time = "nil"
	} else {
		meme.Time = time.String()
	}
  rand.Seed(time.UnixNano())
  var forceErr error
  if len(message) >= 2 {
    meme.Random, forceErr = strconv.Atoi(message[1])
  }
  if len(message) == 1 || forceErr != nil {
  	meme.Random = rand.Intn(1000 - 0 + 1)
  }
	switch {
    //Speical cases
		case meme.Random == 0: {
			meme.Type = "ERROR"
      rand.Seed(time.UnixNano())
      if rand.Intn(5 - 0 + 1) == 1 {
        meme.URL = "https://cdn.discordapp.com/attachments/702153501480058890/703721282708963358/meme_explode.mp4"
      } else {
        meme.URL = "https://cdn.discordapp.com/attachments/702153501480058890/702576438993223711/meme_sorry.mp4"
      }
		}
    case meme.Random == 666: {
      meme.Type = "This is Illegal, you know?"
      meme.URL = "https://cdn.discordapp.com/attachments/702153501480058890/702649753212551278/meme_illegal.mp4"
    }
    case meme.Random == 699: {
      meme.Type = "OUTRAGEOUS"
      meme.URL = "https://cdn.discordapp.com/attachments/702153501480058890/703707360816005170/Meme_outrageous.mp4"
    }
    case meme.Random == 720: {
      meme.Type = "VERY ILLEGAL"
      meme.URL = "https://cdn.discordapp.com/attachments/702153501480058890/703704843184767006/meme_VERY_ILLEGAL.mp4"
    }
    case meme.Random == 999: {
      meme.Type = "MEMES BANNED"
      meme.URL = "https://cdn.discordapp.com/attachments/702153501480058890/703705102862516234/meme_EVERYTHING_ILLEGAL.mp4"
    }
    case meme.Random == 420: {
      meme.Type = "YES"
      meme.URL = "https://cdn.discordapp.com/attachments/702153501480058890/703718574887534733/meme_yes.mp4"
    }
    case meme.Random == 101: {
      meme.Type = "IS THAT A JOJO REFRENCE?"
      meme.URL = "https://cdn.discordapp.com/attachments/702153501480058890/703719482555629649/meme_IS_THAT_A_JOJO.mp4"
    }
    case meme.Random == 102: {
      meme.Type =  "IS THAT A JOJO REFRENCE?"
      meme.URL = "https://cdn.discordapp.com/attachments/702153501480058890/703720502795042987/meme_IS_THAT_A_JOJO_2.mp4"
    }
    case meme.Random == 300: {
      meme.Type = "Persona Contract Gained"
      meme.URL = "https://cdn.discordapp.com/attachments/702153501480058890/704011280167600128/meme_persona.mp4"
    }
    case meme.Random == 333: {
      meme.Type = "Unsure" //Is actually approved
      meme.URL = "https://cdn.discordapp.com/attachments/702153501480058890/704012667450425444/meme_limbo_.mp4"
      me.Type = "Approved"
      me.Edit = true
      me.Delay = 15
    }
    case meme.Random == 421: {
      meme.Type = "Meme"
      meme.URL = "https://cdn.discordapp.com/attachments/614851241406627914/704020711416529046/meme.mp4"
    }
    case meme.Random == 900 : {
      meme.Type = "Memes of production stolen"
      meme.URL = "https://cdn.discordapp.com/attachments/702153501480058890/704013328913268736/meme_production_stolen.mp4"
    }
    case meme.Random == 301: {
    meme.Type = "Meme Approved"
    meme.URL = "https://cdn.discordapp.com/attachments/702153501480058890/704528203179360296/meme_persona_2.mp4"
    }
    //Ranges
		case meme.Random > 0 && meme.Random < 334: {
				meme.Type = "Approved"
        rand.Seed(time.UnixNano())
        randNum := rand.Intn(10 - 0 + 1)
        switch randNum {
          case 1: meme.URL = "https://cdn.discordapp.com/attachments/702153501480058890/703723341604716664/meme_approved_2.mp4"
          case 2: meme.URL =  "https://cdn.discordapp.com/attachments/614851241406627914/704020695469785228/meme_approved_3.mp4"
          default: meme.URL = "https://cdn.discordapp.com/attachments/702153501480058890/702153568689455154/Meme_Approved-1.mp4"
        }
		}
		case meme.Random > 333 && meme.Random < 666: {
			meme.Type = "Unsure"
			meme.URL = "https://cdn.discordapp.com/attachments/702153501480058890/702153595583201369/Meme_Limbo-1.mp4"
		}
		case meme.Random > 666: {
			meme.Type = "Denied"
			meme.URL = "https://cdn.discordapp.com/attachments/702153501480058890/702153644732317706/Meme_Denied-1.mp4"
		}
	}
	videoEmbed := discordgo.MessageEmbedVideo{
		URL: meme.URL,
	}
	memeReviewEmbed := &discordgo.MessageEmbed{
		Color:       0x880000,
		Description: "",
		Video: &videoEmbed,
		Title:     meme.Type,
	}

	memeEmbed, err := s.ChannelMessageSendEmbed(m.ChannelID, memeReviewEmbed)
  if err != nil {
    log.Errorln("Meme embed failed: " + err.Error())
    return
  }
  if me.Edit {
    defer editMeme(me, memeEmbed.ID, s, m)
  }
	s.ChannelMessageSend(m.ChannelID, meme.URL)
	//log.Println("memereview sending " + memeerr.Error())
}

type memeReview struct {
	URL   string
	Type string
	Time string
	Random int
}
type editMemeSt struct {
  Edit bool
  Type string
  Delay time.Duration
}

func editMeme(do editMemeSt, msgid string, s *discordgo.Session, m *discordgo.MessageCreate) {
  memeReviewEmbed := &discordgo.MessageEmbed{
    Color:       0x880000,
    Description: "",
    //Video: &videoEmbed,
    Title:     do.Type,
  }
  time.Sleep(do.Delay*time.Second)
  s.ChannelMessageEditEmbed(m.ChannelID, msgid, memeReviewEmbed)
}
