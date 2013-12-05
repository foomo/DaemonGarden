#!/bin/bash
rsync -avz ~/go foomo@test.foomo:~/ && ssh foomo@test.foomo "export GOPATH=/home/foomo/go;/usr/local/go/bin/go build -o ~/go/src/github.com/foomo/DaemonGarden/daemonGarden-linux ~/go/src/github.com/foomo/DaemonGarden/main.go"

