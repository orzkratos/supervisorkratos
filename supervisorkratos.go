package supervisorkratos

import (
	"path/filepath"
	"slices"
	"strconv"
	"strings"

	"github.com/yyle88/must"
	"github.com/yyle88/printgo"
)

// ProgramConfig single program configuration
// 单个程序配置
type ProgramConfig struct {
	// Basic program information // 基本程序信息
	Name     string // Program name // 程序名称
	UserName string // User to run programs // 运行程序的用户名称
	Root     string // Program root DIR // 程序根目录
	SlogRoot string // Standard output log root DIR // 标准输出日志根目录

	// Environment variables // 环境变量
	Environment map[string]string // Environment variables // 环境变量

	// Process control settings // 进程控制设置
	AutoStart    bool // Auto start on supervisor startup // supervisor 启动时自动启动
	AutoRestart  bool // Auto restart on failure // 失败时自动重启
	StartRetries int  // Max start retry times // 最大启动重试次数
	StartSecs    int  // Seconds to wait before considering start successful // 启动成功前等待秒数

	// Log settings // 日志设置
	LogMaxBytes    string // Max log file size // 最大日志文件大小
	LogBackups     int    // Log backup files count // 日志备份文件数量
	RedirectStderr bool   // Redirect stderr to stdout // 重定向 stderr 到 stdout

	// Advanced process control // 高级进程控制
	StopAsGroup  bool   // Stop all processes as group // 作为组停止所有进程
	StopWaitSecs int    // Graceful stop timeout seconds // 优雅停止超时秒数
	KillAsGroup  bool   // Kill child processes as group // 强制杀死子进程组
	StopSignal   string // Stop signal name // 停止信号名称
	Priority     int    // Start priority // 启动优先级
	ExitCodes    []int  // Expected exit codes // 预期退出码

	// Multi-instance settings // 多实例设置
	NumProcs    int    // Number of process instances // 进程实例数量
	ProcessName string // Process name template // 进程名称模板
}

// GroupConfig supervisor group configuration
// supervisor 组配置
type GroupConfig struct {
	Name     string           // Group name // 组名称
	Programs []*ProgramConfig // Program configs // 程序配置列表
}

// NewProgramConfig create new ProgramConfig with required fields
// Name, Root, UserName, SlogRoot are required parameters
//
// 创建新的 ProgramConfig，需要提供必填字段
// Name、Root、UserName、SlogRoot 是必填参数
func NewProgramConfig(name string, root string, userName string, slogRoot string) *ProgramConfig {
	return &ProgramConfig{
		// Basic program information // 基本程序信息
		Name:     must.Nice(name),
		UserName: must.Nice(userName),
		Root:     must.Nice(root),
		SlogRoot: must.Nice(slogRoot),

		// Environment variables // 环境变量
		Environment: make(map[string]string),

		// Set supervisor official default values
		// 设置 supervisor 官方默认值

		// Process control settings // 进程控制设置
		AutoStart:    true,
		AutoRestart:  true, // "unexpected" but we use bool, true is closest
		StartRetries: 3,
		StartSecs:    1,

		// Log settings // 日志设置
		LogMaxBytes:    "50MB",
		LogBackups:     10,
		RedirectStderr: false,

		// Advanced process control defaults
		// 高级进程控制默认值
		StopAsGroup:  false,
		StopWaitSecs: 10,
		KillAsGroup:  false,
		StopSignal:   "TERM",
		Priority:     999,
		ExitCodes:    []int{0},

		// Multi-instance defaults
		// 多实例默认值
		NumProcs:    1,
		ProcessName: "%(program_name)s",
	}
}

// NewGroupConfig create new GroupConfig
// 创建新的 GroupConfig
func NewGroupConfig(name string) *GroupConfig {
	return &GroupConfig{
		Name:     must.Nice(name),
		Programs: make([]*ProgramConfig, 0),
	}
}

// AddProgram add program to group
// 添加程序到组
func (g *GroupConfig) AddProgram(program *ProgramConfig) *GroupConfig {
	g.Programs = append(g.Programs, program)
	return g
}

// ProgramConfig chain methods for configuration customization
// ProgramConfig 链式配置方法

// WithAutoStart set auto start flag
// 设置自动启动标志
func (p *ProgramConfig) WithAutoStart(autoStart bool) *ProgramConfig {
	p.AutoStart = autoStart
	return p
}

// WithAutoRestart set auto restart flag
// 设置自动重启标志
func (p *ProgramConfig) WithAutoRestart(autoRestart bool) *ProgramConfig {
	p.AutoRestart = autoRestart
	return p
}

