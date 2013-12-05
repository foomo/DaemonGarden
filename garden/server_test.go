package garden

import (
	//	"fmt"
	"strings"
	"testing"
)

type Kindergarden struct {
}

func NewKindergarden() *Kindergarden {
	return new(Kindergarden)
}

func (garden *Kindergarden) Spawn(name string, cmdFile string, arguments []string) (reply string, err error) {
	reply = "Spawn " + name + " " + cmdFile + " " + strings.Join(arguments, " ")
	return
}

func (garden *Kindergarden) Kill(name string) (reply string, err error) {
	reply = "Kill " + name
	return
}

func (garden *Kindergarden) Status() (reply string, err error) {
	reply = "Status"
	return
}

func testACommand(t *testing.T, rawCall []string, expectedReply string) {
	garden := NewKindergarden()
	reply, _ := runCommand(garden, rawCall)
	if reply != expectedReply {
		t.Fatal("expected", expectedReply, "got", reply)
	}

}

func TestSpawn(t *testing.T) {
	testACommand(t, []string{"cmd", "spawn", "hansi", "ls", "-la", "/tmp"}, "Spawn hansi ls -la /tmp")
}

func TestKill(t *testing.T) {
	testACommand(t, []string{"cmd", "kill", "hansi"}, "Kill hansi")
}

func TestStatus(t *testing.T) {
	testACommand(t, []string{"status"}, "Status")
}
