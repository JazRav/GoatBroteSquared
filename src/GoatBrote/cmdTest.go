package main

import "github.com/bwmarrin/discordgo"

func init() {
	makeCmd("test", cmdTest).helpText("this is a test").add()
}
func cmdTest(message []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "Works")
}
