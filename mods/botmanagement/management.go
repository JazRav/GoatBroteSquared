package manage

import (
    "strconv"
    "strings"
    "sort"

    log "github.com/Sirupsen/logrus"
    "github.com/bwmarrin/discordgo"
    "github.com/go-ini/ini"


    "github.com/dokvis/goatbrotesquared/cmd"
    "github.com/dokvis/goatbrotesquared/util/gvars"
    "github.com/dokvis/goatbrotesquared/util/tools"
    "github.com/dokvis/goatbrotesquared/util/tools/discord"
)

//Load - Loads the Management plugin
func Load() {
  log.Println("Loading Manage Plugin")
  manageCat := "Management"
  cmd.Make("devmode", manageCat, cmdDevModeToggle).Owner().Add()
  cmd.Make("logmode", manageCat, cmdLogModeToggle).Owner().Add()
  cmd.Make("cfgreload", manageCat, cmdCfgRefresh).Owner().Add()
  cmd.Make("listguilds", manageCat, cmdGuildList).Owner().Add()
  cmd.Make("listchans", manageCat, cmdChanList).Owner().Add()
  cmd.Make("msgchan", manageCat, cmdMsgChan).Owner().Add()
  cmd.Make("makeinvite", manageCat, cmdMakeInvite).Owner().Add()
  cmd.Make("status", manageCat, cmdChangeStatus).HelpText("`status message` `type` `url` `reset?`\nStatus Message: message to show as status\nType: 0 - None, 1 - Playing, 2 - Listening to, 3 - Streaming").Owner().Add()
  //User Commands
  cmd.Make("about", "Info", cmdVersion).HelpText("gets version infomation of bot").Add()
  cmd.Make("help", "Info", cmdHelp).HelpText("i dunno what this does").Add()
  cmd.Make("owner", "Info", cmdOwnerTag).HelpText("Pings owner").Add()
}

