package adapters

import (
	"fmt"
	"os"
	"strings"
)

type LsAdapter struct{}

func (LsAdapter) Execute() (*string, error) {
	files, err := os.ReadDir(".")

	if err != nil {
		return nil, fmt.Errorf("error to list files")
	}

	var names []string
	var fileName string

	for _, file := range files {
		if file.IsDir() {
			fileName = fmt.Sprintf("\033[38;2;0;173;216m %s\033[0m", file.Name())
		} else {
			fileName = file.Name()
		}

		names = append(names, fileName)
	}

	filesToPrint := strings.Join(names, " - ")

	return &filesToPrint, nil
}

func (LsAdapter) GetDescription() string {
	instruction := "This command will print all existing files in the current path"

	return instruction
}
