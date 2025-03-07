package ports

type BaseCommandPort interface {
	GetDescription() string
}

type CommandPort interface {
	Execute(params CommandParams) (*string, error)
	BaseCommandPort
}

type CommandWithoutParams interface {
	Execute() (*string, error)
	BaseCommandPort
}
