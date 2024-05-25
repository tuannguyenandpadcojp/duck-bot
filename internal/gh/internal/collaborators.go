package internal

import (
	"context"
	"flag"
	"fmt"
	"slices"
	"strings"

	dbi "github.com/tuannguyenandpadcojp/duck-bot/internal"
	"github.com/tuannguyenandpadcojp/duck-bot/internal/clients"
	"github.com/tuannguyenandpadcojp/duck-bot/pkg/logging"
)

var collaboratorHelpFn = func(c dbi.Command) func() {
	return func() {
		fmt.Printf(`Usage: %s [arguments] owner repository
Example: %s tuannguyenandpadcojp duck-bot

Arguments:
 -g string  | Optional | Filter collaborators by name
`, c.FullCommand(), c.FullCommand())
	}
}

type CollaboratorCmd struct {
	dbi.BaseCmd
	c GithubClient
}

func NewCollaboratorCmd(c GithubClient) *CollaboratorCmd {
	memberFlg := flag.NewFlagSet("collaborators", flag.ExitOnError)
	memberFlg.String("repo", "", "Repository name")
	memberFlg.String("owner", "", "Repository owner")
	memberFlg.String("g", "", "Filter collaborators by name")

	cmd := &CollaboratorCmd{
		BaseCmd: dbi.BaseCmd{
			FlagSet:     memberFlg,
			Key:         "collaborators",
			Description: "List collaborators in the Duck Bot project",
			Commands:    make(map[string]dbi.Command),
		},
		c: c,
	}
	cmd.Help = collaboratorHelpFn(cmd)

	return cmd
}

func (c *CollaboratorCmd) Execute(args ...string) error {
	logging.Debug("Executing CollaboratorCmd")
	defer logging.Debug("Finished executing CollaboratorCmd")

	c.FlagSet.Parse(args)
	if c.FlagSet.NArg() < 2 {
		c.Help()
		return nil
	}
	owner := c.FlagSet.Arg(0)
	repo := c.FlagSet.Arg(1)
	if owner == "" || repo == "" {
		c.Help()
		return nil
	}
	collaborators, err := c.c.Collaborators(context.Background(), clients.CollaboratorRequest{
		Owner: owner,
		Repo:  repo,
	})
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return err
	}
	grep := c.FlagSet.Lookup("g").Value.String()
	if grep != "" {
		collaborators = slices.DeleteFunc(collaborators, func(collaborator string) bool {
			return !strings.Contains(collaborator, grep)
		})
	}
	fmt.Printf("Collaborators: %s\n", strings.Join(collaborators, ", "))
	// TOTO: Generate a slack message with the collaborators and send it to the slack channel
	return nil
}
