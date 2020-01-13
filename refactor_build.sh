#!/bin/bash

#mac 版的binary
go build -o build/refactor/mac/gdeyamlOperator cmd/refactor2/refactor_main.go

#Linux版的binary
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/refactor/linux/gdeyamlOperator cmd/refactor2/refactor_main.go
