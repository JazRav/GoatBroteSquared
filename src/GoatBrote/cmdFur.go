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
	e621HelpMessage := "gives you a e621\\e926 image\ne621 in NSFW channels\ne926 in SFW channels\nput booru tags after command\nin DMs, add `NSFW` at end of tags for NSFW"
	makeCmd("fur", cmdFurTrash).helpText(e621HelpMessage).add()
	makeCmd("e621", cmdFurTrash).helpText(e621HelpMessage).add()
	makeCmd("e926", cmdFurTrash).helpText(e621HelpMessage).add()
	makeCmd("furid", cmdFurIDLookup).helpText("sends image with the ID provided\ne621 in NSFW channels, e926 in SFW channels").add()

//Fur subcommands
	makeCmd("ralsei", cmdFurRalsei).helpText("sends image of best goat\nadd booru tags at the end\nALWAYS SFW, you monster").add()
	makeCmd("treeboi", cmdFurRalsei).helpText("sends image of best tree\nadd booru tags at the end\nALWAYS SFW, you monster").add()
	makeCmd("katia", cmdFurKatia).helpText("sends image of best cat\n" + e621HelpMessage).add()
	makeCmd("legoshi", cmdFurLegoshi).helpText("sends image of best wolf\n"+e621HelpMessage).add()
	makeCmd("legosi", cmdFurLegoshi).helpText("sends image of best wolf\n"+e621HelpMessage).add()
	makeCmd("centi", cmdFurCenti).helpText("sends image of centi\n"+e621HelpMessage).add()
	makeCmd("centipeetle", cmdFurCenti).helpText("sends image of centi\n"+e621HelpMessage).add()
	makeCmd("isabelle", cmdFurIsabelle).helpText("sends image of Isabelle from Animal Crossing\n"+e621HelpMessage).add()
//End of fur subcommands
}

func ralseiAntiLewd() string{
	rand.Seed(time.Now().UnixNano())
	ralseiNoLewd := []string{"1700281" , "1874162", "2031072", "2064695"}
	return ralseiNoLewd[rand.Intn(len(ralseiNoLewd))]
}

func cmdFurTrash(message []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	search := strings.TrimPrefix(m.Content, message[0]+" ")
	e621EmbedMessage(search, false, "", false, "", "", s, m)
}


func cmdFurIDLookup(message []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	if len(message) >= 2 {
		e621EmbedMessage(message[1], true, "", false, "", "", s, m)
	} else {
		s.ChannelMessageSend(m.ChannelID, "You need to put in an ID")
	}
}

//Fur Subcommands

func cmdFurRalsei(message []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	search := strings.TrimPrefix(m.Content, message[0]+" ")
	whatBoi := "GOAT"

	prefix := globalPrefix
	if !useGlobalPrefix {
		prefix = "" //either do this, or make this less shit when per-server-prefix is added
	}

	if message[0] == prefix+"treeboi" {
		whatBoi = "TREEBOI"
	}


	e621EmbedMessage(search, false, "Ralsei", true, "NO LEWD " + whatBoi, ralseiAntiLewd(), s, m)
}

func cmdFurKatia(message []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	search := strings.TrimPrefix(m.Content, message[0]+" ")
	e621EmbedMessage(search, false, "Katia_Managan", false, "", "", s, m)
}

func cmdFurCenti(message []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	search := strings.TrimPrefix(m.Content, message[0]+" ")
	e621EmbedMessage(search, false, "Centipeetle", false, "", "", s, m)
}

func cmdFurLegoshi(message []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	search := strings.TrimPrefix(m.Content, message[0]+" ")
	e621EmbedMessage(search, false, "Legoshi_(Beastars)", false, "", "", s, m)
}
func cmdFurIsabelle(message []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	search := strings.TrimPrefix(m.Content, message[0]+" ")
	e621EmbedMessage(search, false, "isabelle_(animal_crossing)", false, "", "", s, m)
}
//End of Fur Subcommands


