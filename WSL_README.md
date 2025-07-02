# WSL 环境使用说明

## 环境适配

本项目已针对 WSL (Windows Subsystem for Linux) 环境进行了适配，主要解决了以下问题：

### 1. 路径问题
- 自动检测操作系统类型
- Linux/WSL 环境使用相对路径
- macOS 环境使用绝对路径

### 2. Swagger 浏览器问题
- WSL 环境下无法直接打开浏览器
- 提供了手动访问的 URL 信息
- 支持通过 WSL IP 地址访问

## 常用命令

### 生成 API 代码
```bash
make api
```

### 生成 Model 代码
```bash
make model
```

### 生成 Swagger 文档
```bash
make swagger
```

### 运行 Swagger UI (推荐)
```bash
make swagger-serve
```
启动后会在终端显示访问地址，通常在 Windows 浏览器中访问：
- `http://localhost:8080`
- 如果无法访问，尝试：`http://<WSL_IP>:8080`

### 构建项目
```bash
make build
```

### 运行项目
```bash
make run
```

### 完整流程
```bash
make all  # 清理、生成代码、构建并运行
```

## 故障排除

### 1. 端口访问问题
如果 `localhost:8080` 无法访问，请尝试：
```bash
# 查看 WSL IP 地址
hostname -I

# 使用 WSL IP 地址访问
# 例如：http://172.17.0.1:8080
```

### 2. 防火墙问题
确保 Windows 防火墙允许 WSL 的网络访问。

### 3. 模板路径问题
如果遇到模板路径错误，请检查：
```bash
ls -la api/internal/template/
```

### 4. Swagger 安装问题
如果 swagger 命令不存在：
```bash
make swagger-install
```

## 开发建议

1. **使用 `make swagger-serve` 而不是 `make swagger-run`**，避免浏览器打开失败的错误
2. **在 Windows 浏览器中访问 Swagger UI**，而不是在 WSL 终端中
3. **使用 `make help` 查看所有可用命令**
4. **定期运行 `make clean` 清理生成的文件**

## 网络配置

如果遇到网络访问问题，可以在 Windows PowerShell 中运行：
```powershell
# 查看 WSL 网络配置
wsl --list --verbose
```

或者在 WSL 中运行：
```bash
# 查看网络接口
ip addr show
``` 