unpacked_project/
├── client                            # （打包后的文件）
├── go.mod                            # Go 依赖管理定义文件
├── go.sum                            # Go 模块校验和文件
├── GoUnusedProtection__.go          # 用于防止未使用代码编译失败（Thrift 自动生成）
├── log_service.go                   # 核心接口及客户端定义（由 Thrift 生成）
├── log_service-consts.go           # 常量定义（由 Thrift 生成）
├── log_service-remote/
│   └── log_service-remote.go        # 服务端通信处理（由 Thrift 生成）
