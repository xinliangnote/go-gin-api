#!/bin/bash
printf "\nRegenerating gqlgen files\n"
rm -f internal/graph/generated/generated.go \
    internal/graph/model/generated.go \
    internal/graph/resolvers/generated/generated.go
go get github.com/99designs/gqlgen
time gqlgen
printf "\nDone.\n\n"