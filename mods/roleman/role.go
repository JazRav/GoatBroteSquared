package roleman

import (
	log "github.com/Sirupsen/logrus"

  "github.com/ashfennix/goatbrotesquared/cmd"
  //"github.com/ashfennix/goatbrotesquared/util/gvars"
)

//Load - Roleman loader
func Load() {
  log.Println("Loading Role Mangament plugin")
  cmd.Make("ra","Role", cmdRoleAdd).HelpText("Adds Role to you or someone else").Add()

}
