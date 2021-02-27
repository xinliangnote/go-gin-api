#!/bin/bash
printf "\nRegenerating handler file\n\n"
time go run -v ./cmd/handlergen/main.go  -handler $1

printf "\nFormatting code\n\n"
time go run -v github.com/koketama/mfmt

printf "\nDone.\n\n"
