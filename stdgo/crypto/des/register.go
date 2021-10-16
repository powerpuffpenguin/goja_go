package des

import (
	"crypto/des"

	"github.com/powerpuffpenguin/goja"
)

func (f *factory) register() {
	f.Accessor(`BlockSize`, f.getBlockSize, nil)

	f.Set(`NewCipher`, des.NewCipher)
	f.Set(`NewTripleDESCipher`, des.NewTripleDESCipher)

	f.Set(`isKeySizeError`, isKeySizeError)
}
func (f *factory) getBlockSize(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(des.BlockSize)
}
func isKeySizeError(i interface{}) bool {
	_, result := i.(des.KeySizeError)
	return result
}
