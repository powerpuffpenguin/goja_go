declare module "stdgo/crypto/x509/pkix" {
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
    import * as asn1 from "stdgo/encoding/asn1";
    import * as time from "stdgo/time";
    import * as big from "stdgo/math/big";

    interface AlgorithmIdentifier extends Native {
        readonly __AlgorithmIdentifier: AlgorithmIdentifier
        Algorithm: asn1.ObjectIdentifier
        Parameters: asn1.RawValue
    }
    interface AttributeTypeAndValue extends Native {
        readonly __AttributeTypeAndValue: AttributeTypeAndValue
        Type: asn1.ObjectIdentifier
        Value: any
    }
    interface AttributeTypeAndValueSET extends Native {
        readonly __AttributeTypeAndValueSET: AttributeTypeAndValueSET
        Type: asn1.ObjectIdentifier
        Value: Slice<Slice<AttributeTypeAndValue>>
    }
    interface CertificateListPointer extends Native {
        readonly __CertificateListPointer: CertificateListPointer
        TBSCertList: TBSCertificateListPointer
        SignatureAlgorithm: AlgorithmIdentifier
        SignatureValue: asn1.BitString

        HasExpired(now: time.Time): boolean
    }
    interface Extension extends Native {
        readonly __Extension: Extension
        Id: asn1.ObjectIdentifier
        Critical: boolean
        Value: Bytes
    }
    interface Name extends Native {
        readonly __Name: Name
        Country: Array<string>
        Organization: Array<string>
        OrganizationalUnit: Array<string>
        Locality: Array<string>
        Province: Array<string>
        StreetAddress: Array<string>
        PostalCode: Array<string>
        SerialNumber: string
        CommonName: string

        // Names contains all parsed attributes. When parsing distinguished names,
        // this can be used to extract non-standard attributes that are not parsed
        // by this package. When marshaling to RDNSequences, the Names field is
        // ignored, see ExtraNames.
        Names: Array<AttributeTypeAndValue>

        // ExtraNames contains attributes to be copied, raw, into any marshaled
        // distinguished names. Values override any attributes with the same OID.
        // The ExtraNames field is not populated when parsing, see Names.
        ExtraNames: Array<AttributeTypeAndValue>

        FillFromRDNSequence(rdns: RDNSequencePointer): void
        String(): string
        ToRDNSequence(): RDNSequence
    }
    interface RDNSequence extends Slice<RelativeDistinguishedNameSET> {
        readonly __RDNSequence: RDNSequence
        String(): string
    }
    interface RelativeDistinguishedNameSET extends Slice<AttributeTypeAndValue> {
        readonly __RelativeDistinguishedNameSET: RelativeDistinguishedNameSET
    }
    interface RevokedCertificate extends Native {
        readonly __RevokedCertificate: RevokedCertificate
        SerialNumber: big.IntPointer
        RevocationTime: time.Time
        Extensions: Array<Extension>
    }
    interface TBSCertificateListPointer extends Native {
        readonly __TBSCertificateListPointer: TBSCertificateListPointer
        Raw: asn1.RawContent
        Version: Int
        Signature: AlgorithmIdentifier
        Issuer: RDNSequence
        ThisUpdate: time.Time
        NextUpdate: time.Time
        RevokedCertificates: Array<RevokedCertificate>
        Extensions: Array<Extension>
    }

    function isAlgorithmIdentifier(v: any): v is AlgorithmIdentifier
    function isAttributeTypeAndValue(v: any): v is AttributeTypeAndValue
    function isAttributeTypeAndValueSET(v: any): v is AttributeTypeAndValueSET
    function isCertificateListPointer(v: any): v is CertificateListPointer
    function isExtension(v: any): v is Extension
    function isName(v: any): v is Name
    function isRDNSequence(v: any): v is RDNSequence
    function isRelativeDistinguishedNameSET(v: any): v is RelativeDistinguishedNameSET
    function isRevokedCertificate(v: any): v is RevokedCertificate
    function isTBSCertificateListPointer(v: any): v is TBSCertificateListPointer
}