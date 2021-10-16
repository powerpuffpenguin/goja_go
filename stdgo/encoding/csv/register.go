package csv

import (
	"encoding/csv"

	"github.com/powerpuffpenguin/goja"
)

func (f *factory) register() {
	f.Accessor(`ErrTrailingComma`, f.getErrTrailingComma, nil)
	f.Accessor(`ErrBareQuote`, f.getErrBareQuote, nil)
	f.Accessor(`ErrQuote`, f.getErrQuote, nil)
	f.Accessor(`ErrFieldCount`, f.getErrFieldCount, nil)

	f.Set(`NewReader`, csv.NewReader)
	f.Set(`NewWriter`, csv.NewWriter)

	f.Set(`isParseErrorPointer`, isParseErrorPointer)
	f.Set(`isReaderPointer`, isReaderPointer)
	f.Set(`isWriterPointer`, isWriterPointer)
}
func (f *factory) getErrTrailingComma(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(csv.ErrTrailingComma)
}
func (f *factory) getErrBareQuote(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(csv.ErrBareQuote)
}
func (f *factory) getErrQuote(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(csv.ErrQuote)
}
func (f *factory) getErrFieldCount(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(csv.ErrFieldCount)
}
func isParseErrorPointer(i interface{}) bool {
	_, result := i.(*csv.ParseError)
	return result
}
func isReaderPointer(i interface{}) bool {
	_, result := i.(*csv.Reader)
	return result
}
func isWriterPointer(i interface{}) bool {
	_, result := i.(*csv.Writer)
	return result
}
