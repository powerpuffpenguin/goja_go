declare module "stdgo/crypto/ed25519" {
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
    import * as io from "stdgo/io";

    // PublicKeySize is the size, in bytes, of public keys as used in this package.
    const PublicKeySize = 32
    // PrivateKeySize is the size, in bytes, of private keys as used in this package.
    const PrivateKeySize = 64
    // SignatureSize is the size, in bytes, of signatures generated and verified by this package.
    const SignatureSize = 64
    // SeedSize is the size, in bytes, of private key seeds. These are the private key representations used by RFC 8032.
    const SeedSize = 32

    /** return (PublicKey, PrivateKey, error) */
    function GenerateKey(rand: io.Reader): [PublicKey, PrivateKey]
    function Sign(privateKey: PrivateKey, message: Bytes): Bytes
    function Verify(publicKey: PublicKey, message: Bytes, sig: Bytes): boolean

    function NewKeyFromSeed(seed: Bytes): PrivateKey
    interface PrivateKey extends Bytes {
        readonly __PrivateKey: PrivateKey

        Equal(x: crypto.PrivateKey): boolean
        Public(): crypto.PublicKey
        Seed(): Bytes
        Sign(rand: io.Reader, message: Bytes, opts: crypto.SignerOpts): Bytes
    }

    interface PublicKey extends Bytes {
        readonly __PublicKey: PublicKey

        Equal(x: crypto.PublicKey): boolean
    }
    function isPrivateKey(v: any): v is PrivateKey
    function isPublicKey(v: any): v is PublicKey
}