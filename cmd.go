package simpleexec

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"

	"github.com/google/shlex"
)

type Cmd struct {
	Parent *Cmd
	*exec.Cmd
}

func ParseCmd(command string) *Cmd {
	cmdArray, err := shlex.Split(command)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return nil
	}

	eCmd := exec.Command(cmdArray[0], cmdArray[1:]...)
	eCmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

	cmd := &Cmd{}
	cmd.Cmd = eCmd
	return cmd
}

func (cmd *Cmd) Pipe(command string) *Cmd {
	pCmd := ParseCmd(command)

	// Although cmd.StdoutPipe() returns errors, the documentation does
	// not say _when_ said errors occur
	cOut, _ := cmd.StdoutPipe()

	pCmd.Stdin = cOut
	pCmd.Parent = cmd
	return pCmd
}

func (cmd *Cmd) Start() (err error) {
	err = cmd.Cmd.Start()
	if err != nil {
		return
	}
	if cmd.Parent != nil {
		err = cmd.Parent.Start()
	}
	return
}
