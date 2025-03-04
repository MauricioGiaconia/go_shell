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
