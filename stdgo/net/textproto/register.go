package textproto

import "net/textproto"

func (f *factory) register() {
	f.Set(`CanonicalMIMEHeaderKey`, textproto.CanonicalMIMEHeaderKey)
	f.Set(`TrimBytes`, textproto.TrimBytes)
	f.Set(`TrimString`, textproto.TrimString)

	f.Set(`Dial`, textproto.Dial)
	f.Set(`NewConn`, textproto.NewConn)
	f.Set(`NewReader`, textproto.NewReader)
	f.Set(`NewWriter`, textproto.NewWriter)
	f.Set(`isConnPointer`, isConnPointer)
	f.Set(`isErrorPointer`, isErrorPointer)
	f.Set(`isMIMEHeader`, isMIMEHeader)
	f.Set(`isPipelinePointer`, isPipelinePointer)
	f.Set(`isProtocolError`, isProtocolError)
	f.Set(`isReaderPointer`, isReaderPointer)
	f.Set(`isWriterPointer`, isWriterPointer)
}
func isConnPointer(i interface{}) bool {
	_, result := i.(*textproto.Conn)
	return result
}
func isErrorPointer(i interface{}) bool {
	_, result := i.(*textproto.Error)
	return result
}
func isMIMEHeader(i interface{}) bool {
	_, result := i.(textproto.MIMEHeader)
	return result
}
func isPipelinePointer(i interface{}) bool {
	_, result := i.(*textproto.Pipeline)
	return result
}
func isProtocolError(i interface{}) bool {
	_, result := i.(textproto.ProtocolError)
	return result
}
func isReaderPointer(i interface{}) bool {
	_, result := i.(*textproto.Reader)
	return result
}
func isWriterPointer(i interface{}) bool {
	_, result := i.(*textproto.Writer)
	return result
}
