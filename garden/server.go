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

type ServerReply struct {
	httpCode int
	body string
	contentType string
}

func NewServerReply() *ServerReply {
	reply := new(ServerReply)
	reply.httpCode = http.StatusOK
	reply.contentType = "text/plain; charset=utf-8;"
	reply.body = ""
	return reply
}

func (server *Server) Start(addr string) (ok bool, err error) {
	handler := NewHandler(server)
	err = http.ListenAndServe(addr, handler)
	if err == nil {
		ok = true
	} else {
		ok = false
	}
	return
}

func runCommand(gardener Gardener, rawCall []string) (reply *ServerReply, err error) {
	reply = NewServerReply();
	reply.httpCode = 404;
	reply.body = "supported commands: \n  /status\n  /cmd/spawn/name/cmdFile/arg/arg/...\n  /cmd/kill/name\n"
	if len(rawCall) > 0 {
		if rawCall[0] == "cmd" {
			if rawCall[1] == "spawn" {
				if len(rawCall) > 3 && rawCall[3] != "" {
					reply, err = gardener.Spawn(rawCall[2], rawCall[3], rawCall[4:len(rawCall)])
				} else {
					reply.body = "can not spawn ..."
				}
			} else if rawCall[1] == "kill" {
				reply, err = gardener.Kill(rawCall[2])
			} else if rawCall[1] == "restart" {
				reply.body = "restart"
			} else {
				reply.body = "wtf is that command"
			}
		} else if rawCall[0] == "status" {
			reply, err = gardener.Status()
		}
	}
	return
}

func (server *Server) Execute(rawCall []string) (reply *ServerReply) {
	reply, err := runCommand(server.garden, rawCall)
	if err != nil {
		reply.httpCode = http.StatusInternalServerError
		reply.body = "oh body that did not work"
	}
	return
}
