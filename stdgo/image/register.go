package image

import (
	"image"

	"github.com/powerpuffpenguin/goja"
)

func (f *factory) register() {
	f.Accessor(`Black`, f.getBlack, nil)
	f.Accessor(`White`, f.getWhite, nil)
	f.Accessor(`Transparent`, f.getTransparent, nil)
	f.Accessor(`Opaque`, f.getOpaque, nil)

	f.Accessor(`ErrFormat`, f.getErrFormat, nil)

	f.Set(`RegisterFormat`, image.RegisterFormat)

	f.Set(`NewAlpha`, image.NewAlpha)
	f.Set(`NewAlpha16`, image.NewAlpha16)
	f.Set(`NewCMYK`, image.NewCMYK)
	f.Set(`DecodeConfig`, image.DecodeConfig)
	f.Set(`NewGray`, image.NewGray)
	f.Set(`NewGray16`, image.NewGray16)
	f.Set(`Decode`, image.Decode)
	f.Set(`NewNRGBA`, image.NewNRGBA)
	f.Set(`NewNRGBA64`, image.NewNRGBA64)
	f.Set(`NewNYCbCrA`, image.NewNYCbCrA)
	f.Set(`NewPaletted`, image.NewPaletted)
	f.Set(`Pt`, image.Pt)
	f.Set(`NewRGBA`, image.NewRGBA)
	f.Set(`NewRGBA64`, image.NewRGBA64)
	f.Set(`Rect`, image.Rect)
	f.Set(`NewUniform`, image.NewUniform)
	f.Set(`NewYCbCr`, image.NewYCbCr)
	f.Set(`YCbCrSubsampleRatio`, YCbCrSubsampleRatio)
	f.Accessor(`YCbCrSubsampleRatio444`, f.getYCbCrSubsampleRatio444, nil)
	f.Accessor(`YCbCrSubsampleRatio422`, f.getYCbCrSubsampleRatio422, nil)
	f.Accessor(`YCbCrSubsampleRatio420`, f.getYCbCrSubsampleRatio420, nil)
	f.Accessor(`YCbCrSubsampleRatio440`, f.getYCbCrSubsampleRatio440, nil)
	f.Accessor(`YCbCrSubsampleRatio411`, f.getYCbCrSubsampleRatio411, nil)
	f.Accessor(`YCbCrSubsampleRatio410`, f.getYCbCrSubsampleRatio410, nil)
	f.Set(`isAlphaPointer`, isAlphaPointer)
	f.Set(`isAlpha16Pointer`, isAlpha16Pointer)
	f.Set(`isCMYKPointer`, isCMYKPointer)
	f.Set(`isConfig`, isConfig)
	f.Set(`isGrayPointer`, isGrayPointer)
	f.Set(`isGray16Pointer`, isGray16Pointer)
	f.Set(`isImage`, isImage)
	f.Set(`isNRGBAPointer`, isNRGBAPointer)
	f.Set(`isNRGBA64Pointer`, isNRGBA64Pointer)
	f.Set(`isNYCbCrAPointer`, isNYCbCrAPointer)
	f.Set(`isPalettedPointer`, isPalettedPointer)
	f.Set(`isPalettedImage`, isPalettedImage)
	f.Set(`isPoint`, isPoint)
	f.Set(`isRGBAPointer`, isRGBAPointer)
	f.Set(`isRGBA64Pointer`, isRGBA64Pointer)
	f.Set(`isRectangle`, isRectangle)
	f.Set(`isUniformPointer`, isUniformPointer)
	f.Set(`isYCbCrPointer`, isYCbCrPointer)
	f.Set(`isYCbCrSubsampleRatio`, isYCbCrSubsampleRatio)
}
func isAlphaPointer(i interface{}) bool {
	_, result := i.(*image.Alpha)
	return result
}
func isAlpha16Pointer(i interface{}) bool {
	_, result := i.(*image.Alpha16)
	return result
}
func isCMYKPointer(i interface{}) bool {
	_, result := i.(*image.CMYK)
	return result
}
func isConfig(i interface{}) bool {
	_, result := i.(image.Config)
	return result
}
func isGrayPointer(i interface{}) bool {
	_, result := i.(*image.Gray)
	return result
}
func isGray16Pointer(i interface{}) bool {
	_, result := i.(*image.Gray16)
	return result
}
func isImage(i interface{}) bool {
	_, result := i.(image.Image)
	return result
}
func isNRGBAPointer(i interface{}) bool {
	_, result := i.(*image.NRGBA)
	return result
}
func isNRGBA64Pointer(i interface{}) bool {
	_, result := i.(*image.NRGBA64)
	return result
}
func isNYCbCrAPointer(i interface{}) bool {
	_, result := i.(*image.NYCbCrA)
	return result
}
func isPalettedPointer(i interface{}) bool {
	_, result := i.(*image.Paletted)
	return result
}
func isPalettedImage(i interface{}) bool {
	_, result := i.(image.PalettedImage)
	return result
}
func isPoint(i interface{}) bool {
	_, result := i.(image.Point)
	return result
}
func isRGBAPointer(i interface{}) bool {
	_, result := i.(*image.RGBA)
	return result
}
func isRGBA64Pointer(i interface{}) bool {
	_, result := i.(*image.RGBA64)
	return result
}
func isRectangle(i interface{}) bool {
	_, result := i.(image.Rectangle)
	return result
}
func isUniformPointer(i interface{}) bool {
	_, result := i.(*image.Uniform)
	return result
}
func isYCbCrPointer(i interface{}) bool {
	_, result := i.(*image.YCbCr)
	return result
}
func isYCbCrSubsampleRatio(i interface{}) bool {
	_, result := i.(image.YCbCrSubsampleRatio)
	return result
}
func YCbCrSubsampleRatio(v int) image.YCbCrSubsampleRatio {
	return image.YCbCrSubsampleRatio(v)
}
func (f *factory) getYCbCrSubsampleRatio444(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(image.YCbCrSubsampleRatio444)
}
func (f *factory) getYCbCrSubsampleRatio422(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(image.YCbCrSubsampleRatio422)
}
func (f *factory) getYCbCrSubsampleRatio420(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(image.YCbCrSubsampleRatio420)
}
func (f *factory) getYCbCrSubsampleRatio440(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(image.YCbCrSubsampleRatio440)
}
func (f *factory) getYCbCrSubsampleRatio411(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(image.YCbCrSubsampleRatio411)
}
func (f *factory) getYCbCrSubsampleRatio410(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(image.YCbCrSubsampleRatio410)
}
func (f *factory) getBlack(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(image.Black)
}
func (f *factory) getWhite(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(image.White)
}
func (f *factory) getTransparent(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(image.Transparent)
}
func (f *factory) getOpaque(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(image.Opaque)
}
func (f *factory) getErrFormat(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(image.ErrFormat)
}
