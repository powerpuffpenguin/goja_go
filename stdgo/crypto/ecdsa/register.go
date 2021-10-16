package ecdsa

import "crypto/ecdsa"

func (f *factory) register() {
	f.Set(`Sign`, ecdsa.Sign)
	f.Set(`SignASN1`, ecdsa.SignASN1)
	f.Set(`Verify`, ecdsa.Verify)
	f.Set(`VerifyASN1`, ecdsa.VerifyASN1)

	f.Set(`GenerateKey`, ecdsa.GenerateKey)
	f.Set(`isPrivateKeyPointer`, isPrivateKeyPointer)
	f.Set(`isPublicKeyPointer`, isPublicKeyPointer)
}
func isPrivateKeyPointer(i interface{}) bool {
	_, result := i.(*ecdsa.PrivateKey)
	return result
}
func isPublicKeyPointer(i interface{}) bool {
	_, result := i.(*ecdsa.PublicKey)
	return result
}
