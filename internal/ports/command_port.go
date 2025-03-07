package ports

type CommandPort interface {
	Execute(params CommandParams) (*string, error)
	GetDescription() string
}

type CommandWithoutParams interface {
	Execute() (*string, error)
	GetDescription() string
}
