package hello

import (
  //Imported
  "github.com/bwmarrin/discordgo"
  log "github.com/Sirupsen/logrus"

  //Project
  "github.com/dokvis/goatbrotesquared/cmd"
)

func Load() {
    log.Println("Loading Hello World commands")
    cmd.Make("hello", "Hello World", cmdHelloWorld).HelpText("Hello World").Add()
}

func cmdHelloWorld(message []string, s *discordgo.Session, m *discordgo.MessageCreate) {
  log.Println("Ran Hello World")
  s.ChannelMessageSend(m.ChannelID, "Hello G-")
}
