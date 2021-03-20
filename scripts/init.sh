#!/bin/bash

shellExit()
{
if [ $1 -eq 1 ]; then
    printf "\nfailed!!!\n\n"
    exit 1
fi
}

printf "\nInit db\n\n"

time go run -v ./cmd/init/db/main.go  -addr $1 -user $2 -pass $3 -name $4
shellExit $?

printf "\nDone.\n\n"
