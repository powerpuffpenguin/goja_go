package base32

import (
	"encoding/base32"

	"github.com/powerpuffpenguin/goja"
)

func (f *factory) register() {
	f.Accessor(`StdPadding`, f.getStdPadding, nil)
	f.Accessor(`NoPadding`, f.getNoPadding, nil)

	f.Accessor(`HexEncoding`, f.getHexEncoding, nil)
	f.Accessor(`StdEncoding`, f.getStdEncoding, nil)

	f.Set(`NewDecoder`, base32.NewDecoder)
	f.Set(`NewEncoder`, base32.NewEncoder)

	f.Set(`NewEncoding`, base32.NewEncoding)

	f.Set(`isCorruptInputError`, isCorruptInputError)
	f.Set(`isEncodingPointer`, isEncodingPointer)
}
func isCorruptInputError(i interface{}) bool {
	_, result := i.(base32.CorruptInputError)
	return result
}
func isEncodingPointer(i interface{}) bool {
	_, result := i.(*base32.Encoding)
	return result
}
func (f *factory) getHexEncoding(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(base32.HexEncoding)
}
func (f *factory) getStdEncoding(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(base32.StdEncoding)
}
func (f *factory) getStdPadding(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(base32.StdPadding)
}
func (f *factory) getNoPadding(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(base32.NoPadding)
}
