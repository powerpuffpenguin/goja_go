package aes

import (
	"crypto/aes"

	"github.com/powerpuffpenguin/goja"
)

func (f *factory) register() {
	f.Accessor(`BlockSize`, f.getBlockSize, nil)

	f.Set(`NewCipher`, aes.NewCipher)

	f.Set(`isKeySizeError`, isKeySizeError)
}
func (f *factory) getBlockSize(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(aes.BlockSize)
}
func isKeySizeError(i interface{}) bool {
	_, result := i.(aes.KeySizeError)
	return result
}
