package garden

import (
	"encoding/json"
	"errors"
	"log"
)

var LogDir string = ""

type Gardener interface {
	Spawn(name string, cmdFile string, arguments []string) (reply *ServerReply, err error)
	Kill(name string) (reply *ServerReply, err error)
	Status() (reply *ServerReply, err error)
}

type Garden struct {
	daemons map[string]*Daemon
}

func toJSON(data interface{}) (encodedJSON string, err error) {
	jsonBytes, err := json.MarshalIndent(data, "", "\t")
	if err == nil {
		encodedJSON = string(jsonBytes)
	}
	return
}

func (garden *Garden) removeDaemon(daemon *Daemon) {
	daemon.Kill()
	delete(garden.daemons, daemon.Name)
}

func (garden *Garden) Spawn(name string, cmdFile string, args []string) (reply *ServerReply, err error) {
	reply = NewServerReply()
	log.Println("garden.Spawn", name, cmdFile, args)
	if runningDaemon, ok := garden.daemons[name]; ok {
		log.Println("  oh, there already is a daemon with that name")
		if runningDaemon.state == nil {
			reply.body = "process is already running"
			log.Println("    denying Spawn there already is a running daemon")
			err = errors.New(reply.body)
			return
		} else {
			log.Println("    looks like that is a dead one - removing it")
			garden.removeDaemon(runningDaemon)
		}
	}
	daemon := NewDaemon(name, cmdFile, args)
	err = daemon.Spawn()
	if err == nil {
		garden.daemons[name] = daemon
		reply.body = "daemon spawned"
	} else {
		reply.body = "could not spawn deamon " + err.Error()
	}
	log.Println("  " + reply.body)
	return
}

func (garden *Garden) Status() (reply *ServerReply, err error) {
	reply = NewServerReply()
	statusMap := make(map[string]*DaemonStatus)
	for name, daemon := range garden.daemons {
		statusMap[name] = daemon.Status()
	}
	reply.contentType = "application/json"
	reply.body, err = toJSON(statusMap)
	log.Println("garden.Status")
	log.Print(reply.body)
	log.Println("  end of garden.Status")
	return
}

func (garden *Garden) Kill(name string) (reply *ServerReply, err error) {
	reply = NewServerReply()
	log.Println("garden.Kill", name)
	if daemon, ok := garden.daemons[name]; ok {
		garden.removeDaemon(daemon)
		reply.body = "killed " + name
	} else {
		reply.body = "daemon not found"
	}
	log.Println("  " + reply.body)
	return
}

func NewGarden() *Garden {
	garden := new(Garden)
	garden.daemons = make(map[string]*Daemon)
	return garden
}
