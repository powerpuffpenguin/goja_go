package crc64

import (
	"hash/crc64"

	"github.com/powerpuffpenguin/goja"
	"github.com/powerpuffpenguin/goja_go/stdgo/builtin"
)

func (f *factory) register() {
	f.Accessor(`ISO`, f.getISO, nil)
	f.Accessor(`ECMA`, f.getECMA, nil)

	f.Accessor(`Size`, f.getSize, nil)

	f.Set(`Checksum`, crc64.Checksum)
	f.Set(`New`, crc64.New)
	f.Set(`Update`, crc64.Update)

	f.Set(`MakeTable`, crc64.MakeTable)

	f.Set(`isTablePointer`, isTablePointer)
}
func isTablePointer(i interface{}) bool {
	_, result := i.(*crc64.Table)
	return result
}
func (f *factory) getSize(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Uint64(crc64.Size))
}
func (f *factory) getISO(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Uint64(crc64.ISO))
}
func (f *factory) getECMA(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Uint64(crc64.ECMA))
}
