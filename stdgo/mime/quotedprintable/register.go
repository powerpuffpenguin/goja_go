package quotedprintable

import "mime/quotedprintable"

func (f *factory) register() {
	f.Set(`NewReader`, quotedprintable.NewReader)
	f.Set(`NewWriter`, quotedprintable.NewWriter)
	f.Set(`isReaderPointer`, isReaderPointer)
	f.Set(`isWriterPointer`, isWriterPointer)
}
func isReaderPointer(i interface{}) bool {
	_, result := i.(*quotedprintable.Reader)
	return result
}
func isWriterPointer(i interface{}) bool {
	_, result := i.(*quotedprintable.Writer)
	return result
}
