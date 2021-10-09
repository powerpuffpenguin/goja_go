package stdgo

import (
	"github.com/powerpuffpenguin/goja"
	"github.com/powerpuffpenguin/goja/require"

	errors "github.com/powerpuffpenguin/goja_go/stdgo/errors"

	io "github.com/powerpuffpenguin/goja_go/stdgo/io"
	io_fs "github.com/powerpuffpenguin/goja_go/stdgo/io/fs"
	io_ioutil "github.com/powerpuffpenguin/goja_go/stdgo/io/ioutil"

	os "github.com/powerpuffpenguin/goja_go/stdgo/os"

	path "github.com/powerpuffpenguin/goja_go/stdgo/path"
	path_filepath "github.com/powerpuffpenguin/goja_go/stdgo/path/filepath"

	time "github.com/powerpuffpenguin/goja_go/stdgo/time"
)

var modules = []module{
	{errors.ModuleID, errors.Require},

	{io.ModuleID, io.Require},
	{io_fs.ModuleID, io_fs.Require},
	{io_ioutil.ModuleID, io_ioutil.Require},

	{os.ModuleID, os.Require},

	{path.ModuleID, path.Require},
	{path_filepath.ModuleID, path_filepath.Require},

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
