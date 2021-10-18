package md5

import (
	"crypto/md5"

	"github.com/powerpuffpenguin/goja"
)

func (f *factory) register() {
	f.Accessor(`BlockSize`, f.getBlockSize, nil)
	f.Accessor(`Size`, f.getSize, nil)

	f.Set(`New`, md5.New)
	f.Set(`Sum`, md5.Sum)
}
func (f *factory) getBlockSize(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(md5.BlockSize)
}
func (f *factory) getSize(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(md5.Size)
}
