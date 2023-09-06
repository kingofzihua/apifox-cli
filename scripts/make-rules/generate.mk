##@ Generate

.PHONY: generate.install
generate.install: ## install generate tools
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest


.PHONY: generate
generate: ## go generate
	go mod tidy
	go generate ./...
