package rsa

import (
	"crypto"
	"crypto/rsa"
	"math/big"

	"github.com/powerpuffpenguin/goja"
	"github.com/powerpuffpenguin/goja_go/stdgo/builtin"
)

func (f *factory) register() {
	f.Accessor(`PSSSaltLengthAuto`, f.getPSSSaltLengthAuto, nil)
	f.Accessor(`PSSSaltLengthEqualsHash`, f.getPSSSaltLengthEqualsHash, nil)

	f.Accessor(`ErrDecryption`, f.getErrDecryption, nil)
	f.Accessor(`ErrMessageTooLong`, f.getErrMessageTooLong, nil)
	f.Accessor(`ErrVerification`, f.getErrVerification, nil)

	f.Set(`DecryptOAEP`, rsa.DecryptOAEP)
	f.Set(`DecryptPKCS1v15`, rsa.DecryptPKCS1v15)
	f.Set(`DecryptPKCS1v15SessionKey`, rsa.DecryptPKCS1v15SessionKey)
	f.Set(`EncryptOAEP`, rsa.EncryptOAEP)
	f.Set(`EncryptPKCS1v15`, rsa.EncryptPKCS1v15)
	f.Set(`SignPKCS1v15`, rsa.SignPKCS1v15)
	f.Set(`SignPSS`, rsa.SignPSS)
	f.Set(`VerifyPKCS1v15`, rsa.VerifyPKCS1v15)
	f.Set(`VerifyPSS`, rsa.VerifyPSS)

	f.Set(`CRTValue`, CRTValue)
	f.Set(`OAEPOptions`, OAEPOptions)
	f.Set(`PKCS1v15DecryptOptions`, PKCS1v15DecryptOptions)
	f.Set(`PSSOptions`, PSSOptions)
	f.Set(`GenerateKey`, rsa.GenerateKey)
	f.Set(`GenerateMultiPrimeKey`, rsa.GenerateMultiPrimeKey)
	f.Set(`PublicKey`, PublicKey)
}
func (f *factory) getErrDecryption(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(rsa.ErrDecryption)
}
func (f *factory) getErrMessageTooLong(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(rsa.ErrMessageTooLong)
}
func (f *factory) getErrVerification(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(rsa.ErrVerification)
}
func (f *factory) getPSSSaltLengthAuto(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(rsa.PSSSaltLengthAuto))
}
func (f *factory) getPSSSaltLengthEqualsHash(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(rsa.PSSSaltLengthEqualsHash))
}
func CRTValue(exp, coeff, r *big.Int) rsa.CRTValue {
	return rsa.CRTValue{
		Exp:   exp,
		Coeff: coeff,
		R:     r,
	}
}
func OAEPOptions(hash crypto.Hash, label []byte) *rsa.OAEPOptions {
	return &rsa.OAEPOptions{
		Hash:  hash,
		Label: label,
	}
}
func PKCS1v15DecryptOptions(sessionKeyLen int) *rsa.PKCS1v15DecryptOptions {
	return &rsa.PKCS1v15DecryptOptions{
		SessionKeyLen: sessionKeyLen,
	}
}
func PSSOptions(saltLength int, hash crypto.Hash) *rsa.PSSOptions {
	return &rsa.PSSOptions{
		SaltLength: saltLength,
		Hash:       hash,
	}
}
func PublicKey(n *big.Int, e int) *rsa.PublicKey {
	return &rsa.PublicKey{
		N: n,
		E: e,
	}
}
