package shell

import (
	"fmt"
	"strings"

	"go_shell/internal/adapters"
	"go_shell/internal/ports"
)

var CommandRegistry = map[string]ports.CommandPort{
	"echo": adapters.EchoAdapter{},
	"ls":   adapters.LsAdapter{},
}

func ExecuteCommand(input string) {
	args := strings.Fields(input)

	if len(args) == 0 {
		return
	}

	command, exists := CommandRegistry[args[0]]
	if !exists {
		fmt.Println("Command not found:", args[0])
		return
	}

	fmt.Println(command.Execute(args[1:]))
}
