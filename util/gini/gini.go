package gini

import (
  "github.com/go-ini/ini"
  log "github.com/Sirupsen/logrus"

  "github.com/ashfennix/goatbrotesquared/util/gvars"
)
//Init - Starts global
func Init() (initErr error){
  var errINI error
  gvars.CFG, errINI = ini.Load(gvars.ConfigFile)
	if errINI != nil {
		log.Printf("FAILED TO READ BOT INI FILE: %v\n\n\n\nMAKE SURE ITS IN BOTDIR\\config\\bot.ini", errINI)
		return errINI
	}
  var err error
  gvars.BotToken = gvars.CFG.Section("auth").Key("bot_token").String()
  gvars.Owner = gvars.CFG.Section("auth").Key("owner_id").String()
  gvars.LogAll, err = gvars.CFG.Section("bot").Key("logall").Bool()
  if err != nil {
    log.Errorln("logall is not a boll, setting to false")
    gvars.LogAll = false
  }
  err = nil
  gvars.Prefix = gvars.CFG.Section("bot").Key("globalPrefix").String()
  gvars.StatusMessage = gvars.CFG.Section("bot").Key("statusMessage").String()
  gvars.StatusType, err = gvars.CFG.Section("bot").Key("statusType").Int()
  if err != nil {
    log.Errorln("statusType is not an int, setting to 1")
  }
  gvars.StatusURL = gvars.CFG.Section("bot").Key("statusURL").String()
  return
}
