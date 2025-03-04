package adapters

import (
	"fmt"
	"go_shell/internal/ports"
	"os"
)

type GoToAdapter struct{}

func (GoToAdapter) Execute(params ports.CommandParams) (*string, error) {
	err := os.Chdir(params.Args[0])

	if err != nil {
		return nil, fmt.Errorf("error to change directory: %s", err.Error())
	}

	*params.CurrentPath, err = os.Getwd()

	if err != nil {
		*params.CurrentPath = ""
	}

	return nil, nil
}
