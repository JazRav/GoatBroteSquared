package e621

import (
  //Built in
  "strconv"
  "strings"

  //Imported
  log "github.com/Sirupsen/logrus"
  "github.com/bwmarrin/discordgo"

  //Project
  "github.com/dokvis/goatbrotesquared/cmd"
  "github.com/dokvis/goatbrotesquared/util/gvars"
)
//Load - Loads mod
func Load(){

  log.Println("Loading e621 plugin")
  e621HelpMessage := "gives you a e621\\e926 image\ne621 in NSFW channels\ne926 in SFW channels\nput booru tags after command\nin DMs, add `NSFW` at end of tags for NSFW"
  cmd.Make("fur", "Furry", cmdE621).HelpText(e621HelpMessage).Add()
}

func cmdE621(message []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	search := ""
	if len(message) > 1 {
		search = strings.TrimPrefix(m.Content, message[0]+" ")
	}
	e621EmbedMessage(search, false, "", false, "", "", s, m)
}


func e621EmbedMessage(search string, idlookup bool, forcesearch string, nolewd bool, nolewdmessage string, nolewdid string, s *discordgo.Session, m *discordgo.MessageCreate){
	chanInfo, _ := s.Channel(m.ChannelID)
	if nolewd && chanInfo.NSFW {
		search = "id:"+nolewdid
		forcesearch = ""
	} else if (strings.Contains(search, " ralsei") && (chanInfo.NSFW || (chanInfo.GuildID == "" && strings.HasSuffix(search, " NSFW") ) ) ) {
		search = "id:"//+ralseiAntiLewd()
		forcesearch = ""
		nolewdmessage = "NO LEWDING THE GOAT YOU FUCK, AT ALL"
		nolewd = true
	}

	forceID := false
	if idlookup {
		search = "id:" +search
		forceID = true
	}
	nsfw := chanInfo.NSFW
	//log.Println("GID: "+chanInfo.GuildID + " CID: " +chanInfo.ID)
	//Should be a DM if GuildID is blank
	if chanInfo.GuildID == "" && strings.HasSuffix(search, "NSFW"){
		search = strings.TrimSuffix(search, " NSFW")
		nsfw = true
	}
	eStuff, err := e621Handler(search, forceID, forcesearch, nsfw, nolewd, "")
	if eStuff.Rating == "e" || eStuff.Rating == "q" {
		for a := 0; a < len(eStuff.Tags.Character); a++ {
			if eStuff.Tags.Character[a] == "ralsei" {
				search = "id:"+ralseiAntiLewd()
				forcesearch = ""
				nolewdmessage = "LEWD WITH GOAT DETECTED"
				if eStuff.Rating == "q" {
					nolewdmessage = "SEMI-LEWD WITH GOAT DETECTED"
				}
				nolewd = true
				eStuff, err = e621Handler(search, forceID, forcesearch, nsfw, nolewd, "")
				if err != nil {
					log.Println("fuck me it broke with error: " + err.Error())
					s.ChannelMessageSend(m.ChannelID, "fuck me it broke with error: "+err.Error())
					return
				}
			}
		}
		//Overkill, as the blacklist should catch cub content, but just in case, have FBI
		for a := 0; a < len(eStuff.Tags.General); a++ {
			if (eStuff.Tags.General[a] == "cub" || eStuff.Tags.General[a] == "young") {
				search = "id:2161983"
				forcesearch = ""
				nolewdmessage = "CHILD DETECTED, CONTACTING FBI"
				nolewd = true
				eStuff, err = e621Handler(search, forceID, forcesearch, nsfw, nolewd, "")
				if err != nil {
					log.Println("fuck me it broke with error: " + err.Error())
					s.ChannelMessageSend(m.ChannelID, "fuck me it broke with error: "+err.Error())
					return
				}
			}
		}
	}
	if err != nil {
		log.Println("fuck me it broke with error: " + err.Error())
		s.ChannelMessageSend(m.ChannelID, "fuck me it broke with error: "+err.Error())
		return
	}
	var link string
	var title string
	e6ORe9 := "e926"
	if chanInfo.NSFW {
		e6ORe9 = "e621"
		if nolewd {
			e6ORe9 = nolewdmessage
		}
	}
	if eStuff.Source != "" {
		link = eStuff.Source
		title = "Source"
	} else {
		link = eStuff.Page
		title = ""
	}
	var imageEmbed discordgo.MessageEmbedImage
	var videoEmbed discordgo.MessageEmbedVideo
	var clickMessage string
	if (eStuff.EXT == "webm") {
		videoEmbed = discordgo.MessageEmbedVideo{
			URL: eStuff.URL,
		}
	} else if (eStuff.EXT == "swf") {
		clickMessage = "\n\nFile is SWF, please click Source or " + e6ORe9 + " to view"
	} else {
		imageEmbed = discordgo.MessageEmbedImage{
			URL: eStuff.URL,
		}
	}
	soundWarning := ""
	if eStuff.SoundWarning {
		soundWarning = "\n⚠WARNING SOUND MIGHT BE LOUD⚠"
	}
	if eStuff.Page != "" {
		e621embed := &discordgo.MessageEmbed{
			Color:       0x0055ff,
			Description: "Artist: " + eStuff.Artist + "\nRating: " + eStuff.Rating + " Score: " + strconv.Itoa(eStuff.Score) + "\nID: " + strconv.Itoa(eStuff.ID) +clickMessage+soundWarning,
			URL:         link,
			Author: &discordgo.MessageEmbedAuthor{
				URL:     eStuff.Page,
				Name:    e6ORe9,
				IconURL: "https://i.imgur.com/dbKpPIs.png",
			},
			Image: &imageEmbed,
			Video: &videoEmbed,
			Title:     title,
			Timestamp: eStuff.TimeStamp,
		}
		s.ChannelMessageSendEmbed(m.ChannelID, e621embed)
		if eStuff.EXT == "webm" {
			s.ChannelMessageSend(m.ChannelID, eStuff.URL)
		}
		if gvars.DevMode {
				//s.ChannelMessageSend(m.ChannelID, "URL of Image:" + eStuff.URL)
		}
		return
	}
	if idlookup {
		nsfwMessage := ""
		if !chanInfo.NSFW {
			nsfwMessage = " or if its NSFW, make sure you are in a NSFW channel"
		}
		s.ChannelMessageSend(m.ChannelID, "Found nothing for `" + search + "`\nMake sure its an actual ID, and its not blacklisted" + nsfwMessage)
		//s.ChannelMessageSend(m.ChannelID, "DEBUG: ID:" + strconv.Itoa(eStuff.ID)+" URL: " + eStuff.URL)
	} else {
		s.ChannelMessageSend(m.ChannelID, "Found nothing for `"+search+"`\nMake sure names with spaces, like Katia Managan is spelt like `Katia_Managan`")
		//s.ChannelMessageSend(m.ChannelID, "DEBUG: ID:" + strconv.Itoa(eStuff.ID)+" URL: " + eStuff.URL)
	}

}
