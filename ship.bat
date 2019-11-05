@echo off
echo Building GoatBroteSquared multiple platforms
set GOPATH=%cd%

REM Windows
echo --------------------------------------------

echo Windows x86-64
set GOARCH=amd64
set GOOS=windows
set /p Version=<version.txt
echo %Version%
FOR /F "tokens=* USEBACKQ" %%F IN (`git rev-parse HEAD`) DO ( SET GitHash=%%F )
echo Deleting old Windows x86-64 build
del /f /s /q "ship\windows-amd64\"
echo Compiling %GOOS%_%GOARCH%
@echo on
go get goatbrote
set timetime=%TIME: =0%
go install -ldflags "-X main.Version=%version% -X main.BinaryOS=%GOOS% -X main.BinaryArch=%GOARCH% -X main.GitHash=%githash% -X main.BuildTime=%date:~0,2%-%date:~3,2%-%date:~6,2%T%timetime%" goatbrote
@echo off
cmd /c echo F | xcopy "bin\goatbrote.exe" "ship\%GOOS%-%GOARCH%\goatbrote.exe"/Y
cmd /c echo F | xcopy "config\example_bot.ini" "ship\%GOOS%-%GOARCH%\config\example_bot.ini" /Y
cmd /c echo F | xcopy "README.md" "ship\%GOOS%-%GOARCH%\README.md" /Y
cmd /c echo F | xcopy "images" "ship\%GOOS%-%GOARCH%\images\" /Y
cd "ship\%GOOS%-%GOARCH%\"
"C:\Program Files\7-Zip\7z.exe" a "..\GoatBroteSquared-%version%-%GOOS%-%GOARCH%.zip" "*"
cd %GOPATH%


echo --------------------------------------------

echo Windows x86
set GOARCH=386
set GOOS=windows
set /p Version=<version.txt
FOR /F "tokens=* USEBACKQ" %%F IN (`git rev-parse HEAD`) DO ( SET GitHash=%%F )
echo Deleting old Windows x86 build
del /f /s /q "ship\windows-386\"
echo Compiling %GOOS%_%GOARCH%
@echo on
go get goatbrote
set timetime=%TIME: =0%
go install -ldflags "-X main.Version=%version% -X main.BinaryOS=%GOOS% -X main.BinaryArch=%GOARCH% -X main.GitHash=%githash% -X main.BuildTime=%date:~0,2%-%date:~3,2%-%date:~6,2%T%timetime%" goatbrote
@echo off
cmd /c echo F | xcopy "bin\%GOOS%_%GOARCH%\goatbrote.exe" "ship\%GOOS%-%GOARCH%\goatbrote.exe" /Y
cmd /c echo F | xcopy "config\example_bot.ini" "ship\%GOOS%-%GOARCH%\config\example_bot.ini" /Y
cmd /c echo F | xcopy "README.md" "ship\%GOOS%-%GOARCH%\README.md" /Y
cmd /c echo F | xcopy "images" "ship\%GOOS%-%GOARCH%\images\" /Y
cd "ship\%GOOS%-%GOARCH%\"
"C:\Program Files\7-Zip\7z.exe" a "..\GoatBroteSquared-%version%-%GOOS%-%GOARCH%.zip" "*"
cd %GOPATH%


REM Linux
echo --------------------------------------------

echo Linux x86-64
set GOARCH=amd64
set GOOS=linux
set /p Version=<version.txt
FOR /F "tokens=* USEBACKQ" %%F IN (`git rev-parse HEAD`) DO ( SET GitHash=%%F )
echo Deleting old Linux x86-64 build
del /f /s /q "ship\linux-amd64\"
echo Compiling %GOOS%_%GOARCH%
@echo on
go get goatbrote
set timetime=%TIME: =0%
go install -ldflags "-X main.Version=%version% -X main.BinaryOS=%GOOS% -X main.BinaryArch=%GOARCH% -X main.GitHash=%githash% -X main.BuildTime=%date:~0,2%-%date:~3,2%-%date:~6,2%T%timetime%" goatbrote
@echo off
cmd /c echo F | xcopy "bin\%GOOS%_%GOARCH%\goatbrote" "ship\%GOOS%-%GOARCH%\goatbrote" /Y
cmd /c echo F | xcopy "config\example_bot.ini" "ship\%GOOS%-%GOARCH%\config\example_bot.ini" /Y
cmd /c echo F | xcopy "README.md" "ship\%GOOS%-%GOARCH%\README.md" /Y
cmd /c echo F | xcopy "images" "ship\%GOOS%-%GOARCH%\images\" /Y
cd "ship\%GOOS%-%GOARCH%\"
"C:\Program Files\7-Zip\7z.exe" a "..\GoatBroteSquared-%version%-%GOOS%-%GOARCH%.zip" "*"
cd %GOPATH%

echo --------------------------------------------

echo Linux x86
set GOARCH=386
set GOOS=linux
set /p Version=<version.txt
FOR /F "tokens=* USEBACKQ" %%F IN (`git rev-parse HEAD`) DO ( SET GitHash=%%F )
echo Deleting old Linux x86 build
del /f /s /q "ship\linux-386\"
echo Compiling %GOOS%_%GOARCH%
@echo on
go get goatbrote
set timetime=%TIME: =0%
go install -ldflags "-X main.Version=%version% -X main.BinaryOS=%GOOS% -X main.BinaryArch=%GOARCH% -X main.GitHash=%githash% -X main.BuildTime=%date:~0,2%-%date:~3,2%-%date:~6,2%T%timetime%" goatbrote
@echo off
cmd /c echo F | xcopy "bin\%GOOS%_%GOARCH%\goatbrote" "ship\%GOOS%-%GOARCH%\goatbrote" /Y
cmd /c echo F | xcopy "config\example_bot.ini" "ship\%GOOS%-%GOARCH%\config\example_bot.ini" /Y
cmd /c echo F | xcopy "README.md" "ship\%GOOS%-%GOARCH%\README.md" /Y
cmd /c echo F | xcopy "images" "ship\%GOOS%-%GOARCH%\images\" /Y
cd "ship\%GOOS%-%GOARCH%\"
"C:\Program Files\7-Zip\7z.exe" a "..\GoatBroteSquared-%version%-%GOOS%-%GOARCH%.zip" "*"
cd %GOPATH%

echo --------------------------------------------

echo Linux Arm (Raspberry Pi)
set GOARCH=arm
set GOOS=linux
set GOARM=5
set /p Version=<version.txt
FOR /F "tokens=* USEBACKQ" %%F IN (`git rev-parse HEAD`) DO ( SET GitHash=%%F )
echo Deleting old RPi build
del /f /s /q "ship\linux-arm\"
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