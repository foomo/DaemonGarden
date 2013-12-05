package garden

import (
	"encoding/json"
	"errors"
	// "fmt"
)

var logDir string = "/var/log/daemonGarden"

type Gardener interface {
	Spawn(name string, cmdFile string, arguments []string) (reply string, err error)
	Kill(name string) (reply string, err error)
	Status() (reply string, err error)
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

func (garden *Garden) Spawn(name string, cmdFile string, args []string) (reply string, err error) {
	if runningDaemon, ok := garden.daemons[name]; ok {
		if runningDaemon.state == nil {
			reply = "process is already running"
			err = errors.New(reply)
			return
		} else {
			garden.removeDaemon(runningDaemon)
		}
	}
	daemon := NewDaemon(name, cmdFile, args)
	err = daemon.Spawn()
	if err == nil {
		garden.daemons[name] = daemon
		reply = "daemon spawned"
	} else {
		reply = "could not spawn deamon " + err.Error()
	}
	return
}

func (garden *Garden) Status() (reply string, err error) {
	statusMap := make(map[string]*DaemonStatus)
	for name, daemon := range garden.daemons {
		statusMap[name] = daemon.Status()
	}
	reply, err = toJSON(statusMap)
	return
}

func (garden *Garden) Kill(name string) (reply string, err error) {
	if daemon, ok := garden.daemons[name]; ok {
		garden.removeDaemon(daemon)
		reply = "killed " + name
	} else {
		reply = "daemon not found"
	}
	return
}

func NewGarden() *Garden {
	garden := new(Garden)
	garden.daemons = make(map[string]*Daemon)
	return garden
}
