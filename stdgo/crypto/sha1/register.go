package sha1

import (
	"crypto/sha1"

	"github.com/powerpuffpenguin/goja"
)

func (f *factory) register() {
	f.Accessor(`BlockSize`, f.getBlockSize, nil)
	f.Accessor(`Size`, f.getSize, nil)

	f.Set(`New`, sha1.New)
	f.Set(`Sum`, sha1.Sum)
}
func (f *factory) getBlockSize(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(sha1.BlockSize)
}
func (f *factory) getSize(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(sha1.Size)
}