type e621 struct {
	Posts []Post `json:"posts"`
}
//Post from e621
type Post struct {
	ID            int         `json:"id"`
	CreatedAt     string        `json:"created_at"`
	UpdatedAt     string        `json:"updated_at"`
	File          File          `json:"file"`
	Preview       Preview       `json:"preview"`
	Sample        Sample        `json:"sample"`
	Score         Score         `json:"score"`
	Tags          Tags          `json:"tags,omitempty"`
	LockedTags    []interface{} `json:"locked_tags,omitempty"`
	ChangeSeq     int         `json:"change_seq,omitempty"`
	Flags         Flags         `json:"flags,omitempty"`
	Rating        string        `json:"rating"`
	FavCount      int         `json:"fav_count"`
	Sources       []string      `json:"sources,omitempty"`
	Pools         []interface{} `json:"pools,omitempty"`
	Relationships Relationships `json:"relationships,omitempty"`
	ApproverID    int         `json:"approver_id"`
	UploaderID    int         `json:"uploader_id"`
	Description   string        `json:"description,omitempty"`
	CommentCount  int         `json:"comment_count"`
	IsFavorited   bool          `json:"is_favorited,omitempty"`
}
//File from e621.Post
type File struct {
	Width  int  `json:"width"`
	Height int  `json:"height"`
	EXT    string `json:"ext"`
	Size   int  `json:"size"`
	Md5    string `json:"md5"`
	URL    string `json:"url"`
}
//Flags from e621.Post
type Flags struct {
	Pending      bool `json:"pending,omitempty"`
	Flagged      bool `json:"flagged,omitempty"`
	NoteLocked   bool `json:"note_locked,omitempty"`
	StatusLocked bool `json:"status_locked,omitempty"`
	RatingLocked bool `json:"rating_locked,omitempty"`
	Deleted      bool `json:"deleted,omitempty"`
}
//Preview from e621.Post
type Preview struct {
	Width  int  `json:"width,omitempty"`
	Height int  `json:"height,omitempty"`
	URL    string `json:"url,omitempty"`
}
//Relationships from e621.Post
type Relationships struct {
	ParentID          interface{}   `json:"parent_id,omitempty"`
	HasChildren       bool          `json:"has_children,omitempty"`
	HasActiveChildren bool          `json:"has_active_children,omitempty"`
	Children          []interface{} `json:"children,omitempty"`
}
//Sample from e621.Post
type Sample struct {
	Has    bool   `json:"has,omitempty"`
	Height int  `json:"height,omitempty"`
	Width  int  `json:"width,omitempty"`
	URL    string `json:"url,omitempty"`
}
//Score from e621.Post
type Score struct {
	Up    int `json:"up"`
	Down  int `json:"down"`
	Total int `json:"total"`
}
//Tags from e621.Post
type Tags struct {
	General   []string      `json:"general,omitempty"`
	Species   []string      `json:"species,omitempty"`
	Character []string      `json:"character,omitempty"`
	Copyright []string      `json:"copyright,omitempty"`
	Artist    []string      `json:"artist,omitempty"`
	Invalid   []interface{} `json:"invalid,omitempty"`
	Lore      []interface{} `json:"lore,omitempty"`
	Meta      []interface{} `json:"meta,omitempty"`
}

type eImage struct {
	URL       string
	Page      string
	Artist    string
	Source    string
	Score     int
	Tags struct{
		General   []string
		Species   []string
		Character []string
	}
	Rating    string
	TimeStamp string
	ID 				int
	EXT				string
	SoundWarning bool
}

