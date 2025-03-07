package adapters

type ClearAdapter struct{}

func (ClearAdapter) Execute() (*string, error) {

	clearCommand := "\033[H\033[2J"

	return &clearCommand, nil
}

func (ClearAdapter) GetDescription() string {
	instruction := "This command will delete all instructions written in the console"

	return instruction
}
