package main

import (
  //Built In
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
  "flag"
	"fmt"

  //Imported
	log "github.com/Sirupsen/logrus"
	"github.com/bwmarrin/discordgo"

  //Project
  "github.com/ashfennix/goatbrotesquared/cmd/handler"
  "github.com/ashfennix/goatbrotesquared/util/gvars"
  "github.com/ashfennix/goatbrotesquared/util/gini"
  "github.com/ashfennix/goatbrotesquared/util/guildini"
	"github.com/ashfennix/goatbrotesquared/util/tools"
	"github.com/ashfennix/goatbrotesquared/util/tools/discord"
)
func main() {
	//Starts uptime shit
	tools.StartTheTimer()
	consoleFormat()
	verInfoSet()
  CFGFile := flag.String("c", "", "Config file name")
	flag.Parse()
	if *CFGFile != "" {
		gvars.ConfigFile = "data/config/" + *CFGFile
		gvars.ConfigFileName = strings.TrimSuffix(*CFGFile, ".ini")
		log.Println("Loading custom config: " + gvars.ConfigFile)
	} else {
		log.Println("Loading normal config")
		gvars.ConfigFile = "data/config/bot.ini"
		gvars.ConfigFileName = "bot"
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
  cmdHandle.Load()
	dg.AddHandler(ready)
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

func ready(s *discordgo.Session, event *discordgo.Ready) {

}

func guildCreate(s *discordgo.Session, event *discordgo.GuildCreate) {
	if event.Guild.Unavailable {
		log.Println("Guild " + event.Guild.ID + " is unavailable")
		return
	}
	guildINI.MakeGuildIni(event.Guild.ID)
	log.Printf("Connected to guild: " + event.Guild.Name + " (" + event.Guild.ID + ")")
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
  message := strings.Fields(m.Content)
	if m.Content != "" {
		if strings.HasPrefix(message[0], gvars.Prefix) {
			cmdHandle.Handle(message, s, m)
		}
	}
	if gvars.LogAll == true {
		discordTools.LogThatShit(s, m)
		log.Println(discordTools.GetNameFromGID(m.Message.GuildID, s) + " (" + discordTools.GetNameFromCID(m.Message.ChannelID, s) + "): " + m.Author.Username + ": " + m.Content)
	}
}
//PlainFormatter - Stolen code from some stackoverflow shit - https://stackoverflow.com/questions/43022607/how-do-i-disable-field-name-in-logrus-while-logging-to-file
type PlainFormatter struct {
    TimestampFormat string
    LevelDesc []string
}

func verInfoSet(){
	HostNameCmd := exec.Command("hostname")
	HostNameSTD, HostNameErr := HostNameCmd.Output()
	if HostNameErr != nil {
		gvars.HostName = "unknown"
	} else {
		gvars.HostName = string(HostNameSTD)
	}
	gvars.Version = Version
	gvars.BinaryOS = BinaryOS
	gvars.BinaryArch = BinaryArch
	gvars.BuildTime = BuildTime
	gvars.GitHash = GitHash
	log.Println("Loading \u001B[35mGoatBrote\033[0m²")
	log.Println("Version: "+gvars.Version + "-" + gvars.BinaryOS + "-" + gvars.BinaryArch)
	log.Println("GitHash: " + gvars.GitHash)
	log.Println("Build Time:" + gvars.BuildTime)
}
//Format - aaaaaa
func (f *PlainFormatter) Format(entry *log.Entry) ([]byte, error) {
    timestamp := fmt.Sprintf(entry.Time.Format(f.TimestampFormat))
		var level string
		switch f.LevelDesc[entry.Level] {
			case "ERRO" : level = "\033[1;31m"+f.LevelDesc[entry.Level]+"\033[0m"
			case "FATL" : level = "\033[1;32m"+f.LevelDesc[entry.Level]+"\033[0m"
			case "DEBG" : level = "\u001B[35m"+f.LevelDesc[entry.Level]+"\033[0m"
			default: level = "\033[1;34m"+f.LevelDesc[entry.Level]+"\033[0m"
		}
    return []byte(fmt.Sprintf("%s %s %s\n", level, timestamp, "\t"+entry.Message)), nil
}

func consoleFormat() {
	plainFormatter := new(PlainFormatter)
	plainFormatter.TimestampFormat = "2006-01-02 15:04:05"
	plainFormatter.LevelDesc = []string{"PANC", "FATL", "ERRO", "WARN", "INFO", "DEBG"}
	log.SetFormatter(plainFormatter)
}

var (
	//Don't use this shit

	//Version - a
	Version   = "dev"
	//BinaryOS - a
	BinaryOS = "dev"
	//BinaryArch - a
	BinaryArch = "dev"
	//BuildTime - a
	BuildTime = "before time"
	//GitHash - a
	GitHash   = "what the fuck is a git?"
)
