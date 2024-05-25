package internal

import (
	"context"

	"github.com/tuannguyenandpadcojp/duck-bot/internal/clients"
)

type GithubClient interface {
	Collaborators(ctx context.Context, req clients.CollaboratorRequest) ([]string, error)
}
