#!/bin/bash
printf "\nRegenerating swagger doc\n\n"
time go run -v github.com/swaggo/swag/cmd/swag init
printf "\nDone.\n\n"