# Help
# This will parse the Makefile and generate help commands
.PHONEY: help

help: ## Show Makefile help
    @awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

install:
	go get gotest.tools/gotestsum
	go get github.com/stretchr/testify/assert

build:
	go build

test:
	gotestsum --format short-verbose

test-race:
	gotestsum --format short-verbose -- -race

test-bench:
	gotestsum --format short-verbose -- -bench=. -run=^$$

test-ci:
	gotestsum --junitfile reports/unit-tests.xml -- -bench=. -run=^$$

test-ci-bench:
	go test -run=^$$ -bench=. ./... | tee reports/bench.txt

test-ci-race:
	gotestsum --junitfile reports/race-tests.xml -- -race
