package shell

import (
	"fmt"
	"strings"

	"go_shell/internal/commands"
	"go_shell/internal/executor"
)

var CommandRegistry = commands.SetupCommands()

func ExecuteCommand(input string, currentPath *string) {

	if input == "" {
		return
	}

	args := strings.Fields(input)

	result, err := executor.ExecuteCommand(args, currentPath)

	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	if result != nil {
		fmt.Println(*result)
	}

}
