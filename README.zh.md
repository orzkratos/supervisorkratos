# supervisorkratos

用于为 Kratos 微服务生成 supervisor 配置文件的 Go 包。

---

## ENGLISH README

[English Documentation](README.md)

## 核心特性

🎯 **流畅配置 API**: 链式方法调用，直观的配置构建体验  
⚡ **Kratos 集成**: 专为 Kratos 微服务部署模式优化  
🔄 **组管理**: 多程序组的集中配置管理  
🌍 **生产就绪**: 经过实战验证的高性能服务配置模板  
📋 **类型安全**: 强类型配置，内置合理默认值

## 安装

```bash
go get github.com/orzkratos/supervisorkratos
```

## 使用方法

### 单程序配置

```go
package main

import (
    "fmt"
    "github.com/orzkratos/supervisorkratos"
)

func main() {
    // 创建程序配置，提供必需参数
    program := supervisorkratos.NewProgramConfig(
        "myapp",           // 程序名称
        "/opt/myapp",      // 程序根目录
        "deploy",          // 运行用户
        "/var/log/myapp",  // 日志目录
    ).WithStartRetries(10).
      WithEnvironment(map[string]string{
          "APP_ENV": "production",
      })

    // 生成配置
    config := supervisorkratos.GenerateProgramConfig(program)
    fmt.Println(config)
}
```

### 组配置

```go
// 创建多个程序
apiServer := supervisorkratos.NewProgramConfig(
    "api-server", "/opt/api-server", "deploy", "/var/log/services",
).WithStartRetries(3)

worker := supervisorkratos.NewProgramConfig(
    "worker", "/opt/worker", "deploy", "/var/log/services",
).WithAutoStart(false)

// 创建程序组
group := supervisorkratos.NewGroupConfig("microservices").
    AddProgram(apiServer).
    AddProgram(worker)

config := supervisorkratos.GenerateGroupConfig(group)
```

### 高级配置

```go
// 高性能服务配置
program := supervisorkratos.NewProgramConfig(
    "high-perf", "/opt/high-perf", "performance", "/var/log/perf",
).WithStartRetries(100).
  WithStopWaitSecs(60).
  WithLogMaxBytes("500MB").
  WithLogBackups(50).
  WithPriority(1)
```

### 多实例部署

```go
// 多实例 Web 服务器
program := supervisorkratos.NewProgramConfig(
    "web-server", "/opt/web-server", "deploy", "/var/log/cluster",
).WithNumProcs(3).
  WithProcessName("%(program_name)s_%(process_num)02d").
  WithEnvironment(map[string]string{
      "PORT_BASE": "8080",
  })
```

## 配置选项

### 进程控制
- `WithAutoStart(bool)` - supervisor 启动时自动启动
- `WithAutoRestart(bool)` - 失败时自动重启  
- `WithStartRetries(int)` - 最大启动重试次数
- `WithStartSecs(int)` - 启动成功前等待秒数

### 日志设置
- `WithLogMaxBytes(string)` - 最大日志文件大小（如："50MB", "1GB"）
- `WithLogBackups(int)` - 日志备份文件数量
- `WithRedirectStderr(bool)` - 重定向 stderr 到 stdout

### 进程管理
- `WithStopWaitSecs(int)` - 优雅停止超时秒数
- `WithStopSignal(string)` - 停止信号名称（TERM, INT, QUIT）
- `WithKillAsGroup(bool)` - 作为组强制杀死子进程
- `WithPriority(int)` - 启动优先级（数字越小优先级越高）

### 多实例
- `WithNumProcs(int)` - 进程实例数量
- `WithProcessName(string)` - 进程名称模板

### 环境变量
- `WithEnvironment(map[string]string)` - 环境变量设置
- `WithExitCodes([]int)` - 期望的退出码

## 推荐工作流程

```bash
# 1. 生成配置文件
go run main.go > /etc/supervisor/conf.d/myapp.conf

# 2. 更新 supervisor
sudo supervisorctl reread
sudo supervisorctl update

# 3. 控制服务
sudo supervisorctl start myapp
sudo supervisorctl status
```

<!-- TEMPLATE (ZH) BEGIN: STANDARD PROJECT FOOTER -->

## 📄 许可证类型

MIT 许可证。详见 [LICENSE](LICENSE)。

---

## 🤝 项目贡献

非常欢迎贡献代码！报告 BUG、建议功能、贡献代码：

- 🐛 **发现问题？** 在 GitHub 上提交问题并附上重现步骤
- 💡 **功能建议？** 创建 issue 讨论您的想法
- 📖 **文档疑惑？** 报告问题，帮助我们改进文档
- 🚀 **需要功能？** 分享使用场景，帮助理解需求
- ⚡ **性能瓶颈？** 报告慢操作，帮助我们优化性能
- 🔧 **配置困扰？** 询问复杂设置的相关问题
- 📢 **关注进展？** 关注仓库以获取新版本和功能
- 🌟 **成功案例？** 分享这个包如何改善工作流程
- 💬 **意见反馈？** 欢迎所有建议和宝贵意见

---

## 🔧 代码贡献

新代码贡献，请遵循此流程：

1. **Fork**：在 GitHub 上 Fork 仓库（使用网页界面）
2. **克隆**：克隆 Fork 的项目（`git clone https://github.com/yourname/repo-name.git`）
3. **导航**：进入克隆的项目（`cd repo-name`）
4. **分支**：创建功能分支（`git checkout -b feature/xxx`）
5. **编码**：实现您的更改并编写全面的测试
6. **测试**：（Golang 项目）确保测试通过（`go test ./...`）并遵循 Go 代码风格约定
7. **文档**：为面向用户的更改更新文档，并使用有意义的提交消息
8. **暂存**：暂存更改（`git add .`）
9. **提交**：提交更改（`git commit -m "Add feature xxx"`）确保向后兼容的代码
10. **推送**：推送到分支（`git push origin feature/xxx`）
11. **PR**：在 GitHub 上打开 Pull Request（在 GitHub 网页上）并提供详细描述

请确保测试通过并包含相关的文档更新。

---

## 🌟 项目支持

非常欢迎通过提交 Pull Request 和报告问题来为此项目做出贡献。

**项目支持：**

- ⭐ **给予星标**如果项目对您有帮助
- 🤝 **分享项目**给团队成员和（golang）编程朋友
- 📝 **撰写博客**关于开发工具和工作流程 - 我们提供写作支持
- 🌟 **加入生态** - 致力于支持开源和（golang）开发场景

**使用这个包快乐编程！** 🎉

<!-- TEMPLATE (ZH) END: STANDARD PROJECT FOOTER -->

---

## GitHub 标星点赞

[![Stargazers](https://starchart.cc/orzkratos/supervisorkratos.svg?variant=adaptive)](https://starchart.cc/orzkratos/supervisorkratos)