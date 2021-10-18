package pkix

import (
	"crypto/x509/pkix"
)

func (f *factory) register() {
	f.Set(`isAlgorithmIdentifier`, isAlgorithmIdentifier)
	f.Set(`isAttributeTypeAndValue`, isAttributeTypeAndValue)
	f.Set(`isAttributeTypeAndValueSET`, isAttributeTypeAndValueSET)
	f.Set(`isCertificateListPointer`, isCertificateListPointer)
	f.Set(`isExtension`, isExtension)
	f.Set(`isName`, isName)
	f.Set(`isRDNSequence`, isRDNSequence)
	f.Set(`isRelativeDistinguishedNameSET`, isRelativeDistinguishedNameSET)
	f.Set(`isRevokedCertificate`, isRevokedCertificate)
	f.Set(`isTBSCertificateListPointer`, isTBSCertificateListPointer)
}
func isAlgorithmIdentifier(i interface{}) bool {
	_, result := i.(pkix.AlgorithmIdentifier)
	return result
}
func isAttributeTypeAndValue(i interface{}) bool {
	_, result := i.(pkix.AttributeTypeAndValue)
	return result
}
func isAttributeTypeAndValueSET(i interface{}) bool {
	_, result := i.(pkix.AttributeTypeAndValueSET)
	return result
}
func isCertificateListPointer(i interface{}) bool {
	_, result := i.(*pkix.CertificateList)
	return result
}
func isExtension(i interface{}) bool {
	_, result := i.(pkix.Extension)
	return result
}
func isName(i interface{}) bool {
	_, result := i.(pkix.Name)
	return result
}
func isRDNSequence(i interface{}) bool {
	_, result := i.(pkix.RDNSequence)
	return result
}
func isRelativeDistinguishedNameSET(i interface{}) bool {
	_, result := i.(pkix.RelativeDistinguishedNameSET)
	return result
}
func isRevokedCertificate(i interface{}) bool {
	_, result := i.(pkix.RevokedCertificate)
	return result
}
func isTBSCertificateListPointer(i interface{}) bool {
	_, result := i.(*pkix.TBSCertificateList)
	return result
}
