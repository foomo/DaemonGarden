package garden

import (
	"net/http"
)

type CommandDescription struct {
	Name string   `json:"name"`
	Args []string `json:"args"`
}

type Server struct {
	numRequests int64
	garden      *Garden
}

func NewServer() *Server {
	server := new(Server)
	server.garden = NewGarden()
	server.numRequests = 0
	return server
}

func (server *Server) Start(addr string) (ok bool, err error) {
	handler := NewHandler(server)
	err = http.ListenAndServe("127.0.0.1:8080", handler)
	if err != nil {
		ok = false
	} else {
		ok = true
	}
	return
}

func runCommand(gardener Gardener, rawCall []string) (reply string, err error) {
	if rawCall[0] == "cmd" {
		if rawCall[1] == "spawn" {
			if len(rawCall) > 3 && rawCall[3] != "" {
				reply, err = gardener.Spawn(rawCall[2], rawCall[3], rawCall[4:len(rawCall)])
			} else {
				reply = "can not spawn ..."
			}
		} else if rawCall[1] == "kill" {
			reply, err = gardener.Kill(rawCall[2])
		} else if rawCall[1] == "restart" {
			reply = "restart"
		} else {
			reply = "wtf is that command"
		}
	} else if rawCall[0] == "status" {
		reply, err = gardener.Status()
	} else {
		reply = "wtf"
	}
	return
}

func (server *Server) Execute(rawCall []string) (reply string, statusCode int) {
	reply = ""
	statusCode = http.StatusOK
	reply, err := runCommand(server.garden, rawCall)
	if err != nil {
		statusCode = http.StatusInternalServerError
	}
	return
}
