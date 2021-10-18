declare module "stdgo/crypto/rsa" {
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
    import * as io from "stdgo/io";
    import * as hash from "stdgo/hash";

    // PSSSaltLengthAuto causes the salt in a PSS signature to be as large
    // as possible when signing, and to be auto-detected when verifying.
    const PSSSaltLengthAuto = Int(0)
    // PSSSaltLengthEqualsHash causes the salt length to equal the length
    // of the hash used in the signature.
    const PSSSaltLengthEqualsHash = Int(-1)

    const ErrDecryption: Error
    const ErrMessageTooLong: Error
    const ErrVerification: Error

    function DecryptOAEP(hash: hash.Hash, random: io.Reader, priv: PrivateKeyPointer, ciphertext: Bytes, label: Bytes): Bytes
    function DecryptPKCS1v15(rand: io.Reader, priv: PrivateKeyPointer, ciphertext: Bytes): Bytes
    function DecryptPKCS1v15SessionKey(rand: io.Reader, priv: PrivateKeyPointer, ciphertext: Bytes, key: Bytes): void
    function EncryptOAEP(hash: hash.Hash, random: io.Reader, pub: PublicKeyPointer, msg: Bytes, label: Bytes): Bytes
    function EncryptPKCS1v15(rand: io.Reader, pub: PublicKeyPointer, msg: Bytes): Bytes
    function SignPKCS1v15(rand: io.Reader, priv: PrivateKeyPointer, hash: crypto.Hash, hashed: Bytes): Bytes
    function SignPSS(rand: io.Reader, priv: PrivateKeyPointer, hash: crypto.Hash, digest: Bytes, opts: PSSOptionsPointer): Bytes
    function VerifyPKCS1v15(pub: PublicKeyPointer, hash: crypto.Hash, hashed: Bytes, sig: Bytes): void
    function VerifyPSS(pub: PublicKeyPointer, hash: crypto.Hash, digest: Bytes, sig: Bytes, opts: PSSOptionsPointer): void

    function CRTValue(exp: big.IntPointer, coeff: big.IntPointer, r: big.IntPointer): CRTValue
    interface CRTValue extends Native {
        readonly __CRTValue: CRTValue

        Exp: big.IntPointer // D mod (prime-1).
        Coeff: big.IntPointer // R·Coeff ≡ 1 mod Prime.
        R: big.IntPointer // product of primes prior to this (inc p and q).
    }

    function OAEPOptions(hash: crypto.Hash, label: Bytes): OAEPOptionsPointer
    interface OAEPOptionsPointer extends Native {
        readonly __OAEPOptionsPointer: OAEPOptionsPointer
        // Hash is the hash function that will be used when generating the mask.
        Hash: crypto.Hash
        // Label is an arbitrary byte string that must be equal to the value
        // used when encrypting.
        Label: Bytes
    }

    function PKCS1v15DecryptOptions(sessionKeyLen: Int | NumberLike): PKCS1v15DecryptOptionsPointer
    interface PKCS1v15DecryptOptionsPointer extends Native {
        readonly __PKCS1v15DecryptOptionsPointer: PKCS1v15DecryptOptionsPointer
        // SessionKeyLen is the length of the session key that is being
        // decrypted. If not zero, then a padding error during decryption will
        // cause a random plaintext of this length to be returned rather than
        // an error. These alternatives happen in constant time.
        SessionKeyLen: Int
    }

    function PSSOptions(saltLength: Int | NumberLike, hash: crypto.Hash): PSSOptionsPointer
    interface PSSOptionsPointer extends Native {
        readonly __PSSOptionsPointer: PSSOptionsPointer
        // SaltLength controls the length of the salt used in the PSS
        // signature. It can either be a number of bytes, or one of the special
        // PSSSaltLength constants.
        SaltLength: Int

        // Hash is the hash function used to generate the message digest. If not
        // zero, it overrides the hash function passed to SignPSS. It's required
        // when using PrivateKey.Sign.
        Hash: crypto.Hash

        HashFunc(): crypto.Hash
    }

    interface PrecomputedValues extends Native {
        readonly __PrecomputedValues: PrecomputedValues
        Dp: big.IntPointer
        Dq: big.IntPointer // D mod (P-1) (or mod Q-1)
        Qinv: big.IntPointer // Q^-1 mod P

        // CRTValues is used for the 3rd and subsequent primes. Due to a
        // historical accident, the CRT for the first two primes is handled
        // differently in PKCS #1 and interoperability is sufficiently
        // important that we mirror this.
        CRTValues: Array<CRTValue>
    }

    function GenerateKey(random: io.Reader, bits: Int | NumberLike): PrivateKeyPointer
    function GenerateMultiPrimeKey(random: io.Reader, nprimes: Int | NumberLike, bits: Int | NumberLike): PrivateKeyPointer
    interface PrivateKeyPointer extends PublicKeyPointer {
        readonly __PrivateKeyPointer: PrivateKeyPointer
        PublicKey: PublicKeyPointer          // public part.
        D: big.IntPointer   // private exponent
        Primes: Array<big.IntPointer> // prime factors of N, has >= 2 elements.

        // Precomputed contains precomputed values that speed up private
        // operations, if available.
        Precomputed: PrecomputedValues

        Decrypt(rand: io.Reader, ciphertext: Bytes, opts: crypto.DecrypterOpts): Bytes
        Equal(x: crypto.PrivateKey): boolean
        Precompute(): void
        Public(): crypto.PublicKey
        Sign(rand: io.Reader, digest: Bytes, opts: crypto.SignerOpts): Bytes
        Validate(): void
    }

    function PublicKey(n: big.IntPointer, e: Int): PublicKeyPointer
    interface PublicKeyPointer extends Native {
        readonly __PublicKeyPointer: PublicKeyPointer
        N: big.IntPointer // modulus
        E: Int      // public exponent

        Equal(x: crypto.PublicKey): boolean
        Size(): Int
    }
    // CRTValue
    // OAEPOptionsPointer
    // PKCS1v15DecryptOptionsPointer
    // PSSOptionsPointer
    // PrecomputedValues

    // PublicKeyPointer
}