func e621Handler(search string, forceID bool, forcesearch string, nsfw bool, nolewd bool, blacklist string) (eStuff eImage, err error) {
	var e621s e621
	search = strings.Replace(search, " ", "+", -1)
	//cub begone!
	search = strings.Replace(search, ";", "", -1)
	filter := ""
	eLink := ""
	if !forceID {
		if e6Filter {
			filter = "score:>="+e6FilterScore
		}
		eLink = "https://e621.net/posts.json?tags=" + filter +"+"+ search + "+"+forcesearch+blacklist+ "+rating:s&limit=320&page="
		if nsfw {
			if !nolewd {
				 //log.Println("did a thing 1")
					eLink = "https://e621.net/posts.json?tags=" + filter +"+"+ search + "+"+forcesearch+blacklist+ "+-cub+-young+-rating:s&limit=320&page="
			} else {
				  //log.Println("did a thing 1")
				  eLink = "https://e621.net/posts.json?tags=" + filter +"+"+ search + "+"+forcesearch+blacklist+ "+rating:s&limit=320&page="
			}

		}
	} else {
			eLink = "https://e621.net/posts.json?tags="+search+forcesearch+"+rating:s&limit=320&page="
			if nsfw {
				if !nolewd {
					  //log.Println("did thing 1")
						eLink = "https://e621.net/posts.json?tags=" + filter +"+"+ search + "+"+forcesearch+blacklist+ "+-cub+-young+-rating:s&limit=320&page="
				} else {
					 	//log.Println("did thing 2")
					  eLink = "https://e621.net/posts.json?tags=" + filter +"+"+ search + "+"+forcesearch+blacklist+ "+rating:s&limit=320&page="
				}
			}
	}
	//log.Println("Json: "+eLink)
	rand.Seed(time.Now().UnixNano())
	//fileGetter(eLink, "temp/e621.json")
	client := &http.Client{}
	req, err := http.NewRequest("GET", eLink, nil)
	if err != nil {
		log.Println("Borked: Request Error: " + err.Error())
		return eStuff, err
	}
	req.Header.Set("User-Agent", "GoatBroteSquared_DiscordGo_Bot/"+Version)
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Borked: Resp Error: " + err.Error())
		return eStuff, err
	}
	//log.Println(resp )
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Borked: Read Error:" + err.Error())
		return eStuff, err
	}
	if string(body) == "" {
		log.Println("Borked: body is blank")
		return eStuff, err
	}
	fuckMeErr := json.Unmarshal(body, &e621s)
	if fuckMeErr != nil {
			log.Println("Borked: unmarshaling failed" + fuckMeErr.Error())
			return eStuff, err
	}
	var maxE621 int
	maxE621 = len(e621s.Posts)
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

	eStuff.Rating = e621s.Posts[numE621].Rating
	//Having some issues with loading of images, using sample instead
	if !e6Sample {
		eStuff.URL = e621s.Posts[numE621].File.URL
	} else if e621s.Posts[numE621].File.EXT == "webm" {
		eStuff.URL = e621s.Posts[numE621].File.URL
	} else {
		eStuff.URL = e621s.Posts[numE621].Sample.URL
	}
	eStuff.URL = strings.Replace(eStuff.URL, " ", "%20", -1)
	eStuff.Tags.Character = e621s.Posts[numE621].Tags.Character
	eStuff.Tags.General = e621s.Posts[numE621].Tags.General
	eStuff.Tags.Species = e621s.Posts[numE621].Tags.Species
	eStuff.Score = e621s.Posts[numE621].Score.Total
	if nsfw {
		eStuff.Page = "https://e621.net/post/show/" + strconv.Itoa(e621s.Posts[numE621].ID)
	} else {
		eStuff.Page = "https://e926.net/post/show/" + strconv.Itoa(e621s.Posts[numE621].ID)
	}
	eStuff.SoundWarning = false
	for a := 0; a < len(e621s.Posts[numE621].Tags.Artist); a++ {
		if e621s.Posts[numE621].Tags.Artist[a] == "sound_warning" {
			eStuff.SoundWarning = true
		}
	}
	if len(e621s.Posts[numE621].Tags.Artist) == 1 {
		eStuff.Artist = e621s.Posts[numE621].Tags.Artist[0]
	} else if len(e621s.Posts[numE621].Tags.Artist) == 2 {
		if e621s.Posts[numE621].Tags.Artist[1] == "sound_warning" {
			eStuff.Artist = e621s.Posts[numE621].Tags.Artist[0]
		} else {
			eStuff.Artist = e621s.Posts[numE621].Tags.Artist[0] + " & " + e621s.Posts[numE621].Tags.Artist[1]
		}
	} else if len(e621s.Posts[numE621].Tags.Artist) > 2 {
		eStuff.Artist = e621s.Posts[numE621].Tags.Artist[0] + " & "+strconv.Itoa(len(e621s.Posts[numE621].Tags.Artist)-1) + " more"
	} else {
		eStuff.Artist = "unknown artist"
	}
	if len(e621s.Posts[numE621].Sources) > 0 {
			eStuff.Source = e621s.Posts[numE621].Sources[0]
	}
	eStuff.TimeStamp = e621s.Posts[numE621].CreatedAt

	eStuff.EXT = e621s.Posts[numE621].File.EXT
	eStuff.ID = e621s.Posts[numE621].ID

	return eStuff, err
}

func e621EmbedMessage(search string, idlookup bool, forcesearch string, nolewd bool, nolewdmessage string, nolewdid string, s *discordgo.Session, m *discordgo.MessageCreate){
	chanInfo, _ := s.Channel(m.ChannelID)
	if nolewd && chanInfo.NSFW {
		search = "id:"+nolewdid
		forcesearch = ""
	} else if (strings.Contains(search, " ralsei") && (chanInfo.NSFW || (chanInfo.GuildID == "" && strings.HasSuffix(search, " NSFW") ) ) ) {
		search = "id:"+ralseiAntiLewd()
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
	if eStuff.Rating == "e" {
		for a := 0; a < len(eStuff.Tags.Character); a++ {
			if eStuff.Tags.Character[a] == "ralsei" {
				search = "id:"+ralseiAntiLewd()
				forcesearch = ""
				nolewdmessage = "LEWD WITH GOAT DETECTED"
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
		//s.ChannelMessageSend(m.ChannelID, "DEBUG: ID:" + strconv.Itoa(eStuff.ID)+" URL: " + eStuff.URL)
	} else {
		s.ChannelMessageSend(m.ChannelID, "We found nothing for `"+search+"`\nMake sure names with spaces, like Katia Managan is spelt like `Katia_Managan`")
		//s.ChannelMessageSend(m.ChannelID, "DEBUG: ID:" + strconv.Itoa(eStuff.ID)+" URL: " + eStuff.URL)
	}

}
