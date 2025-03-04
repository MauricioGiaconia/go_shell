package ports

type CommandPort interface {
	Execute(params CommandParams) (*string, error)
}
