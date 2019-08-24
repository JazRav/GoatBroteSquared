package main

import "github.com/go-ini/ini"

var (
	Version   = "undefined"
	BuildTime = "undefined"
	GitHash   = "undefined"

	botToken string
	devMode  bool
	logAll   bool

	ownerID string

	cfg     *ini.File
	cfgFile string

	noPat        string
	noPatMessage string
	prefix       string
	dankmemes    bool

	guildID []string

	commands = make(map[string]command)
)
