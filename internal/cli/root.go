package cli

import (
	"fmt"
	"strings"
)

type Command interface {
	Name() string
	Init(args ...string) error
	Run() error
}

func Run(args []string) error {
	cmds := []Command{
		NewModelCommand(),
		NewServiceCommand(),
	}

	if len(args) < 1 {
		availableCmds := make([]string, len(cmds))
		for i, cmd := range cmds {
			availableCmds[i] = cmd.Name()
		}
		return fmt.Errorf("sub-command not specified (available: %v)", strings.Join(availableCmds, ", "))
	}

	subcommand := args[0]
	for _, cmd := range cmds {
		if strings.EqualFold(cmd.Name(), subcommand) {
			cmd.Init(args[1:]...)
			return cmd.Run()
		}
	}

	return fmt.Errorf("unknown subcommand: %s", subcommand)
}
