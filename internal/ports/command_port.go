package ports

type CommandPort interface {
	Execute(args []string) string
}
