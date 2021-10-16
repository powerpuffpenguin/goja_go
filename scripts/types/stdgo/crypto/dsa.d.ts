declare module "stdgo/crypto/dsa" {
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
    import * as io from "stdgo/io";

    const ErrInvalidPublicKey: Error

    function GenerateKey(priv: PrivateKeyPointer, rand: io.Reader): void
    function GenerateParameters(params: ParametersPointer, rand: io.Reader, sizes: ParameterSizes): void
    /** return (r, s * big.Int, err error) */
    function Sign(rand: io.Reader, priv: PrivateKeyPointer, hash: Bytes): [big.IntPointer, big.IntPointer]
    function Verify(pub: PublicKeyPointer, hash: Bytes, r: big.IntPointer, s: big.IntPointer): boolean

    interface ParameterSizes extends Number {
        readonly __ParameterSizes: ParameterSizes
    }
    const L1024N160: ParameterSizes
    const L2048N224: ParameterSizes
    const L2048N256: ParameterSizes
    const L3072N256: ParameterSizes

    function Parameters(): ParametersPointer
    interface ParametersPointer extends Native {
        P: big.IntPointer
        Q: big.IntPointer
        G: big.IntPointer
    }
    function PrivateKey(): PrivateKeyPointer
    interface PrivateKeyPointer extends PublicKeyPointer {
        X: big.IntPointer
    }
    function PublicKey(): PublicKeyPointer
    interface PublicKeyPointer extends ParametersPointer {
        Y: big.IntPointer
    }
    function isParametersPointer(v: any): v is ParametersPointer
    function isPrivateKeyPointer(v: any): v is PrivateKeyPointer
    function isPublicKeyPointer(v: any): v is PublicKeyPointer
}