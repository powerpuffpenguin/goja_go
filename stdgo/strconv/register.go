package strconv

import (
	"strconv"

	"github.com/powerpuffpenguin/goja"
)

func (f *factory) register() {
	f.Accessor(`IntSize`, f.getIntSize, nil)
	f.Accessor(`ErrRange`, f.getErrRange, nil)
	f.Accessor(`ErrSyntax`, f.getErrSyntax, nil)

	f.Set(`AppendBool`, strconv.AppendBool)
	f.Set(`AppendFloat`, strconv.AppendFloat)
	f.Set(`AppendInt`, strconv.AppendInt)
	f.Set(`AppendQuote`, strconv.AppendQuote)
	f.Set(`AppendQuoteRune`, strconv.AppendQuoteRune)
	f.Set(`AppendQuoteRuneToASCII`, strconv.AppendQuoteRuneToASCII)
	f.Set(`AppendQuoteRuneToGraphic`, strconv.AppendQuoteRuneToGraphic)
	f.Set(`AppendQuoteToASCII`, strconv.AppendQuoteToASCII)
	f.Set(`AppendQuoteToGraphic`, strconv.AppendQuoteToGraphic)
	f.Set(`AppendUint`, strconv.AppendUint)
	f.Set(`Atoi`, strconv.Atoi)
	f.Set(`CanBackquote`, strconv.CanBackquote)
	f.Set(`FormatBool`, strconv.FormatBool)
	f.Set(`FormatComplex`, strconv.FormatComplex)
	f.Set(`FormatFloat`, strconv.FormatFloat)
	f.Set(`FormatInt`, strconv.FormatInt)
	f.Set(`FormatUint`, strconv.FormatUint)
	f.Set(`IsGraphic`, strconv.IsGraphic)
	f.Set(`IsPrint`, strconv.IsPrint)
	f.Set(`Itoa`, strconv.Itoa)
	f.Set(`ParseBool`, strconv.ParseBool)
	f.Set(`ParseComplex`, strconv.ParseComplex)
	f.Set(`ParseFloat`, strconv.ParseFloat)
	f.Set(`ParseInt`, strconv.ParseInt)
	f.Set(`ParseUint`, strconv.ParseUint)
	f.Set(`Quote`, strconv.Quote)
	f.Set(`QuoteRune`, strconv.QuoteRune)
	f.Set(`QuoteRuneToASCII`, strconv.QuoteRuneToASCII)
	f.Set(`QuoteRuneToGraphic`, strconv.QuoteRuneToGraphic)
	f.Set(`QuoteToASCII`, strconv.QuoteToASCII)
	f.Set(`QuoteToGraphic`, strconv.QuoteToGraphic)
	f.Set(`Unquote`, strconv.Unquote)
	f.Set(`UnquoteChar`, strconv.UnquoteChar)

	f.Set(`isNumErrorPointer`, isNumErrorPointer)
}
func isNumErrorPointer(i interface{}) bool {
	_, result := i.(*strconv.NumError)
	return result
}
func (f *factory) getIntSize(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(strconv.IntSize)
}
func (f *factory) getErrRange(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(strconv.ErrRange)
}
func (f *factory) getErrSyntax(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(strconv.ErrSyntax)
}
