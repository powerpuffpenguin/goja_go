package rc4

import "crypto/rc4"

func (f *factory) register() {
	f.Set(`NewCipher`, rc4.NewCipher)

	f.Set(`isCipherPointer`, isCipherPointer)
	f.Set(`isKeySizeError`, isKeySizeError)
}
func isCipherPointer(i interface{}) bool {
	_, result := i.(*rc4.Cipher)
	return result
}
func isKeySizeError(i interface{}) bool {
	_, result := i.(rc4.KeySizeError)
	return result
}
