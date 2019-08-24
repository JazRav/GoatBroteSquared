package main

import (
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/bwmarrin/discordgo"
	"github.com/go-ini/ini"
)

func init() {
	//Owner Commands
	makeCmd("devmode", cmdDevModeToggle).owner().add()
	makeCmd("logmode", cmdLogModeToggle).owner().add()
	makeCmd("cfgreload", cmdCfgRefresh).owner().add()
	makeCmd("listguilds", cmdGuildList).owner().add()
	makeCmd("listchans", cmdChanList).owner().add()
	makeCmd("msgchan", cmdMsgChan).owner().add()
	makeCmd("dankmemes", cmdMemeToggle).owner().add()
	makeCmd("e6Sample", cmde621SampleToggle).owner().add()
	makeCmd("makeinvite", cmdMakeInvite).owner().add()
	makeCmd("e6filter", cmde621FilterToggle).owner().add()
	makeCmd("e6filterscore", cmde621FilterScore).owner().add()

	//User Commands
	makeCmd("ver", cmdVersion).helpText("gets version infomation of bot").add()
	makeCmd("help", cmdHelp).helpText("i dunno what this does").add()
}

func cmdDevModeToggle(message []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	if devMode == false {
		devMode = true
		cfg.Section("bot").Key("dev_mode").SetValue("true")
		cfg.SaveTo(cfgFile)
		s.ChannelMessageSend(m.ChannelID, "DEV MODE ENABLED")
		log.Println("DEV MODE ENABLED")
	} else {
		devMode = false
		cfg.Section("bot").Key("dev_mode").SetValue("false")
		cfg.SaveTo(cfgFile)
		s.ChannelMessageSend(m.ChannelID, "DEV MODE DISABLED")
		log.Println("DEV MODE DISABLED")
	}
}

func cmdLogModeToggle(message []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	if logAll == false {
		logAll = true
		cfg.Section("bot").Key("logall").SetValue("true")
		cfg.SaveTo(cfgFile)
		s.ChannelMessageSend(m.ChannelID, "LOG ENABLED")
		log.Println("LOGALL ENABLED")
	} else {
		logAll = false
		cfg.Section("bot").Key("logall").SetValue("false")
		cfg.SaveTo(cfgFile)
		s.ChannelMessageSend(m.ChannelID, "LOG DISABLED")
		log.Println("LOGALL DISABLED")
	}
}

func cmdCfgRefresh(message []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	var err error
	cfg, err = ini.Load(cfgFile)
	if err != nil {
		log.Printf("FAILED TO READ BOT INI FILE: %v\n\n\n\nMAKE SURE ITS IN BOTDIR\\config\\bot.ini", err)
		s.ChannelMessageSend(m.ChannelID, "Failed to refresh config")
		return
	}
	s.ChannelMessageSend(m.ChannelID, "Refreshed config file")
}

func cmdGuildList(message []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	guildIDLen := len(guildID)
	msg := ""
	var guild *discordgo.Guild
	var err error
	if guildID[0] != "" {
		guild, err = s.Guild(guildID[0])
		if err != nil {
			return
		}
		msg = guild.Name + " (" + guildID[0] + ")"
	}
	for i := 1; i < guildIDLen; i++ {
		guild, err = s.Guild(guildID[i])
		if err != nil {
			return
		}
		msg = msg + "\n" + guild.Name + " (" + guildID[i] + ")"
	}
	s.ChannelMessageSend(m.ChannelID, msg)
}

func cmdChanList(message []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	if len(message) < 2 {
		s.ChannelMessageSend(m.ChannelID, "Not enough params")
		return
	}
	gid := message[1]
	chans, err := s.GuildChannels(gid)
	var msg string
	var guild *discordgo.Guild
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Not vaild")
		return
	}
	var c *discordgo.Channel
	for _, c = range chans {
		var chanType string
		if c.Type == discordgo.ChannelTypeGuildText {
			chanType = "text"
		} else {
			chanType = "voice"
		}
		msg = msg + "\n" + c.Name + " (" + c.ID + ") Type: " + chanType

	}
	guild, err = s.Guild(gid)
	s.ChannelMessageSend(m.ChannelID, guild.Name+" has the channels\n-------------"+msg)
}
func cmdMsgChan(message []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	if len(message) <= 2 {
		s.ChannelMessageSend(m.ChannelID, "Not enough params")
		return
	}
	chanid := message[1]
	msg := strings.TrimPrefix(m.Content, message[0]+" "+chanid)
	s.ChannelMessageSend(chanid, msg)
	s.ChannelMessageSend(m.ChannelID, chanid+" was sent the message: "+msg)
}

