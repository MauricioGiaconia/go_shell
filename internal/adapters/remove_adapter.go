package adapters

import (
	"fmt"
	"go_shell/internal/ports"
	"os"
)

type RemoveAdapter struct{}

func (RemoveAdapter) Execute(params ports.CommandParams) (*string, error) {

	if len(params.Args) < 1 {
		return nil, fmt.Errorf("invalid number of arguments")
	}

	err := os.Remove(params.Args[0])

	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (RemoveAdapter) GetDescription() string {
	instruction := "This command will delete a specified file"

	return instruction
}
