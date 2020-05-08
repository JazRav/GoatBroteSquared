package e621

import (
  //Built in
  "strconv"
  "strings"
  "math/rand"
  "time"
  //Imported
  log "github.com/Sirupsen/logrus"
  "github.com/bwmarrin/discordgo"

  //Project
  "github.com/dokvis/goatbrotesquared/cmd"
  "github.com/dokvis/goatbrotesquared/util/gvars"

  "github.com/dokvis/goatbrotesquared/mods/e621/handler"
)
//Load - Loads mod
func Load(){
  //e621 commands
  furCat := "Furry"
  log.Println("Loading e621 plugin")
  e621HelpMessage := "gives you a e621\\e926 image\ne621 in NSFW channels\ne926 in SFW channels\nput booru tags after command\nin DMs, add `NSFW` at end of tags for NSFW"
  cmd.Make("fur", furCat, cmdE621).HelpText(e621HelpMessage).Hidden().Add()
  cmd.Make("e621", furCat, cmdE621).HelpText(e621HelpMessage).Hidden().Add()
  cmd.Make("e926", furCat, cmdE621).HelpText(e621HelpMessage).Hidden().Add()
  cmd.Make("furid", furCat, cmdE621ID).HelpText("sends image with the ID provided\ne621 in NSFW channels, e926 in SFW channels").Add()

  //e621 subcommands
  cmd.Make("ralsei", furCat, cmdFurRalsei).HelpText("sends image of best goat\nadd booru tags at the end\nALWAYS SFW, you monster").Add()
  cmd.Make("treeboi", furCat, cmdFurRalsei).HelpText("sends image of best tree\nadd booru tags at the end\nALWAYS SFW, you monster").Hidden().Add()
  cmd.Make("katia", furCat, cmdFurKatia).HelpText("sends image of best cat\n" + e621HelpMessage).Add()
  cmd.Make("legoshi", furCat, cmdFurLegoshi).HelpText("sends image of best wolf\n"+e621HelpMessage).Add()
  cmd.Make("legosi", furCat, cmdFurLegoshi).HelpText("sends image of best wolf\n"+e621HelpMessage).Hidden().Add()
  cmd.Make("centi", furCat, cmdFurCenti).HelpText("sends image of centi\n"+e621HelpMessage).Add()
  cmd.Make("centipeetle", furCat, cmdFurCenti).HelpText("sends image of centi\n"+e621HelpMessage).Hidden().Add()
  cmd.Make("isabelle", furCat, cmdFurIsabelle).HelpText("sends image of Isabelle from Animal Crossing\n"+e621HelpMessage).Add()

  //Manage
  cmd.Make("e6sample", furCat, cmde621SampleToggle).Owner().Add()
	cmd.Make("e6filter", furCat, cmde621FilterToggle).Owner().Add()
	cmd.Make("e6filterscore", furCat, cmde621FilterScore).Owner().Add()

  loadINI()
}

func cmdE621(message []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	search := ""
	if len(message) > 1 {
		search = strings.TrimPrefix(m.Content, message[0]+" ")
	}
	e621EmbedMessage(search, false, "", false, "", "", s, m)
}
func cmdE621ID(message []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	if len(message) >= 2 {
		e621EmbedMessage(message[1], true, "", false, "", "", s, m)
	} else {
		s.ChannelMessageSend(m.ChannelID, "You need to put in an ID")
	}
}
func cmdFurRalsei(message []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	search := ""
	if len(message) > 1 {
		search = strings.TrimPrefix(m.Content, message[0]+" ")
	}
	whatBoi := "GOAT"

	prefix := gvars.Prefix

	if message[0] == prefix+"treeboi" {
		whatBoi = "TREEBOI"
	}


	e621EmbedMessage(search, false, "Ralsei", true, "NO LEWD " + whatBoi, ralseiAntiLewd(), s, m)
}

func cmdFurKatia(message []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	search := ""
	if len(message) > 1 {
		search = strings.TrimPrefix(m.Content, message[0]+" ")
	}
	e621EmbedMessage(search, false, "Katia_Managan", false, "", "", s, m)
}

func cmdFurCenti(message []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	search := ""
	if len(message) > 1 {
		search = strings.TrimPrefix(m.Content, message[0]+" ")
	}
	e621EmbedMessage(search, false, "Centipeetle", false, "", "", s, m)
}

