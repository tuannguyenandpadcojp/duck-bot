## Project Structure

The project follows a standard directory structure:
```bash
.
├── LICENSE
├── Makefile               # Makefile for building, testing, and running the project
├── README.md              # The top-level README for developers using this project
├── cmd
│   └── main.go            # The entry point for the application
├── config                 # Configuration files
│   └── sample.env         # Sample environment variables
├── docs                   # Documentation files
│   ├── skeleton.md        # Project structure
├── go.mod                 # Go module file
├── go.sum                 # Go sum file
├── internal
│   ├── clients            # External clients for interacting with external services
│   │   ├── gh.go
│   │   └── slack.go
│   ├── config             # Config structure and method to load config
│   │   └── config.go
│   ├── errors             # Error definitions and handling
│   │   └── errors.go
│   ├── gh                 # GitHub package commands for interacting with GitHub
│   │   ├── internal
│   │   │   ├── client.go             # GitHub client interface
│   │   │   ├── collaborators.go      # GitHub collaborators command implementation
│   │   │   └── root.go               # GitHub root command implementation and registration of subcommands `gh <command> ...`
│   │   └── register.go               # Register GitHub commands
│   ├── jira
│   │   ├── internal                  # Jira package commands for interacting with Jira
│   │   └── register.go               # Register Jira commands
│   ├── command.go                    # Define `command` interface and implementation for basic command structure
│   └── manager.go                    # Command manager to manage all commands
└── pkg
    ├── logging                       # Logging package for logging messages and errors in the application
    │   ├── logger.go                 # Logger interface
    │   ├── noop_logger.go            # Noop logger implementation
    │   ├── slog_logger.go            # Standard logger implementation
    │   └── zero_logger.go            # Zero logger implementation
    └── utils                         # Utility functions
        └── http.go                   # HTTP utility functions
```
