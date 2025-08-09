package supervisorkratos

import (
	"path/filepath"
	"strings"

	"github.com/yyle88/must"
	"github.com/yyle88/printgo"
)

type Config struct {
	Group    string   // 程序组名称
	Roots    []string // 程序根目录列表
	UserName string   // 运行程序的用户名称
	SlogRoot string   // 标准输出日志根目录
}

func GenerateSupervisorConfig(cfg *Config) string {
	must.Full(cfg)
	must.Nice(cfg.Group)
	must.Have(cfg.Roots)
	must.Nice(cfg.UserName)
	must.Nice(cfg.SlogRoot)

	ptx := printgo.NewPTX()
	ptx.Println(`[group:` + cfg.Group + `]`)
	ptx.Println(`programs=` + strings.Join(programsNames(cfg.Roots), ","))
	ptx.Println()

	for _, root := range cfg.Roots {
		name := filepath.Base(root)

		ptx.Println(`
[program:` + name + `]
user            = ` + cfg.UserName + `
directory       = ` + root + `
command         = ` + filepath.Join(root, "bin", name) + `

autostart       = true
autorestart     = true
startretries    = 30
startsecs       = 1

stdout_logfile  = ` + filepath.Join(cfg.SlogRoot, name+`.log`) + `
stdout_logfile_maxbytes = 30MB
stdout_logfile_backups = 10

stderr_logfile  = ` + filepath.Join(cfg.SlogRoot, name+`.err`) + `
stderr_logfile_maxbytes = 30MB
stderr_logfile_backups = 10
redirect_stderr = false

stopasgroup     = true
`)
	}

	return ptx.String()
}

func programsNames(roots []string) []string {
	var names = make([]string, 0, len(roots))
	for _, root := range roots {
		names = append(names, filepath.Base(root))
	}
	return names
}
