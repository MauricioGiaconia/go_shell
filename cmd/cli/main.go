package main

import (
	"bufio"
	"fmt"
	"go_shell/internal/shell"
	"os"
	"strings"
)

func main() {
	fmt.Println("Welcome to go shell! Please, write a command:")

	reader := bufio.NewReader(os.Stdin)

	currentPath, err := os.Getwd()

	if err != nil {
		currentPath = ""
	}

	for {
		fmt.Printf("\033[38;2;0;255;255m%s\033[0m Let's GO $ ", currentPath)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			fmt.Println("GO out...")
			break
		}

		shell.ExecuteCommand(input, &currentPath)
	}
}
