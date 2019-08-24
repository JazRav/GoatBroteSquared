package main

import (
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/bwmarrin/discordgo"
)

/*Below is where you put active commands*/
func Cmds() {

}

type command struct {
	Name  string
	Help  string
	Owner bool
	Exec  func([]string, *discordgo.Session, *discordgo.MessageCreate)
	Alias []string
}

func cmdhandle(message []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	cmd := strings.TrimPrefix(message[0], prefix)
	cmd = strings.ToLower(cmd)

	if command, ok := commands[cmd]; ok && (cmd == strings.ToLower(command.Name)) {

		if devMode {
			location := "borked"
			guild, err := s.Guild(m.GuildID)
			if err != nil {
				log.Println("No Guild ID from command location")
			} else {
				location = "(" + guild.Name + ")"
			}
			channel, err := s.Channel(m.ChannelID)
			if err != nil {
				log.Println("No Channel ID from command location")
			} else {
				location = channel.Name + location
			}

			log.Println("User " + m.Author.Username + "#" + m.Author.Discriminator + " ran command " + cmd + " in " + location)
		}
		isOwner := m.Author.ID == ownerID
		if !command.Owner || isOwner {
			command.Exec(message, s, m)
			return
		}
		s.ChannelMessageSend(m.ChannelID, "You don't have the correct permissions to run this!")
		return
	}
}

func (cmd command) add() command {
	commands[strings.ToLower(cmd.Name)] = cmd
	return cmd
}

func makeCmd(name string, fun func([]string, *discordgo.Session, *discordgo.MessageCreate)) command {
	return command{
		Name: name,
		Exec: fun,
	}
}

func (cmd command) owner() command {
	cmd.Owner = true
	return cmd
}

func (cmd command) helpText(helpText string) command {
	cmd.Help = helpText
	return cmd
}

func (cmd command) alias(alias []string) command {
	cmd.Alias = alias
	return cmd
}
