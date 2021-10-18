package rand

import (
	"crypto/rand"

	"github.com/powerpuffpenguin/goja"
)

func (f *factory) register() {
	f.Accessor(`Reader`, f.getReader, nil)

	f.Set(`Int`, rand.Int)
	f.Set(`Prime`, rand.Prime)
	f.Set(`Read`, rand.Read)
}
func (f *factory) getReader(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(rand.Reader)
}
