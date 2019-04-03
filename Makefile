# Help
# This will parse the Makefile and generate help commands
.PHONEY: help

help: ## Show Makefile help
    @awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

build:
	go build

test:
	go test ./...

test-race:
	go test -race ./...

test-bench:
	go test -bench=.
