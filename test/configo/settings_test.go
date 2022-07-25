package configo_test

import (
	"os"
	"testing"

	"github.com/arthur0/configo"
)

var settings = configo.Settings{}


func TestSettings(t *testing.T) {
	settings.Load("../data/settings.toml", "default")
	if settings.Get("empty") != "" {
		t.Error("Error to get str EMPTY on default env")
	}
	if settings.Get("server") != "sandbox.configo.com" {
		t.Error("Error to get str SERVER on default env")
	}
	if settings.Int("port") != 3000 {
		t.Error("Error to get int PORT on default env")
	}
	os.Setenv("PORT", "5000")
	defer os.Unsetenv("PORT")
	if settings.Int("port") != 5000 {
		t.Error("Error to get int PORT on os env")
	}
	envs := []string{"local", "dev", "dev.sandbox", "prod"} 
	if len(settings.Strings("envs")) != len(envs) {
		t.Error("Error to get string slice ENVS default env")
	}
}
