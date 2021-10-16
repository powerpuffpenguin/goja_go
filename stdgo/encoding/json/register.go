package json

import "encoding/json"

func (f *factory) register() {
	f.Set(`Compact`, json.Compact)
	f.Set(`HTMLEscape`, json.HTMLEscape)
	f.Set(`Indent`, json.Indent)
	f.Set(`Marshal`, json.Marshal)
	f.Set(`MarshalIndent`, json.MarshalIndent)
	f.Set(`Unmarshal`, json.Unmarshal)
	f.Set(`Valid`, json.Valid)

	f.Set(`NewDecoder`, json.NewDecoder)
	f.Set(`NewEncoder`, json.NewEncoder)

	f.Set(`isDecoderPointer`, isDecoderPointer)
	f.Set(`isDelim`, isDelim)
	f.Set(`isEncoderPointer`, isEncoderPointer)
	f.Set(`isInvalidUnmarshalErrorPointer`, isInvalidUnmarshalErrorPointer)
	f.Set(`isMarshaler`, isMarshaler)
	f.Set(`isMarshalerErrorPointer`, isMarshalerErrorPointer)
	f.Set(`isNumber`, isNumber)
	f.Set(`isRawMessage`, isRawMessage)
	f.Set(`isSyntaxErrorPointer`, isSyntaxErrorPointer)
	f.Set(`isToken`, isToken)
	f.Set(`isUnmarshalTypeErrorPointer`, isUnmarshalTypeErrorPointer)
	f.Set(`isUnmarshaler`, isUnmarshaler)
	f.Set(`isUnsupportedTypeErrorPointer`, isUnsupportedTypeErrorPointer)
	f.Set(`isUnsupportedValueErrorPointer`, isUnsupportedValueErrorPointer)
}
func isDecoderPointer(i interface{}) bool {
	_, result := i.(*json.Decoder)
	return result
}
func isDelim(i interface{}) bool {
	_, result := i.(json.Delim)
	return result
}
func isEncoderPointer(i interface{}) bool {
	_, result := i.(*json.Encoder)
	return result
}
func isInvalidUnmarshalErrorPointer(i interface{}) bool {
	_, result := i.(*json.InvalidUnmarshalError)
	return result
}
func isMarshaler(i interface{}) bool {
	_, result := i.(json.Marshaler)
	return result
}
func isMarshalerErrorPointer(i interface{}) bool {
	_, result := i.(*json.MarshalerError)
	return result
}
func isNumber(i interface{}) bool {
	_, result := i.(json.Number)
	return result
}
func isRawMessage(i interface{}) bool {
	_, result := i.(json.RawMessage)
	return result
}
func isSyntaxErrorPointer(i interface{}) bool {
	_, result := i.(*json.SyntaxError)
	return result
}
func isToken(i interface{}) bool {
	_, result := i.(json.Token)
	return result
}
func isUnmarshalTypeErrorPointer(i interface{}) bool {
	_, result := i.(*json.UnmarshalTypeError)
	return result
}
func isUnmarshaler(i interface{}) bool {
	_, result := i.(json.Unmarshaler)
	return result
}
func isUnsupportedTypeErrorPointer(i interface{}) bool {
	_, result := i.(*json.UnsupportedTypeError)
	return result
}
func isUnsupportedValueErrorPointer(i interface{}) bool {
	_, result := i.(*json.UnsupportedValueError)
	return result
}
