package garden

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

type DaemonStatus struct {
	Name    string `json:"name"`
	Pid     int    `json:"pid"`
	Running bool   `json:"running"`
	state   *os.ProcessState
}

type Daemon struct {
	Name    string
	command *exec.Cmd
	cmdFile string
	args    []string
	logDir  string
	state   *os.ProcessState
	files   struct {
		err *os.File
		out *os.File
	}
}

func NewDaemon(name string, cmdFile string, args []string) *Daemon {
	daemon := new(Daemon)
	daemon.Name = name
	daemon.cmdFile = cmdFile
	daemon.args = args
	daemon.logDir = logDir
	return daemon
}

func getFileToAppend(filename string) (file *os.File, err error) {
	file, err = os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	return
}

func (daemon *Daemon) wireStreams() (err error) {
	stdout, err := daemon.command.StdoutPipe()
	if err != nil {
		return
	}
	stderr, err := daemon.command.StderrPipe()
	if err != nil {
		return
	}

	fileErr, err := getFileToAppend(logDir + "/" + daemon.Name + ".errors")
	if err != nil {
		return
	}

	fileOut, err := getFileToAppend(logDir + "/" + daemon.Name + ".log")
	if err != nil {
		fileErr.Close()
		return
	}

	daemon.files.err = fileErr
	daemon.files.out = fileOut

	go io.Copy(daemon.files.out, stdout)
	go io.Copy(daemon.files.err, stderr)

	return
}

func (daemon *Daemon) Spawn() (err error) {
	daemon.command = exec.Command(daemon.cmdFile, daemon.args...)
	err = daemon.wireStreams()
	if err == nil {
		err = daemon.command.Start()
		fmt.Println("process started")
		go func() {
			daemonState, waitErr := daemon.command.Process.Wait()
			if waitErr == nil {
				daemon.state = daemonState
			}
			daemon.files.err.Close()
			daemon.files.out.Close()
		}()
	}
	return
}

func (daemon *Daemon) Kill() (err error) {
	err = daemon.command.Process.Kill()
	return
}

func (daemon *Daemon) Status() (status *DaemonStatus) {
	status = new(DaemonStatus)
	status.Pid = daemon.command.Process.Pid
	status.Name = daemon.Name
	status.state = daemon.state
	if status.state != nil {
		status.Running = !status.state.Exited()
	} else {
		status.Running = true
	}
	return
}
