package loader

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

const envRegex = `^\[.*\]$`

type TOMLloader struct{}

func (l *TOMLloader) Load(file *os.File) map[string]map[string]string {
	var configs = map[string]map[string]string{}
	scanner := bufio.NewScanner(file)

	re := regexp.MustCompile(envRegex)
	var currentEnv string
	for scanner.Scan() {
		line := scanner.Text()
		env := re.FindString(line)
		env = cleanEnv(env)
		if env != "" {
			currentEnv = env
			configs[currentEnv] = map[string]string{}
		}
		key, value := findKeyPair(line)
		if key != "" {
			configs[currentEnv][key] = value
		}
	}

	return configs
}

func findKeyPair(line string) (string, string) {
	keyPair := strings.SplitN(line, "=", 2)
	if len(keyPair) == 2 {
		return cleanKey(keyPair[0]), cleanValue(keyPair[1])
	}
	return "", ""
}

func cleanEnv(env string) string {
	env = strings.TrimSpace(env)
	env = strings.TrimLeft(env, "[")
	env = strings.TrimRight(env, "]")
	return env
}

func cleanKey(key string) string {
	key = strings.ToUpper(strings.TrimSpace(key))
	return key
}

func cleanValue(value string) string {
	value = strings.TrimSpace(value)
	value = strings.TrimLeft(value, `"`)
	value = strings.TrimLeft(value, `'`)
	value = strings.TrimRight(value, `"`)
	value = strings.TrimRight(value, `'`)
	return value

}
