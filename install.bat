@echo off
rem run this script as admin

if not exist simple-ims.exe (
    echo simple-ims.exe not found
    goto :exit
)

sc create simple-ims binpath= "%CD%\simple-ims.exe" start= auto DisplayName= "simple-ims"
sc description simple-ims "simple-ims.exe"
sc start simple-ims
sc query simple-ims

:exit