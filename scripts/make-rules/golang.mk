
.PHONY: go.build
go.build:## 构建 go 二进制文件
	mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./...
