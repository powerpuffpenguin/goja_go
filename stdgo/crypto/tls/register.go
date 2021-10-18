package tls

import (
	"crypto/tls"
	"net"

	"github.com/powerpuffpenguin/goja"
	"github.com/powerpuffpenguin/goja_go/stdgo/builtin"
)

func (f *factory) register() {
	f.Accessor(`TLS_RSA_WITH_RC4_128_SHA`, f.getTLS_RSA_WITH_RC4_128_SHA, nil)
	f.Accessor(`TLS_RSA_WITH_3DES_EDE_CBC_SHA`, f.getTLS_RSA_WITH_3DES_EDE_CBC_SHA, nil)
	f.Accessor(`TLS_RSA_WITH_AES_128_CBC_SHA`, f.getTLS_RSA_WITH_AES_128_CBC_SHA, nil)
	f.Accessor(`TLS_RSA_WITH_AES_256_CBC_SHA`, f.getTLS_RSA_WITH_AES_256_CBC_SHA, nil)
	f.Accessor(`TLS_RSA_WITH_AES_128_CBC_SHA256`, f.getTLS_RSA_WITH_AES_128_CBC_SHA256, nil)
	f.Accessor(`TLS_RSA_WITH_AES_128_GCM_SHA256`, f.getTLS_RSA_WITH_AES_128_GCM_SHA256, nil)
	f.Accessor(`TLS_RSA_WITH_AES_256_GCM_SHA384`, f.getTLS_RSA_WITH_AES_256_GCM_SHA384, nil)
	f.Accessor(`TLS_ECDHE_ECDSA_WITH_RC4_128_SHA`, f.getTLS_ECDHE_ECDSA_WITH_RC4_128_SHA, nil)
	f.Accessor(`TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA`, f.getTLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA, nil)
	f.Accessor(`TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA`, f.getTLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA, nil)
	f.Accessor(`TLS_ECDHE_RSA_WITH_RC4_128_SHA`, f.getTLS_ECDHE_RSA_WITH_RC4_128_SHA, nil)
	f.Accessor(`TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA`, f.getTLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA, nil)
	f.Accessor(`TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA`, f.getTLS_ECDHE_RSA_WITH_AES_128_CBC_SHA, nil)
	f.Accessor(`TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA`, f.getTLS_ECDHE_RSA_WITH_AES_256_CBC_SHA, nil)
	f.Accessor(`TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256`, f.getTLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256, nil)
	f.Accessor(`TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256`, f.getTLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256, nil)
	f.Accessor(`TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256`, f.getTLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256, nil)
	f.Accessor(`TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256`, f.getTLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256, nil)
	f.Accessor(`TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384`, f.getTLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384, nil)
	f.Accessor(`TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384`, f.getTLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384, nil)
	f.Accessor(`TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256`, f.getTLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256, nil)
	f.Accessor(`TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256`, f.getTLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256, nil)
	f.Accessor(`TLS_AES_128_GCM_SHA256`, f.getTLS_AES_128_GCM_SHA256, nil)
	f.Accessor(`TLS_AES_256_GCM_SHA384`, f.getTLS_AES_256_GCM_SHA384, nil)
	f.Accessor(`TLS_CHACHA20_POLY1305_SHA256`, f.getTLS_CHACHA20_POLY1305_SHA256, nil)
	f.Accessor(`TLS_FALLBACK_SCSV`, f.getTLS_FALLBACK_SCSV, nil)
	f.Accessor(`TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305`, f.getTLS_ECDHE_RSA_WITH_CHACHA20_POLY1305, nil)
	f.Accessor(`TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305`, f.getTLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305, nil)

	f.Accessor(`VersionTLS10`, f.getVersionTLS10, nil)
	f.Accessor(`VersionTLS11`, f.getVersionTLS11, nil)
	f.Accessor(`VersionTLS12`, f.getVersionTLS12, nil)
	f.Accessor(`VersionTLS13`, f.getVersionTLS13, nil)
	f.Accessor(`VersionSSL30`, f.getVersionSSL30, nil)

	f.Set(`CipherSuiteName`, tls.CipherSuiteName)
	f.Set(`Listen`, tls.Listen)
	f.Set(`NewListener`, tls.NewListener)

	f.Set(`LoadX509KeyPair`, tls.LoadX509KeyPair)
	f.Set(`X509KeyPair`, tls.X509KeyPair)
	f.Set(`CipherSuites`, tls.CipherSuites)
	f.Set(`InsecureCipherSuites`, tls.InsecureCipherSuites)
	f.Set(`ClientAuthType`, ClientAuthType)
	f.Accessor(`NoClientCert`, f.getNoClientCert, nil)
	f.Accessor(`RequestClientCert`, f.getRequestClientCert, nil)
	f.Accessor(`RequireAnyClientCert`, f.getRequireAnyClientCert, nil)
	f.Accessor(`VerifyClientCertIfGiven`, f.getVerifyClientCertIfGiven, nil)
	f.Accessor(`RequireAndVerifyClientCert`, f.getRequireAndVerifyClientCert, nil)
	f.Set(`NewLRUClientSessionCache`, tls.NewLRUClientSessionCache)
	f.Set(`Config`, Config)
	f.Set(`CurveID`, CurveID)
	f.Accessor(`CurveP256`, f.getCurveP256, nil)
	f.Accessor(`CurveP384`, f.getCurveP384, nil)
	f.Accessor(`CurveP521`, f.getCurveP521, nil)
	f.Accessor(`X25519`, f.getX25519, nil)
	f.Set(`Dialer`, Dialer)
	f.Set(`RenegotiationSupport`, RenegotiationSupport)
	f.Accessor(`RenegotiateNever`, f.getRenegotiateNever, nil)
	f.Accessor(`RenegotiateOnceAsClient`, f.getRenegotiateOnceAsClient, nil)
	f.Accessor(`RenegotiateFreelyAsClient`, f.getRenegotiateFreelyAsClient, nil)
	f.Set(`SignatureScheme`, SignatureScheme)
	f.Accessor(`PKCS1WithSHA256`, f.getPKCS1WithSHA256, nil)
	f.Accessor(`PKCS1WithSHA384`, f.getPKCS1WithSHA384, nil)
	f.Accessor(`PKCS1WithSHA512`, f.getPKCS1WithSHA512, nil)
	f.Accessor(`PSSWithSHA256`, f.getPSSWithSHA256, nil)
	f.Accessor(`PSSWithSHA384`, f.getPSSWithSHA384, nil)
	f.Accessor(`PSSWithSHA512`, f.getPSSWithSHA512, nil)
	f.Accessor(`ECDSAWithP256AndSHA256`, f.getECDSAWithP256AndSHA256, nil)
	f.Accessor(`ECDSAWithP384AndSHA384`, f.getECDSAWithP384AndSHA384, nil)
	f.Accessor(`ECDSAWithP521AndSHA512`, f.getECDSAWithP521AndSHA512, nil)
	f.Accessor(`Ed25519`, f.getEd25519, nil)
	f.Accessor(`PKCS1WithSHA1`, f.getPKCS1WithSHA1, nil)
	f.Accessor(`ECDSAWithSHA1`, f.getECDSAWithSHA1, nil)
	f.Set(`isCertificate`, isCertificate)
	f.Set(`isCertificatePointer`, isCertificatePointer)
	f.Set(`isCertificateRequestInfo`, isCertificateRequestInfo)
	f.Set(`isCertificateRequestInfoPointer`, isCertificateRequestInfoPointer)
	f.Set(`isCipherSuitePointer`, isCipherSuitePointer)
	f.Set(`isClientAuthType`, isClientAuthType)
	f.Set(`isClientHelloInfoPointer`, isClientHelloInfoPointer)
	f.Set(`isClientSessionStatePointer`, isClientSessionStatePointer)
	f.Set(`isConfigPointer`, isConfigPointer)
	f.Set(`isConnectionState`, isConnectionState)
	f.Set(`isCurveID`, isCurveID)
	f.Set(`isDialerPointer`, isDialerPointer)
	f.Set(`isRecordHeaderError`, isRecordHeaderError)
	f.Set(`isRenegotiationSupport`, isRenegotiationSupport)
	f.Set(`isSignatureScheme`, isSignatureScheme)
}
func (f *factory) getVersionTLS10(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(tls.VersionTLS10))
}
func (f *factory) getVersionTLS11(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(tls.VersionTLS11))
}
func (f *factory) getVersionTLS12(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(tls.VersionTLS12))
}
func (f *factory) getVersionTLS13(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(tls.VersionTLS13))
}
func (f *factory) getVersionSSL30(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(tls.VersionSSL30))
}
func (f *factory) getTLS_RSA_WITH_RC4_128_SHA(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Uint16(tls.TLS_RSA_WITH_RC4_128_SHA))
}
func (f *factory) getTLS_RSA_WITH_3DES_EDE_CBC_SHA(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Uint16(tls.TLS_RSA_WITH_3DES_EDE_CBC_SHA))
}
func (f *factory) getTLS_RSA_WITH_AES_128_CBC_SHA(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Uint16(tls.TLS_RSA_WITH_AES_128_CBC_SHA))
}
func (f *factory) getTLS_RSA_WITH_AES_256_CBC_SHA(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Uint16(tls.TLS_RSA_WITH_AES_256_CBC_SHA))
}
func (f *factory) getTLS_RSA_WITH_AES_128_CBC_SHA256(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Uint16(tls.TLS_RSA_WITH_AES_128_CBC_SHA256))
}
func (f *factory) getTLS_RSA_WITH_AES_128_GCM_SHA256(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Uint16(tls.TLS_RSA_WITH_AES_128_GCM_SHA256))
}
func (f *factory) getTLS_RSA_WITH_AES_256_GCM_SHA384(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Uint16(tls.TLS_RSA_WITH_AES_256_GCM_SHA384))
}
func (f *factory) getTLS_ECDHE_ECDSA_WITH_RC4_128_SHA(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Uint16(tls.TLS_ECDHE_ECDSA_WITH_RC4_128_SHA))
}
func (f *factory) getTLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Uint16(tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA))
}
func (f *factory) getTLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Uint16(tls.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA))
}
func (f *factory) getTLS_ECDHE_RSA_WITH_RC4_128_SHA(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Uint16(tls.TLS_ECDHE_RSA_WITH_RC4_128_SHA))
}
func (f *factory) getTLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Uint16(tls.TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA))
}
func (f *factory) getTLS_ECDHE_RSA_WITH_AES_128_CBC_SHA(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Uint16(tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA))
}
func (f *factory) getTLS_ECDHE_RSA_WITH_AES_256_CBC_SHA(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Uint16(tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA))
}
func (f *factory) getTLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Uint16(tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256))
}
func (f *factory) getTLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Uint16(tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256))
}
func (f *factory) getTLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Uint16(tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256))
}
func (f *factory) getTLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Uint16(tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256))
}
func (f *factory) getTLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Uint16(tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384))
}
func (f *factory) getTLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Uint16(tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384))
}
func (f *factory) getTLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Uint16(tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256))
}
func (f *factory) getTLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Uint16(tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256))
}
func (f *factory) getTLS_AES_128_GCM_SHA256(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Uint16(tls.TLS_AES_128_GCM_SHA256))
}
func (f *factory) getTLS_AES_256_GCM_SHA384(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Uint16(tls.TLS_AES_256_GCM_SHA384))
}
func (f *factory) getTLS_CHACHA20_POLY1305_SHA256(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Uint16(tls.TLS_CHACHA20_POLY1305_SHA256))
}
func (f *factory) getTLS_FALLBACK_SCSV(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Uint16(tls.TLS_FALLBACK_SCSV))
}
func (f *factory) getTLS_ECDHE_RSA_WITH_CHACHA20_POLY1305(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Uint16(tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305))
}
func (f *factory) getTLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Uint16(tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305))
}
func isCertificate(i interface{}) bool {
	_, result := i.(tls.Certificate)
	return result
}
func isCertificatePointer(i interface{}) bool {
	_, result := i.(*tls.Certificate)
	return result
}
func isCertificateRequestInfo(i interface{}) bool {
	_, result := i.(tls.CertificateRequestInfo)
	return result
}
func isCertificateRequestInfoPointer(i interface{}) bool {
	_, result := i.(*tls.CertificateRequestInfo)
	return result
}
func isCipherSuitePointer(i interface{}) bool {
	_, result := i.(*tls.CipherSuite)
	return result
}
func isClientAuthType(i interface{}) bool {
	_, result := i.(tls.ClientAuthType)
	return result
}
func isClientHelloInfoPointer(i interface{}) bool {
	_, result := i.(*tls.ClientHelloInfo)
	return result
}
func isClientSessionStatePointer(i interface{}) bool {
	_, result := i.(*tls.ClientSessionState)
	return result
}
func isConfigPointer(i interface{}) bool {
	_, result := i.(*tls.Config)
	return result
}
func isConnectionState(i interface{}) bool {
	_, result := i.(tls.ConnectionState)
	return result
}
func isCurveID(i interface{}) bool {
	_, result := i.(tls.CurveID)
	return result
}
func isDialerPointer(i interface{}) bool {
	_, result := i.(*tls.Dialer)
	return result
}
func isRecordHeaderError(i interface{}) bool {
	_, result := i.(tls.RecordHeaderError)
	return result
}
func isRenegotiationSupport(i interface{}) bool {
	_, result := i.(tls.RenegotiationSupport)
	return result
}
func isSignatureScheme(i interface{}) bool {
	_, result := i.(tls.SignatureScheme)
	return result
}
func (f *factory) getPKCS1WithSHA256(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(tls.PKCS1WithSHA256)
}
func (f *factory) getPKCS1WithSHA384(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(tls.PKCS1WithSHA384)
}
func (f *factory) getPKCS1WithSHA512(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(tls.PKCS1WithSHA512)
}
func (f *factory) getPSSWithSHA256(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(tls.PSSWithSHA256)
}
func (f *factory) getPSSWithSHA384(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(tls.PSSWithSHA384)
}
func (f *factory) getPSSWithSHA512(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(tls.PSSWithSHA512)
}
func (f *factory) getECDSAWithP256AndSHA256(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(tls.ECDSAWithP256AndSHA256)
}
func (f *factory) getECDSAWithP384AndSHA384(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(tls.ECDSAWithP384AndSHA384)
}
func (f *factory) getECDSAWithP521AndSHA512(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(tls.ECDSAWithP521AndSHA512)
}
func (f *factory) getEd25519(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(tls.Ed25519)
}
func (f *factory) getPKCS1WithSHA1(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(tls.PKCS1WithSHA1)
}
func (f *factory) getECDSAWithSHA1(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(tls.ECDSAWithSHA1)
}
func SignatureScheme(v uint16) tls.SignatureScheme {
	return tls.SignatureScheme(v)
}
func (f *factory) getRenegotiateNever(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(tls.RenegotiateNever)
}
func (f *factory) getRenegotiateOnceAsClient(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(tls.RenegotiateOnceAsClient)
}
func (f *factory) getRenegotiateFreelyAsClient(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(tls.RenegotiateFreelyAsClient)
}
func RenegotiationSupport(v int) tls.RenegotiationSupport {
	return tls.RenegotiationSupport(v)
}
func Dialer(dialer *net.Dialer, config *tls.Config) *tls.Dialer {
	return &tls.Dialer{
		NetDialer: dialer,
		Config:    config,
	}
}
func (f *factory) getCurveP256(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(tls.CurveP256)
}
func (f *factory) getCurveP384(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(tls.CurveP384)
}
func (f *factory) getCurveP521(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(tls.CurveP521)
}
func (f *factory) getX25519(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(tls.X25519)
}
func CurveID(v uint16) tls.CurveID {
	return tls.CurveID(v)
}
func Config() *tls.Config {
	return &tls.Config{}
}
func (f *factory) getNoClientCert(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(tls.NoClientCert)
}
func (f *factory) getRequestClientCert(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(tls.RequestClientCert)
}
func (f *factory) getRequireAnyClientCert(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(tls.RequireAnyClientCert)
}
func (f *factory) getVerifyClientCertIfGiven(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(tls.VerifyClientCertIfGiven)
}
func (f *factory) getRequireAndVerifyClientCert(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(tls.RequireAndVerifyClientCert)
}
func ClientAuthType(v int) tls.ClientAuthType {
	return tls.ClientAuthType(v)
}
