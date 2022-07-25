package configo

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/arthur0/configo/loader"
)

const (
	ErrUnsupportedExtensionMsg = "Currently we only support %s extensions"
	ErrKeyNotFoundMsg          = "The key %s was not found in the settings"
	ErrParseValueMsg           = "The key %s has value %s and cannot be parsed to %s"
	ErrIOMsg                   = "Unable to open settings file\n %w"
)

type Settings struct {
	config map[string]string
}

func (s *Settings) Load(path string, env string) error {
	var configLoader loader.ConfigLoader
	var settingsErr error
	extension := filepath.Ext(path)
	switch extension {
	case ".toml":
		configLoader = &loader.TOMLloader{}
	default:
		settingsErr = fmt.Errorf(ErrUnsupportedExtensionMsg, loader.AllowedExtensions)
	}

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	loaderConfig := configLoader.Load(file)
	s.config = loaderConfig["default"]
	for key, value := range loaderConfig[env] {
		s.config[key] = value
	}
	return settingsErr
}

func (s *Settings) Get(key string) string {
	fromEnv, found := os.LookupEnv(strings.ToUpper(key))
	if found {
		return fromEnv

	}
	if value, found := s.config[strings.ToUpper(key)]; found {
		return value
	}
	log.Fatalf(ErrKeyNotFoundMsg, key)
	return ""
}

func (s *Settings) Int(key string) int {
	value := s.Get(key)
	intValue, parseValueErr := strconv.Atoi(value)
	logParseError(parseValueErr, key, value, "int")
	return intValue
}

func (s *Settings) Bool(key string) bool {
	value := s.Get(key)
	boolValue, parseValueErr := strconv.ParseBool(value)
	logParseError(parseValueErr, key, value, "bool")
	return boolValue
}

func (s *Settings) Float32(key string) float32 {
	value := s.Get(key)
	floatValue, parseValueErr := strconv.ParseFloat(value, 32)
	logParseError(parseValueErr, key, value, "float32")
	return float32(floatValue)
}

func (s *Settings) Float64(key string) float64 {
	value := s.Get(key)
	floatValue, parseValueErr := strconv.ParseFloat(value, 64)
	logParseError(parseValueErr, key, value, "float64")
	return floatValue
}

func (s *Settings) Strings(key string) []string {
	value := s.Get(key)
	withoutBrackets := value[1 : len(value)-1]
	return strings.Split(withoutBrackets, ",")
}

func (s *Settings) Map(key string) map[string]string {
	value := s.Get(key)
	var result map[string]string
	parseValueErr := json.Unmarshal([]byte(value), &result)
	logParseError(parseValueErr, key, value, "map[string]string")
	return result
}

func logParseError(parseValueErr error, key string, value string, expectedType string) {
	if parseValueErr != nil {
		log.Printf(ErrParseValueMsg, key, value, expectedType)
	}
}
