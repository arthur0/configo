package loader_test

import (
	"log"
	"os"
	"testing"

	"github.com/arthur0/configo/loader"
)

var defaultConfig = map[string]string{
	"EMPTY":   "",
	"SERVER":  "sandbox.configo.com",
	"DEBUG":   "true",
	"PORT":    "3000",
	"VERSION": "0.1",
	"ENVS":    `["local", "dev", "dev.sandbox", "prod"]`,
	"INDEX":   "{}",
	"YOLO":    `You only live once\n`,
}

var devConfig = map[string]string{
	"SERVER": "dev.configo.com",
	"INDEX":  `{"env": "dev", "msg": "Hi from dev"}`,
}

var prdConfig = map[string]string{
	"SERVER":    "configo.com",
	"DEBUG":     "false",
	"PORT":      "80",
	"INDEX":     `{"env": "prd", "msg": "Hello from Production"}`,
	"ONLY_PROD": "true",
}


func Load(t *testing.T, path string, l loader.ConfigLoader) {
	expectedConf := map[string]map[string]string{
		"default": defaultConfig,
		"dev":     devConfig,
		"prd":     prdConfig,
	}
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	const errMsg = `
Env %s Key %s
Expected value: %s
Got value     : %s 
`
	got := l.Load(file)
	for env := range expectedConf {
		for key, value := range expectedConf[env] {
			if value != got[env][key] {
				t.Errorf(errMsg, env, key, value, got[env][key])
			}
		}
	}
}
