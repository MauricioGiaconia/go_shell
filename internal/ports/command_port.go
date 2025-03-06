package ports

type CommandPort interface {
	Execute(params CommandParams) (*string, error)
}

type CommandWithoutParams interface {
	Execute() (*string, error)
}
