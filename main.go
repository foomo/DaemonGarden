package main

import (
	"fmt"
	"github.com/foomo/DaemonGarden/garden"
	"log"
)

func main() {
	addr := "127.0.0.1:8080"
	server := garden.NewServer()
	fmt.Println("starting server", addr)
	ok, err := server.Start(addr)
	if ok == false {
		log.Fatal(err)
	}
}
