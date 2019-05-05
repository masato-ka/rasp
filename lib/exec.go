package lib

import (
	"io"
	"os"
	"os/exec"
)

func Execution(commandName string, args []string, dst string, stdout, stderr io.Writer) (*os.ProcessState, error) {

	cmd := exec.Command(commandName, args...)
	cmd.Dir = dst
	childStdout, _ := cmd.StdoutPipe()
	childStderr, _ := cmd.StderrPipe()

	go io.Copy(stdout, childStdout)
	go io.Copy(stderr, childStderr)

	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	return cmd.ProcessState, nil
}
