package executor

import (
	"fmt"
	"go_shell/internal/commands"
	"go_shell/internal/ports"
	"strings"
)

var CommandRegistry = commands.SetupCommands()

func ExecuteCommand(args []string, currentPath *string) (*string, error) {

	if len(args) == 0 {
		return nil, fmt.Errorf("invalid command")
	}

	command, exists := CommandRegistry[strings.ToLower(args[0])]

	if !exists {
		return nil, fmt.Errorf("command not found %s", args[0])
	}

	var result *string
	var err error

	switch cmd := command.(type) {

	case ports.CommandWithoutParams:
		result, err = cmd.Execute()

		if err != nil {
			return nil, fmt.Errorf("error to execute command %s", err.Error())
		}

	case ports.CommandPort:
		params := ports.CommandParams{
			Args:        args[1:],
			CurrentPath: currentPath,
		}

		result, err = cmd.Execute(params)

		if err != nil {
			return nil, fmt.Errorf("error to execute command %s", err.Error())
		}

	default:
		return nil, fmt.Errorf("command not found %s", args[0])
	}

	return result, nil
}
