package modules

import (
	"fmt"

	"github.com/disperze/wasmx/database"
	"github.com/disperze/wasmx/modules/wasm"
	"github.com/forbole/juno/v3/modules/registrar"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	junomod "github.com/forbole/juno/v3/modules"
	junoremote "github.com/forbole/juno/v3/node/remote"
)

// ModulesRegistrar represents the modules.Registrar that allows to register all custom modules
type ModulesRegistrar struct {
}

// NewModulesRegistrar allows to build a new ModulesRegistrar instance
func NewModulesRegistrar() *ModulesRegistrar {
	return &ModulesRegistrar{}
}

// BuildModules implements modules.Registrar
func (r *ModulesRegistrar) BuildModules(ctx registrar.Context) junomod.Modules {
	remoteCfg, ok := ctx.JunoConfig.Node.Details.(*junoremote.Details)
	if !ok {
		panic(fmt.Errorf("invalid remote grpc config"))
	}

	grpcConnection, err := junoremote.CreateGrpcConnection(remoteCfg.GRPC)
	if err != nil {
		panic(err)
	}

	client := wasmtypes.NewQueryClient(grpcConnection)
	wasmDb := database.Cast(ctx.Database)

	return []junomod.Module{
		wasm.NewModule(wasmDb, client),
	}
}
