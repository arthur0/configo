package loader_test

import (
	"testing"

	"github.com/arthur0/configo/loader"
)
func TestLoadTOML(t *testing.T) {
	loader := loader.TOMLloader{}
	Load(t, "../../data/settings.toml", &loader)
}