func cmdMemeToggle(message []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	if dankmemes == false {
		dankmemes = true
		cfg.Section("bot").Key("dank_memes").SetValue("true")
		cfg.SaveTo(cfgFile)
		s.ChannelMessageSend(m.ChannelID, "DANK MEMES ENABLED")
		log.Println("DANK MEMES ENABLED")
	} else {
		dankmemes = false
		cfg.Section("bot").Key("dank_memes").SetValue("false")
		cfg.SaveTo(cfgFile)
		s.ChannelMessageSend(m.ChannelID, "DANK MEMES DISABLED")
		log.Println("DANK MEMES DISABLED")
	}
}

func cmde621SampleToggle(message []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	if e6Sample == false {
		e6Sample = true
		cfg.Section("bot").Key("e621Sample").SetValue("true")
		cfg.SaveTo(cfgFile)
		s.ChannelMessageSend(m.ChannelID, "e621/e926 SAMPLE ENABLED")
		log.Println("e621 SAMPLE ENABLED")
	} else {
		e6Sample = false
		cfg.Section("bot").Key("e621Sample").SetValue("false")
		cfg.SaveTo(cfgFile)
		s.ChannelMessageSend(m.ChannelID, "e621/e926 SAMPLE DISABLED")
		log.Println("e621 SAMPLE DISABLED")
	}
}

func cmde621FilterToggle(message []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	if e6Filter == false {
		e6Filter = true
		cfg.Section("bot").Key("e621Filter").SetValue("true")
		cfg.SaveTo(cfgFile)
		s.ChannelMessageSend(m.ChannelID, "e621/e926 FILTER ENABLED")
		log.Println("e621 FILTER ENABLED")
	} else {
		e6Filter = false
		cfg.Section("bot").Key("e621Filter").SetValue("false")
		cfg.SaveTo(cfgFile)
		s.ChannelMessageSend(m.ChannelID, "e621/e926 FILTER DISABLED")
		log.Println("e621 FILTER DISABLED")
	}
}

func cmde621FilterScore(message []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	score := strings.TrimPrefix(m.Content, message[0] + " ")
	if e6Filter == false {
		s.ChannelMessageSend(m.ChannelID, "FILTER DISABLED, PLEASE ENABLE FILTER WITH ``" + prefix + "e6filter`")
		log.Println("e621 FILTER ENABLED")
	} else {
		var extraMessage string
		extraMessage = ""
		if score == "<e6filterscore"	{
			score = "2"
			extraMessage = "NOTHING SET, SETTING "
		}
		e6FilterScore = score
		cfg.Section("bot").Key("e621FilterScore").SetValue(e6FilterScore)
		s.ChannelMessageSend(m.ChannelID, extraMessage +"FILTER SCORE TO " + e6FilterScore)
		log.Println("e621 FILTER SCORE SET TO " + e6FilterScore)
	}
}

func cmdMakeInvite(message []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	i := &discordgo.Invite{
		MaxAge: 0,
		Uses:   0,
	}
	var err error
	if len(message) >= 2 {

		i, err = s.ChannelInviteCreate(message[1], *i)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Invite creation failed")
			return
		}

	} else {
		i, err = s.ChannelInviteCreate(m.ChannelID, *i)
	}

	s.ChannelMessageSend(m.ChannelID, "discord.gg/"+i.Code)
}

func cmdVersion(message []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "Version: "+Version+"\nBuild Time: "+BuildTime+"\nGitHash: "+GitHash)
}

func cmdHelp(message []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	if len(message) == 2 {
		if cmd, ok := commands[strings.ToLower(message[1])]; ok {
			//log.Println("cmd = " + cmd.Name + " help = " + cmd.Help)
			s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
				Color: 0xffffff,
				Author: &discordgo.MessageEmbedAuthor{
					Name: "Help",
				},
				Description: cmd.Help,
				Title:       prefix + cmd.Name,
				//Timestamp:   time.Now().Format(time.RFC3339),
			})
		}
		return
	}
	var cmds []string
	for _, cmd := range commands {
		if m.Author.ID == ownerID || !cmd.Owner {
			cmds = append(cmds, prefix+cmd.Name)
		}
	}
	_, err := s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Color:       0xffffff,
		Title:       "List of Commands",
		Description: strings.Join(cmds, ", "),
	})
	if err != nil {
		log.Println("Help Embed Error: " + err.Error())
	}

}
