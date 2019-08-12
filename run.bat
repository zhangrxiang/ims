@echo off

(uninstall.bat || go build && install.bat) && go build && install.bat