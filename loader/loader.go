package loader

import "os"

type ConfigLoader interface {
	Load(file *os.File) map[string]map[string]string
}

var AllowedExtensions []string = []string{".toml"}
