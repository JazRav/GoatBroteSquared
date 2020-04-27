package cmd

import (
  "strings"
  "github.com/bwmarrin/discordgo"
  log "github.com/Sirupsen/logrus"
)

var (
  //Commands - The commands
  Commands = make(map[string]command)
)

type command struct {
	Name  string
	Help  string
	IsOwner bool
	IsDMDisabled bool
	Exec  func([]string, *discordgo.Session, *discordgo.MessageCreate)
  Category string
  IsHidden bool
}



func (cmd command) Add() command {
	Commands[strings.ToLower(cmd.Name)] = cmd
	return cmd
}

//Make - Makes new command
func Make(name string,cat string, fun func([]string, *discordgo.Session, *discordgo.MessageCreate)) command {
  log.Println("Loaded Command '" + name + "' added in the '" + cat + "' category")
  return command{
		Name: name,
		Exec: fun,
    Category: cat,
	}
}

func (cmd command) Owner() command {
	cmd.IsOwner = true
	return cmd
}

func (cmd command) Hidden() command {
	cmd.IsHidden = true
	return cmd
}
func (cmd command) DisableDM() command {
	cmd.IsDMDisabled = true
	return cmd
}

func (cmd command) HelpText(helpText string) command {
	cmd.Help = helpText
	return cmd
}
