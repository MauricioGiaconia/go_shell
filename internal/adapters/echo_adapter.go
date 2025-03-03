package adapters

import (
	"strings"
)

type EchoAdapter struct{}

func (EchoAdapter) Execute(args []string) string {
	return strings.Join(args, " ")
}
