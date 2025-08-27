package supervisorkratos_test

import (
	"testing"

	"github.com/orzkratos/supervisorkratos"
	"github.com/stretchr/testify/require"
)

func TestSingleProgramConfig(t *testing.T) {
	// Test single program config without group
	// 测试单个程序配置（不含组）
	cfg := supervisorkratos.NewProgramConfig(
		"myapp",
		"/opt/myapp",
		"deploy",
		"/var/log/myapp",
	).WithStartRetries(10).
		WithEnvironment(map[string]string{
			"APP_ENV": "production",
		})

	// Generate config for single program
	// 生成单个程序配置
	content := supervisorkratos.GenerateProgramConfig(cfg)

	t.Log("=== Single Program Configuration ===")
	t.Log(content)

	const expected = `[program:myapp]
user            = deploy
directory       = /opt/myapp
command         = /opt/myapp/bin/myapp
environment     = APP_ENV=production

startretries    = 10

stdout_logfile  = /var/log/myapp/myapp.log

stderr_logfile  = /var/log/myapp/myapp.err

`

	require.Equal(t, expected, content)
}

func TestNewGroupConfig(t *testing.T) {
	// Test new GroupConfig structure with multiple programs
	// 测试新的 GroupConfig 结构与多个程序
	program1 := supervisorkratos.NewProgramConfig(
		"api-server",
		"/opt/api-server",
		"deploy",
		"/var/log/services",
	).WithStartRetries(3)

	program2 := supervisorkratos.NewProgramConfig(
		"worker",
		"/opt/worker",
		"deploy",
		"/var/log/services",
	).WithAutoStart(false)

	group := supervisorkratos.NewGroupConfig("microservices").
		AddProgram(program1).
		AddProgram(program2)

	content := supervisorkratos.GenerateGroupConfig(group)
	t.Log("=== New GroupConfig Structure ===")
	t.Log(content)

	const expected = `[group:microservices]
programs=api-server,worker


[program:api-server]
user            = deploy
directory       = /opt/api-server
command         = /opt/api-server/bin/api-server

stdout_logfile  = /var/log/services/api-server.log

stderr_logfile  = /var/log/services/api-server.err

[program:worker]
user            = deploy
directory       = /opt/worker
command         = /opt/worker/bin/worker

autostart       = false

stdout_logfile  = /var/log/services/worker.log

stderr_logfile  = /var/log/services/worker.err
`

	require.Equal(t, expected, content)
}

func TestAdvancedProgramConfig(t *testing.T) {
	// Test advanced program configuration options
	// 测试高级程序配置选项
	program := supervisorkratos.NewProgramConfig(
		"advanced-service",
		"/opt/advanced",
		"performance",
		"/var/log/advanced",
	).WithStopWaitSecs(60).
		WithStopSignal("INT").
		WithPriority(100).
		WithKillAsGroup(true).
		WithExitCodes([]int{0, 1, 2})

	content := supervisorkratos.GenerateProgramConfig(program)
	t.Log("=== Advanced Program Configuration ===")
	t.Log(content)

	const expected = `[program:advanced-service]
user            = performance
directory       = /opt/advanced
command         = /opt/advanced/bin/advanced-service

stdout_logfile  = /var/log/advanced/advanced-service.log

stderr_logfile  = /var/log/advanced/advanced-service.err

stopwaitsecs    = 60
killasgroup     = true
stopsignal      = INT
priority        = 100
exitcodes       = 0,1,2
`

	require.Equal(t, expected, content)
}

func TestWithCustomization(t *testing.T) {
	// Test customization (from old version git diff)
	// 测试定制化配置（来自旧版本 git diff）
	program := supervisorkratos.NewProgramConfig(
		"service1",
		"/opt/service1",
		"deploy",
		"/var/log/services",
	).WithStartRetries(50).
		WithLogMaxBytes("100MB").
		WithRedirectStderr(true)

	content := supervisorkratos.GenerateProgramConfig(program)
	t.Log("=== Required parameters + chain customization ===")
	t.Log(content)

	// Using exact format from old version git diff
	// 使用旧版本 git diff 中的确切格式
	const expected = `[program:service1]
user            = deploy
directory       = /opt/service1
command         = /opt/service1/bin/service1

startretries    = 50

stdout_logfile  = /var/log/services/service1.log
stdout_logfile_maxbytes = 100MB

stderr_logfile  = /var/log/services/service1.err
stderr_logfile_maxbytes = 100MB
redirect_stderr = true

`

	require.Equal(t, expected, content)
}

