package main

import (
  //Built In
	"os"
	"os/signal"
	"strings"
	"syscall"
  "flag"

  //Imported
	log "github.com/Sirupsen/logrus"
	"github.com/bwmarrin/discordgo"

  //Project
  "github.com/dokvis/goatbrotesquared/cmd/handler"
  "github.com/dokvis/goatbrotesquared/data/gvars"
  "github.com/dokvis/goatbrotesquared/data/gini"
)

func main() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
  CFGFile := flag.String("c", "", "Config file name")
	flag.Parse()
	if *CFGFile != "" {
		gvars.ConfigFile = "config/" + *CFGFile
		log.Println("Loading custom config: " + gvars.ConfigFile)
	} else {
		log.Println("Loading normal config")
		gvars.ConfigFile = "config/bot.ini"
	}
  aerr := gini.Init()
  if aerr != nil {
    return
  }
  dg, err := discordgo.New("Bot " + gvars.BotToken)
  if err != nil {
		log.Println("Error starting Discord session: ", err)
		return
	}
  gvars.Prefix = "<"
  cmdHandle.Load()
	dg.AddHandler(Ready)
	dg.AddHandler(messageCreate)
	dg.AddHandler(guildCreate)
	//dg.AddHandler(comesFromDM)

	err = dg.Open()
	if err != nil {
		log.Println("Error starting Discord session: ", err)
	}
  //log.Println(cmddo.Commands)
	//Kills bot with CTRL - C
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-c
	tempRemoveErr := os.RemoveAll("temp")
	if tempRemoveErr != nil {
		log.Println("Temp folder failed to delete: " + tempRemoveErr.Error())
	} else {
		log.Println("Deleted temp folder")
	}
	log.Println("KILL SIGNAL DETECTED! Closing Discord Session")
	dg.Close()
	log.Println("Closed Discord Session")
}

func Ready(s *discordgo.Session, event *discordgo.Ready) {

}

func guildCreate(s *discordgo.Session, event *discordgo.GuildCreate) {
	if event.Guild.Unavailable {
		log.Println("Guild " + event.Guild.ID + " is unavailable")
		return
	}
	log.Printf("Connected to guild: " + event.Guild.Name + " (" + event.Guild.ID + ")")
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
  message := strings.Fields(m.Content)
	if m.Content != "" {
		if strings.HasPrefix(message[0], gvars.Prefix) {
			cmdHandle.Handle(message, s, m)
		}
	}
}
