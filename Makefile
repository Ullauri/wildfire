WF_MAIN := cmd/wildfire/main.go
BINARY_NAME := wildfire

.PHONY: help go_build run

help:
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

go_build: ## build project binary
	go build -o ${BINARY_NAME} ${WF_MAIN}

run: go_build ## run server locally
	./${BINARY_NAME}

test: ## run tests
	go test ./... -v -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html
