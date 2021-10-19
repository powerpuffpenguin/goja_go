package multipart

import (
	"mime/multipart"

	"github.com/powerpuffpenguin/goja"
)

func (f *factory) register() {
	f.Accessor(`ErrMessageTooLarge`, f.getErrMessageTooLarge, nil)
	f.Set(`NewReader`, multipart.NewReader)
	f.Set(`NewWriter`, multipart.NewWriter)
	f.Set(`isFile`, isFile)
	f.Set(`isFileHeaderPointer`, isFileHeaderPointer)
	f.Set(`isFormPointer`, isFormPointer)
	f.Set(`isPartPointer`, isPartPointer)
	f.Set(`isReaderPointer`, isReaderPointer)
	f.Set(`isWriterPointer`, isWriterPointer)
}
func (f *factory) getErrMessageTooLarge(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(multipart.ErrMessageTooLarge)
}
func isFile(i interface{}) bool {
	_, result := i.(multipart.File)
	return result
}
func isFileHeaderPointer(i interface{}) bool {
	_, result := i.(*multipart.FileHeader)
	return result
}
func isFormPointer(i interface{}) bool {
	_, result := i.(*multipart.Form)
	return result
}
func isPartPointer(i interface{}) bool {
	_, result := i.(*multipart.Part)
	return result
}
func isReaderPointer(i interface{}) bool {
	_, result := i.(*multipart.Reader)
	return result
}
func isWriterPointer(i interface{}) bool {
	_, result := i.(*multipart.Writer)
	return result
}
