package commands

import (
	"go_shell/internal/adapters"
	"go_shell/internal/ports"
)

func SetupCommands() map[string]ports.BaseCommandPort {
	commands := map[string]ports.BaseCommandPort{
		"echo":  adapters.EchoAdapter{},
		"ls":    adapters.LsAdapter{},
		"clear": adapters.ClearAdapter{},
		"goto":  adapters.GoToAdapter{},
		"pwd":   adapters.PwdAdapter{},
		"rm":    adapters.RemoveAdapter{},
		"help":  adapters.HelpAdapter{},
	}
	return commands
}
