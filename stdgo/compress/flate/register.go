package flate

import (
	"compress/flate"

	"github.com/powerpuffpenguin/goja"
)

func (f *factory) register() {
	f.Accessor(`NoCompression`, f.getNoCompression, nil)
	f.Accessor(`BestSpeed`, f.getBestSpeed, nil)
	f.Accessor(`BestCompression`, f.getBestCompression, nil)
	f.Accessor(`DefaultCompression`, f.getDefaultCompression, nil)
	f.Accessor(`HuffmanOnly`, f.getHuffmanOnly, nil)

	f.Set(`NewReader`, flate.NewReader)
	f.Set(`NewReaderDict`, flate.NewReaderDict)

	f.Set(`NewWriter`, flate.NewWriter)
	f.Set(`NewWriterDict`, flate.NewWriterDict)

	f.Set(`isCorruptInputError`, isCorruptInputError)
	f.Set(`isInternalError`, isInternalError)
	f.Set(`isReader`, isReader)
	f.Set(`isResetter`, isResetter)
	f.Set(`isWriterPointer`, isWriterPointer)
}
func isWriterPointer(i interface{}) bool {
	_, result := i.(*flate.Writer)
	return result
}
func isCorruptInputError(i interface{}) bool {
	_, result := i.(flate.CorruptInputError)
	return result
}
func isInternalError(i interface{}) bool {
	_, result := i.(flate.InternalError)
	return result
}
func isReader(i interface{}) bool {
	_, result := i.(flate.Reader)
	return result
}
func isResetter(i interface{}) bool {
	_, result := i.(flate.Resetter)
	return result
}
func (f *factory) getNoCompression(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(flate.NoCompression)
}
func (f *factory) getBestSpeed(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(flate.BestSpeed)
}
func (f *factory) getBestCompression(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(flate.BestCompression)
}
func (f *factory) getDefaultCompression(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(flate.DefaultCompression)
}
func (f *factory) getHuffmanOnly(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(flate.HuffmanOnly)
}
