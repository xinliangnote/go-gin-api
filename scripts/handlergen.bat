@echo off
chcp 65001
echo.
echo Regenerating handler file
echo.
go run -v .\cmd\handlergen\main.go -handler %1
if %errorlevel% == 1 (
echo.
echo failed!!!
exit 1
)
echo.
echo Formatting code
echo.
go run -v .\cmd\mfmt\main.go
if %errorlevel% == 1 (
echo.
echo failed!!!
exit 1
)
echo.
echo Done.
