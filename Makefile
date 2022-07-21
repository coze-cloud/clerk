.PHONY: help
help: ##Show the help information
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

.PHONY: test
test: ##Run all tests
	go test `go list ./...` -timeout 15s -count=1

.PHONY: lint
lint: ## Run linters
	golangci-lint run --enable gofmt
