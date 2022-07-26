package configo_test

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/arthur0/configo"
)

var settings = configo.Settings{}

const errMsg = "Error to get %v in %v env. \nExpected: %v\nGot     : %v"

func TestDefaultSettings(t *testing.T) {
	settings.Load("../data/settings.toml", "default")
	var expected, got interface{}

	expected = ""
	got = settings.Get("empty")
	if expected != got {
		t.Errorf(errMsg, "EMPTY", "default", expected, got)
	}

	expected = "sandbox.configo.com"
	got = settings.Get("server")
	if expected != got {
		t.Errorf(errMsg, "SERVER", "default", expected, got)
	}

	expected = true
	got = settings.Bool("debug")
	if expected != got {
		t.Errorf(errMsg, "DEBUG", "default", expected, got)
	}

	expected = 3000
	got = settings.Int("PORT")
	if expected != got {
		t.Errorf(errMsg, "PORT", "default", expected, got)
	}

	expected = 0.1
	got = settings.Float32("VERSION")
	if fmt.Sprintf("%.1f", expected) != fmt.Sprintf("%.1f", got) {
		t.Errorf(errMsg, "VERSION", "default", expected, got)
	}

	expected = 0.1
	got = settings.Float64("VERSION")
	if fmt.Sprintf("%.1f", expected) != fmt.Sprintf("%.1f", got) {
		t.Errorf(errMsg, "VERSION", "default", expected, got)
	}

	expectedSlice := []string{"local", "dev", "dev.sandbox", "prod"}
	gotSlice := settings.Strings("envs")
	if Equal(expectedSlice, gotSlice) {
		t.Errorf(errMsg, "ENVS", "default", expected, got)
	}

	expected = map[string]string{}
	got = settings.Map("index")
	if fmt.Sprintf("%v", expected) != fmt.Sprintf("%v", got) {
		t.Errorf(errMsg, "INDEX", "default", expected, got)
	}

	settings.Load("../data/settings.toml", "prd")

	expected = "YOLO means You only live once"
	got = settings.Get("yolo")
	if expected != got {
		t.Errorf(errMsg, "YOLO", "prd", expected, got)
	}

	expected = "configo.com"
	got = settings.Get("server")
	if expected != got {
		t.Errorf(errMsg, "SERVER", "prd", expected, got)
	}

	expected = true
	got = settings.Bool("only_prod")
	if expected != got {
		t.Errorf(errMsg, "ONLY_PROD", "prd", expected, got)
	}

	expected = map[string]string{"env": "prd", "msg": "Hello from Production"}
	got = settings.Map("index")
	if fmt.Sprintf("%v", expected) != fmt.Sprintf("%v", got) {
		t.Errorf(errMsg, "INDEX", "default", expected, got)
	}
}

func Equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestUnsupportedFormatErr(t *testing.T) {
	err := settings.Load("../data/settings.yml", "default")
	if err == nil {
		t.Error("Expect unsupported format error")
	}
}

func TestNonexistentFileErr(t *testing.T) {
	err := settings.Load("../data/nonexistent.toml", "default")
	if err == nil {
		t.Error("Expect IO error")
	}
}

func TestKeyNotFoundErr(t *testing.T) {
	settings.Load("../data/settings.toml", "default")
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()
	settings.Get("INEXISTENT_KEY")
	output := buf.String()
	if output == "" {
		t.Error("Expect inexistent key error log")
	}
	t.Log(output)
}

func TestParseErr(t *testing.T) {
	settings.Load("../data/settings.toml", "default")
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()
	settings.Bool("YOLO")
	output := buf.String()
	if output == "" {
		t.Error("Expect parse value error log")
	}
	t.Log(output)
}
