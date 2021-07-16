@echo off
chcp 65001
echo.
echo Init db
echo.
go run -v .\cmd\init\db\main.go  -addr %1 -user %2 -pass %3 -name %4
if %errorlevel% == 1 (
echo.
echo failed!!!
exit 1
)
echo.
echo Done.