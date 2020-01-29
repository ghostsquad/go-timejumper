# Help
# This will parse the Makefile and generate help commands
.PHONY: help
help: ## Show Makefile help
    @awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

.PHONY: install
install:
	go get gotest.tools/gotestsum
	go get github.com/stretchr/testify/assert

.PHONY: build
build:
	go build

.PHONY: test
test:
	gotestsum --format short-verbose

.PHONY: test-race
test-race:
	gotestsum --format short-verbose -- -race

.PHONY: test-bench
test-bench:
	gotestsum --format short-verbose -- -bench=. -run=^$$

.PHONY: test-ci
test-ci:
	gotestsum --junitfile reports/unit-tests.xml -- -bench=. -run=^$$

.PHONY: test-ci-bench
test-ci-bench:
	go test -run=^$$ -bench=. ./... | tee reports/bench.txt

.PHONY: test-ci-race
test-ci-race:
	gotestsum --junitfile reports/race-tests.xml -- -race
