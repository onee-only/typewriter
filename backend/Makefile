# General
WORKDIR = $(PWD)
# Go parameters
GOCMD = go
GOTEST = $(GOCMD) test -v

lint:
	@golangci-lint run

test:
	@$(GOTEST) -race ./...