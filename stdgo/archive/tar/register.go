package tar

import (
	"archive/tar"

	"github.com/powerpuffpenguin/goja"
	"github.com/powerpuffpenguin/goja_go/stdgo/builtin"
)

func (f *factory) register() {
	f.Accessor(`TypeReg`, f.getTypeReg, nil)
	f.Accessor(`TypeRegA`, f.getTypeRegA, nil)
	f.Accessor(`TypeLink`, f.getTypeLink, nil)
	f.Accessor(`TypeSymlink`, f.getTypeSymlink, nil)
	f.Accessor(`TypeChar`, f.getTypeChar, nil)
	f.Accessor(`TypeBlock`, f.getTypeBlock, nil)
	f.Accessor(`TypeDir`, f.getTypeDir, nil)
	f.Accessor(`TypeFifo`, f.getTypeFifo, nil)
	f.Accessor(`TypeCont`, f.getTypeCont, nil)
	f.Accessor(`TypeXHeader`, f.getTypeXHeader, nil)
	f.Accessor(`TypeXGlobalHeader`, f.getTypeXGlobalHeader, nil)
	f.Accessor(`TypeGNUSparse`, f.getTypeGNUSparse, nil)
	f.Accessor(`TypeGNULongName`, f.getTypeGNULongName, nil)
	f.Accessor(`TypeGNULongLink`, f.getTypeGNULongLink, nil)

	f.Accessor(`ErrHeader`, f.getErrHeader, nil)
	f.Accessor(`ErrWriteTooLong`, f.getErrWriteTooLong, nil)
	f.Accessor(`ErrFieldTooLong`, f.getErrFieldTooLong, nil)
	f.Accessor(`ErrWriteAfterClose`, f.getErrWriteAfterClose, nil)

	f.Set(`FileInfoHeader`, tar.FileInfoHeader)

	f.Set(`NewReader`, tar.NewReader)
	f.Set(`NewWriter`, tar.NewWriter)

	f.Set(`isFormat`, isFormat)
	f.Set(`isHeaderPointer`, isHeaderPointer)
	f.Set(`isReaderPointer`, isReaderPointer)
	f.Set(`isWriterPointer`, isWriterPointer)
}
func isFormat(i interface{}) bool {
	_, result := i.(tar.Format)
	return result
}
func isHeaderPointer(i interface{}) bool {
	_, result := i.(*tar.Header)
	return result
}
func isReaderPointer(i interface{}) bool {
	_, result := i.(*tar.Reader)
	return result
}
func isWriterPointer(i interface{}) bool {
	_, result := i.(*tar.Writer)
	return result
}
func (f *factory) getErrHeader(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(tar.ErrHeader)
}
func (f *factory) getErrWriteTooLong(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(tar.ErrWriteTooLong)
}
func (f *factory) getErrFieldTooLong(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(tar.ErrFieldTooLong)
}
func (f *factory) getErrWriteAfterClose(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(tar.ErrWriteAfterClose)
}
func (f *factory) getTypeReg(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Uint8(tar.TypeReg))
}
func (f *factory) getTypeRegA(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Uint8(tar.TypeRegA))
}
func (f *factory) getTypeLink(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Uint8(tar.TypeLink))
}
func (f *factory) getTypeSymlink(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Uint8(tar.TypeSymlink))
}
func (f *factory) getTypeChar(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Uint8(tar.TypeChar))
}
func (f *factory) getTypeBlock(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Uint8(tar.TypeBlock))
}
func (f *factory) getTypeDir(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Uint8(tar.TypeDir))
}
func (f *factory) getTypeFifo(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Uint8(tar.TypeFifo))
}
func (f *factory) getTypeCont(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Uint8(tar.TypeCont))
}
func (f *factory) getTypeXHeader(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Uint8(tar.TypeXHeader))
}
func (f *factory) getTypeXGlobalHeader(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Uint8(tar.TypeXGlobalHeader))
}
func (f *factory) getTypeGNUSparse(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Uint8(tar.TypeGNUSparse))
}
func (f *factory) getTypeGNULongName(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Uint8(tar.TypeGNULongName))
}
func (f *factory) getTypeGNULongLink(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Uint8(tar.TypeGNULongLink))
}
