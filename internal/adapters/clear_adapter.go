package adapters

import "go_shell/internal/ports"

type ClearAdapter struct{}

func (ClearAdapter) Execute(params ports.CommandParams) (*string, error) {

	clearCommand := "\033[H\033[2J"

	return &clearCommand, nil
}
