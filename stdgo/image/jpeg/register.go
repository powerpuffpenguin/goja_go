package jpeg

import (
	"image/jpeg"

	"github.com/powerpuffpenguin/goja"
	"github.com/powerpuffpenguin/goja_go/stdgo/builtin"
)

func (f *factory) register() {
	f.Accessor(`DefaultQuality`, f.getDefaultQuality, nil)

	f.Set(`Decode`, jpeg.Decode)
	f.Set(`DecodeConfig`, jpeg.DecodeConfig)
	f.Set(`Encode`, jpeg.Encode)

	f.Set(`Options`, Options)

	f.Set(`isFormatError`, isFormatError)
	f.Set(`isOptionsPointer`, isOptionsPointer)
	f.Set(`isUnsupportedError`, isUnsupportedError)
}
func (f *factory) getDefaultQuality(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(jpeg.DefaultQuality))
}
func Options(quality int) *jpeg.Options {
	return &jpeg.Options{
		Quality: quality,
	}
}
func isFormatError(i interface{}) bool {
	_, result := i.(jpeg.FormatError)
	return result
}
func isOptionsPointer(i interface{}) bool {
	_, result := i.(*jpeg.Options)
	return result
}
func isUnsupportedError(i interface{}) bool {
	_, result := i.(jpeg.UnsupportedError)
	return result
}
