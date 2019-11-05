@echo off
echo Building GoatBroteSquared multiple platforms
set GOPATH=%cd%

echo --------------------------------------------

echo Linux Arm (Raspberry Pi)
set GOARCH=arm
set GOOS=linux
set GOARM=5
set /p Version=<version.txt
FOR /F "tokens=* USEBACKQ" %%F IN (`git rev-parse HEAD`) DO ( SET GitHash=%%F )
echo Deleting old RPi build
del /f /s /q "ship\Win32\"
echo Compiling %GOOS%_%GOARCH%
@echo on
go get goatbrote
set timetime=%TIME: =0%
go install -ldflags "-X main.Version=%version% -X main.BinaryOS=RPi-%GOOS% -X main.BinaryArch=%GOARCH% -X main.GitHash=%githash% -X main.BuildTime=%date:~0,2%-%date:~3,2%-%date:~6,2%T%timetime%" goatbrote
@echo off
cmd /c echo F | xcopy "bin\%GOOS%_%GOARCH%\goatbrote" "ship\%GOOS%-%GOARCH%\goatbrote" /Y
cmd /c echo F | xcopy "config\example_bot.ini" "ship\%GOOS%-%GOARCH%\config\example_bot.ini" /Y
cmd /c echo F | xcopy "README.md" "ship\%GOOS%-%GOARCH%\README.md" /Y
cmd /c echo F | xcopy "images" "ship\%GOOS%-%GOARCH%\images\" /Y
cd "ship\%GOOS%-%GOARCH%\"
"C:\Program Files\7-Zip\7z.exe" a "..\GoatBroteSquared-%version%-RPi-%GOOS%-%GOARCH%.zip" "*"
cd %GOPATH%