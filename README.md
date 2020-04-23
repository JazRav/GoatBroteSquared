# GoatBrote²

A Discord bot that does things and is poorly coded, please don't use it if you want to stay sane.

# Ways to get this shit

## Download a binary

Go to releases up top and download the latest

## Normal source code way

If you want just the source code added to your **normal** Go

`go get github.com/DrBrobot/GoatBroteSquared/src/GoatBrote`

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

Check `build-dumb.bat` to see how I did it

## The dumb way™

If you want to work with the way I set it up *weirdly*

`go get github.com/DrBrobot/GoatBroteSquared/`

You can update your packages if you do it my way with the batch file

`Build-dumb.bat` will Build and Run the bot with the `bot_dev.ini`, which you can make your own by checking in `config` for the `example_bot.ini`, same with `config/twitter` and `example_twitter.ini`

# It Just Works

![alt text][ToddHoward]

[ToddHoward]: https://upload.wikimedia.org/wikipedia/commons/thumb/5/59/ToddHoward2010sm.jpg/220px-ToddHoward2010sm.jpg "Todd 'Godd' Howard"
