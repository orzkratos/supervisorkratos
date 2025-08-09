package supervisorkratos_test

import (
	"testing"

	"github.com/orzkratos/supervisorkratos"
)

func TestGenerateSupervisorConfig(t *testing.T) {
	content := supervisorkratos.GenerateSupervisorConfig(&supervisorkratos.Config{
		Group: "demokratos",
		Roots: []string{
			"/Users/admin/go-projects/orzkratos/demokratos/demo1kratos",
			"/Users/admin/go-projects/orzkratos/demokratos/demo2kratos",
		},
		UserName: "admin",
		SlogRoot: "/data/logs/demokratos/",
	})
	t.Log(content)
}
