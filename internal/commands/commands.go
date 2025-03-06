package commands

import "go_shell/internal/adapters"

func SetupCommands() map[string]interface{} {
	commands := map[string]interface{}{
		"echo":  adapters.EchoAdapter{},
		"ls":    adapters.LsAdapter{},
		"clear": adapters.ClearAdapter{},
		"goto":  adapters.GoToAdapter{},
		"pwd":   adapters.PwdAdapter{},
		"rm":    adapters.RemoveAdapter{},
	}
	return commands
}
