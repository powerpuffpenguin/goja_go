package sha512

import (
	"crypto/sha512"

	"github.com/powerpuffpenguin/goja"
)

func (f *factory) register() {
	f.Accessor(`Size`, f.getSize, nil)
	f.Accessor(`Size224`, f.getSize224, nil)
	f.Accessor(`Size256`, f.getSize256, nil)
	f.Accessor(`Size384`, f.getSize384, nil)
	f.Accessor(`BlockSize`, f.getBlockSize, nil)

	f.Set(`New`, sha512.New)
	f.Set(`New384`, sha512.New384)
	f.Set(`New512_224`, sha512.New512_224)
	f.Set(`New512_256`, sha512.New512_256)
	f.Set(`Sum384`, sha512.Sum384)
	f.Set(`Sum512`, sha512.Sum512)
	f.Set(`Sum512_224`, sha512.Sum512_224)
	f.Set(`Sum512_256`, sha512.Sum512_256)
}
func (f *factory) getSize(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(sha512.Size)
}
func (f *factory) getSize224(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(sha512.Size224)
}
func (f *factory) getSize256(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(sha512.Size256)
}
func (f *factory) getSize384(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(sha512.Size384)
}
func (f *factory) getBlockSize(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(sha512.BlockSize)
}
