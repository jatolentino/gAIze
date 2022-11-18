@echo off
SETLOCAL
SET PATH=%PATH%;MAXINE-AR-SDK\samples\external\opencv\bin;NVIDIA_AR_SDK_0.8.2.0\bin;
SET NVAR_MODEL_DIR=NVIDIA_AR_SDK_0.8.2.0\bin\models
GazeRedirect.exe --offline_mode --split_screen_view=false  --in=%1