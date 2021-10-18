package x509

import (
	"crypto/x509"

	"github.com/powerpuffpenguin/goja"
)

func (f *factory) register() {
	f.Accessor(`ErrUnsupportedAlgorithm`, f.getErrUnsupportedAlgorithm, nil)
	f.Accessor(`IncorrectPasswordError`, f.getIncorrectPasswordError, nil)

	f.Set(`CreateCertificate`, x509.CreateCertificate)
	f.Set(`CreateCertificateRequest`, x509.CreateCertificateRequest)
	f.Set(`CreateRevocationList`, x509.CreateRevocationList)
	f.Set(`MarshalECPrivateKey`, x509.MarshalECPrivateKey)
	f.Set(`MarshalPKCS1PrivateKey`, x509.MarshalPKCS1PrivateKey)
	f.Set(`MarshalPKCS1PublicKey`, x509.MarshalPKCS1PublicKey)
	f.Set(`MarshalPKCS8PrivateKey`, x509.MarshalPKCS8PrivateKey)
	f.Set(`MarshalPKIXPublicKey`, x509.MarshalPKIXPublicKey)
	f.Set(`ParseCRL`, x509.ParseCRL)
	f.Set(`ParseDERCRL`, x509.ParseDERCRL)
	f.Set(`ParseECPrivateKey`, x509.ParseECPrivateKey)
	f.Set(`ParsePKCS1PrivateKey`, x509.ParsePKCS1PrivateKey)
	f.Set(`ParsePKCS1PublicKey`, x509.ParsePKCS1PublicKey)
	f.Set(`ParsePKCS8PrivateKey`, x509.ParsePKCS8PrivateKey)
	f.Set(`ParsePKIXPublicKey`, x509.ParsePKIXPublicKey)

	f.Set(`NewCertPool`, x509.NewCertPool)
	f.Set(`SystemCertPool`, x509.SystemCertPool)

	f.Set(`ParseCertificate`, x509.ParseCertificate)
	f.Set(`ParseCertificates`, x509.ParseCertificates)

	f.Set(`ParseCertificateRequest`, x509.ParseCertificateRequest)

	f.Set(`ExtKeyUsage`, ExtKeyUsage)
	f.Accessor(`ExtKeyUsageAny`, f.getExtKeyUsageAny, nil)
	f.Accessor(`ExtKeyUsageServerAuth`, f.getExtKeyUsageServerAuth, nil)
	f.Accessor(`ExtKeyUsageClientAuth`, f.getExtKeyUsageClientAuth, nil)
	f.Accessor(`ExtKeyUsageCodeSigning`, f.getExtKeyUsageCodeSigning, nil)
	f.Accessor(`ExtKeyUsageEmailProtection`, f.getExtKeyUsageEmailProtection, nil)
	f.Accessor(`ExtKeyUsageIPSECEndSystem`, f.getExtKeyUsageIPSECEndSystem, nil)
	f.Accessor(`ExtKeyUsageIPSECTunnel`, f.getExtKeyUsageIPSECTunnel, nil)
	f.Accessor(`ExtKeyUsageIPSECUser`, f.getExtKeyUsageIPSECUser, nil)
	f.Accessor(`ExtKeyUsageTimeStamping`, f.getExtKeyUsageTimeStamping, nil)
	f.Accessor(`ExtKeyUsageOCSPSigning`, f.getExtKeyUsageOCSPSigning, nil)
	f.Accessor(`ExtKeyUsageMicrosoftServerGatedCrypto`, f.getExtKeyUsageMicrosoftServerGatedCrypto, nil)
	f.Accessor(`ExtKeyUsageNetscapeServerGatedCrypto`, f.getExtKeyUsageNetscapeServerGatedCrypto, nil)
	f.Accessor(`ExtKeyUsageMicrosoftCommercialCodeSigning`, f.getExtKeyUsageMicrosoftCommercialCodeSigning, nil)
	f.Accessor(`ExtKeyUsageMicrosoftKernelCodeSigning`, f.getExtKeyUsageMicrosoftKernelCodeSigning, nil)

	f.Set(`InvalidReason`, InvalidReason)
	f.Accessor(`NotAuthorizedToSign`, f.getNotAuthorizedToSign, nil)
	f.Accessor(`Expired`, f.getExpired, nil)
	f.Accessor(`CANotAuthorizedForThisName`, f.getCANotAuthorizedForThisName, nil)
	f.Accessor(`TooManyIntermediates`, f.getTooManyIntermediates, nil)
	f.Accessor(`IncompatibleUsage`, f.getIncompatibleUsage, nil)
	f.Accessor(`NameMismatch`, f.getNameMismatch, nil)
	f.Accessor(`NameConstraintsWithoutSANs`, f.getNameConstraintsWithoutSANs, nil)
	f.Accessor(`UnconstrainedName`, f.getUnconstrainedName, nil)
	f.Accessor(`TooManyConstraints`, f.getTooManyConstraints, nil)
	f.Accessor(`CANotAuthorizedForExtKeyUsage`, f.getCANotAuthorizedForExtKeyUsage, nil)

	f.Set(`KeyUsage`, KeyUsage)
	f.Accessor(`KeyUsageDigitalSignature`, f.getKeyUsageDigitalSignature, nil)
	f.Accessor(`KeyUsageContentCommitment`, f.getKeyUsageContentCommitment, nil)
	f.Accessor(`KeyUsageKeyEncipherment`, f.getKeyUsageKeyEncipherment, nil)
	f.Accessor(`KeyUsageDataEncipherment`, f.getKeyUsageDataEncipherment, nil)
	f.Accessor(`KeyUsageKeyAgreement`, f.getKeyUsageKeyAgreement, nil)
	f.Accessor(`KeyUsageCertSign`, f.getKeyUsageCertSign, nil)
	f.Accessor(`KeyUsageCRLSign`, f.getKeyUsageCRLSign, nil)
	f.Accessor(`KeyUsageEncipherOnly`, f.getKeyUsageEncipherOnly, nil)
	f.Accessor(`KeyUsageDecipherOnly`, f.getKeyUsageDecipherOnly, nil)

	f.Set(`PEMCipher`, PEMCipher)
	f.Accessor(`PEMCipherDES`, f.getPEMCipherDES, nil)
	f.Accessor(`PEMCipher3DES`, f.getPEMCipher3DES, nil)
	f.Accessor(`PEMCipherAES128`, f.getPEMCipherAES128, nil)
	f.Accessor(`PEMCipherAES192`, f.getPEMCipherAES192, nil)
	f.Accessor(`PEMCipherAES256`, f.getPEMCipherAES256, nil)

	f.Set(`PublicKeyAlgorithm`, PublicKeyAlgorithm)
	f.Accessor(`UnknownPublicKeyAlgorithm`, f.getUnknownPublicKeyAlgorithm, nil)
	f.Accessor(`RSA`, f.getRSA, nil)
	f.Accessor(`DSA`, f.getDSA, nil)
	f.Accessor(`ECDSA`, f.getECDSA, nil)
	f.Accessor(`Ed25519`, f.getEd25519, nil)

	f.Set(`SignatureAlgorithm`, SignatureAlgorithm)
	f.Accessor(`UnknownSignatureAlgorithm`, f.getUnknownSignatureAlgorithm, nil)
	f.Accessor(`MD2WithRSA`, f.getMD2WithRSA, nil)
	f.Accessor(`MD5WithRSA`, f.getMD5WithRSA, nil)
	f.Accessor(`SHA1WithRSA`, f.getSHA1WithRSA, nil)
	f.Accessor(`SHA256WithRSA`, f.getSHA256WithRSA, nil)
	f.Accessor(`SHA384WithRSA`, f.getSHA384WithRSA, nil)
	f.Accessor(`SHA512WithRSA`, f.getSHA512WithRSA, nil)
	f.Accessor(`DSAWithSHA1`, f.getDSAWithSHA1, nil)
	f.Accessor(`DSAWithSHA256`, f.getDSAWithSHA256, nil)
	f.Accessor(`ECDSAWithSHA1`, f.getECDSAWithSHA1, nil)
	f.Accessor(`ECDSAWithSHA256`, f.getECDSAWithSHA256, nil)
	f.Accessor(`ECDSAWithSHA384`, f.getECDSAWithSHA384, nil)
	f.Accessor(`ECDSAWithSHA512`, f.getECDSAWithSHA512, nil)
	f.Accessor(`SHA256WithRSAPSS`, f.getSHA256WithRSAPSS, nil)
	f.Accessor(`SHA384WithRSAPSS`, f.getSHA384WithRSAPSS, nil)
	f.Accessor(`SHA512WithRSAPSS`, f.getSHA512WithRSAPSS, nil)
	f.Accessor(`PureEd25519`, f.getPureEd25519, nil)

	f.Set(`VerifyOptions`, VerifyOptions)

	f.Set(`isCertPoolPointer`, isCertPoolPointer)
	f.Set(`isCertificatePointer`, isCertificatePointer)
	f.Set(`isCertificateInvalidErrorPointer`, isCertificateInvalidErrorPointer)
	f.Set(`isCertificateRequestPointer`, isCertificateRequestPointer)
	f.Set(`isConstraintViolationError`, isConstraintViolationError)
	f.Set(`isExtKeyUsage`, isExtKeyUsage)
	f.Set(`isHostnameError`, isHostnameError)
	f.Set(`isInsecureAlgorithmError`, isInsecureAlgorithmError)
	f.Set(`isInvalidReason`, isInvalidReason)
	f.Set(`isKeyUsage`, isKeyUsage)
	f.Set(`isPEMCipher`, isPEMCipher)
	f.Set(`isPublicKeyAlgorithm`, isPublicKeyAlgorithm)
	f.Set(`isRevocationListPointer`, isRevocationListPointer)
	f.Set(`isSignatureAlgorithm`, isSignatureAlgorithm)
	f.Set(`isSystemRootsError`, isSystemRootsError)
	f.Set(`isUnhandledCriticalExtension`, isUnhandledCriticalExtension)
	f.Set(`isUnknownAuthorityError`, isUnknownAuthorityError)
	f.Set(`isVerifyOptions`, isVerifyOptions)
}
func (f *factory) getErrUnsupportedAlgorithm(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(x509.ErrUnsupportedAlgorithm)
}
func (f *factory) getIncorrectPasswordError(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(x509.IncorrectPasswordError)
}
func isCertPoolPointer(i interface{}) bool {
	_, result := i.(*x509.CertPool)
	return result
}
func isCertificatePointer(i interface{}) bool {
	_, result := i.(*x509.Certificate)
	return result
}
func isCertificateInvalidErrorPointer(i interface{}) bool {
	_, result := i.(*x509.CertificateInvalidError)
	return result
}
func isCertificateRequestPointer(i interface{}) bool {
	_, result := i.(*x509.CertificateRequest)
	return result
}
func isConstraintViolationError(i interface{}) bool {
	_, result := i.(x509.ConstraintViolationError)
	return result
}
func isExtKeyUsage(i interface{}) bool {
	_, result := i.(x509.ExtKeyUsage)
	return result
}
func isHostnameError(i interface{}) bool {
	_, result := i.(x509.HostnameError)
	return result
}
func isInsecureAlgorithmError(i interface{}) bool {
	_, result := i.(x509.InsecureAlgorithmError)
	return result
}
func isInvalidReason(i interface{}) bool {
	_, result := i.(x509.InvalidReason)
	return result
}
func isKeyUsage(i interface{}) bool {
	_, result := i.(x509.KeyUsage)
	return result
}
func isPEMCipher(i interface{}) bool {
	_, result := i.(x509.PEMCipher)
	return result
}
func isPublicKeyAlgorithm(i interface{}) bool {
	_, result := i.(x509.PublicKeyAlgorithm)
	return result
}
func isRevocationListPointer(i interface{}) bool {
	_, result := i.(*x509.RevocationList)
	return result
}
func isSignatureAlgorithm(i interface{}) bool {
	_, result := i.(x509.SignatureAlgorithm)
	return result
}
func isSystemRootsError(i interface{}) bool {
	_, result := i.(x509.SystemRootsError)
	return result
}
func isUnhandledCriticalExtension(i interface{}) bool {
	_, result := i.(x509.UnhandledCriticalExtension)
	return result
}
func isUnknownAuthorityError(i interface{}) bool {
	_, result := i.(x509.UnknownAuthorityError)
	return result
}
func isVerifyOptions(i interface{}) bool {
	_, result := i.(x509.VerifyOptions)
	return result
}
func VerifyOptions() x509.VerifyOptions {
	return x509.VerifyOptions{}
}
func (f *factory) getUnknownSignatureAlgorithm(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(x509.UnknownSignatureAlgorithm)
}
func (f *factory) getMD2WithRSA(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(x509.MD2WithRSA)
}
func (f *factory) getMD5WithRSA(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(x509.MD5WithRSA)
}
func (f *factory) getSHA1WithRSA(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(x509.SHA1WithRSA)
}
func (f *factory) getSHA256WithRSA(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(x509.SHA256WithRSA)
}
func (f *factory) getSHA384WithRSA(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(x509.SHA384WithRSA)
}
func (f *factory) getSHA512WithRSA(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(x509.SHA512WithRSA)
}
func (f *factory) getDSAWithSHA1(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(x509.DSAWithSHA1)
}
func (f *factory) getDSAWithSHA256(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(x509.DSAWithSHA256)
}
func (f *factory) getECDSAWithSHA1(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(x509.ECDSAWithSHA1)
}
func (f *factory) getECDSAWithSHA256(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(x509.ECDSAWithSHA256)
}
func (f *factory) getECDSAWithSHA384(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(x509.ECDSAWithSHA384)
}
func (f *factory) getECDSAWithSHA512(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(x509.ECDSAWithSHA512)
}
func (f *factory) getSHA256WithRSAPSS(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(x509.SHA256WithRSAPSS)
}
func (f *factory) getSHA384WithRSAPSS(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(x509.SHA384WithRSAPSS)
}
func (f *factory) getSHA512WithRSAPSS(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(x509.SHA512WithRSAPSS)
}
func (f *factory) getPureEd25519(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(x509.PureEd25519)
}
func SignatureAlgorithm(v int) x509.SignatureAlgorithm {
	return x509.SignatureAlgorithm(v)
}
func (f *factory) getUnknownPublicKeyAlgorithm(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(x509.UnknownPublicKeyAlgorithm)
}
func (f *factory) getRSA(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(x509.RSA)
}
func (f *factory) getDSA(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(x509.DSA)
}
func (f *factory) getECDSA(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(x509.ECDSA)
}
func (f *factory) getEd25519(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(x509.Ed25519)
}
func PublicKeyAlgorithm(v int) x509.PublicKeyAlgorithm {
	return x509.PublicKeyAlgorithm(v)
}
func (f *factory) getPEMCipherDES(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(x509.PEMCipherDES)
}
func (f *factory) getPEMCipher3DES(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(x509.PEMCipher3DES)
}
func (f *factory) getPEMCipherAES128(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(x509.PEMCipherAES128)
}
func (f *factory) getPEMCipherAES192(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(x509.PEMCipherAES192)
}
func (f *factory) getPEMCipherAES256(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(x509.PEMCipherAES256)
}
func PEMCipher(v int) x509.PEMCipher {
	return x509.PEMCipher(v)
}
func (f *factory) getKeyUsageDigitalSignature(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(x509.KeyUsageDigitalSignature)
}
func (f *factory) getKeyUsageContentCommitment(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(x509.KeyUsageContentCommitment)
}
func (f *factory) getKeyUsageKeyEncipherment(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(x509.KeyUsageKeyEncipherment)
}
func (f *factory) getKeyUsageDataEncipherment(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(x509.KeyUsageDataEncipherment)
}
func (f *factory) getKeyUsageKeyAgreement(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(x509.KeyUsageKeyAgreement)
}
func (f *factory) getKeyUsageCertSign(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(x509.KeyUsageCertSign)
}
func (f *factory) getKeyUsageCRLSign(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(x509.KeyUsageCRLSign)
}
func (f *factory) getKeyUsageEncipherOnly(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(x509.KeyUsageEncipherOnly)
}
func (f *factory) getKeyUsageDecipherOnly(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(x509.KeyUsageDecipherOnly)
}
func KeyUsage(v int) x509.KeyUsage {
	return x509.KeyUsage(v)
}
func (f *factory) getNotAuthorizedToSign(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(x509.NotAuthorizedToSign)
}
func (f *factory) getExpired(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(x509.Expired)
}
func (f *factory) getCANotAuthorizedForThisName(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(x509.CANotAuthorizedForThisName)
}
func (f *factory) getTooManyIntermediates(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(x509.TooManyIntermediates)
}
func (f *factory) getIncompatibleUsage(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(x509.IncompatibleUsage)
}
func (f *factory) getNameMismatch(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(x509.NameMismatch)
}
func (f *factory) getNameConstraintsWithoutSANs(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(x509.NameConstraintsWithoutSANs)
}
func (f *factory) getUnconstrainedName(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(x509.UnconstrainedName)
}
func (f *factory) getTooManyConstraints(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(x509.TooManyConstraints)
}
func (f *factory) getCANotAuthorizedForExtKeyUsage(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(x509.CANotAuthorizedForExtKeyUsage)
}
func InvalidReason(v int) x509.InvalidReason {
	return x509.InvalidReason(v)
}
func (f *factory) getExtKeyUsageAny(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(x509.ExtKeyUsageAny)
}
func (f *factory) getExtKeyUsageServerAuth(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(x509.ExtKeyUsageServerAuth)
}
func (f *factory) getExtKeyUsageClientAuth(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(x509.ExtKeyUsageClientAuth)
}
func (f *factory) getExtKeyUsageCodeSigning(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(x509.ExtKeyUsageCodeSigning)
}
func (f *factory) getExtKeyUsageEmailProtection(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(x509.ExtKeyUsageEmailProtection)
}
func (f *factory) getExtKeyUsageIPSECEndSystem(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(x509.ExtKeyUsageIPSECEndSystem)
}
func (f *factory) getExtKeyUsageIPSECTunnel(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(x509.ExtKeyUsageIPSECTunnel)
}
func (f *factory) getExtKeyUsageIPSECUser(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(x509.ExtKeyUsageIPSECUser)
}
func (f *factory) getExtKeyUsageTimeStamping(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(x509.ExtKeyUsageTimeStamping)
}
func (f *factory) getExtKeyUsageOCSPSigning(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(x509.ExtKeyUsageOCSPSigning)
}
func (f *factory) getExtKeyUsageMicrosoftServerGatedCrypto(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(x509.ExtKeyUsageMicrosoftServerGatedCrypto)
}
func (f *factory) getExtKeyUsageNetscapeServerGatedCrypto(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(x509.ExtKeyUsageNetscapeServerGatedCrypto)
}
func (f *factory) getExtKeyUsageMicrosoftCommercialCodeSigning(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(x509.ExtKeyUsageMicrosoftCommercialCodeSigning)
}
func (f *factory) getExtKeyUsageMicrosoftKernelCodeSigning(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(x509.ExtKeyUsageMicrosoftKernelCodeSigning)
}
func ExtKeyUsage(v int) x509.ExtKeyUsage {
	return x509.ExtKeyUsage(v)
}
