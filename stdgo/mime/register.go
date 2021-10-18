package mime

import (
	"mime"

	"github.com/powerpuffpenguin/goja"
)

func (f *factory) register() {
	f.Accessor(`BEncoding`, f.getBEncoding, nil)
	f.Accessor(`QEncoding`, f.getQEncoding, nil)
	f.Accessor(`ErrInvalidMediaParameter`, f.getErrInvalidMediaParameter, nil)

	f.Set(`AddExtensionType`, mime.AddExtensionType)
	f.Set(`ExtensionsByType`, mime.ExtensionsByType)
	f.Set(`FormatMediaType`, mime.FormatMediaType)
	f.Set(`ParseMediaType`, mime.ParseMediaType)
	f.Set(`TypeByExtension`, mime.TypeByExtension)

	f.Set(`WordEncoder`, WordEncoder)
	f.Set(`isWordDecoderPointer`, isWordDecoderPointer)
	f.Set(`isWordEncoder`, isWordEncoder)
}
func (f *factory) getBEncoding(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(mime.BEncoding)
}
func (f *factory) getQEncoding(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(mime.QEncoding)
}
func (f *factory) getErrInvalidMediaParameter(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(mime.ErrInvalidMediaParameter)
}
func WordEncoder(v byte) mime.WordEncoder {
	return mime.WordEncoder(v)
}
func isWordDecoderPointer(i interface{}) bool {
	_, result := i.(*mime.WordDecoder)
	return result
}
func isWordEncoder(i interface{}) bool {
	_, result := i.(mime.WordEncoder)
	return result
}
