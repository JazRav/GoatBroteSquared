@echo off
echo Setting Go Paths to current folder temporary
set GOPATH=%cd%
set GOBIN=%cd%\bin
echo Getting Discord Go
go get -u github.com/bwmarrin/discordgo
echo Getting Logrus
go get -u github.com/Sirupsen/logrus
echo Getting Go-ini
go get -u github.com/go-ini/ini
echo Getting some twitter shit
go get -u github.com/ChimeraCoder/anaconda
echo Done, you can now close this