package crc32

import (
	"hash/crc32"

	"github.com/powerpuffpenguin/goja"
	"github.com/powerpuffpenguin/goja_go/stdgo/builtin"
)

func (f *factory) register() {
	f.Accessor(`IEEE`, f.getIEEE, nil)
	f.Accessor(`Castagnoli`, f.getCastagnoli, nil)
	f.Accessor(`Koopman`, f.getKoopman, nil)

	f.Accessor(`Size`, f.getSize, nil)
	f.Accessor(`IEEETable`, f.getIEEETable, nil)

	f.Set(`Checksum`, crc32.Checksum)
	f.Set(`ChecksumIEEE`, crc32.ChecksumIEEE)
	f.Set(`New`, crc32.New)
	f.Set(`NewIEEE`, crc32.NewIEEE)
	f.Set(`Update`, crc32.Update)

	f.Set(`MakeTable`, crc32.MakeTable)

	f.Set(`isTablePointer`, isTablePointer)
}
func isTablePointer(i interface{}) bool {
	_, result := i.(*crc32.Table)
	return result
}
func (f *factory) getIEEETable(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(crc32.IEEETable)
}
func (f *factory) getSize(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(crc32.Size)
}
func (f *factory) getIEEE(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Uint32(crc32.IEEE))
}
func (f *factory) getCastagnoli(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Uint32(crc32.Castagnoli))
}
func (f *factory) getKoopman(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Uint32(crc32.Koopman))
}
