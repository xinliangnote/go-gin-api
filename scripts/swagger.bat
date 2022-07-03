@echo off
chcp 65001
echo.
echo Regenerating swagger doc
echo.
go install github.com/swaggo/swag/cmd/swag@v1.7.4
swag init
echo.
echo Done.