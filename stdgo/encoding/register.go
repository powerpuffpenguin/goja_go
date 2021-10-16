package encoding

import "encoding"

func (f *factory) register() {
	f.Set(`isBinaryMarshaler`, isBinaryMarshaler)
	f.Set(`isBinaryUnmarshaler`, isBinaryUnmarshaler)
	f.Set(`isTextMarshaler`, isTextMarshaler)
	f.Set(`isTextUnmarshaler`, isTextUnmarshaler)
}

func isBinaryMarshaler(i interface{}) bool {
	_, result := i.(encoding.BinaryMarshaler)
	return result
}
func isBinaryUnmarshaler(i interface{}) bool {
	_, result := i.(encoding.BinaryUnmarshaler)
	return result
}
func isTextMarshaler(i interface{}) bool {
	_, result := i.(encoding.TextMarshaler)
	return result
}
func isTextUnmarshaler(i interface{}) bool {
	_, result := i.(encoding.TextUnmarshaler)
	return result
}
