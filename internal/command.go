package internal

import (
	"flag"
	"strings"

	"github.com/tuannguyenandpadcojp/duck-bot/internal/errors"
)

type Command interface {
	Cmd() string
	Execute(args ...string) error
	Command(cmd string) Command
	FullCommand() string
	SetParent(parent Command)
}

type BaseCmd struct {
	*flag.FlagSet
	Key         string
	Description string
	Parent      Command
	Commands    map[string]Command
	Help        func()
}

func (c *BaseCmd) Cmd() string {
	return c.Key
}

func (c *BaseCmd) Execute(args ...string) error {
	if len(args) == 0 {
		c.Help()
		return nil
	}
	subKey := args[0]
	cmd := c.Command(subKey)
	if cmd == nil {
		c.Help()
		return nil
	}
	return cmd.Execute(args[1:]...)
}

func (c *BaseCmd) RegisterCommand(cmd Command) error {
	if _, ok := c.Commands[cmd.Cmd()]; ok {
		return errors.ErrCommandAlreadyExists
	}
	cmd.SetParent(c)
	c.Commands[cmd.Cmd()] = cmd
	return nil
}

func (c *BaseCmd) Command(cmd string) Command {
	if cmd, ok := c.Commands[cmd]; ok {
		return cmd
	}
	return nil
}

func (c *BaseCmd) FullCommand() string {
	cmpPath := []string{}
	if c.Parent != nil {
		cmpPath = append(cmpPath, c.Parent.FullCommand())
	}
	cmpPath = append(cmpPath, c.Key)
	if c.FlagSet != nil && len(c.Commands) == 0 {
		c.FlagSet.Visit(func(f *flag.Flag) {
			cmpPath = append(cmpPath, "-"+f.Name, f.Value.String())
		})
	}
	return strings.Join(cmpPath, " ")
}

func (c *BaseCmd) SetParent(parent Command) {
	c.Parent = parent
}
