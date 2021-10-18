declare module "stdgo/crypto/x509" {
    import {
        Float32, Float64,
        Int64, Int32, Int16, Int8, Int,
        Uint64, Uint32, Uint16, Uint8, Uint,
        Number, NumberLike,
        Byte, Bytes, Rune, Runes,
        Float32Slice, Float64Slice,
        Int64Slice, Int32Slice, Int16Slice, Int8Slice, IntSlice,
        Uint64Slice, Uint32Slice, Uint16Slice, Uint8Slice, UintSlice,
        Error,
        ReadChannel, WriteChannel, Channel,
        Slice, Map,
        Uintptr, Native,
    } from "stdgo/builtin";
    import * as big from "stdgo/math/big";
    import * as crypto from "stdgo/crypto";
    import * as ecdsa from "stdgo/crypto/ecdsa";
    import * as rsa from "stdgo/crypto/rsa";
    import * as pkix from "stdgo/crypto/x509/pkix";
    import * as asn1 from "stdgo/encoding/asn1";
    import * as net from "stdgo/net";
    import * as url from "stdgo/net/url";
    import * as io from "stdgo/io";
    import * as time from "stdgo/time";

    const ErrUnsupportedAlgorithm: Error
    const IncorrectPasswordError: Error

    function CreateCertificate(rand: io.Reader, template: CertificatePointer, parent: Certificate, pub: any, priv: any): Bytes
    function CreateCertificateRequest(rand: io.Reader, template: CertificateRequestPointer, priv: any): Bytes
    function CreateRevocationList(rand: io.Reader, template: RevocationListPointer, issuer: CertificatePointer, priv: crypto.Signer): Bytes
    function MarshalECPrivateKey(key: ecdsa.PrivateKeyPointer): Bytes
    function MarshalPKCS1PrivateKey(key: rsa.PrivateKeyPointer): Bytes
    function MarshalPKCS1PublicKey(key: rsa.PublicKeyPointer): Bytes
    function MarshalPKCS8PrivateKey(key: any): Bytes
    function MarshalPKIXPublicKey(pub: any): Bytes
    function ParseCRL(crlBytes: Bytes): pkix.CertificateListPointer
    function ParseDERCRL(derBytes: Bytes): pkix.CertificateListPointer
    function ParseECPrivateKey(der: Bytes): ecdsa.PrivateKeyPointer
    function ParsePKCS1PrivateKey(der: Bytes): rsa.PrivateKeyPointer
    function ParsePKCS1PublicKey(der: Bytes): rsa.PublicKeyPointer
    function ParsePKCS8PrivateKey(der: Bytes): any
    function ParsePKIXPublicKey(derBytes: Bytes): any
    
    function NewCertPool(): CertPoolPointer
    function SystemCertPool(): CertPoolPointer
    interface CertPoolPointer extends Native {
        readonly __CertPoolPointer: CertPoolPointer

        AddCert(cert: CertificatePointer): void
        AppendCertsFromPEM(pemCerts: Bytes): boolean
        Subjects(): Slice<Bytes>
    }

    function ParseCertificate(der: Bytes): CertificatePointer
    function ParseCertificates(der: Bytes): Array<CertificatePointer>
    interface CertificatePointer extends Native {
        readonly __CertificatePointer: CertificatePointer
        Raw: Bytes               // Complete ASN.1 DER content (certificate, signature algorithm and signature).
        RawTBSCertificate: Bytes // Certificate part of raw ASN.1 DER content.
        RawSubjectPublicKeyInfo: Bytes // DER encoded SubjectPublicKeyInfo.
        RawSubject: Bytes // DER encoded Subject
        RawIssuer: Bytes // DER encoded Issuer

        Signature: Bytes
        SignatureAlgorithm: SignatureAlgorithm

        PublicKeyAlgorithm: PublicKeyAlgorithm
        PublicKey: any

        Version: Int
        SerialNumber: big.IntPointer
        Issuer: pkix.Name
        Subject: pkix.Name
        NotBefore: time.Time
        NotAfter: time.Time // Validity bounds.
        KeyUsage: KeyUsage

        // Extensions contains raw X.509 extensions. When parsing certificates,
        // this can be used to extract non-critical extensions that are not
        // parsed by this package. When marshaling certificates, the Extensions
        // field is ignored, see ExtraExtensions.
        Extensions: pkix.Extension

        // ExtraExtensions contains extensions to be copied, raw, into any
        // marshaled certificates. Values override any extensions that would
        // otherwise be produced based on the other fields. The ExtraExtensions
        // field is not populated when parsing certificates, see Extensions.
        ExtraExtensions: pkix.Extension

        // UnhandledCriticalExtensions contains a list of extension IDs that
        // were not (fully) processed when parsing. Verify will fail if this
        // slice is non-empty, unless verification is delegated to an OS
        // library which understands all the critical extensions.
        //
        // Users can access these extensions using Extensions and can remove
        // elements from this slice if they believe that they have been
        // handled.
        UnhandledCriticalExtensions: Array<asn1.ObjectIdentifier>

        ExtKeyUsage: Array<ExtKeyUsage>           // Sequence of extended key usages.
        UnknownExtKeyUsage: Array<asn1.ObjectIdentifier>// Encountered extended key usages unknown to this package.

        // BasicConstraintsValid indicates whether IsCA, MaxPathLen,
        // and MaxPathLenZero are valid.
        BasicConstraintsValid: boolean
        IsCA: boolean

        // MaxPathLen and MaxPathLenZero indicate the presence and
        // value of the BasicConstraints' "pathLenConstraint".
        //
        // When parsing a certificate, a positive non-zero MaxPathLen
        // means that the field was specified, -1 means it was unset,
        // and MaxPathLenZero being true mean that the field was
        // explicitly set to zero. The case of MaxPathLen==0 with MaxPathLenZero==false
        // should be treated equivalent to -1 (unset).
        //
        // When generating a certificate, an unset pathLenConstraint
        // can be requested with either MaxPathLen == -1 or using the
        // zero value for both MaxPathLen and MaxPathLenZero.
        MaxPathLen: Int
        // MaxPathLenZero indicates that BasicConstraintsValid==true
        // and MaxPathLen==0 should be interpreted as an actual
        // maximum path length of zero. Otherwise, that combination is
        // interpreted as MaxPathLen not being set.
        MaxPathLenZero: boolean

        SubjectKeyId: Bytes
        AuthorityKeyId: Bytes

        // RFC 5280, 4.2.2.1 (Authority Information Access)
        OCSPServer: Array<string>
        IssuingCertificateURL: Array<string>

        // Subject Alternate Name values. (Note that these values may not be valid
        // if invalid values were contained within a parsed certificate. For
        // example, an element of DNSNames may not be a valid DNS domain name.)
        DNSNames: Array<string>
        EmailAddresses: Array<string>
        IPAddresses: Array<net.IP>
        URIs: Array<url.URLPointer>

        // Name constraints
        PermittedDNSDomainsCritical: boolean // if true then the name constraints are marked critical.
        PermittedDNSDomains: Array<string>
        ExcludedDNSDomains: Array<string>
        PermittedIPRanges: Array<net.IPNetPointer>
        ExcludedIPRanges: Array<net.IPNetPointer>
        PermittedEmailAddresses: Array<string>
        ExcludedEmailAddresses: Array<string>
        PermittedURIDomains: Array<string>
        ExcludedURIDomains: Array<string>

        // CRL Distribution Points
        CRLDistributionPoints: Array<string>

        PolicyIdentifiers: Array<asn1.ObjectIdentifier>

        CheckCRLSignature(crl: pkix.CertificateListPointer): void
        CheckSignature(algo: SignatureAlgorithm, signed: Bytes, signature: Bytes): void
        CheckSignatureFrom(parent: CertificatePointer): void
        CreateCRL(rand: io.Reader, priv: any, revokedCerts: Array<pkix.RevokedCertificate>, now: time.Time, expiry: time.Time): Bytes
        Equal(other: CertificatePointer): boolean
        Verify(opts: VerifyOptions): Slice<Slice<Certificate>>
        VerifyHostname(h: string): void
    }
    interface CertificateInvalidErrorPointer extends Error {
        readonly __CertificateInvalidErrorPointer: CertificateInvalidErrorPointer

        Cert: CertificatePointer
        Reason: InvalidReason
        Detail: string
    }

    function ParseCertificateRequest(asn1Data: Bytes): CertificateRequestPointer
    interface CertificateRequestPointer extends Native {
        readonly __CertificateRequestPointer: CertificateRequestPointer
        Raw: Bytes // Complete ASN.1 DER content (CSR, signature algorithm and signature).
        RawTBSCertificateRequest: Bytes // Certificate request info part of raw ASN.1 DER content.
        RawSubjectPublicKeyInfo: Bytes // DER encoded SubjectPublicKeyInfo.
        RawSubject: Bytes // DER encoded Subject.

        Version: Int
        Signature: Bytes
        SignatureAlgorithm: SignatureAlgorithm

        PublicKeyAlgorithm: PublicKeyAlgorithm
        PublicKey: any

        Subject: pkix.Name

        // Attributes contains the CSR attributes that can parse as
        // pkix.AttributeTypeAndValueSET.
        //
        // Deprecated: Use Extensions and ExtraExtensions instead for parsing and
        // generating the requestedExtensions attribute.
        Attributes: Array<pkix.AttributeTypeAndValueSET>

        // Extensions contains all requested extensions, in raw form. When parsing
        // CSRs, this can be used to extract extensions that are not parsed by this
        // package.
        Extensions: Array<pkix.Extension>

        // ExtraExtensions contains extensions to be copied, raw, into any CSR
        // marshaled by CreateCertificateRequest. Values override any extensions
        // that would otherwise be produced based on the other fields but are
        // overridden by any extensions specified in Attributes.
        //
        // The ExtraExtensions field is not populated by ParseCertificateRequest,
        // see Extensions instead.
        ExtraExtensions: Array<pkix.Extension>

        // Subject Alternate Name values.
        DNSNames: Array<string>
        EmailAddresses: Array<string>
        IPAddresses: Array<net.IP>
        URIs: Array<url.URLPointer>

        CheckSignature(): void
    }
    interface ConstraintViolationError extends Error {
        readonly __ConstraintViolationError: ConstraintViolationError
    }

    function ExtKeyUsage(v: Int | NumberLike): ExtKeyUsage
    interface ExtKeyUsage extends Number {
        readonly __ExtKeyUsage: ExtKeyUsage
    }
    const ExtKeyUsageAny = ExtKeyUsage(0)
    const ExtKeyUsageServerAuth = ExtKeyUsage(1)
    const ExtKeyUsageClientAuth = ExtKeyUsage(2)
    const ExtKeyUsageCodeSigning = ExtKeyUsage(3)
    const ExtKeyUsageEmailProtection = ExtKeyUsage(4)
    const ExtKeyUsageIPSECEndSystem = ExtKeyUsage(5)
    const ExtKeyUsageIPSECTunnel = ExtKeyUsage(6)
    const ExtKeyUsageIPSECUser = ExtKeyUsage(7)
    const ExtKeyUsageTimeStamping = ExtKeyUsage(8)
    const ExtKeyUsageOCSPSigning = ExtKeyUsage(9)
    const ExtKeyUsageMicrosoftServerGatedCrypto = ExtKeyUsage(10)
    const ExtKeyUsageNetscapeServerGatedCrypto = ExtKeyUsage(11)
    const ExtKeyUsageMicrosoftCommercialCodeSigning = ExtKeyUsage(12)
    const ExtKeyUsageMicrosoftKernelCodeSigning = ExtKeyUsage(13)

    interface HostnameError extends Error {
        readonly __HostnameError: HostnameError
        Certificate: CertificatePointer
        Host: string
    }
    interface InsecureAlgorithmError extends Error {
        readonly __InsecureAlgorithmError: InsecureAlgorithmError
    }

    function InvalidReason(v: Int | NumberLike): InvalidReason
    interface InvalidReason extends Number {
        readonly __InvalidReason: InvalidReason
    }
    const NotAuthorizedToSign = InvalidReason(0)
    const Expired = InvalidReason(1)
    const CANotAuthorizedForThisName = InvalidReason(2)
    const TooManyIntermediates = InvalidReason(3)
    const IncompatibleUsage = InvalidReason(4)
    const NameMismatch = InvalidReason(5)
    const NameConstraintsWithoutSANs = InvalidReason(6)
    const UnconstrainedName = InvalidReason(7)
    const TooManyConstraints = InvalidReason(8)
    const CANotAuthorizedForExtKeyUsage = InvalidReason(9)

    function KeyUsage(v: Int | NumberLike): KeyUsage
    interface KeyUsage extends Number {
        readonly __KeyUsage: KeyUsage
    }
    const KeyUsageDigitalSignature = KeyUsage(1 << 0)
    const KeyUsageContentCommitment = KeyUsage(1 << 1)
    const KeyUsageKeyEncipherment = KeyUsage(1 << 2)
    const KeyUsageDataEncipherment = KeyUsage(1 << 3)
    const KeyUsageKeyAgreement = KeyUsage(1 << 4)
    const KeyUsageCertSign = KeyUsage(1 << 5)
    const KeyUsageCRLSign = KeyUsage(1 << 6)
    const KeyUsageEncipherOnly = KeyUsage(1 << 7)
    const KeyUsageDecipherOnly = KeyUsage(1 << 8)

    function PEMCipher(v: Int | NumberLike): PEMCipher
    interface PEMCipher extends Number {
        readonly __PEMCipher: PEMCipher
    }
    const PEMCipherDES: PEMCipher
    const PEMCipher3DES: PEMCipher
    const PEMCipherAES128: PEMCipher
    const PEMCipherAES192: PEMCipher
    const PEMCipherAES256: PEMCipher

    function PublicKeyAlgorithm(v: Int | NumberLike): PublicKeyAlgorithm
    interface PublicKeyAlgorithm extends Number {
        readonly __PublicKeyAlgorithm: PublicKeyAlgorithm

        String(): string
    }
    const UnknownPublicKeyAlgorithm = PublicKeyAlgorithm(0)
    const RSA = PublicKeyAlgorithm(1)
    const DSA = PublicKeyAlgorithm(2) // Unsupported.
    const ECDSA = PublicKeyAlgorithm(3)
    const Ed25519 = PublicKeyAlgorithm(4)

    interface RevocationListPointer extends Native {
        // SignatureAlgorithm is used to determine the signature algorithm to be
        // used when signing the CRL. If 0 the default algorithm for the signing
        // key will be used.
        SignatureAlgorithm: SignatureAlgorithm

        // RevokedCertificates is used to populate the revokedCertificates
        // sequence in the CRL, it may be empty. RevokedCertificates may be nil,
        // in which case an empty CRL will be created.
        RevokedCertificates: Array<pkix.RevokedCertificate>

        // Number is used to populate the X.509 v2 cRLNumber extension in the CRL,
        // which should be a monotonically increasing sequence number for a given
        // CRL scope and CRL issuer.
        Number: big.IntPointer
        // ThisUpdate is used to populate the thisUpdate field in the CRL, which
        // indicates the issuance date of the CRL.
        ThisUpdate: time.Time
        // NextUpdate is used to populate the nextUpdate field in the CRL, which
        // indicates the date by which the next CRL will be issued. NextUpdate
        // must be greater than ThisUpdate.
        NextUpdate: time.Time
        // ExtraExtensions contains any additional extensions to add directly to
        // the CRL.
        ExtraExtensions: Array<pkix.Extension>
    }

    function SignatureAlgorithm(v: Int | NumberLike): SignatureAlgorithm
    interface SignatureAlgorithm extends Number {
        readonly __SignatureAlgorithm: SignatureAlgorithm

        String(): string
    }
    const UnknownSignatureAlgorithm = SignatureAlgorithm(0)
    const MD2WithRSA = SignatureAlgorithm(1) // Unsupported.
    const MD5WithRSA = SignatureAlgorithm(2) // Only supported for signing, not verification.
    const SHA1WithRSA = SignatureAlgorithm(3)
    const SHA256WithRSA = SignatureAlgorithm(4)
    const SHA384WithRSA = SignatureAlgorithm(5)
    const SHA512WithRSA = SignatureAlgorithm(6)
    const DSAWithSHA1 = SignatureAlgorithm(7)   // Unsupported.
    const DSAWithSHA256 = SignatureAlgorithm(8) // Unsupported.
    const ECDSAWithSHA1 = SignatureAlgorithm(9)
    const ECDSAWithSHA256 = SignatureAlgorithm(10)
    const ECDSAWithSHA384 = SignatureAlgorithm(11)
    const ECDSAWithSHA512 = SignatureAlgorithm(12)
    const SHA256WithRSAPSS = SignatureAlgorithm(13)
    const SHA384WithRSAPSS = SignatureAlgorithm(14)
    const SHA512WithRSAPSS = SignatureAlgorithm(15)
    const PureEd25519 = SignatureAlgorithm(16)

    interface SystemRootsError extends Error {
        readonly __SystemRootsError: SystemRootsError
        Err: Error

        Unwrap(): Error
    }
    interface UnhandledCriticalExtension extends Error {
        readonly __UnhandledCriticalExtension: UnhandledCriticalExtension
    }
    interface UnknownAuthorityError extends Error {
        Cert: CertificatePointer
    }

    function VerifyOptions(): VerifyOptions
    interface VerifyOptions extends Native {
        // DNSName, if set, is checked against the leaf certificate with
        // Certificate.VerifyHostname or the platform verifier.
        DNSName: string

        // Intermediates is an optional pool of certificates that are not trust
        // anchors, but can be used to form a chain from the leaf certificate to a
        // root certificate.
        Intermediates: CertPoolPointer
        // Roots is the set of trusted root certificates the leaf certificate needs
        // to chain up to. If nil, the system roots or the platform verifier are used.
        Roots: CertPoolPointer

        // CurrentTime is used to check the validity of all certificates in the
        // chain. If zero, the current time is used.
        CurrentTime: time.Time

        // KeyUsages specifies which Extended Key Usage values are acceptable. A
        // chain is accepted if it allows any of the listed values. An empty list
        // means ExtKeyUsageServerAuth. To accept any key usage, include ExtKeyUsageAny.
        KeyUsages: Array<ExtKeyUsage>

        // MaxConstraintComparisions is the maximum number of comparisons to
        // perform when checking a given certificate's name constraints. If
        // zero, a sensible default is used. This limit prevents pathological
        // certificates from consuming excessive amounts of CPU time when
        // validating. It does not apply to the platform verifier.
        MaxConstraintComparisions: Int
    }
    function isCertPoolPointer(v: any): v is CertPoolPointer
    function isCertificatePointer(v: any): v is CertificatePointer
    function isCertificateInvalidErrorPointer(v: any): v is CertificateInvalidErrorPointer
    function isCertificateRequestPointer(v: any): v is CertificateRequestPointer
    function isConstraintViolationError(v: any): v is ConstraintViolationError
    function isExtKeyUsage(v: any): v is ExtKeyUsage
    function isHostnameError(v: any): v is HostnameError
    function isInsecureAlgorithmError(v: any): v is InsecureAlgorithmError
    function isInvalidReason(v: any): v is InvalidReason
    function isKeyUsage(v: any): v is KeyUsage
    function isPEMCipher(v: any): v is PEMCipher
    function isPublicKeyAlgorithm(v: any): v is PublicKeyAlgorithm
    function isRevocationListPointer(v: any): v is RevocationListPointer
    function isSignatureAlgorithm(v: any): v is SignatureAlgorithm
    function isSystemRootsError(v: any): v is SystemRootsError
    function isUnhandledCriticalExtension(v: any): v is UnhandledCriticalExtension
    function isUnknownAuthorityError(v: any): v is UnknownAuthorityError
    function isVerifyOptions(v: any): v is VerifyOptions
}