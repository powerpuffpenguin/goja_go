package sha256

import (
	"crypto/sha256"

	"github.com/powerpuffpenguin/goja"
)

func (f *factory) register() {
	f.Accessor(`BlockSize`, f.getBlockSize, nil)
	f.Accessor(`Size`, f.getSize, nil)
	f.Accessor(`Size224`, f.getSize224, nil)

	f.Set(`New`, sha256.New)
	f.Set(`New224`, sha256.New224)
	f.Set(`Sum224`, sha256.Sum224)
	f.Set(`Sum256`, sha256.Sum256)
}
func (f *factory) getBlockSize(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(sha256.BlockSize)
}
func (f *factory) getSize(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(sha256.Size)
}
func (f *factory) getSize224(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(sha256.Size224)
}
