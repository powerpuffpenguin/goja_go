package color

import (
	"image/color"

	"github.com/powerpuffpenguin/goja"
)

func (f *factory) register() {
	f.Accessor(`Black`, f.getBlack, nil)
	f.Accessor(`White`, f.getWhite, nil)
	f.Accessor(`Transparent`, f.getTransparent, nil)
	f.Accessor(`Opaque`, f.getOpaque, nil)

	f.Set(`CMYKToRGB`, color.CMYKToRGB)
	f.Set(`RGBToCMYK`, color.RGBToCMYK)
	f.Set(`RGBToYCbCr`, color.RGBToYCbCr)
	f.Set(`YCbCrToRGB`, color.YCbCrToRGB)

	f.Set(`Alpha`, Alpha)
	f.Set(`Alpha16`, Alpha16)
	f.Set(`CMYK`, CMYK)
	f.Set(`Gray`, Gray)
	f.Set(`Gray16`, Gray16)
	f.Set(`ModelFunc`, color.ModelFunc)
	f.Accessor(`RGBAModel`, f.getRGBAModel, nil)
	f.Accessor(`RGBA64Model`, f.getRGBA64Model, nil)
	f.Accessor(`NRGBAModel`, f.getNRGBAModel, nil)
	f.Accessor(`NRGBA64Model`, f.getNRGBA64Model, nil)
	f.Accessor(`AlphaModel`, f.getAlphaModel, nil)
	f.Accessor(`Alpha16Model`, f.getAlpha16Model, nil)
	f.Accessor(`GrayModel`, f.getGrayModel, nil)
	f.Accessor(`Gray16Model`, f.getGray16Model, nil)
	f.Accessor(`CMYKModel`, f.getCMYKModel, nil)
	f.Accessor(`NYCbCrAModel`, f.getNYCbCrAModel, nil)
	f.Accessor(`YCbCrModel`, f.getYCbCrModel, nil)
	f.Set(`NRGBA`, NRGBA)
	f.Set(`NRGBA64`, NRGBA64)
	f.Set(`NYCbCrA`, NYCbCrA)
	f.Set(`Palette`, Palette)
	f.Set(`RGBA`, RGBA)
	f.Set(`RGBA64`, RGBA64)
	f.Set(`YCbCr`, YCbCr)

	f.Set(`isAlpha`, isAlpha)
	f.Set(`isAlpha16`, isAlpha16)
	f.Set(`isCMYK`, isCMYK)
	f.Set(`isColor`, isColor)
	f.Set(`isGray`, isGray)
	f.Set(`isGray16`, isGray16)
	f.Set(`isModel`, isModel)
	f.Set(`isNRGBA`, isNRGBA)
	f.Set(`isNRGBA64`, isNRGBA64)
	f.Set(`isNYCbCrA`, isNYCbCrA)
	f.Set(`isPalette`, isPalette)
	f.Set(`isRGBA`, isRGBA)
	f.Set(`isRGBA64`, isRGBA64)
	f.Set(`isYCbCr`, isYCbCr)
}
func (f *factory) getBlack(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(color.Black)
}
func (f *factory) getWhite(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(color.White)
}
func (f *factory) getTransparent(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(color.Transparent)
}
func (f *factory) getOpaque(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(color.Opaque)
}
func isAlpha(i interface{}) bool {
	_, result := i.(color.Alpha)
	return result
}
func isAlpha16(i interface{}) bool {
	_, result := i.(color.Alpha16)
	return result
}
func isCMYK(i interface{}) bool {
	_, result := i.(color.CMYK)
	return result
}
func isColor(i interface{}) bool {
	_, result := i.(color.Color)
	return result
}
func isGray(i interface{}) bool {
	_, result := i.(color.Gray)
	return result
}
func isGray16(i interface{}) bool {
	_, result := i.(color.Gray16)
	return result
}
func isModel(i interface{}) bool {
	_, result := i.(color.Model)
	return result
}
func isNRGBA(i interface{}) bool {
	_, result := i.(color.NRGBA)
	return result
}
func isNRGBA64(i interface{}) bool {
	_, result := i.(color.NRGBA64)
	return result
}
func isNYCbCrA(i interface{}) bool {
	_, result := i.(color.NYCbCrA)
	return result
}
func isPalette(i interface{}) bool {
	_, result := i.(color.Palette)
	return result
}
func isRGBA(i interface{}) bool {
	_, result := i.(color.RGBA)
	return result
}
func isRGBA64(i interface{}) bool {
	_, result := i.(color.RGBA64)
	return result
}
func isYCbCr(i interface{}) bool {
	_, result := i.(color.YCbCr)
	return result
}
func (f *factory) getRGBAModel(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(color.RGBAModel)
}
func (f *factory) getRGBA64Model(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(color.RGBA64Model)
}
func (f *factory) getNRGBAModel(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(color.NRGBAModel)
}
func (f *factory) getNRGBA64Model(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(color.NRGBA64Model)
}
func (f *factory) getAlphaModel(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(color.AlphaModel)
}
func (f *factory) getAlpha16Model(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(color.Alpha16Model)
}
func (f *factory) getGrayModel(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(color.GrayModel)
}
func (f *factory) getGray16Model(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(color.Gray16Model)
}
func (f *factory) getCMYKModel(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(color.CMYKModel)
}
func (f *factory) getNYCbCrAModel(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(color.NYCbCrAModel)
}
func (f *factory) getYCbCrModel(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(color.YCbCrModel)
}
func Alpha(a uint8) color.Alpha {
	return color.Alpha{a}
}
func Alpha16(a uint16) color.Alpha16 {
	return color.Alpha16{a}
}
func CMYK(c, m, y, k uint8) color.CMYK {
	return color.CMYK{c, m, y, k}
}
func Gray(y uint8) color.Gray {
	return color.Gray{y}
}
func Gray16(y uint16) color.Gray16 {
	return color.Gray16{y}
}
func NRGBA(r, g, b, a uint8) color.NRGBA {
	return color.NRGBA{r, g, b, a}
}
func NRGBA64(r, g, b, a uint16) color.NRGBA64 {
	return color.NRGBA64{r, g, b, a}
}
func NYCbCrA(y, cb, cr, a uint8) color.NYCbCrA {
	return color.NYCbCrA{
		YCbCr(y, cb, cr),
		a,
	}
}
func Palette(v []color.Color) color.Palette {
	return color.Palette(v)
}
func RGBA(r, g, b, a uint8) color.RGBA {
	return color.RGBA{r, g, b, a}
}
func RGBA64(r, g, b, a uint16) color.RGBA64 {
	return color.RGBA64{r, g, b, a}
}
func YCbCr(y, cb, cr uint8) color.YCbCr {
	return color.YCbCr{y, cb, cr}
}
