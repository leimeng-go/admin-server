# 设置变量
BINARY=admin-server
API_FILE=admin.api
API_DIR=./api
TS_DIR=./web
SWAGGER_FILE_NAME=swagger
DOCS_DIR=docs
MODEL_DIR=api/internal/model
TEMPLATE_HOME=./api/internal/template

# 默认目标
.PHONY: all
all: clean api model build run  ## Clean, generate API/model, build and run

# 生成 API 代码
.PHONY: api
api: ## Generate API code
	@echo "Generating API code..."
	goctl api go -api $(API_DIR)/$(API_FILE) -dir $(API_DIR) -style go_zero --home "$(TEMPLATE_HOME)" --test --type-group 

# 生成 model 代码
.PHONY: model
model: ## Generate model code from MySQL DDL
	@echo "Generating model code from MySQL DDL..."
	@find ./internal/model -type d -name sql | while read dir; do \
		goctl model mysql ddl -src="$$dir"/*.sql -dir=$${dir%/sql} -style=go_zero -cache=true --home "$(TEMPLATE_HOME)"; \
	done

# 生成 swagger 文档
.PHONY: swagger
swagger: ## Generate Swagger documentation
	@echo "Generating Swagger documentation..."
	@mkdir -p docs
	goctl api swagger --api $(API_DIR)/$(API_FILE) --dir $(DOCS_DIR) --filename $(SWAGGER_FILE_NAME)

# 安装 swagger
.PHONY: swagger-install
swagger-install: ## Install go-swagger if not present
	@if ! command -v swagger &> /dev/null; then \
		echo "swagger 未安装，正在从 GitHub 安装..."; \
		go install github.com/go-swagger/go-swagger/cmd/swagger@latest; \
		if [ $$? -ne 0 ]; then \
			echo "安装 swagger 失败"; \
			exit 1; \
		fi; \
		echo "swagger 安装成功"; \
	else \
		echo "swagger 已安装"; \
	fi

# 运行swagger (WSL环境适配)
.PHONY: swagger-run
swagger-run: ## Serve Swagger UI for the generated swagger.json
	@echo "启动 Swagger UI 服务器..."
	@echo "在 WSL 环境下，请手动在 Windows 浏览器中访问: http://localhost:8080"
	@echo "如果无法访问，请尝试: http://$(shell hostname -I | awk '{print $$1}'):8080"
	swagger serve -F=swagger --host=0.0.0.0 --port=8080 --no-open $(DOCS_DIR)/$(SWAGGER_FILE_NAME).json

# 构建项目
.PHONY: build
build: ## Build the binary
	@echo "Building $(BINARY)..."
	go build -o $(BINARY) ./api/

# 运行项目
.PHONY: run
run: ## Run the application
	@echo "Running $(BINARY)..."
	./$(BINARY) -f api/etc/admin.yaml

# 清理生成的文件
.PHONY: clean
clean: ## Clean generated files
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
fmt: ## Format code
	@echo "Formatting code..."
	go fmt ./...

# 运行测试
.PHONY: test
test: ## Run tests
	@echo "Running tests..."
	go test -v ./...

# 帮助信息
.PHONY: help
help:
	@echo "Available targets in ./api/:"
	@grep -E '^[a-zA-Z0-9_-]+:.*?## ' Makefile | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-18s\033[0m %s\n", $$1, $$2}'

# 批量生成所有api的swagger文档
.PHONY: swagger-all
swagger-all: ## Generate Swagger documentation for all .api files
	@echo "Generating Swagger documentation for all .api files..."
	@mkdir -p $(DOCS_DIR)
	@for file in $(API_DIR)/*.api; do \
	  name=$$(basename $$file .api); \
	  goctl api swagger --api $$file --dir $(DOCS_DIR) --filename $$name.swagger; \
	done 