// WithStartRetries set start retries count
// 设置启动重试次数
func (p *ProgramConfig) WithStartRetries(startRetries int) *ProgramConfig {
	p.StartRetries = startRetries
	return p
}

// WithStartSecs set start seconds
// 设置启动成功等待时间
func (p *ProgramConfig) WithStartSecs(startSecs int) *ProgramConfig {
	p.StartSecs = startSecs
	return p
}

// WithLogMaxBytes set log file max bytes
// 设置日志文件最大字节数
func (p *ProgramConfig) WithLogMaxBytes(logMaxBytes string) *ProgramConfig {
	p.LogMaxBytes = logMaxBytes
	return p
}

// WithLogBackups set log backup count
// 设置日志备份文件数量
func (p *ProgramConfig) WithLogBackups(logBackups int) *ProgramConfig {
	p.LogBackups = logBackups
	return p
}

// WithRedirectStderr set redirect stderr flag
// 设置重定向标准错误标志
func (p *ProgramConfig) WithRedirectStderr(redirectStderr bool) *ProgramConfig {
	p.RedirectStderr = redirectStderr
	return p
}

// WithStopAsGroup set stop as group flag
// 设置作为组停止标志
func (p *ProgramConfig) WithStopAsGroup(stopAsGroup bool) *ProgramConfig {
	p.StopAsGroup = stopAsGroup
	return p
}

// WithKillAsGroup set kill as group flag
// 设置作为组终止标志
func (p *ProgramConfig) WithKillAsGroup(killAsGroup bool) *ProgramConfig {
	p.KillAsGroup = killAsGroup
	return p
}

// WithStopWaitSecs set stop wait seconds
// 设置停止等待时间
func (p *ProgramConfig) WithStopWaitSecs(stopWaitSecs int) *ProgramConfig {
	p.StopWaitSecs = stopWaitSecs
	return p
}

// WithStopSignal set stop signal
// 设置停止信号
func (p *ProgramConfig) WithStopSignal(stopSignal string) *ProgramConfig {
	p.StopSignal = stopSignal
	return p
}

// WithPriority set process priority
// 设置进程优先级
func (p *ProgramConfig) WithPriority(priority int) *ProgramConfig {
	p.Priority = priority
	return p
}

// WithEnvironment set environment variables
// 设置环境变量
func (p *ProgramConfig) WithEnvironment(environment map[string]string) *ProgramConfig {
	p.Environment = environment
	return p
}

// WithExitCodes set expected exit codes
// 设置期望的退出码
func (p *ProgramConfig) WithExitCodes(exitCodes []int) *ProgramConfig {
	p.ExitCodes = exitCodes
	return p
}

// WithNumProcs set number of processes
// 设置进程数量
func (p *ProgramConfig) WithNumProcs(numProcs int) *ProgramConfig {
	p.NumProcs = numProcs
	return p
}

// WithProcessName set process name pattern
// 设置进程名称模式
func (p *ProgramConfig) WithProcessName(processName string) *ProgramConfig {
	p.ProcessName = processName
	return p
}

// GenerateGroupConfig generate supervisor group configuration
// 生成 supervisor 组配置
func GenerateGroupConfig(group *GroupConfig) string {
	must.Full(group)
	must.Nice(group.Name)
	must.Have(group.Programs)

	ptx := printgo.NewPTX()

	// Generate group header
	// 生成组头部
	ptx.Println(`[group:` + group.Name + `]`)
	programs := make([]string, 0, len(group.Programs))
	for _, p := range group.Programs {
		programs = append(programs, p.Name)
	}
	ptx.Println(`programs=` + strings.Join(programs, ","))
	ptx.Println()

	// Generate each program config
	// 生成每个程序配置
	for _, program := range group.Programs {
		ptx.Println()
		cfs := GenerateProgramConfig(program)
		ptx.Println(strings.TrimSpace(cfs))
	}

	return ptx.String()
}

