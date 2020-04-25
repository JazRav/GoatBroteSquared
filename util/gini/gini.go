package gini

import (
  "github.com/go-ini/ini"
  log "github.com/Sirupsen/logrus"

  "github.com/dokvis/goatbrotesquared/util/gvars"
)
//Init - Starts global
func Init() (initErr error){
  var errINI error
  gvars.CFG, errINI = ini.Load(gvars.ConfigFile)
	if errINI != nil {
		log.Printf("FAILED TO READ BOT INI FILE: %v\n\n\n\nMAKE SURE ITS IN BOTDIR\\config\\bot.ini", errINI)
		return errINI
	}
  gvars.BotToken = gvars.CFG.Section("auth").Key("bot_token").String()
  gvars.Owner = gvars.CFG.Section("auth").Key("owner_id").String()
  return
}
