.PHONY: lint test setup-hooks

lint: ## Run golangci-lint
	golangci-lint run

test: ## Run all tests
	go test ./...

setup-hooks: ## Install Git pre-commit hooks
	git config core.hooksPath .githooks
	@echo "git hooks installed from .githooks/"
