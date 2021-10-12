package bufio

import (
	"bufio"

	"github.com/powerpuffpenguin/goja"
)

func (f *factory) register() {
	f.Accessor(`MaxScanTokenSize`, f.getMaxScanTokenSize, nil)

	f.Accessor(`ErrInvalidUnreadByte`, f.getErrInvalidUnreadByte, nil)
	f.Accessor(`ErrInvalidUnreadRune`, f.getErrInvalidUnreadRune, nil)
	f.Accessor(`ErrBufferFull`, f.getErrBufferFull, nil)
	f.Accessor(`ErrNegativeCount`, f.getErrNegativeCount, nil)

	f.Accessor(`ErrTooLong`, f.getErrTooLong, nil)
	f.Accessor(`ErrNegativeAdvance`, f.getErrNegativeAdvance, nil)
	f.Accessor(`ErrAdvanceTooFar`, f.getErrAdvanceTooFar, nil)
	f.Accessor(`ErrBadReadCount`, f.getErrBadReadCount, nil)

	f.Accessor(`ErrFinalToken`, f.getErrFinalToken, nil)

	f.Set(`ScanBytes`, bufio.ScanBytes)
	f.Set(`ScanLines`, bufio.ScanLines)
	f.Set(`ScanRunes`, bufio.ScanRunes)
	f.Set(`ScanWords`, bufio.ScanWords)

	f.Set(`NewReadWriter`, bufio.NewReadWriter)

	f.Set(`NewReader`, bufio.NewReader)
	f.Set(`NewReaderSize`, bufio.NewReaderSize)

	f.Set(`NewScanner`, bufio.NewScanner)

	f.Set(`NewWriter`, bufio.NewWriter)
	f.Set(`NewWriterSize`, bufio.NewWriterSize)

	f.Set(`isReadWriterPointer`, isReadWriterPointer)
	f.Set(`isReaderPointer`, isReaderPointer)
	f.Set(`isScannerPointer`, isScannerPointer)
	f.Set(`isSplitFunc`, isSplitFunc)
	f.Set(`isWriterPointer`, isWriterPointer)

}
func isReadWriterPointer(i interface{}) bool {
	_, result := i.(*bufio.ReadWriter)
	return result
}
func isReaderPointer(i interface{}) bool {
	_, result := i.(*bufio.Reader)
	return result
}
func isScannerPointer(i interface{}) bool {
	_, result := i.(*bufio.Scanner)
	return result
}
func isSplitFunc(i interface{}) bool {
	_, result := i.(bufio.SplitFunc)
	return result
}
func isWriterPointer(i interface{}) bool {
	_, result := i.(*bufio.Writer)
	return result
}
func (f *factory) getErrFinalToken(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(bufio.ErrFinalToken)
}
func (f *factory) getErrTooLong(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(bufio.ErrTooLong)
}
func (f *factory) getErrNegativeAdvance(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(bufio.ErrNegativeAdvance)
}
func (f *factory) getErrAdvanceTooFar(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(bufio.ErrAdvanceTooFar)
}
func (f *factory) getErrBadReadCount(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(bufio.ErrBadReadCount)
}
func (f *factory) getErrInvalidUnreadByte(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(bufio.ErrInvalidUnreadByte)
}
func (f *factory) getErrInvalidUnreadRune(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(bufio.ErrInvalidUnreadRune)
}
func (f *factory) getErrBufferFull(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(bufio.ErrBufferFull)
}
func (f *factory) getErrNegativeCount(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(bufio.ErrNegativeCount)
}
func (f *factory) getMaxScanTokenSize(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(bufio.MaxScanTokenSize)
}
