package guildINI

import (
  "github.com/dokvis/goatbrotesquared/util/tools"
  "github.com/dokvis/goatbrotesquared/util/gvars"
)

//MakeGuildIni - Makes the ini for the guild
func MakeGuildIni(gid string){
  gIniExist, err := tools.DirExists("data/config/"+gvars.ConfigFileName+"/"+gid)
  if err != nil {
    return
  }
  if gIniExist {

  } else {

  }
}
