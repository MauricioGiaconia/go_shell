package shell

import (
	"fmt"
	"strings"

	"go_shell/internal/commands"
	"go_shell/internal/ports"
)

var CommandRegistry = commands.SetupCommands()

func ExecuteCommand(input string, currentPath *string) {
	var result *string
	var err error
	args := strings.Fields(input)

	if len(args) == 0 {
		return
	}

	command, exists := CommandRegistry[strings.ToLower(args[0])]

	if !exists {
		fmt.Println("Command not found:", args[0])
		return
	}

	switch cmd := command.(type) {

	case ports.CommandWithoutParams:
		result, err = cmd.Execute()

		if err != nil {
			fmt.Println("Error to execute command:", err)
			return
		}

	case ports.CommandPort:
		params := ports.CommandParams{
			Args:        args[1:],
			CurrentPath: currentPath,
		}

		result, err = cmd.Execute(params)

		if err != nil {
			fmt.Println("Error to execute command:", err)
			return
		}

	default:
		fmt.Println("Command not found:", args[0])
		return
	}

	if result != nil {
		fmt.Println(*result)
	}
}
