package mods

import (
  //Imported
   log "github.com/Sirupsen/logrus"

  //Project
  "github.com/dokvis/goatbrotesquared/mods/helloworld"
  "github.com/dokvis/goatbrotesquared/mods/e621"

)

func Load(){
    log.Println("Loading mods package")
    hello.Load()
    e621.Load()
}
