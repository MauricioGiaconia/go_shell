package shell

import (
	"fmt"
	"strings"

	"github.com/MauricioGiaconia/go_shell/internal/adapters"
	"github.com/MauricioGiaconia/go_shell/internal/ports"
)

var CommandRegistry = map[string]ports.CommandPort{
	"echo": adapters.EchoAdapter{},
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
