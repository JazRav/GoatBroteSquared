package cmdHandle

import (
  //Built-in
	"strings"

  //Imported
	"github.com/bwmarrin/discordgo"
  log "github.com/Sirupsen/logrus"

  //Project
  "github.com/dokvis/goatbrotesquared/cmd"
  "github.com/dokvis/goatbrotesquared/mods"
  "github.com/dokvis/goatbrotesquared/data/gvars"
)
func Load(){
  log.Println("Loaded cmds")
  mods.Load()
}
func Handle(message []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	cmds := strings.TrimPrefix(message[0], gvars.Prefix)
	cmds = strings.ToLower(cmds)
	if command, ok := cmd.Commands[cmds]; ok && (cmds == strings.ToLower(command.Name)) {
      //log.Println(cmddo.Commands[cmd])
      //Does the command
      command.Exec(message, s, m)

		return
	}

}
