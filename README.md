# Admin Server

这是一个基于 Go-Zero 框架的后台管理系统。

## API 文档

本项目使用 Swagger 2.0 来生成 API 文档。你可以通过以下方式查看 API 文档：

1. 使用 Swagger UI Docker 镜像（推荐）
   ```bash
   # 启动 Swagger UI
   docker run --name swagger-ui -d \
     -p 8080:8080 \
     -e SWAGGER_JSON=/swagger.json \
     -v $PWD/swagger.json:/swagger.json \
     swaggerapi/swagger-ui

   # 停止服务
   docker stop swagger-ui
   docker rm swagger-ui
   ```
   然后访问 http://localhost:8080

2. 在线查看
   - 访问 [Swagger Editor](https://editor.swagger.io/)
   - 将 `swagger.json` 或 `swagger.yaml` 文件的内容复制到编辑器中

## 开发

### 生成 API 文档

1. 安装 goctl-swagger
   ```bash
   go install github.com/zeromicro/goctl-swagger@latest
   ```

2. 生成 swagger 文档
   ```bash
   # 生成 JSON 格式
   goctl api plugin -plugin goctl-swagger="swagger -filename swagger.json" -api admin.api -dir .
   ```

### 更新 API 文档

1. 修改 `admin.api` 文件
2. 重新生成 swagger 文档
   ```bash
   goctl api plugin -plugin goctl-swagger="swagger -filename swagger.json" -api admin.api -dir .
   ```

### 查看 API 文档

1. 启动 Swagger UI
   ```bash
   docker run --name swagger-ui -d \
     -p 8080:8080 \
     -e SWAGGER_JSON=/swagger.json \
     -v $PWD/swagger.json:/swagger.json \
     swaggerapi/swagger-ui
   ```

2. 访问文档
   - 打开浏览器访问 http://localhost:8080
   - 可以看到所有的 API 接口文档
   - 支持在线调试 API（需要后端服务在运行）

3. 关闭服务
   ```bash
   docker stop swagger-ui
   docker rm swagger-ui
   ``` 