##@ General

help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Development

tools: ## Install tools required for development
	@echo "Installing tools..."
	@go get goa.design/model/cmd/mdl
	@go install goa.design/model/cmd/mdl
	@go get goa.design/model/cmd/stz
	@go install goa.design/model/cmd/stz
	@echo "Tools installed"

ARGS =
run: ## Run the application
	@export $(shell cat config/.local.env | xargs) && go run cmd/main.go $(ARGS)

grep =
collaborations: ## Generate collaborations diagram
	@export $(shell cat config/.local.env | xargs) && go run cmd/main.go gh collaborations $(grep)

##@ Documentation

gen-docs: ## Start the documentation server
	@echo "Starting generate C4 model"
	@docker-compose -f docs/docker-compose.yml up -d
	@echo "C4 model generated"