func TestMultiInstanceConfig(t *testing.T) {
	// Test multi-instance deployment
	// 测试多实例部署
	program := supervisorkratos.NewProgramConfig(
		"web-server",
		"/opt/web-server",
		"deploy",
		"/var/log/cluster",
	).WithNumProcs(3).
		WithProcessName("%(program_name)s_%(process_num)02d").
		WithEnvironment(map[string]string{
			"PORT_BASE": "8080",
		})

	content := supervisorkratos.GenerateProgramConfig(program)
	t.Log("=== Multi-instance deployment ===")
	t.Log(content)

	const expected = `[program:web-server]
user            = deploy
directory       = /opt/web-server
command         = /opt/web-server/bin/web-server
environment     = PORT_BASE=8080

stdout_logfile  = /var/log/cluster/web-server.log

stderr_logfile  = /var/log/cluster/web-server.err

numprocs        = 3
process_name    = %(program_name)s_%(process_num)02d
`

	require.Equal(t, expected, content)
}

func TestHighPerformanceConfig(t *testing.T) {
	// Test high performance settings
	// 测试高性能设置
	program := supervisorkratos.NewProgramConfig(
		"high-perf",
		"/opt/high-perf",
		"performance",
		"/var/log/perf",
	).WithStartRetries(100).
		WithStopWaitSecs(60).
		WithLogMaxBytes("500MB").
		WithLogBackups(50).
		WithPriority(1)

	content := supervisorkratos.GenerateProgramConfig(program)
	t.Log("=== High performance configuration ===")
	t.Log(content)

	const expected = `[program:high-perf]
user            = performance
directory       = /opt/high-perf
command         = /opt/high-perf/bin/high-perf

startretries    = 100

stdout_logfile  = /var/log/perf/high-perf.log
stdout_logfile_maxbytes = 500MB
stdout_logfile_backups = 50

stderr_logfile  = /var/log/perf/high-perf.err
stderr_logfile_maxbytes = 500MB
stderr_logfile_backups = 50

stopwaitsecs    = 60
priority        = 1
`

	require.Equal(t, expected, content)
}

func TestDevelopmentConfig(t *testing.T) {
	// Test development environment configuration
	// 测试开发环境配置
	program := supervisorkratos.NewProgramConfig(
		"dev-service",
		"/home/dev/service",
		"developer",
		"/tmp/dev-logs",
	).WithAutoStart(false).
		WithAutoRestart(false).
		WithStartRetries(1).
		WithLogMaxBytes("10MB").
		WithLogBackups(3).
		WithRedirectStderr(true).
		WithStopAsGroup(false).
		WithEnvironment(map[string]string{
			"NODE_ENV": "development",
		})

	content := supervisorkratos.GenerateProgramConfig(program)
	t.Log("=== Development environment configuration ===")
	t.Log(content)

	const expected = `[program:dev-service]
user            = developer
directory       = /home/dev/service
command         = /home/dev/service/bin/dev-service
environment     = NODE_ENV=development

autostart       = false
autorestart     = false
startretries    = 1

stdout_logfile  = /tmp/dev-logs/dev-service.log
stdout_logfile_maxbytes = 10MB
stdout_logfile_backups = 3

stderr_logfile  = /tmp/dev-logs/dev-service.err
stderr_logfile_maxbytes = 10MB
stderr_logfile_backups = 3
redirect_stderr = true

`

	require.Equal(t, expected, content)
}

func TestCustomExitCodesConfig(t *testing.T) {
	// Test custom exit codes configuration
	// 测试自定义退出码配置
	program := supervisorkratos.NewProgramConfig(
		"exit-service",
		"/opt/exit-service",
		"exit-user",
		"/var/log/exit",
	).WithExitCodes([]int{0, 1, 2, 130}).
		WithStopSignal("QUIT").
		WithKillAsGroup(false)

	content := supervisorkratos.GenerateProgramConfig(program)
	t.Log("=== Custom exit codes configuration ===")
	t.Log(content)

	const expected = `[program:exit-service]
user            = exit-user
directory       = /opt/exit-service
command         = /opt/exit-service/bin/exit-service

stdout_logfile  = /var/log/exit/exit-service.log

stderr_logfile  = /var/log/exit/exit-service.err

stopsignal      = QUIT
exitcodes       = 0,1,2,130
`

	require.Equal(t, expected, content)
}

func TestLargeScaleGroupConfig(t *testing.T) {
	// Test large-scale group configuration
	// 测试大规模组配置
	group := supervisorkratos.NewGroupConfig("mega-cluster")

	// Create multiple programs with different configurations
	// 创建多个不同配置的程序
	for i := 1; i <= 3; i++ {
		name := "service" + string(rune('0'+i))
		program := supervisorkratos.NewProgramConfig(
			name,
			"/opt/"+name,
			"cluster-user",
			"/var/log/cluster",
		).WithPriority(50).
			WithNumProcs(2).
			WithProcessName("%(program_name)s-%(process_num)02d").
			WithEnvironment(map[string]string{
				"CLUSTER_MODE": "production",
			})

		group.AddProgram(program)
	}

	content := supervisorkratos.GenerateGroupConfig(group)
	t.Log("=== Large-scale group configuration ===")
	t.Log(content)

	const expected = `[group:mega-cluster]
programs=service1,service2,service3


[program:service1]
user            = cluster-user
directory       = /opt/service1
command         = /opt/service1/bin/service1
environment     = CLUSTER_MODE=production

stdout_logfile  = /var/log/cluster/service1.log

stderr_logfile  = /var/log/cluster/service1.err

priority        = 50
numprocs        = 2
process_name    = %(program_name)s-%(process_num)02d

[program:service2]
user            = cluster-user
directory       = /opt/service2
command         = /opt/service2/bin/service2
environment     = CLUSTER_MODE=production

stdout_logfile  = /var/log/cluster/service2.log

stderr_logfile  = /var/log/cluster/service2.err

priority        = 50
numprocs        = 2
process_name    = %(program_name)s-%(process_num)02d

[program:service3]
user            = cluster-user
directory       = /opt/service3
command         = /opt/service3/bin/service3
environment     = CLUSTER_MODE=production

stdout_logfile  = /var/log/cluster/service3.log

stderr_logfile  = /var/log/cluster/service3.err

priority        = 50
numprocs        = 2
process_name    = %(program_name)s-%(process_num)02d
`

	require.Equal(t, expected, content)
}

