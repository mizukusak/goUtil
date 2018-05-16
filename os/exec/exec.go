package execUtil

import (
	"os/exec"
	"io/ioutil"
	"context"
)

func Execute(cmd *exec.Cmd, ctx context.Context) []byte {
	var out []byte
	stdout, _ := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	cmdDone := make(chan bool)
	cmdStart := make(chan bool)

	go func() {
		cmd.Start()
		cmdStart <- true
		out, _ = ioutil.ReadAll(stdout)
		cmdDone <- true
	}()

	<-cmdStart
	select {
	case <- cmdDone:
		go func() {<-ctx.Done()}()
	case <- ctx.Done():
		cmd.Process.Kill()
		<-cmdDone
	}
	return out
}

