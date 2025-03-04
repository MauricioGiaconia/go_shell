package shell

import (
	"fmt"
	"strings"

	"go_shell/internal/adapters"
	"go_shell/internal/ports"
)

var CommandRegistry = map[string]ports.CommandPort{
	"echo":  adapters.EchoAdapter{},
	"ls":    adapters.LsAdapter{},
	"clear": adapters.ClearAdapter{},
	"goto":  adapters.GoToAdapter{},
}

func ExecuteCommand(input string, currentPath *string) {
	args := strings.Fields(input)

	if len(args) == 0 {
		return
	}

	command, exists := CommandRegistry[args[0]]
	if !exists {
		fmt.Println("Command not found:", args[0])
		return
	}

	params := ports.CommandParams{
		Args:        args[1:],
		CurrentPath: currentPath,
	}

	result, err := command.Execute(params)

	if err != nil {
		fmt.Println("Error to execute command:", err)
		return
	}

	if result != nil {
		fmt.Println(*result)
	}
}
