@echo off
REM Launch Chrome with remote debugging enabled

set CHROME_PATH=%CHROME_PATH%
if "%CHROME_PATH%"=="" set CHROME_PATH="C:\Program Files\Google\Chrome\Application\chrome.exe"
set PORT=%DEVTOOLS_PORT%
if "%PORT%"=="" set PORT=9222

start "" %CHROME_PATH% --remote-debugging-port=%PORT% --user-data-dir="%USERPROFILE%\AppData\Local\Google\Chrome\User Data" --no-first-run --no-default-browser-check %*
