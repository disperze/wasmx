package main

import (
	"log"

	"github.com/disperze/wasmx/database"
	"github.com/disperze/wasmx/modules"
	"github.com/disperze/wasmx/types/config"
	junocmd "github.com/forbole/juno/v2/cmd"
	parsecmd "github.com/forbole/juno/v2/cmd/parse"
)

func main() {
	// Setup the config
	parseCfg := parsecmd.NewConfig().
		WithRegistrar(modules.NewModulesRegistrar()).
		WithEncodingConfigBuilder(config.MakeEncodingConfig).
		WithDBBuilder(database.Builder)

	cfg := junocmd.NewConfig("wasmx").
		WithParseConfig(parseCfg)

	// Run the commands and panic on any error
	executor := junocmd.BuildDefaultExecutor(cfg)
	err := executor.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
