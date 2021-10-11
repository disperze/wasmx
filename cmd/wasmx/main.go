package main

import (
	"log"

	junocmd "github.com/desmos-labs/juno/cmd"
	parsecmd "github.com/desmos-labs/juno/cmd/parse"
	"github.com/disperze/wasmx/database"
	"github.com/disperze/wasmx/modules"
	"github.com/disperze/wasmx/types/config"
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
