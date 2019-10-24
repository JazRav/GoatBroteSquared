package main

import (
	"flag"
	"io"
	"net/http"
	"os"
	"os/signal"
	"os/exec"
	"strconv"
	"strings"
	"syscall"

	log "github.com/Sirupsen/logrus"
	"github.com/bwmarrin/discordgo"
	"github.com/go-ini/ini"
)

func main() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	log.Println("Bot is starting, Version: " + Version + " Build Time: " + BuildTime + " GitHash: " + GitHash)
	CFGFile := flag.String("c", "", "Config file name")
	flag.Parse()
	if *CFGFile != "" {
		cfgFile = "config/" + *CFGFile
		log.Println("Loading custom config" + cfgFile)
	} else {
		log.Println("Loading normal config")
		cfgFile = "config/bot.ini"
	}
	//Reads INI
	var errINI error
	cfg, errINI = ini.Load(cfgFile)
	if errINI != nil {
		log.Printf("FAILED TO READ BOT INI FILE: %v\n\n\n\nMAKE SURE ITS IN BOTDIR\\config\\bot.ini", errINI)
		return
	}

	//Gets Token
	botToken = cfg.Section("auth").Key("bot_token").String()

	//Gets OwnerID
	ownerID = cfg.Section("auth").Key("owner_id").String()
	if ownerID == "" {
		log.Println("Failed to get Owner ID, using hard coded creator ID for Dok#3678")
		ownerID = "84695583679315968" //ID of Dok#3678
	}
	//Gets Devmode status
	var Err error
	devMode, Err = cfg.Section("bot").Key("dev_mode").Bool()
	if Err != nil {
		log.Print("INI ERROR: dev_mode not set to BOOL value, setting false")
		devMode = false
	}

	//Gets LogALL status
	logAll, Err = cfg.Section("bot").Key("logall").Bool()
	if Err != nil {
		log.Print("INI ERROR: logALL not set to BOOL value, setting false")
		logAll = false
	}
	//Checks if memes are dank
	dankmemes, Err = cfg.Section("bot").Key("dank_memes").Bool()
	if Err != nil {
		log.Print("INI ERROR: logALL not set to BOOL value, setting false")
		dankmemes = false
	}

	e6Sample, Err = cfg.Section("bot").Key("e621Sample").Bool()
	if Err != nil {
		log.Print("INI ERROR: e621Sample not set to BOOL value, setting true")
		e6Sample = true
	}

	e6Filter, Err = cfg.Section("bot").Key("e621Filter").Bool()
	if Err != nil {
		log.Print("INI ERROR: e621Fitler not set to BOOL value, setting false")
		e6Filter = false
	}

	e6FilterScore = cfg.Section("bot").Key("e621FilterScore").String()
	if e6FilterScore == "" {
		log.Print("INI ERROR: e621FitlerScore not set, setting to 2")
		e6FilterScore = "2"
	}
	//attempts to start a discord session
	dg, err := discordgo.New("Bot " + botToken)
	if err != nil {
		log.Println("Error starting Discord session: ", err)
		return
	}

	//Adds handlers
	HostNameCmd := exec.Command("hostname")
	HostNameSTD, HostNameErr := HostNameCmd.Output()
	if HostNameErr != nil {
		HostName = "unknown"
	} else {
		HostName = string(HostNameSTD)
	}
	dg.AddHandler(Ready)
	dg.AddHandler(messageCreate)
	dg.AddHandler(guildCreate)
	//dg.AddHandler(comesFromDM)

	err = dg.Open()
	if err != nil {
		log.Println("Error starting Discord session: ", err)
	}
	log.Println("Loading Commands")
	Cmds()
	log.Println("Bot has started, use CTRL-C to kill")
	log.Println("DEVMODE: \t\t" + strconv.FormatBool(devMode))

	log.Println("LogAll (why???): \t" + strconv.FormatBool(logAll))
	//Kills bot with CTRL - C
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-c
	log.Println("KILL SIGNAL DETECTED! Closing Discord Session")
	dg.Close()
	log.Println("Closed Discord Session")
}

