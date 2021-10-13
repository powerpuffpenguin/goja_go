package adler32

import (
	"hash/adler32"

	"github.com/powerpuffpenguin/goja"
)

func (f *factory) register() {
	f.Accessor(`Size`, f.getSize, nil)

	f.Set(`Checksum`, adler32.Checksum)
	f.Set(`New`, adler32.New)
}
func (f *factory) getSize(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(adler32.Size)
}