func TestDefaultValues(t *testing.T) {
	// Test minimal configuration with only defaults (from old version)
	// 测试仅使用默认值的最小配置（来自旧版本）
	program := supervisorkratos.NewProgramConfig(
		"basic-service",
		"/opt/basic-service",
		"deploy",
		"/var/log/basic",
	)

	content := supervisorkratos.GenerateProgramConfig(program)
	t.Log("=== Required parameters with defaults ===")
	t.Log(content)

	// Use exact format from old version git diff
	// 使用旧版本 git diff 中的确切格式
	const expected = `[program:basic-service]
user            = deploy
directory       = /opt/basic-service
command         = /opt/basic-service/bin/basic-service

stdout_logfile  = /var/log/basic/basic-service.log

stderr_logfile  = /var/log/basic/basic-service.err

`

	require.Equal(t, expected, content)
}

func TestMicroserviceCluster(t *testing.T) {
	// Test microservice cluster with different service types
	// 测试微服务集群，包含不同类型的服务
	gateway := supervisorkratos.NewProgramConfig(
		"api-gateway",
		"/opt/gateway",
		"deploy",
		"/var/log/cluster",
	).WithPriority(1).
		WithNumProcs(2).
		WithProcessName("%(program_name)s-%(process_num)02d").
		WithEnvironment(map[string]string{
			"SERVICE_TYPE": "gateway",
		})

	userService := supervisorkratos.NewProgramConfig(
		"user-service",
		"/opt/user-service",
		"deploy",
		"/var/log/cluster",
	).WithStartRetries(5).
		WithStopWaitSecs(30)

	orderService := supervisorkratos.NewProgramConfig(
		"order-service",
		"/opt/order-service",
		"deploy",
		"/var/log/cluster",
	).WithAutoRestart(false).
		WithLogMaxBytes("200MB")

	cluster := supervisorkratos.NewGroupConfig("microservice-cluster").
		AddProgram(gateway).
		AddProgram(userService).
		AddProgram(orderService)

	content := supervisorkratos.GenerateGroupConfig(cluster)
	t.Log("=== Microservice cluster configuration ===")
	t.Log(content)

	const expected = `[group:microservice-cluster]
programs=api-gateway,user-service,order-service


[program:api-gateway]
user            = deploy
directory       = /opt/gateway
command         = /opt/gateway/bin/api-gateway
environment     = SERVICE_TYPE=gateway

stdout_logfile  = /var/log/cluster/api-gateway.log

stderr_logfile  = /var/log/cluster/api-gateway.err

priority        = 1
numprocs        = 2
process_name    = %(program_name)s-%(process_num)02d

[program:user-service]
user            = deploy
directory       = /opt/user-service
command         = /opt/user-service/bin/user-service

startretries    = 5

stdout_logfile  = /var/log/cluster/user-service.log

stderr_logfile  = /var/log/cluster/user-service.err

stopwaitsecs    = 30

[program:order-service]
user            = deploy
directory       = /opt/order-service
command         = /opt/order-service/bin/order-service

autorestart     = false

stdout_logfile  = /var/log/cluster/order-service.log
stdout_logfile_maxbytes = 200MB

stderr_logfile  = /var/log/cluster/order-service.err
stderr_logfile_maxbytes = 200MB
`

	require.Equal(t, expected, content)
}

func TestZeroConfigProgram(t *testing.T) {
	// Test program with absolutely minimal config using pure defaults
	// 测试使用纯默认值的最小配置程序
	program := supervisorkratos.NewProgramConfig(
		"zero-config",
		"/opt/zero-config",
		"deploy",
		"/var/log/minimal",
	)

	content := supervisorkratos.GenerateProgramConfig(program)
	t.Log("=== Zero customization program configuration ===")
	t.Log(content)

	const expected = `[program:zero-config]
user            = deploy
directory       = /opt/zero-config
command         = /opt/zero-config/bin/zero-config

stdout_logfile  = /var/log/minimal/zero-config.log

stderr_logfile  = /var/log/minimal/zero-config.err

`

	require.Equal(t, expected, content)
}
