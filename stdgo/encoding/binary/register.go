package binary

import (
	"encoding/binary"

	"github.com/powerpuffpenguin/goja"
)

func (f *factory) register() {
	f.Accessor(`MaxVarintLen16`, f.getMaxVarintLen16, nil)
	f.Accessor(`MaxVarintLen32`, f.getMaxVarintLen32, nil)
	f.Accessor(`MaxVarintLen64`, f.getMaxVarintLen64, nil)

	f.Accessor(`BigEndian`, f.getBigEndian, nil)
	f.Accessor(`LittleEndian`, f.getLittleEndian, nil)

	f.Set(`PutUvarint`, binary.PutUvarint)
	f.Set(`PutVarint`, binary.PutVarint)
	f.Set(`Read`, binary.Read)
	f.Set(`ReadUvarint`, binary.ReadUvarint)
	f.Set(`ReadVarint`, binary.ReadVarint)
	f.Set(`Size`, binary.Size)
	f.Set(`Uvarint`, binary.Uvarint)
	f.Set(`Varint`, binary.Varint)
	f.Set(`Write`, binary.Write)

	f.Set(`isByteOrder`, isByteOrder)
}
func (f *factory) getBigEndian(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(binary.BigEndian)
}
func (f *factory) getLittleEndian(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(binary.LittleEndian)
}
func (f *factory) getMaxVarintLen16(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(binary.MaxVarintLen16)
}
func (f *factory) getMaxVarintLen32(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(binary.MaxVarintLen32)
}
func (f *factory) getMaxVarintLen64(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(binary.MaxVarintLen64)
}
func isByteOrder(i interface{}) bool {
	_, result := i.(binary.ByteOrder)
	return result
}
