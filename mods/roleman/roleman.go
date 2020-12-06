package roleman

import (
    "github.com/bwmarrin/discordgo"
)

func cmdRoleAdd(message []string, s *discordgo.Session, m *discordgo.MessageCreate) {
  s.ChannelMessageSend(m.ChannelID, "a")
}
