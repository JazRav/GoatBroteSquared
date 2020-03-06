package main

import (
	"github.com/go-ini/ini"
	"time"
)

var (
	Version   = "undefined"
	BinaryOS = "undefiend"
	BinaryArch = "undefiend"
	BuildTime = "undefined"
	GitHash   = "undefined"
	HostName = "None"

	botToken string
	devMode  bool
	logAll   bool

	ownerID string

	cfg     *ini.File
	cfgFile string

	noPat        string
	noPatMessage string
	useGlobalPrefix bool
	globalPrefix string
	dankmemes    bool

	guildID []string

	e6Sample bool

	e6Filter bool
	e6FilterScore string

	commands = make(map[string]command)

	selfUpdate bool

	twit twitAPIKeys

)

type twitAPIKeys struct {
	DefaultConfig string
	CurrentConfg string
	ConsumerKey string
	ConsumerSecret string
	AccessToken string
	AccessTokenSecret string
	Delay time.Duration
}
