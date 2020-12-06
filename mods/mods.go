package mods

import (
  //Imported
   log "github.com/Sirupsen/logrus"

  //Project
  //"github.com/ashfennix/goatbrotesquared/mods/helloworld"
  "github.com/ashfennix/goatbrotesquared/mods/e621"
  "github.com/ashfennix/goatbrotesquared/mods/twitter"
  "github.com/ashfennix/goatbrotesquared/mods/meme-review"
  "github.com/ashfennix/goatbrotesquared/mods/headpat"
  "github.com/ashfennix/goatbrotesquared/mods/botmanagement"
  "github.com/ashfennix/goatbrotesquared/mods/roleman"
  "github.com/ashfennix/goatbrotesquared/mods/swebclone"
)

//Load = Loads mods listed
func Load(){
    log.Println("Loading mods package")
    //hello.Load()
    e621.Load()
    tweeter.Load()
    memereview.Load()
    pat.Load()
    manage.Load()
    roleman.Load()
    sweb.Load()
}
