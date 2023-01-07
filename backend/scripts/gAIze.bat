@echo off
SETLOCAL
SET PATH=%PATH%;..\..\..\..\samples\external\opencv\bin;..\..\..\..\MAXINE\bin;
SET NVAR_MODEL_DIR=..\..\..\..\MAXINE\bin\models
GazeRedirect.exe --offline_mode --split_screen_view=false  --in=%1