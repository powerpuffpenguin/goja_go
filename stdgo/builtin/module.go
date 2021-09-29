package builtin

import (
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
	"github.com/powerpuffpenguin/goja_go/core/utils"
)

func init() {
	require.RegisterNativeModule(ModuleID, Require)
}

const (
	ModuleID = `stdgo/builtin`
)

func Require(runtime *goja.Runtime, module *goja.Object) {
	exports := module.Get(`exports`).(*goja.Object)
	factory := newFactory(runtime, exports)
	factory.register()
	factory.registerNumber()
}

type factory struct {
	utils.Definer
	runtime *goja.Runtime
	exports *goja.Object
}

func newFactory(runtime *goja.Runtime, exports *goja.Object) *factory {
	return &factory{
		Definer: utils.MakeDefiner(runtime, exports),
		runtime: runtime,
		exports: exports,
	}
}
