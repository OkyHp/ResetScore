@echo off
REM bld.bat - For Windows builds

REM Create the target directories
if not exist "%PREFIX%" mkdir "%PREFIX%"

REM Copy the DLL and plugin file
copy ResetScore.dll "%PREFIX%\" || exit 1
copy ResetScore.pplugin "%PREFIX%\" || exit 1

exit 0
