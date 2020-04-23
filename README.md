# GoatBrote²

A Discord bot that does things and is poorly coded, please don't use it if you want to stay sane.

# Ways to get this shit

## Download a binary

Go to releases up top and download the latest

## Normal source code way

If you want just the source code added to your **normal** Go

`go get github.com/DrBrobot/GoatBroteSquared/`

Check the `config` folder in this git for the `example_bot.ini` file, rename it to `bot.ini` and change for your info

Same with `config/twitter` with the file `example_twitter.ini` Rename it anything and make to set it in your `bot.ini` file

You need everything in the `images` folder for some commands

These libraries

###### Discord Go
`go get github.com/bwmarrin/discordgo`

###### Logrus
`go get github.com/Sirupsen/logrus`

###### Go Ini
`go get github.com/go-ini/ini`

###### Anaconda (Twitter API)
`go get github.com/ChimeraCoder/anaconda`

**You need set these ldflags**

`main.Version`

`main.BinaryOS`

`main.BinaryArch`

`main.GitHash`

`main.BuildTime`

I think this works this way

```batch
set GOPATH=%cd%
set GOBIN=%cd%\bin
set GOARCH=amd64
set GOOS=windows
go install -ldflags "-X main.Version=%version% -X main.BinaryOS=%GOOS% -X main.BinaryArch=%GOARCH% -X main.GitHash=%githash% -X main.BuildTime=%date:~0,2%-%date:~3,2%-%date:~6,2%T%timetime%" github.com/dokvis/goatbrotesquared
```
# It Just Works

![alt text][ToddHoward]

[ToddHoward]: https://upload.wikimedia.org/wikipedia/commons/thumb/5/59/ToddHoward2010sm.jpg/220px-ToddHoward2010sm.jpg "Todd 'Godd' Howard"
