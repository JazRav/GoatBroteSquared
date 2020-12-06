package memereview

import (
  log "github.com/Sirupsen/logrus"

  "github.com/ashfennix/goatbrotesquared/cmd"
)

//Load - Loads meme review
func Load() {
  log.Println("Loading Meme Review plugin")
	cmd.Make("meme","Fun", cmdMemeReview).HelpText("Reviews meme").Hidden().Add()
  cmd.Make("mr","Fun", cmdMemeReview).HelpText("Reviews meme").Hidden().Add()
  cmd.Make("memereview","Fun", cmdMemeReview).HelpText("Reviews meme").Add()
}
