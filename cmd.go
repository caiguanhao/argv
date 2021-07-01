package argv

import (
	"io"
	"os"
	"os/exec"
)

type Commands []*exec.Cmd

func NewCommands(args ...[]string) (cmds Commands) {
	for _, argv := range args {
		if len(argv) == 0 {
			continue
		}
		cmds = append(cmds, exec.Command(argv[0], argv[1:]...))
	}
	for i := range cmds {
		if i == 0 {
			continue
		}
		cmds[i].Stdin, _ = cmds[i-1].StdoutPipe()
	}
	return
}

func (cmds Commands) Std() Commands {
	return cmds.In(os.Stdin).Out(os.Stdout).Err(os.Stderr)
}

func (cmds Commands) In(stdin io.Reader) Commands {
	cmds[0].Stdin = stdin
	return cmds
}

func (cmds Commands) Out(stdout io.Writer) Commands {
	cmds[len(cmds)-1].Stdout = stdout
	return cmds
}

func (cmds Commands) Err(stderr io.Writer) Commands {
	for i := range cmds {
		cmds[i].Stderr = stderr
	}
	return cmds
}

func (cmds Commands) Run() error {
	for i := len(cmds) - 1; i > -1; i-- {
		if err := cmds[i].Start(); err != nil {
			return err
		}
	}
	for i := range cmds {
		if err := cmds[i].Wait(); err != nil {
			return err
		}
	}
	return nil
}
