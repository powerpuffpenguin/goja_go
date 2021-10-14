package gif

import (
	"image/draw"
	"image/gif"

	"github.com/powerpuffpenguin/goja"
	"github.com/powerpuffpenguin/goja_go/stdgo/builtin"
)

func (f *factory) register() {
	f.Accessor(`DisposalNone`, f.getDisposalNone, nil)
	f.Accessor(`DisposalBackground`, f.getDisposalBackground, nil)
	f.Accessor(`DisposalPrevious`, f.getDisposalPrevious, nil)

	f.Set(`Decode`, gif.Decode)
	f.Set(`DecodeConfig`, gif.DecodeConfig)
	f.Set(`Encode`, gif.Encode)
	f.Set(`EncodeAll`, gif.EncodeAll)

	f.Set(`DecodeAll`, gif.DecodeAll)
	f.Set(`Options`, Options)

	f.Set(`isGIFPointer`, isGIFPointer)
	f.Set(`isOptionsPointer`, isOptionsPointer)
}
func (f *factory) getDisposalNone(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Uint8(gif.DisposalNone))
}
func (f *factory) getDisposalBackground(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Uint8(gif.DisposalBackground))
}
func (f *factory) getDisposalPrevious(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Uint8(gif.DisposalPrevious))
}
func Options(numColors int, quantizer draw.Quantizer, drawer draw.Drawer) *gif.Options {
	return &gif.Options{
		NumColors: numColors,
		Quantizer: quantizer,
		Drawer:    drawer,
	}
}
func isGIFPointer(i interface{}) bool {
	_, result := i.(*gif.GIF)
	return result
}
func isOptionsPointer(i interface{}) bool {
	_, result := i.(*gif.Options)
	return result
}
