/*
Using INI for now
*/

package main

import (
  "os"
  "github.com/go-ini/ini"
  "github.com/bwmarrin/discordgo"
)

var (
  userDataDir = "config/data/users"
)

func init(){
  //Make userDataDir
  if _, userDirErr := os.Stat(userDataDir); os.IsNotExist(userDirErr) {
    os.MkdirAll(userDataDir, os.ModePerm)
  }
  //commands
  makeCmd("setdata", cmdUserDataSet).helpText("Set user data\nSet using one of these\n`Twitter`, `Youtube`, `Steam`, `All`").add()
  makeCmd("getdata", cmdUserDataGet).helpText("Get user data").add()
}

func cmdUserDataGet(message []string, s *discordgo.Session, m *discordgo.MessageCreate) {
  loadUserDataINI(s, m)
}

func cmdUserDataSet(message []string, s *discordgo.Session, m *discordgo.MessageCreate) {
  //Check for which data is saved, if all, have it detect it by URL?
}

type User struct{
  Data struct{
    UserID string
    Username string //Name as of last save
    Discrim string //Discrim as of last save
    Twitter string
    Youtube string
    Steam string
  }
  File struct{
    Location string
  }
}

func loadUserDataINI(s *discordgo.Session, m *discordgo.MessageCreate) (data *userData, err error) {
  err = nil
  userDataFile
  iniErr := ini.Load()
  if iniErr != nil {
    mkErr := os.
    if mkErr != nil {
      break
    }
  }
  return data, err
}

func saveUserDataINI(data User.Data, s *discordgo.Session, m *discordgo.MessageCreate) (err error) {
  err = nil

  return err
}

func embedUserData(data User.Data, s *discordgo.Session, m *discordgo.MessageCreate) (err error) {
  err = nil

  return err
}
