# ConfiGO: Configuration Management for Go

This lib is intended to make easier and readable the configuration management for Go projects

The examples bellow are made using the [settings.toml](test/data/settings.toml) file

## Getting a default config

```go
package main

import (
	"fmt"

	"github.com/arthur0/configo"
)
	
func main() {
	cfg := Settings{}
	cfg.Load("test/data/settings.toml", "default")
	yolo := cfg.Get("YOLO")
	fmt.Println(yolo)
}
// output: You only live once\n
```

## Getting a string slice from the prd env

The `envs` key is defined on `[default]` section, that contains the common configs for all envs


```go
package main

import (
	"fmt"

	"github.com/arthur0/configo"
)
	
func main() {
	cfg := Settings{}
	cfg.Load("test/data/settings.toml", "prd")
	envs := cfg.Strings("envs")
	fmt.Println(envs)
}
// output: ["local"  "dev"  "dev.sandbox"  "prod"]
```

# Complete Example

```go
func main() {
	cfg := Settings{}
	cfg.Load("test/data/settings.toml", "prd")
	server := cfg.Get("server")
	port := cfg.Int("port")
	debug := cfg.Bool("debug")
	version := cfg.Float32("version")
	index_json := cfg.Map("index")
	msg := index_json["msg"]

	fmt.Printf(`
	server running at %s:%d
	version:%.1f debug mode: %t

	%s`, server, port, version, debug, msg)
	
}
	// output
	// server running at configo.com:80
	// version:0.1 debug mode: false

	// Hello from Production
```
