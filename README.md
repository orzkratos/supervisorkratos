# supervisorkratos

Go package for generating supervisor configuration files with Kratos microservices integration.

---

## CHINESE README

[中文说明](README.zh.md)

## Key Features

🎯 **Fluent Configuration API**: Chain methods for intuitive supervisor config building  
⚡ **Kratos Integration**: Optimized for Kratos microservices deployment patterns  
🔄 **Group Management**: Multi-program groups with centralized configuration  
🌍 **Production Ready**: Battle-tested configuration templates for high-performance services  
📋 **Type Safety**: Strongly typed configuration with sensible defaults

## Install

```bash
go get github.com/orzkratos/supervisorkratos
```

## Usage

### Single Program Configuration

```go
package main

import (
    "fmt"
    "github.com/orzkratos/supervisorkratos"
)

func main() {
    // Create program config with required parameters
    program := supervisorkratos.NewProgramConfig(
        "myapp",           // Program name
        "/opt/myapp",      // Program root DIR
        "deploy",          // User name
        "/var/log/myapp",  // Log DIR
    ).WithStartRetries(10).
      WithEnvironment(map[string]string{
          "APP_ENV": "production",
      })

    // Generate configuration
    config := supervisorkratos.GenerateProgramConfig(program)
    fmt.Println(config)
}
```

### Group Configuration

```go
// Create multiple programs
apiServer := supervisorkratos.NewProgramConfig(
    "api-server", "/opt/api-server", "deploy", "/var/log/services",
).WithStartRetries(3)

worker := supervisorkratos.NewProgramConfig(
    "worker", "/opt/worker", "deploy", "/var/log/services",
).WithAutoStart(false)

// Create group
group := supervisorkratos.NewGroupConfig("microservices").
    AddProgram(apiServer).
    AddProgram(worker)

config := supervisorkratos.GenerateGroupConfig(group)
```

### Advanced Configuration

```go
// High-performance service configuration
program := supervisorkratos.NewProgramConfig(
    "high-perf", "/opt/high-perf", "performance", "/var/log/perf",
).WithStartRetries(100).
  WithStopWaitSecs(60).
  WithLogMaxBytes("500MB").
  WithLogBackups(50).
  WithPriority(1)
```

### Multi-Instance Deployment

```go
// Multi-instance web server
program := supervisorkratos.NewProgramConfig(
    "web-server", "/opt/web-server", "deploy", "/var/log/cluster",
).WithNumProcs(3).
  WithProcessName("%(program_name)s_%(process_num)02d").
  WithEnvironment(map[string]string{
      "PORT_BASE": "8080",
  })
```

## Configuration Options

### Process Control
- `WithAutoStart(bool)` - Auto start on supervisor startup
- `WithAutoRestart(bool)` - Auto restart on failure  
- `WithStartRetries(int)` - Max start retry times
- `WithStartSecs(int)` - Seconds to wait before considering start successful

### Logging
- `WithLogMaxBytes(string)` - Max log file size (e.g., "50MB", "1GB")
- `WithLogBackups(int)` - Log backup files count
- `WithRedirectStderr(bool)` - Redirect stderr to stdout

### Process Management
- `WithStopWaitSecs(int)` - Graceful stop timeout seconds
- `WithStopSignal(string)` - Stop signal name (TERM, INT, QUIT)
- `WithKillAsGroup(bool)` - Kill child processes as group
- `WithPriority(int)` - Start priority (lower numbers start first)

### Multi-Instance
- `WithNumProcs(int)` - Number of process instances
- `WithProcessName(string)` - Process name template

### Environment
- `WithEnvironment(map[string]string)` - Environment variables
- `WithExitCodes([]int)` - Expected exit codes

## Recommended Workflow

```bash
# 1. Generate config file
go run main.go > /etc/supervisor/conf.d/myapp.conf

# 2. Update supervisor
sudo supervisorctl reread
sudo supervisorctl update

# 3. Control services
sudo supervisorctl start myapp
sudo supervisorctl status
```

<!-- TEMPLATE (EN) BEGIN: STANDARD PROJECT FOOTER -->

## 📄 License

MIT License. See [LICENSE](LICENSE).

---

## 🤝 Contributing

Contributions are welcome! Report bugs, suggest features, and contribute code:

- 🐛 **Found a bug?** Open an issue on GitHub with reproduction steps
- 💡 **Have a feature idea?** Create an issue to discuss the suggestion
- 📖 **Documentation confusing?** Report it so we can improve
- 🚀 **Need new features?** Share your use cases to help us understand requirements
- ⚡ **Performance issue?** Help us optimize by reporting slow operations
- 🔧 **Configuration problem?** Ask questions about complex setups
- 📢 **Follow project progress?** Watch the repo for new releases and features
- 🌟 **Success stories?** Share how this package improved your workflow
- 💬 **General feedback?** All suggestions and comments are welcome

---

## 🔧 Development

New code contributions, follow this process:

1. **Fork**: Fork the repo on GitHub (using the webpage interface).
2. **Clone**: Clone the forked project (`git clone https://github.com/yourname/repo-name.git`).
3. **Navigate**: Navigate to the cloned project (`cd repo-name`)
4. **Branch**: Create a feature branch (`git checkout -b feature/xxx`).
5. **Code**: Implement your changes with comprehensive tests
6. **Testing**: (Golang project) Ensure tests pass (`go test ./...`) and follow Go code style conventions
7. **Documentation**: Update documentation for user-facing changes and use meaningful commit messages
8. **Stage**: Stage changes (`git add .`)
9. **Commit**: Commit changes (`git commit -m "Add feature xxx"`) ensuring backward compatible code
10. **Push**: Push to the branch (`git push origin feature/xxx`).
11. **PR**: Open a pull request on GitHub (on the GitHub webpage) with detailed description.

Please ensure tests pass and include relevant documentation updates.

---

## 🌟 Support

Welcome to contribute to this project by submitting pull requests and reporting issues.

**Project Support:**

- ⭐ **Give GitHub stars** if this project helps you
- 🤝 **Share with teammates** and (golang) programming friends
- 📝 **Write tech blogs** about development tools and workflows - we provide content writing support
- 🌟 **Join the ecosystem** - committed to supporting open source and the (golang) development scene

**Happy Coding with this package!** 🎉

<!-- TEMPLATE (EN) END: STANDARD PROJECT FOOTER -->

---

## GitHub Stars

[![Stargazers](https://starchart.cc/orzkratos/supervisorkratos.svg?variant=adaptive)](https://starchart.cc/orzkratos/supervisorkratos)