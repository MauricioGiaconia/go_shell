package adapters

type ClearAdapter struct{}

func (ClearAdapter) Execute(args []string) string {
	return "\033[H\033[2J"
}
