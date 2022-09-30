#!/bin/sh
go mod download
go get golang.org/x/tools/cmd/present@latest
go run golang.org/x/tools/cmd/present@latest
