package stdgo

import (
	"github.com/powerpuffpenguin/goja"
	"github.com/powerpuffpenguin/goja/require"

	io "github.com/powerpuffpenguin/goja_go/stdgo/io"
	io_fs "github.com/powerpuffpenguin/goja_go/stdgo/io/fs"

	os "github.com/powerpuffpenguin/goja_go/stdgo/os"

	time "github.com/powerpuffpenguin/goja_go/stdgo/time"
)

var modules = []module{
	{io.ModuleID, io.Require},
	{io_fs.ModuleID, io_fs.Require},

	{os.ModuleID, os.Require},

	{time.ModuleID, time.Require},
}

type module struct {
	ID      string
	Require func(runtime *goja.Runtime, module *goja.Object)
}

func RegisterNativeModule() {
	for _, m := range modules {
		require.RegisterNativeModule(m.ID, m.Require)
	}
}
func RegisterNativeModuleToRegistry(registry *require.Registry) {
	for _, m := range modules {
		registry.RegisterNativeModule(m.ID, m.Require)
	}
}
