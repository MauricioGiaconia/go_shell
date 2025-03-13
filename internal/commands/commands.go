package commands

import (
	"go_shell/internal/adapters"
	"go_shell/internal/ports"
)

func SetupCommands() map[string]ports.BaseCommandPort {
	commands := map[string]ports.BaseCommandPort{
		"go-echo":  adapters.EchoAdapter{},
		"go-ls":    adapters.LsAdapter{},
		"go-clear": adapters.ClearAdapter{},
		"go-to":    adapters.GoToAdapter{},
		"go-pwd":   adapters.PwdAdapter{},
		"go-rm":    adapters.RemoveAdapter{},
		"go-help":  adapters.HelpAdapter{},
	}
	return commands
}
