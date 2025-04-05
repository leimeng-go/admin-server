# 设置变量
BINARY=admin-server
API_FILE=admin.api
SWAGGER_FILE=docs/swagger.json

# 默认目标
.PHONY: all
all: clean swagger build

# 生成 API 代码
.PHONY: api
api:
	@echo "Cleaning up old generated files..."
	@rm -rf internal/handler
	@rm -rf internal/logic
	@rm -rf internal/svc
	@rm -rf internal/types
	@rm -rf internal/middleware
	@echo "Generating API code..."
	goctl api go -api $(API_FILE) -dir . -style go_zero

# 生成 swagger 文档
.PHONY: swagger
swagger:
	@echo "Generating Swagger documentation..."
	@mkdir -p docs
	goctl api plugin -plugin goctl-swagger="swagger -filename $(SWAGGER_FILE)" -api $(API_FILE)

# 构建项目
.PHONY: build
build:
	@echo "Building $(BINARY)..."
	go build -o $(BINARY)

# 运行项目
.PHONY: run
run:
	@echo "Running $(BINARY)..."
	go run .

# 清理生成的文件
.PHONY: clean
clean:
	@echo "Cleaning..."
	@rm -f $(BINARY)
	@rm -f $(SWAGGER_FILE)
	@rm -rf internal/handler
	@rm -rf internal/logic
	@rm -rf internal/svc
	@rm -rf internal/types
	@rm -rf internal/middleware

# 格式化代码
.PHONY: fmt
fmt:
	@echo "Formatting code..."
	go fmt ./...

# 运行测试
.PHONY: test
test:
	@echo "Running tests..."
	go test -v ./...

# 帮助信息
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  make          - Generate API code, Swagger docs and build binary"
	@echo "  make api      - Generate API code from $(API_FILE)"
	@echo "  make swagger  - Generate Swagger documentation"
	@echo "  make build    - Build the binary"
	@echo "  make run      - Run the application"
	@echo "  make clean    - Clean generated files"
	@echo "  make fmt      - Format code"
	@echo "  make test     - Run tests"
	@echo "  make help     - Show this help message" 