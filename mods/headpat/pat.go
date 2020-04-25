package pat

import (
	log "github.com/Sirupsen/logrus"

  "github.com/dokvis/goatbrotesquared/cmd"
)

//Load - Headpat loader
func Load() {
  log.Println("Head Pat Plugin loaded")
	cmd.Make("headpat","Headpat", cmdPat).HelpText("gives random headpats\nadd a number at the end certian pat").Add()
	cmd.Make("pat","Headpat", cmdPat).HelpText("gives random headpats\nadd a number at the end certian pat").Add()
}
