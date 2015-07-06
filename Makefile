SHELL := /bin/bash

options:
	echo "you can clean | test | build | run | package"
clean:
	rm -f bin/daemon-garden
build: clean
	go build -o bin/daemon-garden
package: build
	pkg/build.sh
test:
	go test -v github.com/foomo/DaemonGarden/server/repo
