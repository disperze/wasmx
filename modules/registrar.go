package modules

import (
	"github.com/disperze/wasmx/database"
	"github.com/disperze/wasmx/modules/wasm"
	"github.com/forbole/juno/v2/modules/registrar"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	junomod "github.com/forbole/juno/v2/modules"
	junoremote "github.com/forbole/juno/v2/node/remote"
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
	grpcCfg := ctx.JunoConfig.Node.Details.(*junoremote.Details)

	grpcConnection, err := junoremote.CreateGrpcConnection(grpcCfg.GRPC)
	if err != nil {
		panic(err)
	}

	client := wasmtypes.NewQueryClient(grpcConnection)
	wasmDb := database.Cast(ctx.Database)

	return []junomod.Module{
		wasm.NewModule(wasmDb, client),
	}
}
