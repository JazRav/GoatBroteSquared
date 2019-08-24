package main

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

	log "github.com/Sirupsen/logrus"
	"github.com/bwmarrin/discordgo"
)

func init() {
	makeCmd("headpat", cmdPat).helpText("gives random headpats\nadd a number at the end certian pat").add()
}

func cmdPat(command []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	patNumber := ""
	if len(command) > 1 {
		patNumber = command[1]
	}
	pat, mattpat, err, patNum, maxPat := headPat(patNumber)
	if err != nil {
		s.ChannelFileSendWithMessage(m.ChannelID, noPatMessage, "mattpat.png", mattpat)
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
		Description: patMessage,
		Image: &discordgo.MessageEmbedImage{
			URL: pat,
		},
		Title: "From Headp.at",
	}
	_, errbed := s.ChannelMessageSendEmbed(m.ChannelID, patBed)
	if devMode && errbed != nil {
		s.ChannelMessageSend(m.ChannelID, "Embed had error: "+errbed.Error()+"\nHeadpat URL "+pat)
	}
}

func headPat(setPatNum string) (url string, file io.Reader, err error, patNum int, maxPat int) {
	//var pats headPats
	rand.Seed(time.Now().UnixNano())
	//patsJsonFile := "temp\\pats.json"
	if strings.ToLower(setPatNum) == "mattpat" {
		img, _ := patError()
		return "images/matpatt.png", bufio.NewReader(img), nil, 0, 0
	}
	patsJsonWeb, err := http.Get("https://headp.at/js/pats.json")
	if err != nil {
		img, _ := patError()
		log.Printf("Failed to get file")
		return "", bufio.NewReader(img), err, 0, 0
	}
	defer patsJsonWeb.Body.Close()

	patsJson, err := ioutil.ReadAll(patsJsonWeb.Body)
	if err != nil {
		img, _ := patError()
		log.Printf("Failed to get file")
		return "", bufio.NewReader(img), err, 0, 0
	}
	var pats []string
	err = json.Unmarshal(patsJson, &pats)
	if err != nil {
		img, _ := patError()
		log.Printf("Failed to unmarshall file")
		return "", bufio.NewReader(img), err, 0, 0
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
	return url, nil, err, patNum + 1, maxPat
}

func patError() (file io.Reader, errOut error) {
	noPat = cfg.Section("bot").Key("noPat").String()
	noPatMessage = cfg.Section("bot").Key("noPatMessage").String()
	if noPat == "" {
		noPat = "404headpatnotfoundsohereisamatpat.png"
		cfg.Section("bot").Key("noPat").SetValue(noPat)
		cfg.SaveTo(cfgFile)
		log.Printf("noPat was not set, setting to default mattpat pat")
	}
	//log.Printf("HEADPAT IS BORKED RIP")
	img, err := os.Open("images/" + noPat)
	if err != nil {
		log.Printf("AND MATTPAT IS MISSING :'(")
		return nil, err
	}
	return bufio.NewReader(img), err
}
