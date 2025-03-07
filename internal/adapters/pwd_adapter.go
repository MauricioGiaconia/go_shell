package adapters

import "os"

type PwdAdapter struct{}

func (PwdAdapter) Execute() (*string, error) {

	currentPath, err := os.Getwd()

	if err != nil {
		return nil, err
	}

	return &currentPath, nil
}

func (PwdAdapter) GetDescription() string {
	instruction := "This command will print the current path"

	return instruction
}
