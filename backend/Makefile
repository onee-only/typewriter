# General
WORKDIR = $(PWD)
# Go parameters
GOCMD = go
GOTEST = $(GOCMD) test -v

.PHONY: lint
lint:
	@golangci-lint run

.PHONY: test
test: 
	@$(GOTEST) -race ./...