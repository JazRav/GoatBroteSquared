package main

import (
	"github.com/go-ini/ini"
	"time"
)

var (
	//Version is whatever is in the version.txt
	Version   = "undefined"
	//BinaryOS is whatever the compiled ver is built for
	BinaryOS = "undefiend"
	//BinaryArch is whatever architecture its built for
	BinaryArch = "undefiend"
	//BuildTime is when it was built
	BuildTime = "undefined"
	//GitHash is whatever the git generates
	GitHash   = "undefined"
	//HostName of the machine is running on currently
	HostName = "None"

	botToken string
	devMode  bool
	logAll   bool

	statusMessage string
	statusType int // 1: Playing, 2: Listening, 3: Streaming
	statusURL string

	ownerID string

	cfg     *ini.File
	cfgFile string

	noPat        string
	noPatMessage string
	useGlobalPrefix bool
	globalPrefix string

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
