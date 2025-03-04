package adapters

import (
	"fmt"
	"go_shell/internal/ports"
	"os"
	"strings"
)

type LsAdapter struct{}

func (LsAdapter) Execute(params ports.CommandParams) (*string, error) {
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
