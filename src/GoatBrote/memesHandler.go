package main

import (
	"bufio"
	"os"
	"regexp"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/bwmarrin/discordgo"
)

func memesHandler(message []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	matched, _ := regexp.MatchString("^sans.*fucked.*(mom|mum)$", strings.ToLower(m.Content))
	if matched {
		log.Printf("sans fucked " + m.Author.Username + " mom")
		img, err := os.Open("images/mom.png")
		if err != nil {
			log.Printf("MOM PNG MISSING! PANIC!")
			s.ChannelMessageSend(m.ChannelID, "FILE MISSING, JOKE RUINED")
			return
		}
		momImage := bufio.NewReader(img)
		if devMode == true && logAll == false {
			log.Println("GID-CID: " + m.Message.GuildID + "-" + m.Message.ChannelID + "\t" + m.Author.Username + "(" + m.Author.ID + "): \"" + m.Content + "\"" + img.Name())
		}

		s.ChannelFileSend(m.ChannelID, "mom.png", momImage)
	}
}
