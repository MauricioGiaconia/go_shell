package ports

type CommandParams struct {
	Args        []string
	CurrentPath *string
	Commands    map[string]interface{}
}
