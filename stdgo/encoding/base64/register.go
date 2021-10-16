package base64

import (
	"encoding/base64"

	"github.com/powerpuffpenguin/goja"
)

func (f *factory) register() {
	f.Accessor(`StdPadding`, f.getStdPadding, nil)
	f.Accessor(`NoPadding`, f.getNoPadding, nil)

	f.Accessor(`RawStdEncoding`, f.getRawStdEncoding, nil)
	f.Accessor(`RawURLEncoding`, f.getRawURLEncoding, nil)
	f.Accessor(`StdEncoding`, f.getStdEncoding, nil)
	f.Accessor(`URLEncoding`, f.getURLEncoding, nil)

	f.Set(`NewDecoder`, base64.NewDecoder)
	f.Set(`NewEncoder`, base64.NewEncoder)

	f.Set(`NewEncoding`, base64.NewEncoding)

	f.Set(`isCorruptInputError`, isCorruptInputError)
	f.Set(`isEncodingPointer`, isEncodingPointer)
}
func isCorruptInputError(i interface{}) bool {
	_, result := i.(base64.CorruptInputError)
	return result
}
func isEncodingPointer(i interface{}) bool {
	_, result := i.(*base64.Encoding)
	return result
}
func (f *factory) getStdPadding(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(base64.StdPadding)
}
func (f *factory) getNoPadding(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(base64.NoPadding)
}
func (f *factory) getRawStdEncoding(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(base64.RawStdEncoding)
}
func (f *factory) getRawURLEncoding(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(base64.RawURLEncoding)
}
func (f *factory) getStdEncoding(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(base64.StdEncoding)
}
func (f *factory) getURLEncoding(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(base64.URLEncoding)
}
