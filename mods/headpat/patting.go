package pat

import (
	"bufio"
	"encoding/json"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

  "github.com/bwmarrin/discordgo"
	log "github.com/Sirupsen/logrus"


  "github.com/ashfennix/goatbrotesquared/util/gvars"
)

func cmdPat(command []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	patNumber := ""
	if len(command) > 1 {
		patNumber = command[1]
	}
	pat, mattpat, patNum, maxPat, err := headPat(patNumber)
	if err != nil {
		//s.ChannelFileSendWithMessage(m.ChannelID, noPatMessage, "mattpat.png", mattpat)
		return
	}
	if mattpat != nil {
		ms := &discordgo.MessageSend{
			Embed: &discordgo.MessageEmbed{
				Image: &discordgo.MessageEmbedImage{
					URL: "attachment://" + pat,
				},
			},
			Files: []*discordgo.File{
				&discordgo.File{
					Name:   pat,
					Reader: mattpat,
				},
			},
		}

		s.ChannelMessageSendComplex(m.ChannelID, ms)
		return
	}
	patMessage := "Pat " + strconv.Itoa(patNum) + " of " + strconv.Itoa(maxPat)
	if maxPat == 0 && patNum == 0 {
		patMessage = noPatMessage
	}
	patBed := &discordgo.MessageEmbed{
		Author:      &discordgo.MessageEmbedAuthor{},
		Color:       0x00a0ff,
		//Description: patMessage,
		Image: &discordgo.MessageEmbedImage{
			URL: pat,
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: patMessage,
		},
		Title: "Headp.at",
		URL: "http://headp.at",
	}
	_, errbed := s.ChannelMessageSendEmbed(m.ChannelID, patBed)
	if gvars.DevMode && errbed != nil {
		s.ChannelMessageSend(m.ChannelID, "Embed had error: "+errbed.Error()+"\nHeadpat URL "+pat)
	}
}

func headPat(setPatNum string) (url string, file io.Reader, patNum int, maxPat int, err error) {
	//var pats headPats
	rand.Seed(time.Now().UnixNano())
	//patsJSONFile := "temp\\pats.json"
	if strings.ToLower(setPatNum) == "mattpat" {
		img, _ := patError()
		return "images/matpatt.png", bufio.NewReader(img), 0, 0, nil
	}
	patsJSONWeb, err := http.Get("https://headp.at/js/pats.json")
	if err != nil {
		img, _ := patError()
		log.Errorln("Failed to get file")
		return "", bufio.NewReader(img), 0, 0, err
	}
	defer patsJSONWeb.Body.Close()

	patsJSON, err := ioutil.ReadAll(patsJSONWeb.Body)
	if err != nil {
		img, _ := patError()
		log.Errorln("Failed to get file")
		return "", bufio.NewReader(img), 0, 0, err
	}
	var pats []string
	err = json.Unmarshal(patsJSON, &pats)
	if err != nil {
		img, _ := patError()
		log.Errorln("Failed to unmarshall file")
		return "", bufio.NewReader(img), 0, 0, err
	}
	maxPat = len(pats)
	if setPatNum != "" {
		patNum, err = strconv.Atoi(setPatNum)
		patNum = patNum - 1
		if err != nil {
			patNum = rand.Intn(maxPat - 1)
		}
		if patNum > maxPat-1 || patNum < 0 {
			patNum = rand.Intn(maxPat - 1)
		}
	} else {
		patNum = rand.Intn(maxPat - 1)
	}
	url = "https://headp.at/pats/" + pats[patNum]
	//fixes url by replaces spaces with URL code for spaces %20
	url = strings.Replace(url, " ", "%20", -1)
	return url, nil, patNum + 1, maxPat, err
}

func patError() (file io.Reader, errOut error) {
	noPat = gvars.CFG.Section("headpat").Key("noPat").String()
	noPatMessage = gvars.CFG.Section("headpat").Key("noPatMessage").String()
	if noPat == "" {
		noPat = "404headpatnotfoundsohereisamatpat.png"
		gvars.CFG.Section("headpat").Key("noPat").SetValue(noPat)
		gvars.CFG.SaveTo(gvars.ConfigFile)
		log.Errorln("noPat was not set, setting to default mattpat pat")
	}
	//log.Printf("HEADPAT IS BORKED RIP")
	img, err := os.Open("data/images/" + noPat)
	if err != nil {
		log.Errorln("AND MATTPAT IS MISSING :'(")
		return nil, err
	}
	return bufio.NewReader(img), err
}
