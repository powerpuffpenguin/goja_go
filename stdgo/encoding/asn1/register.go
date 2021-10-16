package asn1

import (
	"encoding/asn1"

	"github.com/powerpuffpenguin/goja"
	"github.com/powerpuffpenguin/goja_go/stdgo/builtin"
)

func (f *factory) register() {
	f.Accessor(`TagBoolean`, f.getTagBoolean, nil)
	f.Accessor(`TagInteger`, f.getTagInteger, nil)
	f.Accessor(`TagBitString`, f.getTagBitString, nil)
	f.Accessor(`TagOctetString`, f.getTagOctetString, nil)
	f.Accessor(`TagNull`, f.getTagNull, nil)
	f.Accessor(`TagOID`, f.getTagOID, nil)
	f.Accessor(`TagEnum`, f.getTagEnum, nil)
	f.Accessor(`TagUTF8String`, f.getTagUTF8String, nil)
	f.Accessor(`TagSequence`, f.getTagSequence, nil)
	f.Accessor(`TagSet`, f.getTagSet, nil)
	f.Accessor(`TagNumericString`, f.getTagNumericString, nil)
	f.Accessor(`TagPrintableString`, f.getTagPrintableString, nil)
	f.Accessor(`TagT61String`, f.getTagT61String, nil)
	f.Accessor(`TagIA5String`, f.getTagIA5String, nil)
	f.Accessor(`TagUTCTime`, f.getTagUTCTime, nil)
	f.Accessor(`TagGeneralizedTime`, f.getTagGeneralizedTime, nil)
	f.Accessor(`TagGeneralString`, f.getTagGeneralString, nil)
	f.Accessor(`TagBMPString`, f.getTagBMPString, nil)

	f.Accessor(`ClassUniversal`, f.getClassUniversal, nil)
	f.Accessor(`ClassApplication`, f.getClassApplication, nil)
	f.Accessor(`ClassContextSpecific`, f.getClassContextSpecific, nil)
	f.Accessor(`ClassPrivate`, f.getClassPrivate, nil)

	f.Accessor(`NullBytes`, f.getNullBytes, nil)
	f.Accessor(`NullRawValue`, f.getNullRawValue, nil)

	f.Set(`Marshal`, asn1.Marshal)
	f.Set(`MarshalWithParams`, asn1.MarshalWithParams)
	f.Set(`Unmarshal`, asn1.Unmarshal)
	f.Set(`UnmarshalWithParams`, asn1.UnmarshalWithParams)

	f.Set(`BitString`, BitString)
	f.Set(`Enumeravted`, Enumeravted)
	f.Set(`Flag`, Flag)
	f.Set(`ObjectIdentifier`, ObjectIdentifier)
	f.Set(`RawContent`, RawContent)
	f.Set(`RawValue`, RawValue)

	f.Set(`isBitString`, isBitString)
	f.Set(`isEnumerated`, isEnumerated)
	f.Set(`isFlag`, isFlag)
	f.Set(`isObjectIdentifier`, isObjectIdentifier)
	f.Set(`isRawContent`, isRawContent)
	f.Set(`isRawValue`, isRawValue)
	f.Set(`isStructuralError`, isStructuralError)
	f.Set(`isSyntaxError`, isSyntaxError)
}
func (f *factory) getNullRawValue(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(asn1.NullRawValue)
}
func (f *factory) getNullBytes(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Uint8Slice(asn1.NullBytes))
}
func (f *factory) getClassUniversal(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(asn1.ClassUniversal))
}
func (f *factory) getClassApplication(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(asn1.ClassApplication))
}
func (f *factory) getClassContextSpecific(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(asn1.ClassContextSpecific))
}
func (f *factory) getClassPrivate(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(asn1.ClassPrivate))
}
func (f *factory) getTagBoolean(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(asn1.TagBoolean))
}
func (f *factory) getTagInteger(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(asn1.TagInteger))
}
func (f *factory) getTagBitString(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(asn1.TagBitString))
}
func (f *factory) getTagOctetString(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(asn1.TagOctetString))
}
func (f *factory) getTagNull(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(asn1.TagNull))
}
func (f *factory) getTagOID(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(asn1.TagOID))
}
func (f *factory) getTagEnum(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(asn1.TagEnum))
}
func (f *factory) getTagUTF8String(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(asn1.TagUTF8String))
}
func (f *factory) getTagSequence(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(asn1.TagSequence))
}
func (f *factory) getTagSet(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(asn1.TagSet))
}
func (f *factory) getTagNumericString(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(asn1.TagNumericString))
}
func (f *factory) getTagPrintableString(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(asn1.TagPrintableString))
}
func (f *factory) getTagT61String(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(asn1.TagT61String))
}
func (f *factory) getTagIA5String(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(asn1.TagIA5String))
}
func (f *factory) getTagUTCTime(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(asn1.TagUTCTime))
}
func (f *factory) getTagGeneralizedTime(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(asn1.TagGeneralizedTime))
}
func (f *factory) getTagGeneralString(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(asn1.TagGeneralString))
}
func (f *factory) getTagBMPString(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(asn1.TagBMPString))
}
func BitString(bytes []byte, length int) asn1.BitString {
	return asn1.BitString{
		Bytes:     bytes,
		BitLength: length,
	}
}
func Enumeravted(v int) asn1.Enumerated {
	return asn1.Enumerated(v)
}
func Flag(v bool) asn1.Flag {
	return asn1.Flag(v)
}
func ObjectIdentifier(v []int) asn1.ObjectIdentifier {
	return asn1.ObjectIdentifier(v)
}
func RawContent(v []byte) asn1.RawContent {
	return asn1.RawContent(v)
}
func RawValue() asn1.RawValue {
	return asn1.RawValue{}
}
func isBitString(i interface{}) bool {
	_, result := i.(asn1.BitString)
	return result
}
func isEnumerated(i interface{}) bool {
	_, result := i.(asn1.Enumerated)
	return result
}
func isFlag(i interface{}) bool {
	_, result := i.(asn1.Flag)
	return result
}
func isObjectIdentifier(i interface{}) bool {
	_, result := i.(asn1.ObjectIdentifier)
	return result
}
func isRawContent(i interface{}) bool {
	_, result := i.(asn1.RawContent)
	return result
}
func isRawValue(i interface{}) bool {
	_, result := i.(asn1.RawValue)
	return result
}
func isStructuralError(i interface{}) bool {
	_, result := i.(asn1.StructuralError)
	return result
}
func isSyntaxError(i interface{}) bool {
	_, result := i.(asn1.SyntaxError)
	return result
}
