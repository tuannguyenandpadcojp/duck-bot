package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tuannguyenandpadcojp/duck-bot/internal"
	"github.com/tuannguyenandpadcojp/duck-bot/internal/clients"
	"github.com/tuannguyenandpadcojp/duck-bot/internal/config"
	"github.com/tuannguyenandpadcojp/duck-bot/internal/gh"
	"github.com/tuannguyenandpadcojp/duck-bot/pkg/logging"
	"github.com/tuannguyenandpadcojp/duck-bot/pkg/utils"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: duck-bot <command> [arguments]")
		return
	}
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	// init logger
	logging.SetDefaultLogger(
		logging.NewZLogger(logging.ZLoggerConfig{
			Env: cfg.Env,
		}),
	)
	// initialize dependencies
	ghClient := clients.NewGithub(clients.GithubConfig{
		PAT:    cfg.PAT,
		Client: utils.NewHTTPClient(),
	})

	// init command manager
	manager := internal.NewCommandManager()
	manager.RegisterCommand(gh.Commands(&gh.Config{
		GithubClient: ghClient,
	}))

	logging.Debug("Running command")
	if err := manager.Run(os.Args[1], os.Args[2:]...); err != nil {
		logging.Errorf("failed to run command: %v", err)
	}
}
