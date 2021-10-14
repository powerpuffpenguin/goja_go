package palette

import (
	"image/color/palette"

	"github.com/powerpuffpenguin/goja"
)

func (f *factory) register() {
	f.Accessor(`Plan9`, f.getPlan9, nil)
	f.Accessor(`WebSafe`, f.getWebSafe, nil)
}
func (f *factory) getPlan9(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(palette.Plan9)
}
func (f *factory) getWebSafe(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(palette.WebSafe)
}
