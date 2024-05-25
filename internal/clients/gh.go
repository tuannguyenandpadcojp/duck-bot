package clients

import (
	"context"
	"net/http"

	"github.com/google/go-github/v62/github"
)

type GithubConfig struct {
	Client *http.Client
	PAT    string
}

type Github struct {
	c *github.Client
}

func NewGithub(cfg GithubConfig) *Github {
	ghClient := github.NewClient(cfg.Client).WithAuthToken(cfg.PAT)
	return &Github{
		c: ghClient,
	}
}

type CollaboratorRequest struct {
	Owner string
	Repo  string
}

func (g *Github) Collaborators(ctx context.Context, req CollaboratorRequest) ([]string, error) {
	collaborators, _, err := g.c.Repositories.ListCollaborators(context.Background(), req.Owner, req.Repo, nil)
	if err != nil {
		return nil, err
	}

	var cNames []string
	for _, collaborator := range collaborators {
		name := collaborator.GetLogin()
		if name == "" {
			continue
		}
		cNames = append(cNames, name)
	}

	return cNames, nil
}
