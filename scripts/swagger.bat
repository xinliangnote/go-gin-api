@echo off
chcp 65001
echo.
echo Regenerating swagger doc
echo.
go run -v github.com/swaggo/swag/cmd/swag init
echo.
echo Done.