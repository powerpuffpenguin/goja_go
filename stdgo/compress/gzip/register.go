package gzip

import (
	"compress/gzip"

	"github.com/powerpuffpenguin/goja"
)

func (f *factory) register() {
	f.Accessor(`NoCompression`, f.getNoCompression, nil)
	f.Accessor(`BestSpeed`, f.getBestSpeed, nil)
	f.Accessor(`BestCompression`, f.getBestCompression, nil)
	f.Accessor(`DefaultCompression`, f.getDefaultCompression, nil)
	f.Accessor(`HuffmanOnly`, f.getHuffmanOnly, nil)

	f.Accessor(`ErrChecksum`, f.getErrChecksum, nil)
	f.Accessor(`ErrHeader`, f.getErrHeader, nil)

	f.Set(`NewReader`, gzip.NewReader)
	f.Set(`NewWriter`, gzip.NewWriter)
	f.Set(`NewWriterLevel`, gzip.NewWriterLevel)

	f.Set(`isHeader`, isHeader)
	f.Set(`isReaderPointer`, isReaderPointer)
	f.Set(`isWriterPointer`, isWriterPointer)
}
func isHeader(i interface{}) bool {
	_, result := i.(gzip.Header)
	return result
}
func isReaderPointer(i interface{}) bool {
	_, result := i.(*gzip.Reader)
	return result
}
func isWriterPointer(i interface{}) bool {
	_, result := i.(*gzip.Writer)
	return result
}

func (f *factory) getErrChecksum(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(gzip.ErrChecksum)
}
func (f *factory) getErrHeader(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(gzip.ErrHeader)
}
func (f *factory) getNoCompression(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(gzip.NoCompression)
}
func (f *factory) getBestSpeed(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(gzip.BestSpeed)
}
func (f *factory) getBestCompression(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(gzip.BestCompression)
}
func (f *factory) getDefaultCompression(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(gzip.DefaultCompression)
}
func (f *factory) getHuffmanOnly(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(gzip.HuffmanOnly)
}
