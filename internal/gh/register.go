package gh

import (
	"github.com/tuannguyenandpadcojp/duck-bot/internal"
	"github.com/tuannguyenandpadcojp/duck-bot/internal/clients"
	ghi "github.com/tuannguyenandpadcojp/duck-bot/internal/gh/internal"
)

var (
	_ internal.Command = (*ghi.RootCmd)(nil)
	_ internal.Command = (*ghi.CollaboratorCmd)(nil)
	_ ghi.GithubClient = (*clients.Github)(nil)
)

type Config struct {
	GithubClient ghi.GithubClient
}

func Commands(config *Config) internal.Command {
	rootCommand := ghi.NewRootCmd()
	rootCommand.RegisterCommand(
		ghi.NewCollaboratorCmd(config.GithubClient),
	)

	return rootCommand
}
