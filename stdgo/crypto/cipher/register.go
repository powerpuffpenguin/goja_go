package cipher

import "crypto/cipher"

func (f *factory) register() {
	f.Set(`NewGCM`, cipher.NewGCM)
	f.Set(`NewGCMWithNonceSize`, cipher.NewGCMWithNonceSize)
	f.Set(`NewGCMWithTagSize`, cipher.NewGCMWithTagSize)

	f.Set(`NewCBCDecrypter`, cipher.NewCBCDecrypter)
	f.Set(`NewCBCEncrypter`, cipher.NewCBCEncrypter)

	f.Set(`NewCFBDecrypter`, cipher.NewCFBDecrypter)
	f.Set(`NewCFBEncrypter`, cipher.NewCFBEncrypter)
	f.Set(`NewCTR`, cipher.NewCTR)
	f.Set(`NewOFB`, cipher.NewOFB)

	f.Set(`isAEAD`, isAEAD)
	f.Set(`isBlock`, isBlock)
	f.Set(`isBlockMode`, isBlockMode)
	f.Set(`isStream`, isStream)
	f.Set(`isStreamReader`, isStreamReader)
	f.Set(`isStreamWriter`, isStreamWriter)
}
func isAEAD(i interface{}) bool {
	_, result := i.(cipher.AEAD)
	return result
}
func isBlock(i interface{}) bool {
	_, result := i.(cipher.Block)
	return result
}
func isBlockMode(i interface{}) bool {
	_, result := i.(cipher.BlockMode)
	return result
}
func isStream(i interface{}) bool {
	_, result := i.(cipher.Stream)
	return result
}
func isStreamReader(i interface{}) bool {
	_, result := i.(cipher.StreamReader)
	return result
}
func isStreamWriter(i interface{}) bool {
	_, result := i.(cipher.StreamWriter)
	return result
}
