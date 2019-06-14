#!/bin/bash

#mac 版的binary
go build  -o getImageLatestTag-mac sourcecode/*.go

#Linux版的binary
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o getImageLatestTag-linux sourcecode/*.go

