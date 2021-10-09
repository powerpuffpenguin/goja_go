package io

import (
	"io"

	"github.com/powerpuffpenguin/goja"
)

func (f *factory) register() {
	f.Accessor(`SeekStart`, f.getSeekStart, nil)
	f.Accessor(`SeekCurrent`, f.getSeekCurrent, nil)
	f.Accessor(`SeekEnd`, f.getSeekEnd, nil)

	f.Accessor(`EOF`, f.getEOF, nil)
	f.Accessor(`ErrClosedPipe`, f.getErrClosedPipe, nil)
	f.Accessor(`ErrNoProgress`, f.getErrNoProgress, nil)
	f.Accessor(`ErrShortBuffer`, f.getErrShortBuffer, nil)
	f.Accessor(`ErrShortWrite`, f.getErrShortWrite, nil)
	f.Accessor(`ErrUnexpectedEOF`, f.getErrUnexpectedEOF, nil)

	f.Set(`Copy`, io.Copy)
	f.Set(`CopyBuffer`, io.CopyBuffer)
	f.Set(`CopyN`, io.CopyN)
	f.Set(`Pipe`, io.Pipe)
	f.Set(`ReadAll`, io.ReadAll)
	f.Set(`ReadAtLeast`, io.ReadAtLeast)
	f.Set(`ReadFull`, io.ReadFull)
	f.Set(`WriteString`, io.WriteString)

	f.Set(`NopCloser`, io.NopCloser)

	f.Set(`LimitReader`, io.LimitReader)
	f.Set(`MultiReader`, io.MultiReader)
	f.Set(`TeeReader`, io.TeeReader)

	f.Set(`NewSectionReader`, io.NewSectionReader)

	f.Accessor(`Discard`, f.getDiscard, nil)
	f.Set(`MultiWriter`, io.MultiWriter)

	f.Set(`isByteReader`, isByteReader)
	f.Set(`isByteScanner`, isByteScanner)
	f.Set(`isByteWriter`, isByteWriter)
	f.Set(`isCloser`, isCloser)
	f.Set(`isLimitedReaderPointer`, isLimitedReaderPointer)
	f.Set(`isPipeReaderPointer`, isPipeReaderPointer)
	f.Set(`isPipeWriterPointer`, isPipeWriterPointer)
	f.Set(`isReadCloser`, isReadCloser)
	f.Set(`isReadSeekCloser`, isReadSeekCloser)
	f.Set(`isReadSeeker`, isReadSeeker)
	f.Set(`isReadWriteCloser`, isReadWriteCloser)
	f.Set(`isReadWriteSeeker`, isReadWriteSeeker)
	f.Set(`isReadWriter`, isReadWriter)
	f.Set(`isReader`, isReader)
	f.Set(`isReaderAt`, isReaderAt)
	f.Set(`isReaderFrom`, isReaderFrom)
	f.Set(`isRuneReader`, isRuneReader)
	f.Set(`isRuneScanner`, isRuneScanner)
	f.Set(`isSectionReaderPointer`, isSectionReaderPointer)
	f.Set(`isSeeker`, isSeeker)
	f.Set(`isStringWriter`, isStringWriter)
	f.Set(`isWriteCloser`, isWriteCloser)
	f.Set(`isWriteSeeker`, isWriteSeeker)
	f.Set(`isWriter`, isWriter)
	f.Set(`isWriterAt`, isWriterAt)
	f.Set(`isWriterTo`, isWriterTo)
}
func isByteReader(i interface{}) bool {
	_, result := i.(io.ByteReader)
	return result
}
func isByteScanner(i interface{}) bool {
	_, result := i.(io.ByteScanner)
	return result
}
func isByteWriter(i interface{}) bool {
	_, result := i.(io.ByteWriter)
	return result
}
func isCloser(i interface{}) bool {
	_, result := i.(io.Closer)
	return result
}
func isLimitedReaderPointer(i interface{}) bool {
	_, result := i.(*io.LimitedReader)
	return result
}
func isPipeReaderPointer(i interface{}) bool {
	_, result := i.(*io.PipeReader)
	return result
}
func isPipeWriterPointer(i interface{}) bool {
	_, result := i.(*io.PipeWriter)
	return result
}
func isReadCloser(i interface{}) bool {
	_, result := i.(io.ReadCloser)
	return result
}
func isReadSeekCloser(i interface{}) bool {
	_, result := i.(io.ReadSeekCloser)
	return result
}
func isReadSeeker(i interface{}) bool {
	_, result := i.(io.ReadSeeker)
	return result
}
func isReadWriteCloser(i interface{}) bool {
	_, result := i.(io.ReadWriteCloser)
	return result
}
func isReadWriteSeeker(i interface{}) bool {
	_, result := i.(io.ReadWriteSeeker)
	return result
}
func isReadWriter(i interface{}) bool {
	_, result := i.(io.ReadWriter)
	return result
}
func isReader(i interface{}) bool {
	_, result := i.(io.Reader)
	return result
}
func isReaderAt(i interface{}) bool {
	_, result := i.(io.ReaderAt)
	return result
}
func isReaderFrom(i interface{}) bool {
	_, result := i.(io.ReaderFrom)
	return result
}
func isRuneReader(i interface{}) bool {
	_, result := i.(io.RuneReader)
	return result
}
func isRuneScanner(i interface{}) bool {
	_, result := i.(io.RuneScanner)
	return result
}
func isSectionReaderPointer(i interface{}) bool {
	_, result := i.(*io.SectionReader)
	return result
}
func isSeeker(i interface{}) bool {
	_, result := i.(io.Seeker)
	return result
}
func isStringWriter(i interface{}) bool {
	_, result := i.(io.StringWriter)
	return result
}
func isWriteCloser(i interface{}) bool {
	_, result := i.(io.WriteCloser)
	return result
}
func isWriteSeeker(i interface{}) bool {
	_, result := i.(io.WriteSeeker)
	return result
}
func isWriter(i interface{}) bool {
	_, result := i.(io.Writer)
	return result
}
func isWriterAt(i interface{}) bool {
	_, result := i.(io.WriterAt)
	return result
}
func isWriterTo(i interface{}) bool {
	_, result := i.(io.WriterTo)
	return result
}
func (f *factory) getSeekStart(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(io.SeekStart)
}
func (f *factory) getSeekCurrent(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(io.SeekCurrent)
}
func (f *factory) getSeekEnd(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(io.SeekEnd)
}
func (f *factory) getEOF(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(io.EOF)
}
func (f *factory) getErrClosedPipe(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(io.ErrClosedPipe)
}
func (f *factory) getErrNoProgress(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(io.ErrNoProgress)
}
func (f *factory) getErrShortBuffer(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(io.ErrShortBuffer)
}
func (f *factory) getErrShortWrite(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(io.ErrShortWrite)
}
func (f *factory) getErrUnexpectedEOF(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(io.ErrUnexpectedEOF)
}
func (f *factory) getDiscard(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(io.Discard)
}
