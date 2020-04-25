package gvars

import (
  	"github.com/go-ini/ini"
)

var(
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
  //Prefix - The Global Prefix
  Prefix = ">"
  //BotToken - Discord bot token without "bot" at the begining
  BotToken =  "null"
  //Owner - Owners discord ID
  Owner = ""
  //DevMode - Enables dev mods
  DevMode = false
  //ConfigFile - The config file to load
  ConfigFile = "data/config/bot.ini"
  //ConfigFileName - The name of the ini without dir and .ini
  ConfigFileName = "bot"
  //CFG - The global config file itself
  CFG *ini.File
)
