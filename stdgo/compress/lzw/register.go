package lzw

import (
	"compress/lzw"

	"github.com/powerpuffpenguin/goja"
)

func (f *factory) register() {
	f.Set(`NewReader`, lzw.NewReader)
	f.Set(`NewWriter`, lzw.NewWriter)

	f.Accessor(`LSB`, f.getLSB, nil)
	f.Accessor(`MSB`, f.getMSB, nil)

	f.Set(`isOrder`, isOrder)
}

func isOrder(i interface{}) bool {
	_, result := i.(lzw.Order)
	return result
}

func (f *factory) getLSB(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(lzw.LSB)
}
func (f *factory) getMSB(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(lzw.MSB)
}
