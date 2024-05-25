package internal

import (
	"flag"
	"fmt"

	"github.com/tuannguyenandpadcojp/duck-bot/internal"
)

var rootHelpFn = func(args ...any) func() {
	return func() {
		fmt.Printf(`Usage: %s <command> [arguments]

Commands:
collaborations [<arguments>]
  List collaborators of a repository
  Arguments:
    -g string  | Optional | Filter collaborators by name
`, args...)
	}
}

type RootCmd struct {
	*internal.BaseCmd
}

func NewRootCmd() *RootCmd {
	return &RootCmd{
		BaseCmd: &internal.BaseCmd{
			FlagSet:     flag.NewFlagSet("gh", flag.ExitOnError),
			Key:         "gh",
			Description: "Github command",
			Commands:    make(map[string]internal.Command),
			Help:        rootHelpFn("gh"),
		},
	}
}

func (c *RootCmd) RegisterCommand(cmd internal.Command) error {
	return c.BaseCmd.RegisterCommand(cmd)
}