func cmdDevModeToggle(message []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	if gvars.DevMode == false {
		gvars.DevMode = true
		gvars.CFG.Section("bot").Key("dev_mode").SetValue("true")
		gvars.CFG.SaveTo(gvars.ConfigFile)
		s.ChannelMessageSend(m.ChannelID, "DEV MODE ENABLED")
		log.Println("DEV MODE ENABLED")
	} else {
		gvars.DevMode = false
		gvars.CFG.Section("bot").Key("dev_mode").SetValue("false")
		gvars.CFG.SaveTo(gvars.ConfigFile)
		s.ChannelMessageSend(m.ChannelID, "DEV MODE DISABLED")
		log.Println("DEV MODE DISABLED")
	}
}
func cmdLogModeToggle(message []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	if gvars.LogAll == false {
		gvars.LogAll = true
		gvars.CFG.Section("bot").Key("logall").SetValue("true")
		gvars.CFG.SaveTo(gvars.ConfigFile)
		s.ChannelMessageSend(m.ChannelID, "LOG ENABLED")
		log.Println("LOG ENABLED")
	} else {
		gvars.LogAll = false
		gvars.CFG.Section("bot").Key("logall").SetValue("false")
		gvars.CFG.SaveTo(gvars.ConfigFile)
		s.ChannelMessageSend(m.ChannelID, "LOG DISABLED")
		log.Println("LOG DISABLED")
	}
}
func cmdGuildList(message []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	guildIDLen := len(gvars.GuildID)
	msg := ""
	var guild *discordgo.Guild
	var err error
	if gvars.GuildID[0] != "" {
		guild, err = s.Guild(gvars.GuildID[0])
		if err != nil {
			return
		}
		msg = guild.Name + " (" + gvars.GuildID[0] + ")"
	}
	for i := 1; i < guildIDLen; i++ {
		guild, err = s.Guild(gvars.GuildID[i])
		if err != nil {
			return
		}
		msg = msg + "\n" + guild.Name + " (" + gvars.GuildID[i] + ")"
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
	guild, _ = s.Guild(gid)
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
func cmdCfgRefresh(message []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	var err error
	gvars.CFG, err = ini.Load(gvars.ConfigFile)
	if err != nil {
		log.Errorf("FAILED TO READ BOT INI FILE: %v\n\n\n\nMAKE SURE ITS IN BOTDIR\\config\\bot.ini", err)
		s.ChannelMessageSend(m.ChannelID, "Failed to refresh config")
		return
	}
	s.ChannelMessageSend(m.ChannelID, "Refreshed config file")
}

func cmdChangeStatus(message []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	if len(message) == 1 {
		s.ChannelMessageSend(m.ChannelID, "Not enough params\nRun help on this command to see details")
		return
	}
	_message := ""
	_messageType := gvars.StatusType
	_reset := false
	var err error
	if len(message) >= 2 {
		if len(message) > 2 {
			_message = strings.TrimPrefix(m.Content, message[0]+" ")
			_messageType, err = strconv.Atoi(message[len(message)-1])
			if err != nil {
				_messageType = gvars.StatusType
			} else {
				_message = strings.TrimSuffix(_message, message[len(message)-1])
			}
		} else if len(message) == 2 {
			_messageType, err = strconv.Atoi(message[1])
			if err != nil {
				_message = strings.TrimPrefix(m.Content, message[0]+" ")
			}
		}
	}
	discordTools.ChangeStatus(s, _message, _messageType, gvars.StatusURL, _reset)
	gvars.CFG.Section("bot").Key("statusMessage").SetValue(_message)
	gvars.CFG.Section("bot").Key("statusType").SetValue(strconv.Itoa(_messageType))
	gvars.CFG.SaveTo(gvars.ConfigFile)
	s.ChannelMessageSend(m.ChannelID, "Changed status to " + _message )
}

//Might move these

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
		i, _ = s.ChannelInviteCreate(m.ChannelID, *i)
	}

	s.ChannelMessageSend(m.ChannelID, "discord.gg/"+i.Code)
}
func cmdVersion(message []string, s *discordgo.Session, m *discordgo.MessageCreate) {
  s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
      Color: 0xffffff,
      Author: &discordgo.MessageEmbedAuthor{
        Name: "GoatBrote²",
        IconURL: "https://i.imgur.com/XV4v1bT.png",
      },
      Description: "A Shitty Discord Bot",
      Title:      "Github",
      URL: "https://github.com/DokVis/GoatBroteSquared",
      Footer: &discordgo.MessageEmbedFooter{
        Text: "Bot by Dok#3678 | Uptime: " + tools.Uptime().String(),
      },
      Fields: []*discordgo.MessageEmbedField{
      &discordgo.MessageEmbedField{
          Name:   "Version",
          Value:  gvars.Version,
          Inline: true,
      },
      &discordgo.MessageEmbedField{
          Name:   "OS/Arch",
          Value:  gvars.BinaryOS+"/"+gvars.BinaryArch,
          Inline: true,
      },
      &discordgo.MessageEmbedField{
          Name:   "Githash",
          Value:  gvars.GitHash,
          Inline: true,
      },
      &discordgo.MessageEmbedField{
          Name:   "Build Time",
          Value:  gvars.BuildTime,
          Inline: true,
      },
      &discordgo.MessageEmbedField{
          Name:   "Discord Go Version",
          Value:  discordgo.VERSION,
          Inline: true,
      },
      &discordgo.MessageEmbedField{
          Name:   "Host Name",
          Value:  gvars.HostName,
          Inline: true,
      },
    },
  })
}

func cmdOwnerTag(message []string, s *discordgo.Session, m *discordgo.MessageCreate) {
  s.ChannelMessageSend(m.ChannelID, "<@" + gvars.Owner + ">")
}

func cmdHelp(message []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	if len(message) >= 2 {
		if cmdV, ok := cmd.Commands[strings.ToLower(message[1])]; ok {
			//log.Println("cmd = " + cmd.Name + " help = " + cmd.Help)
			s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
				Color: 0xffffff,
				Author: &discordgo.MessageEmbedAuthor{
					Name: "Help",
				},
				Description: cmdV.Help,
				Title:      gvars.Prefix + cmdV.Name+"\n`" + cmdV.Category + "`",
				//Timestamp:   time.Now().Format(time.RFC3339),
			})
		}

		return
	}
  var cats []string
  var des string
  for _, cmdC := range cmd.Commands {
    cats = append(cats, cmdC.Category)
  }
  cat := tools.UniqueSilce(cats)
  sort.Strings(cat)
  var catC []string
  for i := 0; i < len(cat); i++ {
    var cmds []string
    for _, cmdV := range cmd.Commands {
      if (cat[i] == cmdV.Category) && !cmdV.IsHidden && (!cmdV.IsOwner || m.Author.ID == gvars.Owner){
        cmds = append(cmds,gvars.Prefix + cmdV.Name)
      }
    }
    if len(cmds) != 0 {
      sort.Strings(cmds)
      des = "**"+cat[i]+"**" + "\n"
      des = des + strings.Join(cmds, ", ")
      catC = append(catC, des)
    }
  }


	_, err := s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Color:       0xffffff,
		Title:       "List of Commands",
		Description: strings.Join(catC, "\n\n"),
	})
	if err != nil {
		log.Errorln("Help Embed Error: " + err.Error())
	}

}
