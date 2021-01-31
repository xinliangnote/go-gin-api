#!/bin/bash
printf "\nRegenerating gqlgen files\n"
rm -f internal/graph/generated/generated.go \
    internal/graph/model/generated.go \
    internal/graph/resolvers/generated/generated.go
time go run -v github.com/99designs/gqlgen $1
printf "\nDone.\n\n"