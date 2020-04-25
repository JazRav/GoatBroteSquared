package memereview

import (
  "math/rand"
  "strconv"

  "github.com/bwmarrin/discordgo"

  "github.com/dokvis/goatbrotesquared/cmd"
)

//Load - Loads meme review
func Load() {
	cmd.Make("meme","Meme Review", cmdMemeReview).HelpText("Reviews meme").Add()
  cmd.Make("mr","Meme Review", cmdMemeReview).HelpText("Reviews meme").Add()
  cmd.Make("memereview","Meme Review", cmdMemeReview).HelpText("Reviews meme").Add()
}

func cmdMemeReview(message []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	time, err := m.Timestamp.Parse()
	var meme memeReview
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
    //Ranges
		case meme.Random > 0 && meme.Random < 340: {
				meme.Type = "Approved"
        rand.Seed(time.UnixNano())
        if rand.Intn(2 - 0 + 1) == 1 {
          meme.URL = "https://cdn.discordapp.com/attachments/702153501480058890/703723341604716664/meme_approved_2.mp4"
        } else {
           meme.URL = "https://cdn.discordapp.com/attachments/702153501480058890/702153568689455154/Meme_Approved-1.mp4"
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

	s.ChannelMessageSendEmbed(m.ChannelID, memeReviewEmbed)
	s.ChannelMessageSend(m.ChannelID, meme.URL)
	//log.Println("memereview sending " + memeerr.Error())
}

type memeReview struct {
	URL   string
	Type string
	Time string
	Random int
}
