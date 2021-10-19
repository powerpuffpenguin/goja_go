package mail

import (
	"mime"
	"net/mail"

	"github.com/powerpuffpenguin/goja"
)

func (f *factory) register() {
	f.Accessor(`ErrHeaderNotPresent`, f.getErrHeaderNotPresent, nil)

	f.Set(`ParseDate`, mail.ParseDate)

	f.Set(`ParseAddress`, mail.ParseAddress)
	f.Set(`ParseAddressList`, mail.ParseAddressList)
	f.Set(`Header`, Header)
	f.Set(`ReadMessage`, mail.ReadMessage)
	f.Set(`isAddressPointer`, isAddressPointer)
	f.Set(`isAddressParserPointer`, isAddressParserPointer)
	f.Set(`isHeader`, isHeader)
	f.Set(`isMessagePointer`, isMessagePointer)
}
func (f *factory) getErrHeaderNotPresent(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(mail.ErrHeaderNotPresent)
}
func isAddressPointer(i interface{}) bool {
	_, result := i.(*mail.Address)
	return result
}
func isAddressParserPointer(i interface{}) bool {
	_, result := i.(*mail.AddressParser)
	return result
}
func isHeader(i interface{}) bool {
	_, result := i.(mail.Header)
	return result
}
func isMessagePointer(i interface{}) bool {
	_, result := i.(*mail.Message)
	return result
}
func Header() mail.Header {
	return make(mail.Header)
}
func AddressParser(wordDecoder *mime.WordDecoder) *mail.AddressParser {
	return &mail.AddressParser{
		WordDecoder: wordDecoder,
	}
}
