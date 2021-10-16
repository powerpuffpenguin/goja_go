package crypto

import (
	"crypto"

	"github.com/powerpuffpenguin/goja"
)

func (f *factory) register() {
	f.Set(`RegisterHash`, crypto.RegisterHash)

	f.Set(`Hash`, Hash)
	f.Accessor(`MD4`, f.getMD4, nil)
	f.Accessor(`MD5`, f.getMD5, nil)
	f.Accessor(`SHA1`, f.getSHA1, nil)
	f.Accessor(`SHA224`, f.getSHA224, nil)
	f.Accessor(`SHA256`, f.getSHA256, nil)
	f.Accessor(`SHA384`, f.getSHA384, nil)
	f.Accessor(`SHA512`, f.getSHA512, nil)
	f.Accessor(`MD5SHA1`, f.getMD5SHA1, nil)
	f.Accessor(`RIPEMD160`, f.getRIPEMD160, nil)
	f.Accessor(`SHA3_224`, f.getSHA3_224, nil)
	f.Accessor(`SHA3_256`, f.getSHA3_256, nil)
	f.Accessor(`SHA3_384`, f.getSHA3_384, nil)
	f.Accessor(`SHA3_512`, f.getSHA3_512, nil)
	f.Accessor(`SHA512_224`, f.getSHA512_224, nil)
	f.Accessor(`SHA512_256`, f.getSHA512_256, nil)
	f.Accessor(`BLAKE2s_256`, f.getBLAKE2s_256, nil)
	f.Accessor(`BLAKE2b_256`, f.getBLAKE2b_256, nil)
	f.Accessor(`BLAKE2b_384`, f.getBLAKE2b_384, nil)
	f.Accessor(`BLAKE2b_512`, f.getBLAKE2b_512, nil)

	f.Set(`isDecrypter`, isDecrypter)
	f.Set(`isDecrypterOpts`, isDecrypterOpts)
	f.Set(`isHash`, isHash)
	f.Set(`isPrivateKey`, isPrivateKey)
	f.Set(`isPublicKey`, isPublicKey)
	f.Set(`isSigner`, isSigner)
	f.Set(`isSignerOpts`, isSignerOpts)
}
func isDecrypter(i interface{}) bool {
	_, result := i.(crypto.Decrypter)
	return result
}
func isDecrypterOpts(i interface{}) bool {
	_, result := i.(crypto.DecrypterOpts)
	return result
}
func isHash(i interface{}) bool {
	_, result := i.(crypto.Hash)
	return result
}
func isPrivateKey(i interface{}) bool {
	_, result := i.(crypto.PrivateKey)
	return result
}
func isPublicKey(i interface{}) bool {
	_, result := i.(crypto.PublicKey)
	return result
}
func isSigner(i interface{}) bool {
	_, result := i.(crypto.Signer)
	return result
}
func isSignerOpts(i interface{}) bool {
	_, result := i.(crypto.SignerOpts)
	return result
}
func Hash(v uint) crypto.Hash {
	return crypto.Hash(v)
}
func (f *factory) getMD4(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(crypto.MD4)
}
func (f *factory) getMD5(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(crypto.MD5)
}
func (f *factory) getSHA1(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(crypto.SHA1)
}
func (f *factory) getSHA224(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(crypto.SHA224)
}
func (f *factory) getSHA256(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(crypto.SHA256)
}
func (f *factory) getSHA384(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(crypto.SHA384)
}
func (f *factory) getSHA512(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(crypto.SHA512)
}
func (f *factory) getMD5SHA1(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(crypto.MD5SHA1)
}
func (f *factory) getRIPEMD160(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(crypto.RIPEMD160)
}
func (f *factory) getSHA3_224(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(crypto.SHA3_224)
}
func (f *factory) getSHA3_256(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(crypto.SHA3_256)
}
func (f *factory) getSHA3_384(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(crypto.SHA3_384)
}
func (f *factory) getSHA3_512(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(crypto.SHA3_512)
}
func (f *factory) getSHA512_224(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(crypto.SHA512_224)
}
func (f *factory) getSHA512_256(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(crypto.SHA512_256)
}
func (f *factory) getBLAKE2s_256(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(crypto.BLAKE2s_256)
}
func (f *factory) getBLAKE2b_256(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(crypto.BLAKE2b_256)
}
func (f *factory) getBLAKE2b_384(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(crypto.BLAKE2b_384)
}
func (f *factory) getBLAKE2b_512(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(crypto.BLAKE2b_512)
}
