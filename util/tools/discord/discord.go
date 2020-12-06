package discordTools

import (
  "time"
  "os"
  "strconv"
  "github.com/bwmarrin/discordgo"
  	log "github.com/Sirupsen/logrus"
  "github.com/ashfennix/goatbrotesquared/util/tools"
)

//ChangeStatus - Changes status of Discord bot
func ChangeStatus(s *discordgo.Session, message string, messageType int, messageURL string, reset bool) (worked bool) {
	switch messageType {
 	case 1:
		if reset {
				s.UpdateListeningStatus("")
				s.UpdateStreamingStatus(0, "", "")
		}
		s.UpdateStatus(0, message)
	case 2:
		if reset {
			s.UpdateStatus(0, "")
			s.UpdateStreamingStatus(0, "", "")
		}
		s.UpdateListeningStatus(message)
	case 3:
		if reset {
			s.UpdateStatus(0, "")
			s.UpdateListeningStatus("")
		}
		s.UpdateStreamingStatus(0, message, messageURL)
	default:
		s.UpdateStatus(0, "")
		s.UpdateListeningStatus("")
		s.UpdateStreamingStatus(0, "", "")
	}
	return false
}

func LogThatShit(s *discordgo.Session, m *discordgo.MessageCreate) {
	chanInfo, _ := s.Channel(m.ChannelID)
	currentTime := time.Now()
	logPath := "logs/"+ m.GuildID + "("+GetNameFromGID(m.Message.GuildID, s)+")" + "/" + m.ChannelID + "("+GetNameFromCID(m.Message.ChannelID, s)+")/"
	if chanInfo.GuildID == "" {
		logPath = "logs/DM/" + m.Author.ID + "("+m.Author.Username+"#"+m.Author.Discriminator+")/"
	}
	logLocation := logPath + 	currentTime.Format("2006-1-02") + ".log"

	logDirExist, logDirErr := tools.DirExists(logPath)
	if logDirErr != nil {
		  log.Errorf("%v", logDirErr)
	}
	if !logDirExist && logDirErr == nil {
		os.MkdirAll(logPath, os.ModePerm)
	}

	logFile, err := os.OpenFile(logLocation, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)

	if err != nil {
    log.Errorf("error opening file: %v", err)
	}
	defer logFile.Close()
	messTime, _ := m.Message.Timestamp.Parse()
	attachMessage := ""
	for i := 0; i < len(m.Attachments); i++ {
		attachMessage = attachMessage + "\n["+strconv.Itoa(i+1)+"]"+m.Attachments[i].URL
	}
	content := ""
	if m.Content == "" {
		 content = ""
	} else {
		content = m.Content + "\n"
	}
	if len(m.Attachments) > 0 {
		attachMessage = "Attachment(s): " + attachMessage + "\n"
	}
	if _, fileErr := logFile.WriteString(m.Author.Username + "#" + m.Author.Discriminator +" (" + m.Author.ID + ") - " + m.ChannelID + "-" + m.ID + ": " + 	messTime.String() +"\n"+content+attachMessage+"\n"); fileErr != nil {
		log.Println(fileErr)
	}
}

func GetNameFromGID(id string, s *discordgo.Session) (name string){
	guildVar, guilderr :=s.Guild(id)
	if guilderr != nil {
		name = "Error: Name Not Found"
	} else {
		name = guildVar.Name
	}
	return name
}

func GetNameFromCID(id string, s *discordgo.Session) (name string){
	chanVar, chanerr :=s.Channel(id)
	if chanerr != nil {
		name = "Error: Name Not Found"
	} else {
		name = chanVar.Name
	}
	return name
}
