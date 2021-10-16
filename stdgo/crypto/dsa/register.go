package dsa

import (
	"crypto/dsa"

	"github.com/powerpuffpenguin/goja"
)

func (f *factory) register() {
	f.Accessor(`ErrInvalidPublicKey`, f.getErrInvalidPublicKey, nil)

	f.Set(`GenerateKey`, dsa.GenerateKey)
	f.Set(`GenerateParameters`, dsa.GenerateParameters)
	f.Set(`Sign`, dsa.Sign)
	f.Set(`Verify`, dsa.Verify)

	f.Accessor(`L1024N160`, f.getL1024N160, nil)
	f.Accessor(`L2048N224`, f.getL2048N224, nil)
	f.Accessor(`L2048N256`, f.getL2048N256, nil)
	f.Accessor(`L3072N256`, f.getL3072N256, nil)

	f.Set(`Parameters`, Parameters)
	f.Set(`PrivateKey`, PrivateKey)
	f.Set(`PublicKey`, PublicKey)

	f.Set(`isParametersPointer`, isParametersPointer)
	f.Set(`isPrivateKeyPointer`, isPrivateKeyPointer)
	f.Set(`isPublicKeyPointer`, isPublicKeyPointer)
}
func Parameters() *dsa.Parameters {
	return &dsa.Parameters{}
}
func PrivateKey() *dsa.PrivateKey {
	return &dsa.PrivateKey{}
}
func PublicKey() *dsa.PublicKey {
	return &dsa.PublicKey{}
}

func (f *factory) getErrInvalidPublicKey(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(dsa.ErrInvalidPublicKey)
}
func isParametersPointer(i interface{}) bool {
	_, result := i.(*dsa.Parameters)
	return result
}
func isPrivateKeyPointer(i interface{}) bool {
	_, result := i.(*dsa.PrivateKey)
	return result
}
func isPublicKeyPointer(i interface{}) bool {
	_, result := i.(*dsa.PublicKey)
	return result
}
func (f *factory) getL1024N160(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(dsa.L1024N160)
}
func (f *factory) getL2048N224(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(dsa.L2048N224)
}
func (f *factory) getL2048N256(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(dsa.L2048N256)
}
func (f *factory) getL3072N256(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(dsa.L3072N256)
}
