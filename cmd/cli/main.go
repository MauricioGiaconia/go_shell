package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/MauricioGiaconia/go_shell/internal/shell"
)

func main() {
	fmt.Println("Welcome to go shell! Please, write a command:")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Let's GO $ ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			fmt.Println("GO out...")
			break
		}

		shell.ExecuteCommand(input)
	}
}