func cmdFurLegoshi(message []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	search := ""
	if len(message) > 1 {
		search = strings.TrimPrefix(m.Content, message[0]+" ")
	}
	e621EmbedMessage(search, false, "Legoshi_(Beastars)", false, "", "", s, m)
}
func cmdFurIsabelle(message []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	search := ""
	if len(message) > 1 {
		search = strings.TrimPrefix(m.Content, message[0]+" ")
	}
	e621EmbedMessage(search, false, "isabelle_(animal_crossing)", false, "", "", s, m)
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
	eStuff, err := e6.E621Handler(search, forceID, forcesearch, nsfw, nolewd, "")
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
				eStuff, err = e6.E621Handler(search, forceID, forcesearch, nsfw, nolewd, "")
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
				eStuff, err = e6.E621Handler(search, forceID, forcesearch, nsfw, nolewd, "")
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
  bSource := false
	if eStuff.Source != "" {
		link = eStuff.Source
		title = "Source"
    bSource = true
	} else {
		link = eStuff.Page
		title = ""
	}
	var imageEmbed discordgo.MessageEmbedImage
	var videoEmbed discordgo.MessageEmbedVideo
	var clickMessage string
  switch eStuff.EXT {
    case "webm": {
		    videoEmbed = discordgo.MessageEmbedVideo{
			       URL: eStuff.URL,
        }
    }
    case "swf": {
        b := e6ORe9
        if bSource {
          b = title
        }
		    clickMessage = "\n\nFile is Flash, please click **" + b + "** to view"
    }
    default: {
      imageEmbed = discordgo.MessageEmbedImage{
      URL: eStuff.URL,
      }
    }
  }
	soundWarning := ""
	if eStuff.SoundWarning {
		soundWarning = "\n⚠WARNING SOUND MIGHT BE LOUD⚠"
	}
  rating := "unknown"
  switch eStuff.Rating {
    case "s": rating = "Safe"
    case "q": rating = "Questionable"
    case "e": rating = "Explicit"
  }
	if eStuff.Page != "" {
		e621embed := &discordgo.MessageEmbed{
			Color:       0x0055ff,
			Description: clickMessage+soundWarning,
			URL:         link,
			Author: &discordgo.MessageEmbedAuthor{
				URL:     eStuff.Page,
				Name:    e6ORe9,
				IconURL: "https://i.imgur.com/dbKpPIs.png",
			},
      Fields: []*discordgo.MessageEmbedField{
      &discordgo.MessageEmbedField{
          Name:    "Artist",
          Value:  eStuff.Artist,
          Inline: true,
      },
      &discordgo.MessageEmbedField{
          Name:   "Rating",
          Value:  rating,
          Inline: true,
      },
      &discordgo.MessageEmbedField{
          Name:   "Score",
          Value:  strconv.Itoa(eStuff.Score),
          Inline: true,
      },
    },
			Image: &imageEmbed,
			Video: &videoEmbed,
			Title:     title,
      Footer: &discordgo.MessageEmbedFooter{
        Text: "ID: " + strconv.Itoa(eStuff.ID),
      },
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

func cmde621SampleToggle(message []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	if e6.Sample == false {
		e6.Sample = true
		gvars.CFG.Section("e621").Key("sample").SetValue("true")
		gvars.CFG.SaveTo(gvars.ConfigFile)
		s.ChannelMessageSend(m.ChannelID, "e621/e926 SAMPLE ENABLED")
		log.Println("e621 SAMPLE ENABLED")
	} else {
		e6.Sample = false
		gvars.CFG.Section("e621").Key("sample").SetValue("false")
		gvars.CFG.SaveTo(gvars.ConfigFile)
		s.ChannelMessageSend(m.ChannelID, "e621/e926 SAMPLE DISABLED")
		log.Println("e621 SAMPLE DISABLED")
	}
}

func cmde621FilterToggle(message []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	if e6.Filter == false {
		e6.Filter = true
		gvars.CFG.Section("e621").Key("filter").SetValue("true")
		gvars.CFG.SaveTo(gvars.ConfigFile)
		s.ChannelMessageSend(m.ChannelID, "e621/e926 FILTER ENABLED")
		log.Println("e621 FILTER ENABLED")
	} else {
		e6.Filter = false
		gvars.CFG.Section("e621").Key("filter").SetValue("false")
		gvars.CFG.SaveTo(gvars.ConfigFile)
		s.ChannelMessageSend(m.ChannelID, "e621/e926 FILTER DISABLED")
		log.Println("e621 FILTER DISABLED")
	}
}

func cmde621FilterScore(message []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	score := strings.TrimPrefix(m.Content, message[0] + " ")
	if e6.Filter == false {
		s.ChannelMessageSend(m.ChannelID, "FILTER DISABLED, PLEASE ENABLE FILTER WITH ``" +gvars.Prefix + "e6filter`")
		log.Println("e621 FILTER ENABLED")
	} else {
		var extraMessage string
		extraMessage = ""
		if score == "<e6filterscore"	{
			score = "2"
			extraMessage = "NOTHING SET, SETTING "
		}
		e6.FilterScore = score
		gvars.CFG.Section("e621").Key("filterScore").SetValue(e6.FilterScore)
		gvars.CFG.SaveTo(gvars.ConfigFile)
		s.ChannelMessageSend(m.ChannelID, extraMessage +"FILTER SCORE TO " + e6.FilterScore)
		log.Println("e621 FILTER SCORE SET TO " + e6.FilterScore)
	}
}
func loadINI(){
  var err error
  e6.FilterScore = gvars.CFG.Section("e621").Key("filterScore").String()
  e6.Filter, err = gvars.CFG.Section("e621").Key("filter").Bool()
  if err != nil {
    e6.Filter = true
  }
  err = nil
  e6.Sample, err = gvars.CFG.Section("e621").Key("sample").Bool()
  if err != nil {
    e6.Sample = true
  }
}
//Protect the goat
func ralseiAntiLewd() string{
	rand.Seed(time.Now().UnixNano())
	ralseiNoLewd := []string{"1700281" , "2234990", "2031072", "2064695"}
	return ralseiNoLewd[rand.Intn(len(ralseiNoLewd))]
}
