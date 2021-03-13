#!/bin/bash
printf "\nInit db\n\n"
time go run -v ./cmd/init/db/main.go  -addr $1 -user $2 -pass $3 -name $4
printf "\nDone.\n\n"
