package stdgo

import (
	"github.com/powerpuffpenguin/goja"
	"github.com/powerpuffpenguin/goja/require"
	"github.com/powerpuffpenguin/goja_go/stdgo/builtin"
)

func Enable(r *goja.Runtime) {
	require.RegisterNativeModule(builtin.ModuleID, builtin.Require)
}
