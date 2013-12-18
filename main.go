package main

import (
	"flag"
	"github.com/foomo/DaemonGarden/garden"
	"log"
	"os"
)

var address = flag.String("address", "127.0.0.1:8081", "address to bind host:port")
var logDir = flag.String("logDir", "/var/log/daemonGarden", "directory to put my daemon logs to")

func validateLogDir(logDir string) (logFile *os.File) {
	_, err := os.Stat(logDir)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("log directory does not exist")
		}
		if os.IsPermission(err) {
			log.Fatal("log directory is not accessible")
		}
		log.Fatal("there is sth wrong with my log dir")
	}
	logFile, err = os.OpenFile(logDir+"/daemonGarden.serverlog", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("could not open main log file " + err.Error())
	}
	return
}

func main() {
	flag.Parse()

	garden.LogDir = *logDir

	logFile := validateLogDir(garden.LogDir)
	logFile.Close()

	log.Println("starting server", *address)
	server := garden.NewServer()
	ok, err := server.Start(*address)
	if ok == false {
		log.Fatal(err)
	}
}
