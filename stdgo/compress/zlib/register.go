package zlib

import (
	"compress/zlib"

	"github.com/powerpuffpenguin/goja"
)

func (f *factory) register() {
	f.Accessor(`NoCompression`, f.getNoCompression, nil)
	f.Accessor(`BestSpeed`, f.getBestSpeed, nil)
	f.Accessor(`BestCompression`, f.getBestCompression, nil)
	f.Accessor(`DefaultCompression`, f.getDefaultCompression, nil)
	f.Accessor(`HuffmanOnly`, f.getHuffmanOnly, nil)

	f.Accessor(`ErrChecksum`, f.getErrChecksum, nil)
	f.Accessor(`ErrDictionary`, f.getErrDictionary, nil)
	f.Accessor(`ErrHeader`, f.getErrHeader, nil)

	f.Set(`NewReader`, zlib.NewReader)
	f.Set(`NewReaderDict`, zlib.NewReaderDict)

	f.Set(`NewWriter`, zlib.NewWriter)
	f.Set(`NewWriterLevel`, zlib.NewWriterLevel)
	f.Set(`NewWriterLevelDict`, zlib.NewWriterLevelDict)

	f.Set(`isResetter`, isResetter)
	f.Set(`isWriterPointer`, isWriterPointer)
}
func isResetter(i interface{}) bool {
	_, result := i.(zlib.Resetter)
	return result
}
func isWriterPointer(i interface{}) bool {
	_, result := i.(*zlib.Writer)
	return result
}

func (f *factory) getErrChecksum(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(zlib.ErrChecksum)
}
func (f *factory) getErrDictionary(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(zlib.ErrDictionary)
}
func (f *factory) getErrHeader(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(zlib.ErrHeader)
}
func (f *factory) getNoCompression(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(zlib.NoCompression)
}
func (f *factory) getBestSpeed(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(zlib.BestSpeed)
}
func (f *factory) getBestCompression(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(zlib.BestCompression)
}
func (f *factory) getDefaultCompression(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(zlib.DefaultCompression)
}
func (f *factory) getHuffmanOnly(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(zlib.HuffmanOnly)
}
