declare module "stdgo/encoding/asn1" {
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

    const TagBoolean = Int(1)
    const TagInteger = Int(2)
    const TagBitString = Int(3)
    const TagOctetString = Int(4)
    const TagNull = Int(5)
    const TagOID = Int(6)
    const TagEnum = Int(10)
    const TagUTF8String = Int(12)
    const TagSequence = Int(16)
    const TagSet = Int(17)
    const TagNumericString = Int(18)
    const TagPrintableString = Int(19)
    const TagT61String = Int(20)
    const TagIA5String = Int(22)
    const TagUTCTime = Int(23)
    const TagGeneralizedTime = Int(24)
    const TagGeneralString = Int(27)
    const TagBMPString = Int(30)

    const ClassUniversal = Int(0)
    const ClassApplication = Int(1)
    const ClassContextSpecific = Int(2)
    const ClassPrivate = Int(3)

    const NullBytes: Bytes
    const NullRawValue: RawValue

    function Marshal(val: any): Bytes
    function MarshalWithParams(val: any, params: string): Bytes
    function Unmarshal(b: Bytes, val: any): Bytes
    function UnmarshalWithParams(b: Bytes, val: any, params: string): Bytes

    function BitString(bytes: Bytes, length: Int | NumberLike): BitString
    interface BitString extends Native {
        readonly __BitString: BitString

        Bytes: Bytes // bits packed into bytes.
        BitLength: Int    // length in bits.

        At(i: Int | NumberLike): Int
        RightAlign(): Bytes
    }

    function Enumerated(v: Int | NumberLike): Enumerated
    interface Enumerated extends Number {
        readonly __Enumerated: Enumerated
    }

    function Flag(v: boolean): Flag
    interface Flag extends Native {
        readonly __Flag: Flag
    }

    function ObjectIdentifier(v: IntSlice): ObjectIdentifier
    interface ObjectIdentifier extends IntSlice {
        readonly __ObjectIdentifier: ObjectIdentifier

        Equal(other: ObjectIdentifier): boolean
        String(): string
    }

    function RawContent(v: IntSlice): RawContent
    interface RawContent extends Bytes {
        readonly __RawContent: RawContent
    }
    function RawValue(): RawValue
    interface RawValue extends Native {
        readonly __RawValue: RawValue

        Class: Int
        Tag: Int
        IsCompound: boolean
        Bytes: Bytes
        FullBytes: Bytes // includes the tag and length
    }

    interface StructuralError extends Error {
        readonly __StructuralError: StructuralError
    }
    interface SyntaxError extends Error {
        readonly __SyntaxError: SyntaxError
    }

    function isBitString(v: any): v is BitString
    function isEnumerated(v: any): v is Enumerated
    function isFlag(v: any): v is Flag
    function isObjectIdentifier(v: any): v is ObjectIdentifier
    function isRawContent(v: any): v is RawContent
    function isRawValue(v: any): v is RawValue
    function isStructuralError(v: any): v is StructuralError
    function isSyntaxError(v: any): v is SyntaxError
}