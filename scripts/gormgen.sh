#!/bin/bash
printf "\nRegenerating mysql file\n\n"
time go run -v ./cmd/mysqlmd/main.go  -env fat

printf "\nRegenerating code\n\n"
time go build -o gormgen ./cmd/gormgen/main.go
mv gormgen $GOPATH/bin

go generate ./...

printf "\nFormatting code\n\n"
time go run -v github.com/koketama/mfmt

printf "\nDone.\n\n"
