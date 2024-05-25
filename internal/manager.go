package internal

import (
	"sync"

	"github.com/tuannguyenandpadcojp/duck-bot/internal/errors"
)

type CmdManager struct {
	mux      *sync.Mutex
	commands map[string]Command
}

func NewCommandManager() *CmdManager {
	return &CmdManager{
		mux:      &sync.Mutex{},
		commands: make(map[string]Command),
	}
}

func (m *CmdManager) RegisterCommand(cmd Command) error {
	m.mux.Lock()
	defer m.mux.Unlock()

	if _, ok := m.commands[cmd.Cmd()]; ok {
		return errors.ErrCommandAlreadyExists
	}

	m.commands[cmd.Cmd()] = cmd
	return nil
}

func (m *CmdManager) Run(task string, args ...string) error {
	cmd, ok := m.commands[task]
	if !ok {
		return errors.ErrCommandNotFound
	}
	return cmd.Execute(args...)
}
