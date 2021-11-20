@echo off
chcp 65001
echo.
echo Regenerating swagger doc
echo.
go install github.com/swaggo/swag/cmd/swag@latest
swag init
echo.
echo Done.