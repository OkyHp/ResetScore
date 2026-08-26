@echo off
REM bld.bat - For Windows builds

REM Create the target directories
if not exist "%PREFIX%" mkdir "%PREFIX%"

REM Copy the DLL, plugin and translation files
copy resetscore.dll "%PREFIX%\" || exit 1
copy resetscore.pplugin "%PREFIX%\" || exit 1
copy resetscore.yml "%PREFIX%\" || exit 1

exit 0