// GenerateProgramConfig generate single program configuration from ProgramConfig
// 从 ProgramConfig 生成单个程序配置
func GenerateProgramConfig(program *ProgramConfig) string {
	must.Full(program)
	must.Nice(program.Name)
	must.Nice(program.Root)
	must.Nice(program.UserName)
	must.Nice(program.SlogRoot)

	ptx := printgo.NewPTX()

	ptx.Println("[program:" + program.Name + "]")
	ptx.Println("user            = " + program.UserName)
	ptx.Println("directory       = " + program.Root)
	ptx.Println("command         = " + filepath.Join(program.Root, "bin", program.Name))

	if env := combineSsMap(program.Environment, ","); env != "" {
		ptx.Println("environment     = " + env)
	}
	ptx.Println()
	mark := ptx.Len()

	defaultConfig := NewProgramConfig(program.Root, program.Name, program.UserName, program.SlogRoot)

	// Only print non-default values (hardcoded defaults)
	// 只打印非默认值的配置（硬编码默认值）
	if program.AutoStart != defaultConfig.AutoStart {
		ptx.Println("autostart       = " + strconv.FormatBool(program.AutoStart))
	}
	if program.AutoRestart != defaultConfig.AutoRestart {
		ptx.Println("autorestart     = " + strconv.FormatBool(program.AutoRestart))
	}
	if program.StartRetries != defaultConfig.StartRetries {
		ptx.Println("startretries    = " + strconv.Itoa(program.StartRetries))
	}
	if program.StartSecs != defaultConfig.StartSecs {
		ptx.Println("startsecs       = " + strconv.Itoa(program.StartSecs))
	}

	if ptx.Len() > mark {
		ptx.Println()
	}
	mark = ptx.Len()

	// Log settings always show (required for paths)
	// 日志设置始终显示（路径必需）
	ptx.Println("stdout_logfile  = " + filepath.Join(program.SlogRoot, program.Name+".log"))
	if program.LogMaxBytes != defaultConfig.LogMaxBytes {
		ptx.Println("stdout_logfile_maxbytes = " + program.LogMaxBytes)
	}
	if program.LogBackups != defaultConfig.LogBackups {
		ptx.Println("stdout_logfile_backups = " + strconv.Itoa(program.LogBackups))
	}

	if ptx.Len() > mark {
		ptx.Println()
	}
	mark = ptx.Len()

	ptx.Println("stderr_logfile  = " + filepath.Join(program.SlogRoot, program.Name+".err"))
	if program.LogMaxBytes != defaultConfig.LogMaxBytes {
		ptx.Println("stderr_logfile_maxbytes = " + program.LogMaxBytes)
	}
	if program.LogBackups != defaultConfig.LogBackups {
		ptx.Println("stderr_logfile_backups = " + strconv.Itoa(program.LogBackups))
	}
	if program.RedirectStderr != defaultConfig.RedirectStderr {
		ptx.Println("redirect_stderr = " + strconv.FormatBool(program.RedirectStderr))
	}

	if ptx.Len() > mark {
		ptx.Println()
	}
	mark = ptx.Len() //nolint:ineffassign,staticcheck // Keep for future sections

	// Advanced process control - only non-defaults
	// 高级进程控制 - 只显示非默认值
	if program.StopAsGroup != defaultConfig.StopAsGroup {
		ptx.Println("stopasgroup     = " + strconv.FormatBool(program.StopAsGroup))
	}
	if program.StopWaitSecs != defaultConfig.StopWaitSecs {
		ptx.Println("stopwaitsecs    = " + strconv.Itoa(program.StopWaitSecs))
	}
	if program.KillAsGroup != defaultConfig.KillAsGroup {
		ptx.Println("killasgroup     = " + strconv.FormatBool(program.KillAsGroup))
	}
	if program.StopSignal != defaultConfig.StopSignal {
		ptx.Println("stopsignal      = " + program.StopSignal)
	}
	if program.Priority != defaultConfig.Priority {
		ptx.Println("priority        = " + strconv.Itoa(program.Priority))
	}
	if !slices.Equal(program.ExitCodes, defaultConfig.ExitCodes) {
		ptx.Println("exitcodes       = " + combineInts(program.ExitCodes, ","))
	}
	if program.NumProcs != defaultConfig.NumProcs {
		ptx.Println("numprocs        = " + strconv.Itoa(program.NumProcs))
	}
	if program.ProcessName != defaultConfig.ProcessName {
		ptx.Println("process_name    = " + program.ProcessName)
	}

	return ptx.String()
}

func combineInts(items []int, sep string) string {
	if len(items) == 0 {
		return ""
	}
	strs := make([]string, len(items))
	for i, item := range items {
		strs[i] = strconv.Itoa(item)
	}
	return strings.Join(strs, sep)
}

func combineSsMap(items map[string]string, sep string) string {
	if len(items) == 0 {
		return ""
	}
	pairs := make([]string, 0, len(items))
	for key, value := range items {
		pairs = append(pairs, key+"="+value)
	}
	return strings.Join(pairs, sep)
}
