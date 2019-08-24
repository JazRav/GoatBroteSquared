@echo off
set GOPATH=%cd%
set GOBIN=%cd%\bin
set GOARCH=amd64
set GOOS=windows
set /p Version=<version.txt
FOR /F "tokens=* USEBACKQ" %%F IN (`git rev-parse HEAD`) DO ( SET GitHash=%%F )
echo killing goatbrote process
taskkill /F /IM goatbrote.exe
echo deleting old goatbrote
del goatbrote.exe
echo getting and installing goatbrote
@echo on
go get goatbrote
@echo off
echo %githash%
echo %TIME: =0%
set timetime=%TIME: =0%
go install -ldflags "-X main.Version=%version%-windows -X main.GitHash=%githash% -X main.BuildTime=%date:~0,2%-%date:~3,2%-%date:~6,2%T%timetime%" goatbrote
echo copying goatbrote
cmd /c echo F | xcopy bin\goatbrote.exe goatbrote.exe
goatbrote -c bot_dev.ini