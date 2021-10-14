package draw

import (
	"image/draw"

	"github.com/powerpuffpenguin/goja"
)

func (f *factory) register() {
	f.Set(`Draw`, draw.Draw)
	f.Set(`DrawMask`, draw.DrawMask)

	f.Set(`Op`, Op)
	f.Accessor(`Over`, f.getOver, nil)
	f.Accessor(`Src`, f.getSrc, nil)

	f.Set(`isDrawer`, isDrawer)
	f.Set(`isImage`, isImage)
	f.Set(`isQuantizer`, isQuantizer)
}
func Op(v int) draw.Op {
	return draw.Op(v)
}
func (f *factory) getOver(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(draw.Over)
}
func (f *factory) getSrc(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(draw.Src)
}
func isDrawer(i interface{}) bool {
	_, result := i.(draw.Drawer)
	return result
}
func isImage(i interface{}) bool {
	_, result := i.(draw.Image)
	return result
}
func isQuantizer(i interface{}) bool {
	_, result := i.(draw.Quantizer)
	return result
}
