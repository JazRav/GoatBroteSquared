package mods

import (
  //Imported
   log "github.com/Sirupsen/logrus"

  //Project
  "github.com/dokvis/goatbrotesquared/mods/helloworld"
  "github.com/dokvis/goatbrotesquared/mods/e621"
  "github.com/dokvis/goatbrotesquared/mods/twitter"
  "github.com/dokvis/goatbrotesquared/mods/meme-review"
  "github.com/dokvis/goatbrotesquared/mods/headpat"
)

//Load = Loads mods listed
func Load(){
    log.Println("Loading mods package")
    hello.Load()
    e621.Load()
    tweeter.Load()
    memereview.Load()
    pat.Load()
}