func Ready(s *discordgo.Session, event *discordgo.Ready) {
	s.UpdateListeningStatus("Don't Forget")
}

func guildCreate(s *discordgo.Session, event *discordgo.GuildCreate) {
	if event.Guild.Unavailable {
		log.Println("Guild " + event.Guild.ID + " is unavailable")
		return
	}
	log.Printf("Connected to guild: " + event.Guild.Name + " (" + event.Guild.ID + ")")
	invites, err := cfg.Section("bot").Key("make_invites").Bool()
	guildID = append(guildID, event.Guild.ID)
	if err != nil {
		invites = false
	}
	if invites {
		var i discordgo.Invite
		i.MaxAge = 0
		i.MaxUses = 0
		i.Temporary = false
		invite, err := s.ChannelInviteCreate(event.Guild.SystemChannelID, i)
		if err != nil {
			log.Printf("No invite made")
			return
		}
		log.Printf("discord.gg/" + invite.Code)
	}
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	//Logs messages
	globalPrefix = cfg.Section("bot").Key("globalPrefix").String()
	if globalPrefix == "" || globalPrefix == "\\" {
		cfg.Section("bot").Key("globalPrefix").SetValue(">")
		cfg.SaveTo(cfgFile)
		log.Printf("globalPrefix was broken, fixed")

	}
	if logAll == true {
		log.Println(getNameFromGID(m.Message.GuildID, s) + " (" + getNameFromSID(m.Message.ChannelID, s) + "): " + m.Author.Username + ": " + m.Content)
	}
	//If any bot is the author of the message, ignore.
	if m.Author.ID == s.State.User.ID || m.Author.Bot {
		return
	}
	dmCheck, dmErr := comesFromDM(s, m)
	if dmErr != nil {
		log.Printf("DM ERROR FROM "+m.Author.Username+"("+m.Author.ID+"): ", dmErr)
	}
	if dmCheck == true {
		if devMode == true && m.Author.ID != ownerID {
			log.Printf("USER " + m.Author.Username + "(" + m.Author.ID + ") sent a DM: \"" + m.Content + "\"")
			return
		}
		if m.Author.ID != ownerID {
			return
		}
	}
	//Sends messages to commands with prefixs
	message := strings.Fields(m.Content)
	if m.Content != "" {
		if strings.HasPrefix(message[0], globalPrefix) {
			cmdhandle(message, s, m)
		}
	}
	if dankmemes {
		memesHandler(message, s, m)
	}
}

func comesFromDM(s *discordgo.Session, m *discordgo.MessageCreate) (bool, error) {
	channel, err := s.State.Channel(m.ChannelID)
	if err != nil {
		if channel, err = s.Channel(m.ChannelID); err != nil {
			return false, err
		}
	}

	return channel.Type == discordgo.ChannelTypeDM, nil
}
func fileGetter(url string, file string) (err error) {
	log.Println("file get url: " + url)
	mkfile, err := os.Create(file)
	if err != nil {
		return err
	}
	defer mkfile.Close()

	/* Old http get
	data, err := http.Get(url)
	if err != nil {
		return err
	}
	defer data.Body.Close()
	*/
	client := &http.Client{}
	fileGet, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	fileGet.Header.Set("User-Agent", "GoatBroteSquared_DiscordGo_Bot/"+Version)
	fileResp, err := client.Do(fileGet)
	if err != nil {
		log.Printf("Failed to get some shit")
		return err
	}
	defer fileResp.Body.Close()
	log.Println(fileResp.Body)
	_, err = io.Copy(mkfile, fileResp.Body)
	return nil
}

func getNameFromSID(id string, s *discordgo.Session) (name string){
	chanVar, chanerr :=s.Channel(id)
	if chanerr != nil {
		name = "Error: Name Not Found"
	} else {
		name = chanVar.Name
	}
	return name
}

func getNameFromGID(id string, s *discordgo.Session) (name string){
	guildVar, guilderr :=s.Guild(id)
	if guilderr != nil {
		name = "Error: Name Not Found"
	} else {
		name = guildVar.Name
	}
	return name
}
