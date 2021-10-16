package xml

import (
	"encoding/xml"

	"github.com/powerpuffpenguin/goja"
)

func (f *factory) register() {
	f.Accessor(`Header`, f.getHeader, nil)
	f.Accessor(`HTMLAutoClose`, f.getHTMLAutoClose, nil)
	f.Accessor(`HTMLEntity`, f.getHTMLEntity, nil)

	f.Set(`Escape`, xml.Escape)
	f.Set(`EscapeText`, xml.EscapeText)
	f.Set(`Marshal`, xml.Marshal)
	f.Set(`MarshalIndent`, xml.MarshalIndent)
	f.Set(`Unmarshal`, xml.Unmarshal)

	f.Set(`NewDecoder`, xml.NewDecoder)
	f.Set(`NewTokenDecoder`, xml.NewTokenDecoder)
	f.Set(`NewEncoder`, xml.NewEncoder)
	f.Set(`CopyToken`, xml.CopyToken)

	f.Set(`isAttr`, isAttr)
	f.Set(`isCharData`, isCharData)
	f.Set(`isComment`, isComment)
	f.Set(`isDecoderPointer`, isDecoderPointer)
	f.Set(`isDirective`, isDirective)
	f.Set(`isEncoderPointer`, isEncoderPointer)
	f.Set(`isEndElement`, isEndElement)
	f.Set(`isMarshaler`, isMarshaler)
	f.Set(`isMarshalerAttr`, isMarshalerAttr)
	f.Set(`isName`, isName)
	f.Set(`isProcInst`, isProcInst)
	f.Set(`isStartElement`, isStartElement)
	f.Set(`isStartElementPointer`, isStartElementPointer)
	f.Set(`isSyntaxErrorPointer`, isSyntaxErrorPointer)
	f.Set(`isTagPathErrorPointer`, isTagPathErrorPointer)
	f.Set(`isToken`, isToken)
	f.Set(`isTokenReader`, isTokenReader)
	f.Set(`isUnmarshalError`, isUnmarshalError)
	f.Set(`isUnmarshaler`, isUnmarshaler)
	f.Set(`isUnmarshalerAttr`, isUnmarshalerAttr)
	f.Set(`isUnsupportedTypeErrorPointer`, isUnsupportedTypeErrorPointer)
}
func (f *factory) getHeader(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(xml.Header)
}
func (f *factory) getHTMLAutoClose(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(xml.HTMLAutoClose)
}
func (f *factory) getHTMLEntity(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(xml.HTMLEntity)
}
func isAttr(i interface{}) bool {
	_, result := i.(xml.Attr)
	return result
}
func isCharData(i interface{}) bool {
	_, result := i.(xml.CharData)
	return result
}
func isComment(i interface{}) bool {
	_, result := i.(xml.Comment)
	return result
}
func isDecoderPointer(i interface{}) bool {
	_, result := i.(*xml.Decoder)
	return result
}
func isDirective(i interface{}) bool {
	_, result := i.(xml.Directive)
	return result
}
func isEncoderPointer(i interface{}) bool {
	_, result := i.(*xml.Encoder)
	return result
}
func isEndElement(i interface{}) bool {
	_, result := i.(xml.EndElement)
	return result
}
func isMarshaler(i interface{}) bool {
	_, result := i.(xml.Marshaler)
	return result
}
func isMarshalerAttr(i interface{}) bool {
	_, result := i.(xml.MarshalerAttr)
	return result
}
func isName(i interface{}) bool {
	_, result := i.(xml.Name)
	return result
}
func isProcInst(i interface{}) bool {
	_, result := i.(xml.ProcInst)
	return result
}
func isStartElement(i interface{}) bool {
	_, result := i.(xml.StartElement)
	return result
}
func isStartElementPointer(i interface{}) bool {
	_, result := i.(*xml.StartElement)
	return result
}
func isSyntaxErrorPointer(i interface{}) bool {
	_, result := i.(*xml.SyntaxError)
	return result
}
func isTagPathErrorPointer(i interface{}) bool {
	_, result := i.(*xml.TagPathError)
	return result
}
func isToken(i interface{}) bool {
	_, result := i.(xml.Token)
	return result
}
func isTokenReader(i interface{}) bool {
	_, result := i.(xml.TokenReader)
	return result
}
func isUnmarshalError(i interface{}) bool {
	_, result := i.(xml.UnmarshalError)
	return result
}
func isUnmarshaler(i interface{}) bool {
	_, result := i.(xml.Unmarshaler)
	return result
}
func isUnmarshalerAttr(i interface{}) bool {
	_, result := i.(xml.UnmarshalerAttr)
	return result
}
func isUnsupportedTypeErrorPointer(i interface{}) bool {
	_, result := i.(*xml.UnsupportedTypeError)
	return result
}
