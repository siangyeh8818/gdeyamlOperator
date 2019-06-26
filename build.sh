#!/bin/bash

#mac 版的binary
go build  -o gdeyamlOperator-mac sourcecode/*.go

#Linux版的binary
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o gdeyamlOperator-linux sourcecode/*.go

