package main

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/bwmarrin/discordgo"
)

func init() {
	makeCmd("ralsei", cmdFurRalsei).helpText("sends image of best goat\nadd booru tags at the end\nALWAYS SFW, you monster").add()
	makeCmd("treeboi", cmdFurRalsei).helpText("sends image of best tree\nadd booru tags at the end\nALWAYS SFW, you monster").add()
	makeCmd("fur", cmdFurTrash).helpText("gives you a e621\\e926 image\ne621 in NSFW channels\ne926 in SFW channels\nput booru tags after command").add()
	makeCmd("e621", cmdFurTrash).helpText("gives you a e621\\e926 image\ne621 in NSFW channels\ne926 in SFW channels\nput booru tags after command").add()
	makeCmd("e926", cmdFurTrash).helpText("gives you a e621\\e926 image\ne621 in NSFW channels\ne926 in SFW channels\nput booru tags after command").add()
	makeCmd("katia", cmdFurKatia).helpText("sends image of best cat\nadd booru tags at the end\ne621 in NSFW channels, e926 in SFW channels").add()
	makeCmd("furid", cmdFurIDLookup).helpText("sends image with the ID provided\ne621 in NSFW channels, e926 in SFW channels").add()
}

func cmdFurTrash(message []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	search := strings.TrimPrefix(m.Content, message[0])
	e621EmbedMessage(search, false, "", false, "", "", s, m)
}

func cmdFurRalsei(message []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	search := strings.TrimPrefix(m.Content, message[0])
	whatBoi := "GOAT"
	if message[0] == prefix+"treeboi" {
		whatBoi = "TREEBOI"
	}
	e621EmbedMessage(search, false, "Ralsei", true, "NO LEWD " + whatBoi, "1700281", s, m)
}

func cmdFurKatia(message []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	search := strings.TrimPrefix(m.Content, message[0])
	e621EmbedMessage(search, false, "Katia_Managan", false, "", "", s, m)
}

func cmdFurIDLookup(message []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	if len(message) >= 2 {
		e621EmbedMessage(message[1], true, "", false, "", "", s, m)
	} else {
		s.ChannelMessageSend(m.ChannelID, "You need to put in an ID")
	}
}

type e621 struct {
	Rating    string   `json:"rating"`
	URL       string   `json:"file_url"`
	SampleURL string   `json:"sample_url"`
	Tags      string   `json:"tags"`
	Score     int      `json:"score"`
	ID        int      `json:"id"`
	Artists   []string `json:"artist,omitempty"`
	Timestamp string   `json:"created_at"`
	Source    string   `json:"source"`
	PreviewURL string   `json:"preview_url"`
}

//Stolen from arch, use later
type eImage struct {
	URL       string
	Page      string
	Artist    string
	Source    string
	Score     int
	Tags      string
	Rating    string
	TimeStamp string
	ID 				int
}

