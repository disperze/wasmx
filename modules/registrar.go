package modules

import (
	"github.com/desmos-labs/juno/modules/registrar"

	junomod "github.com/desmos-labs/juno/modules"
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

	return []junomod.Module{}
}
