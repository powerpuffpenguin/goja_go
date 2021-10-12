package context

import (
	"context"

	"github.com/powerpuffpenguin/goja"
)

func (f *factory) register() {
	f.Accessor(`Canceled`, f.getCanceled, nil)
	f.Accessor(`DeadlineExceeded`, f.getDeadlineExceeded, nil)

	f.Set(`WithCancel`, context.WithCancel)
	f.Set(`WithDeadline`, context.WithDeadline)
	f.Set(`WithTimeout`, context.WithTimeout)

	f.Set(`Background`, context.Background)
	f.Set(`TODO`, context.TODO)
	f.Set(`WithValue`, context.WithValue)

	f.Set(`isCancelFunc`, isCancelFunc)
	f.Set(`isContext`, isContext)
}
func isCancelFunc(i interface{}) bool {
	_, result := i.(context.CancelFunc)
	return result
}
func isContext(i interface{}) bool {
	_, result := i.(context.Context)
	return result
}
func (f *factory) getCanceled(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(context.Canceled)
}
func (f *factory) getDeadlineExceeded(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(context.DeadlineExceeded)
}
