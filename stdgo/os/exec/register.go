package exec

import (
	"os/exec"

	"github.com/powerpuffpenguin/goja"
)

func (f *factory) register() {
	f.Accessor(`ErrNotFound`, f.getErrNotFound, nil)

	f.Set(`LookPath`, exec.LookPath)

	f.Set(`Command`, exec.Command)
	f.Set(`CommandContext`, exec.CommandContext)

	f.Set(`isCmdPointer`, isCmdPointer)
	f.Set(`isErrorPointer`, isErrorPointer)
	f.Set(`isExitErrorPointer`, isExitErrorPointer)
}
func isCmdPointer(i interface{}) bool {
	_, result := i.(*exec.Cmd)
	return result
}
func isErrorPointer(i interface{}) bool {
	_, result := i.(*exec.Error)
	return result
}
func isExitErrorPointer(i interface{}) bool {
	_, result := i.(*exec.ExitError)
	return result
}
func (f *factory) getErrNotFound(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue((exec.ErrNotFound))
}
