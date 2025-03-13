package adapters

import (
	"fmt"
	"go_shell/internal/ports"
	"sync"
)

type HelpAdapter struct{}

func (HelpAdapter) Execute(params ports.CommandParams) (*string, error) {

	//Obtengo el comando con maxima longitud para formatear las instrucciones a imprimir
	maxKeyLength := 0
	for key := range params.Commands {
		if len(key) > maxKeyLength {
			maxKeyLength = len(key)
		}
	}

	wg := sync.WaitGroup{}
	wg.Add(len(params.Commands))

	for key, value := range params.Commands {
		go func() {
			format := fmt.Sprintf("\033[38;2;0;173;216m %%-%ds\033[0m - %%s\n", maxKeyLength)
			fmt.Printf(format, key, value.GetDescription())
			wg.Done()
		}()
	}

	wg.Wait()

	return nil, nil
}

func (HelpAdapter) GetDescription() string {
	description := "This command print all available commands that can be executed"

	return description
}
