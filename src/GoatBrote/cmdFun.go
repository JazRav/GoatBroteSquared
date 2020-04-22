package main

import (
	"math/rand"
	"github.com/bwmarrin/discordgo"
	//log "github.com/Sirupsen/logrus"
)

func init() {
	makeCmd("meme", cmdMemeReview).helpText("Reviews meme").add()
  makeCmd("memereview", cmdMemeReview).helpText("Reviews meme").add()
  makeCmd("mr", cmdMemeReview).helpText("Reviews meme").add()
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
	meme.Random = rand.Intn(100 - 0 + 1)
	switch {
		case meme.Random == 0: {
			meme.Type = "ERROR"
			meme.URL = "https://cdn.discordapp.com/attachments/702153501480058890/702576438993223711/meme_sorry.mp4"
		}
		case meme.Random > 0 && meme.Random < 34: {
				meme.Type = "Approved"
				meme.URL = "https://cdn.discordapp.com/attachments/702153501480058890/702153568689455154/Meme_Approved-1.mp4"
		}
		case meme.Random > 33 && meme.Random < 66: {
			meme.Type = "Unsure"
			meme.URL = "https://cdn.discordapp.com/attachments/702153501480058890/702153595583201369/Meme_Limbo-1.mp4"
		}
		case meme.Random > 65 && meme.Random < 101: {
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
