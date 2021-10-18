package ed25519

import (
	"crypto/ed25519"

	"github.com/powerpuffpenguin/goja"
)

func (f *factory) register() {
	f.Accessor(`PublicKeySize`, f.getPublicKeySize, nil)
	f.Accessor(`PrivateKeySize`, f.getPrivateKeySize, nil)
	f.Accessor(`SignatureSize`, f.getSignatureSize, nil)
	f.Accessor(`SeedSize`, f.getSeedSize, nil)

	f.Set(`GenerateKey`, ed25519.GenerateKey)
	f.Set(`Sign`, ed25519.Sign)
	f.Set(`Verify`, ed25519.Verify)

	f.Set(`NewKeyFromSeed`, ed25519.NewKeyFromSeed)

	f.Set(`isPrivateKey`, isPrivateKey)
	f.Set(`isPublicKey`, isPublicKey)
}
func (f *factory) getPublicKeySize(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(ed25519.PublicKeySize)
}
func (f *factory) getPrivateKeySize(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(ed25519.PrivateKeySize)
}
func (f *factory) getSignatureSize(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(ed25519.SignatureSize)
}
func (f *factory) getSeedSize(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(ed25519.SeedSize)
}
func isPrivateKey(i interface{}) bool {
	_, result := i.(ed25519.PrivateKey)
	return result
}
func isPublicKey(i interface{}) bool {
	_, result := i.(ed25519.PublicKey)
	return result
}
