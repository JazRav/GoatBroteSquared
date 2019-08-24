@echo off
echo Setting Go Paths to current folder temporary
set GOPATH=%cd%
set GOBIN=%cd%\bin
echo Updating Discord Go
go get github.com/bwmarrin/discordgo
echo Updating Logrus
go get github.com/Sirupsen/logrus
echo Updating Go-ini
go get github.com/go-ini/ini
echo Done, you can now close this