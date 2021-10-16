declare module "stdgo/crypto/ecdsa" {
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
    import * as crypto from "stdgo/crypto";
    import * as elliptic from "stdgo/crypto/elliptic";
    import * as big from "stdgo/math/big";
    import * as io from "stdgo/io";

    /** return (r, s * big.Int, err error) */
    function Sign(rand: io.Reader, priv: PrivateKeyPointer, hash: Bytes): [big.IntPointer, big.IntPointer]
    function SignASN1(rand: io.Reader, priv: PrivateKeyPointer, hash: Bytes): Bytes
    function Verify(pub: PublicKeyPointer, hash: Bytes, r: big.IntPointer, s: big.IntPointer): boolean
    function VerifyASN1(pub: PublicKeyPointer, hash: Bytes, sig: Bytes): boolean

    function GenerateKey(c: elliptic.Curve, rand: io.Reader): PrivateKeyPointer
    interface PrivateKeyPointer extends PublicKeyPointer {
        PublicKey: PublicKeyPointer
        D: big.IntPointer

        Equal(x: crypto.PrivateKey): boolean
        Public(): crypto.PublicKey
        Sign(rand: io.Reader, digest: Bytes, opts: crypto.SignerOpts): Bytes
    }

    interface PublicKeyPointer extends elliptic.Curve {
        readonly __PublicKeyPointer: PublicKeyPointer

        Curve: elliptic.Curve

        X: big.IntPointer
        Y: big.IntPointer

        Equal(x: crypto.PublicKey): boolean
    }

    function isPrivateKeyPointer(v: any): v is PrivateKeyPointer
    function isPublicKeyPointer(v: any): v is PublicKeyPointer
}