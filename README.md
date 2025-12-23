# Go-Nexus

一个功能丰富的 Go 语言工具库集合，提供了日志、Redis、配置管理、数据验证、序列化、加密等常用功能模块。

## 功能特性

### 📦 模块列表

- **log** - 基于 zap 的高性能日志模块
- **redis** - Redis 客户端封装，支持多种数据结构和 Token 管理
- **snowflake** - 雪花算法 ID 生成器
- **validator** - 数据验证工具，支持自定义错误消息
- **viper** - 配置文件管理，支持配置合并
- **serialize** - 序列化工具，支持 JSON、JSONIter、Protobuf
- **model** - 数据模型，支持 MySQL、PostgreSQL、SQLite 三种数据库
- **nexusenum** - 枚举模块，提供状态、标志、方法等枚举类型
- **response** - 统一响应结构，支持泛型
- **utils** - 工具函数集合（文件操作、结构体操作、加密等）

## 安装

```bash
go get github.com/nilchaosky/go-nexus
```

## 模块详细说明

### Log 模块

基于 `zap` 的高性能日志模块，支持：

- 多种日志级别（Debug、Info、Warn、Error）
- 多种编码格式（JSON、Console）
- 日志文件自动轮转
- 自动清理过期日志
- 自定义输出目录

### Redis 模块

完整的 Redis 客户端封装，实现了以下接口：

- **Generic** - 通用操作（Del、Exists、Expire、TTL、Keys、Scan 等）
- **String** - 字符串操作（Get、Set、Cache、Incr、MGet 等）
- **List** - 列表操作（LPush、RPush、LPop、LRange、BLPop 等）
- **Set** - 集合操作（SAdd、SMembers、SInter、SUnion、SPop 等）
- **Hash** - 哈希操作（HGet、HSet、HGetAll、HMGet、HVals 等）
- **ZSet** - 有序集合操作（ZAdd、ZRange、ZScore、ZRank、ZPopMax 等）
- **Token** - Token 管理（SaveToken、GetToken、DeleteToken、RefreshToken 等）

支持自动序列化/反序列化，可直接操作结构体。

### Token 模块

基于 JWT 的 Token 管理模块，支持：

- Token 生成和验证
- Token 刷新机制
- 额外参数存储和获取
- Redis 存储管理

### Snowflake 模块

雪花算法 ID 生成器，支持：

- 分布式唯一 ID 生成
- 数据中心 ID 和 Worker ID 配置
- 序列化类型支持

### Validator 模块

数据验证工具，基于 `go-playground/validator`，支持：

- 丰富的验证标签
- 自定义错误消息格式化
- 支持嵌套结构体和切片验证

### Viper 模块

配置文件管理，支持：

- 多种配置文件格式（YAML、JSON、TOML 等）
- 配置合并（只替换零值字段）
- 泛型支持

### Serialize 模块

序列化工具，支持：

- JSON（标准库）
- JSONIter（高性能）
- Protobuf

### Model 模块

数据模型，支持 MySQL、PostgreSQL、SQLite 三种数据库，提供：

- **ID 模型** - `Snowflake`（雪花 ID 主键）、`AutoIncrement`（自增 ID 主键，MySQL 和 PostgreSQL 支持）
- **时间模型** - `Timestamps`（创建时间、更新时间）、`SoftDelete`（软删除）
- **状态模型** - `Status`（状态字段，使用状态枚举）
- **备注模型** - `Remark`（备注字段）
- **乐观锁** - `OptimisticLock`（版本号字段）
- 支持 GORM 和 JSON 序列化
- 各数据库使用对应的类型标签，确保兼容性

### Response 模块

统一响应结构，支持：

- 泛型响应类型
- 成功/错误响应
- 分页响应
- 类型安全的 nil 响应

### NexusEnum 模块

枚举模块，提供常用的枚举类型：

- **Status** - 状态枚举（禁用、启用）
- **Flag** - 标志枚举（否、是），支持与布尔值转换
- **Method** - 方法枚举（GET、POST、PUT、DELETE、PATCH）

所有枚举类型提供 `String()` 和 `Value()` 方法。

### Utils 模块

工具函数集合：

- **crypto** - 加密工具（bcrypt 密码哈希）
- **file** - 文件操作（目录判断等）
- **struct** - 结构体操作（指针/切片判断等）

## 依赖

主要依赖：

- `github.com/redis/go-redis/v9` - Redis 客户端
- `go.uber.org/zap` - 日志库
- `github.com/spf13/viper` - 配置管理
- `github.com/go-playground/validator/v10` - 数据验证
- `github.com/golang-jwt/jwt/v5` - JWT Token
- `golang.org/x/crypto` - 加密工具
- `github.com/GUAIK-ORG/go-snowflake` - 雪花算法

## 许可证

MIT License

## 贡献

欢迎提交 Issue 和 Pull Request！

