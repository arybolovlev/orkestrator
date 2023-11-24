##@ Development

.PHONY: fmt
fmt: ## Run go fmt against code.
	go fmt ./...

.PHONY: vet
vet: ## Run go vet against code.
	go vet ./...

.PHONY: test
test: fmt vet ## Run tests.
	go test -timeout 5m -v ./...

##@ Build

.PHONY: build
build: fmt vet ## Build orkestrator binary.
	go build -a -trimpath -o bin/orkestrator cmd/main.go