func e621Handler(search string, forceID bool, forcesearch string, nsfw bool, nolewd bool, blacklist string) (eStuff eImage, err error) {
	var e621s []e621
	search = strings.Replace(search, " ", ",", -1)
	rand.Seed(time.Now().UnixNano())
	filter := ""
	eLink := ""
	if !forceID {
		if e6Filter {
			filter = "score:>="+e6FilterScore
		}
		eLink = "https://e926.net/post/index.json?tags=" + filter +","+ search + ","+forcesearch+blacklist+ "&limit=320&page="
		if nsfw && !nolewd {
			eLink = "https://e621.net/post/index.json?tags=" + filter +","+ search + ","+forcesearch+blacklist+ ",-cub,-young,-rating:s&limit=320&page="
		}
	} else {
			eLink = "https://e926.net/post/index.json?tags="+search+forcesearch+"&limit=320&page="
			if nsfw {
				eLink = "https://e621.net/post/index.json?tags="+search+forcesearch+",-cub,-young&limit=320&page="
			}
	}
	log.Println("Json: "+eLink)
	rand.Seed(time.Now().UnixNano())
	//fileGetter(eLink, "temp/e621.json")
	client := &http.Client{}
	req, err := http.NewRequest("GET", eLink, nil)
	if err != nil {
		return eStuff, err
	}
	req.Header.Set("User-Agent", "GoatBroteSquared_DiscordGo_Bot/0.1")
	resp, err := client.Do(req)
	if err != nil {
		return eStuff, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return eStuff, err
	}
	if string(body) == "" {
		return
	}
	json.Unmarshal(body, &e621s)
	maxE621 := len(e621s)
	if maxE621 == 0 {
		return eStuff, err
	}
	/*
		Fixes shit since rand can't be 0 and getting a result of 1, like with sans,
		getting rid of `-1` for rand still crashes, ITS FUCKING MAGIC I TELL YOU
	*/
	if maxE621 == 1 {
		maxE621 = 2
	}
	numE621 := rand.Intn(maxE621 - 1)
	rE621 := e621s[numE621]
	eStuff.Rating = rE621.Rating
	//Having some issues with loading of images, using sample instead
	if !e6Sample {
		eStuff.URL = rE621.URL
	} else {
		eStuff.URL = rE621.SampleURL
	}
	eStuff.URL = strings.Replace(eStuff.URL, " ", "%20", -1)
	eStuff.Tags = rE621.Tags
	eStuff.Score = rE621.Score
	if nsfw {
		eStuff.Page = "https://e621.net/post/show/" + strconv.Itoa(rE621.ID)
	} else {
		eStuff.Page = "https://e926.net/post/show/" + strconv.Itoa(rE621.ID)
	}
	if len(rE621.Artists) > 0 {
		eStuff.Artist = rE621.Artists[0]
	} else {
		eStuff.Artist = "unknown artist"
	}
	eStuff.Source = rE621.Source
	eStuff.TimeStamp = rE621.Timestamp
	eStuff.ID = rE621.ID
	return eStuff, err
}

func e621EmbedMessage(search string, idlookup bool, forcesearch string, nolewd bool, nolewdmessage string, nolewdid string, s *discordgo.Session, m *discordgo.MessageCreate){
	chanInfo, _ := s.Channel(m.ChannelID)
	if nolewd && chanInfo.NSFW {
		search = "id:"+nolewdid
		forcesearch = ""
	}
	forceID := false
	if idlookup {
		search = "id:" +search
		forceID = true
	}
	eStuff, err := e621Handler(search, forceID, forcesearch, chanInfo.NSFW, nolewd, "")
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
	if eStuff.Page != "" {
		e621embed := &discordgo.MessageEmbed{
			Color:       0x0055ff,
			Description: "Artist: " + eStuff.Artist + "\nRating: " + eStuff.Rating + " Score: " + strconv.Itoa(eStuff.Score) + "\nID: " + strconv.Itoa(eStuff.ID),
			URL:         link,
			Author: &discordgo.MessageEmbedAuthor{
				URL:     eStuff.Page,
				Name:    e6ORe9,
				IconURL: "https://i.imgur.com/dbKpPIs.png",
			},
			Image: &discordgo.MessageEmbedImage{
				URL: eStuff.URL,
			},
			Title:     title,
			Timestamp: eStuff.TimeStamp,
		}
		s.ChannelMessageSendEmbed(m.ChannelID, e621embed)
		if devMode {
				//s.ChannelMessageSend(m.ChannelID, "URL of Image:" + eStuff.URL)
		}
		return
	}
	if idlookup {
		nsfwMessage := ""
		if !chanInfo.NSFW {
			nsfwMessage = " or if its NSFW, make sure you are in a NSFW channel"
		}
		s.ChannelMessageSend(m.ChannelID, "We found nothing for `" + search + "`\nMake sure its an actual ID, and its not blacklisted" + nsfwMessage)
	} else {
		s.ChannelMessageSend(m.ChannelID, "We found nothing for `"+search+"`\nMake sure names with spaces, like Katia Managan is spelt like `Katia_Managan`")
	}

}
