package zip

import (
	"archive/zip"

	"github.com/powerpuffpenguin/goja"
)

func (f *factory) register() {
	f.Accessor(`Store`, f.getStore, nil)
	f.Accessor(`Deflate`, f.getDeflate, nil)

	f.Accessor(`ErrFormat`, f.getErrFormat, nil)
	f.Accessor(`ErrAlgorithm`, f.getErrAlgorithm, nil)
	f.Accessor(`ErrChecksum`, f.getErrChecksum, nil)

	f.Set(`RegisterCompressor`, zip.RegisterCompressor)
	f.Set(`RegisterDecompressor`, zip.RegisterDecompressor)

	f.Set(`FileInfoHeader`, zip.FileInfoHeader)
	f.Set(`OpenReader`, zip.OpenReader)
	f.Set(`NewReader`, zip.NewReader)
	f.Set(`NewWriter`, zip.NewWriter)

	f.Set(`isCompressor`, isCompressor)
	f.Set(`isDecompressor`, isDecompressor)
	f.Set(`isFilePointer`, isFilePointer)
	f.Set(`isFileHeaderPointer`, isFileHeaderPointer)
	f.Set(`isReadCloserPointer`, isReadCloserPointer)
	f.Set(`isReaderPointer`, isReaderPointer)
	f.Set(`isWriterPointer`, isWriterPointer)
}
func isFilePointer(i interface{}) bool {
	_, result := i.(*zip.File)
	return result
}
func isFileHeaderPointer(i interface{}) bool {
	_, result := i.(*zip.FileHeader)
	return result
}
func isReadCloserPointer(i interface{}) bool {
	_, result := i.(*zip.ReadCloser)
	return result
}
func isReaderPointer(i interface{}) bool {
	_, result := i.(*zip.Reader)
	return result
}
func isWriterPointer(i interface{}) bool {
	_, result := i.(*zip.Writer)
	return result
}
func isCompressor(i interface{}) bool {
	_, result := i.(zip.Compressor)
	return result
}
func isDecompressor(i interface{}) bool {
	_, result := i.(zip.Decompressor)
	return result
}
func (f *factory) getStore(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(zip.Store)
}
func (f *factory) getDeflate(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(zip.Deflate)
}
func (f *factory) getErrFormat(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(zip.ErrFormat)
}
func (f *factory) getErrAlgorithm(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(zip.ErrAlgorithm)
}
func (f *factory) getErrChecksum(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(zip.ErrChecksum)
}
