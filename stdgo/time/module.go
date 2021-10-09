package time

import (
	"github.com/powerpuffpenguin/goja"
	"github.com/powerpuffpenguin/goja_go/core/utils"
)

const (
	ModuleID = `stdgo/time`
)

func Require(runtime *goja.Runtime, module *goja.Object) {
	exports := module.Get(`exports`).(*goja.Object)
	factory := newFactory(runtime, exports)
	factory.register()
	factory.registerDuration()
	factory.registerTime()
	// factory.registerTimer()
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
