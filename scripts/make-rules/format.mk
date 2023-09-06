##@ Format

.PHONY: format.install
format.install: ## 安装 goimports
	@go install golang.org/x/tools/cmd/goimports@latest
	@echo "goimports@latest install success"
	@go install github.com/reviewdog/reviewdog/cmd/reviewdog@latest
	@echo "reviewdog@latest install success"
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.53.0
	@echo "golangci-lint@v1.53.0 install success"

.PHONY: pretty
pretty: ## 美化 golang 文件
pretty: format.tidy format format.lint

.PHONY: format
format: ## 格式化 go 文件
	sh scripts/go-import-format.sh

.PHONY: format.tidy
format.tidy: ## go mod tidy
	GO111MODULE=on go mod tidy

.PHONY: format.lint
format.lint: ## go lint and reviewdog
	  golangci-lint run --max-same-issues=0 --out-format=line-number ./...
