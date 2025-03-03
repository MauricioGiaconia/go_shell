package adapters

import (
	"fmt"
	"os"
	"strings"
)

type LsAdapter struct{}

func (LsAdapter) Execute(args []string) string {
	files, err := os.ReadDir(".")

	if err != nil {
		return "Error to list files"
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

	return strings.Join(names, " - ")
}
