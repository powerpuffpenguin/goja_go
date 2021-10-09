package console

import (
	"fmt"

	"github.com/powerpuffpenguin/goja"
	"github.com/powerpuffpenguin/goja/require"
)

const (
	ModuleID = `console`
)

func init() {
	require.RegisterNativeModule(ModuleID, Require)
}
func Enable(runtime *goja.Runtime) {
	runtime.Set("console", require.Require(runtime, "console"))
}

func Require(runtime *goja.Runtime, module *goja.Object) {
	obj := module.Get(`exports`).(*goja.Object)
	obj.Set(`trace`, nativeLog)
	obj.Set(`debug`, nativeLog)
	obj.Set(`log`, nativeLog)
	obj.Set(`info`, nativeLog)
	obj.Set(`warn`, nativeLog)
	obj.Set(`error`, nativeLog)
}
func nativeLog(call goja.FunctionCall) goja.Value {
	count := len(call.Arguments)
	if count == 0 {
		fmt.Println()
		return nil
	}
	arrs := make([]interface{}, count)
	for i := 0; i < count; i++ {
		arg := call.Arguments[i]
		export := arg.Export()
		if native, ok := export.(interface {
			String() string
		}); ok {
			arrs[i] = native.String()
		} else {
			arrs[i] = arg.ToString()
		}
	}
	fmt.Println(arrs...)
	return nil
}
