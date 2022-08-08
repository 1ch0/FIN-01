BINARY = server
IMAGE_REGISTRY = harbor.1ch0.com
IMAGE_REPO = 1ch0/server
IMAGE_TAG = latest

.PHONY: vendor build image release

.PHONY: all
all: check build
.PHONY: run
run:
	go run cmd/server/main.go

.PHONY: check
check:
	go fmt ./...
	go vet ./...
.PHONY: build
build:
	@go build -o $(BINARY) cmd/server/main.go
.PHONY: clean
clean:
	rm -f $(BINARY)
.PHONY: lint
lint:
	golangci-lint run -v --config ./.golangci.yml --timeout 5m
.PHONY: test
test:
	go test ./...
.PHONY: cover
cover:
	go test ./... -coverprofile coverage.out
	go tool cover -html=coverage.out
	rm -f coverage.out
.PHONY: vendor
vendor:
	go mod tidy && go mod vendor

build:
	go build -o $(BINARY) cmd/server/main.go
.PHONY: image
image:
	docker build -t $(IMAGE_REGISTRY)/$(IMAGE_REPO):$(IMAGE_TAG) .
.PHONY: release
release: image
	docker push $(IMAGE_REGISTRY)/$(IMAGE_REPO):$(IMAGE_TAG)

## swagger: Generate swagger document.
.PHONY: swagger
swagger:
	swag init -g ./cmd/server/main.go -d ./

.PHONY: help
help:
	@echo "============================================="
	@echo "make all     格式化go代码 并编译生成二进制文件"
	@echo "make build   编译go代码生成二进制文件"
	@echo "make clean   清理中间目标文件"
	@echo "make test    执行测试case"
	@echo "make check   格式化go代码"
	@echo "make cover   检查测试覆盖率"
	@echo "make run     直接运行程序"
	@echo "make lint    执行代码检查"
	@echo "make image   构建docker镜像"
	@echo "make release 推送docker镜像"
	@echo "make swagger 生成 swagger 接口文档"
	@echo "============================================="