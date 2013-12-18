#!/bin/bash
rsync -avz ~/go foomo@test.foomo:~/ && ssh foomo@test.foomo 'date;export GOROOT=/usr/local/go;export GOPATH=/home/foomo/go;echo "GOPATH $GOPATH";/usr/local/go/bin/go build -o ~/go/src/github.com/foomo/DaemonGarden/daemonGarden-linux ~/go/src/github.com/foomo/DaemonGarden/main.go'

