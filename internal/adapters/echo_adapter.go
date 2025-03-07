package adapters

import (
	"go_shell/internal/ports"
	"strings"
)

type EchoAdapter struct{}

func (EchoAdapter) Execute(params ports.CommandParams) (*string, error) {
	textToPrint := strings.Join(params.Args, " ")
	return &textToPrint, nil
}

func (EchoAdapter) GetDescription() string {
	instruction := "This command will print the words typed after the instruction"

	return instruction
}
