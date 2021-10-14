package png

import (
	"image/png"

	"github.com/powerpuffpenguin/goja"
)

func (f *factory) register() {
	f.Set(`Decode`, png.Decode)
	f.Set(`DecodeConfig`, png.DecodeConfig)
	f.Set(`Encode`, png.Encode)

	f.Set(`CompressionLevel`, CompressionLevel)
	f.Accessor(`DefaultCompression`, f.getDefaultCompression, nil)
	f.Accessor(`NoCompression`, f.getNoCompression, nil)
	f.Accessor(`BestSpeed`, f.getBestSpeed, nil)
	f.Accessor(`BestCompression`, f.getBestCompression, nil)

	f.Set(`Encoder`, Encoder)

	f.Set(`isCompressionLevel`, isCompressionLevel)
	f.Set(`isEncoderPointer`, isEncoderPointer)
	f.Set(`isEncoderBufferPointer`, isEncoderBufferPointer)
	f.Set(`isEncoderBufferPool`, isEncoderBufferPool)
	f.Set(`isFormatError`, isFormatError)
	f.Set(`isUnsupportedError`, isUnsupportedError)
}
func isCompressionLevel(i interface{}) bool {
	_, result := i.(png.CompressionLevel)
	return result
}
func isEncoderPointer(i interface{}) bool {
	_, result := i.(*png.Encoder)
	return result
}
func isEncoderBufferPointer(i interface{}) bool {
	_, result := i.(*png.EncoderBuffer)
	return result
}
func isEncoderBufferPool(i interface{}) bool {
	_, result := i.(png.EncoderBufferPool)
	return result
}
func isFormatError(i interface{}) bool {
	_, result := i.(png.FormatError)
	return result
}
func isUnsupportedError(i interface{}) bool {
	_, result := i.(png.UnsupportedError)
	return result
}
func Encoder(level png.CompressionLevel, pool png.EncoderBufferPool) *png.Encoder {
	return &png.Encoder{
		CompressionLevel: level,
		BufferPool:       pool,
	}
}

func CompressionLevel(v int) png.CompressionLevel {
	return png.CompressionLevel(v)
}
func (f *factory) getDefaultCompression(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(png.DefaultCompression)
}
func (f *factory) getNoCompression(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(png.NoCompression)
}
func (f *factory) getBestSpeed(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(png.BestSpeed)
}
func (f *factory) getBestCompression(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(png.BestCompression)
}
