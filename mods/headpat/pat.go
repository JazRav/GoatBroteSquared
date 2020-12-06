package pat

import (
	log "github.com/Sirupsen/logrus"

  "github.com/ashfennix/goatbrotesquared/cmd"
  "github.com/ashfennix/goatbrotesquared/util/gvars"
)

//Load - Headpat loader
func Load() {
  log.Println("Loading Headpat plugin")
	cmd.Make("headpat","Headpat", cmdPat).HelpText("gives random headpats\nadd a number at the end certian pat").Hidden().Add()
	cmd.Make("pat","Headpat", cmdPat).HelpText("gives random headpats\nadd a number at the end certian pat").Add()
	loadINI()
}

func loadINI(){
  noPat = gvars.CFG.Section("headpat").Key("noPat").String()
  noPatMessage = gvars.CFG.Section("headpat").Key("noPatMessage").String()
}
