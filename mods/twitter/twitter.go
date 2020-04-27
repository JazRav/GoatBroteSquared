package tweeter

import (
  log "github.com/Sirupsen/logrus"
  "github.com/go-ini/ini"

  "github.com/dokvis/goatbrotesquared/cmd"
	"github.com/dokvis/goatbrotesquared/mods/twitter/handler"
  "github.com/dokvis/goatbrotesquared/util/gvars"
)

//Load - Twitter Plugin
func Load() {
    log.Println("Loading Twitter Plugin")
    cmd.Make("twitfollow","Twitter", cmdTwitFollow).HelpText("Follows account on global twitter").Owner().Add()
    cmd.Make("twitmassfollow","Twitter", cmdTwitMassFollow).HelpText("Follows accounts on global twitter via uploaded .txt file").Owner().Add()
    cmd.Make("tweet","Twitter", cmdTweet).HelpText("Tweets, can upload a single image, 5mb limit").Add()
    cmd.Make("twit","Twitter", cmdTwitSwitch).HelpText("Manages the twitter config file\n`SET` to set twitter user for global\n`LIST` to list twitter accounts").Owner().Add()
    cmd.Make("twitall","Twitter", cmdTwitForAll).HelpText("Toggles twitter for everyone (global and local)").Owner().Add()
//    cmd.Make("twitlock","Twitter", cmdTwitLock).HelpText("Locks global twitter to non-admins").Owner().Add()
    cmd.Make("chantwitlist","Twitter", cmdTwitListChans).HelpText("List whatever twitter account is tied channel").Owner().Add()
    cmd.Make("chantwitset","Twitter", cmdTwitListChans).HelpText("Set twitter config to this channel").Owner().DisableDM().Add()
    cmd.Make("chantwitremove","Twitter", cmdTwitRemoveChan).HelpText("Removes whatever twitter account is tied channel").Owner().DisableDM().Add()
    loadINI()
}

func loadINI(){
  var err error
  defaultAccount = gvars.CFG.Section("twitter").Key("account").String()
  forAll, err = gvars.CFG.Section("twitter").Key("forall").Bool()
  if err !=nil {
    forAll = false
  }
  twitCfg, twitErr := ini.Load("data/config/twitter/"+defaultAccount+".ini")
  if twitErr != nil {
    log.Errorln("Twitter config error: " +twitErr.Error())
  } else {
    twit.AccessToken = twitCfg.Section("").Key("token").String()
    twit.AccessTokenSecret = twitCfg.Section("").Key("tokenSecret").String()
    twit.ConsumerKey = twitCfg.Section("").Key("consumer").String()
    twit.ConsumerSecret = twitCfg.Section("").Key("consumerSecret").String()
  }
}
