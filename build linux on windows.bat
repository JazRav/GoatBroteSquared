@echo off
set GOPATH=%cd%
set GOARCH=amd64
set GOOS=linux
set /p Version=<version.txt
FOR /F "tokens=* USEBACKQ" %%F IN (`git rev-parse HEAD`) DO ( SET GitHash=%%F )
echo deleting old goatbrote
del goatbrote
echo getting and installing goatbrote
@echo on
go get goatbrote
set timetime=%TIME: =0%
go install -ldflags "-X main.Version=%version%-linux -X main.GitHash=%githash% -X main.BuildTime=%date:~0,2%-%date:~3,2%-%date:~6,2%T%timetime%" goatbrote
@echo off
echo copying goatbrote
cmd /c echo F | xcopy bin\linux_amd64\goatbrote goatbrote