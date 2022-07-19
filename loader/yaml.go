package loader

import (
	"os"
)

type YAMLloader struct{}

func (l *YAMLloader) Load(file *os.File) map[string]map[string]string {
	panic("Yaml loader was not implemented yet")
}
