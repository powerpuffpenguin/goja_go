declare module "stdgo/crypto" {
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
    import * as io from "stdgo/io";
    import * as hash from "stdgo/hash";

    function RegisterHash(h: Hash, f: () => hash.Hash): void

    interface Decrypter extends Native {
        Public(): PublicKey
        Decrypt(rand: io.Reader, msg: Bytes, opts: DecrypterOpts): Bytes
    }
    interface DecrypterOpts extends Native { }

    function Hash(v: uint): Hash
    interface Hash extends Number {
        readonly __Hash: Hash

        Available(): boolean
        HashFunc(): Hash
        New(): hash.Hash
        Size(): Int
        String(): string
    }
    const MD4 = Hash(1)       // import golang.org/x/crypto/md4
    const MD5 = Hash(2)                         // import crypto/md5
    const SHA1 = Hash(3)                        // import crypto/sha1
    const SHA224 = Hash(4)                      // import crypto/sha256
    const SHA256 = Hash(5)                      // import crypto/sha256
    const SHA384 = Hash(6)                      // import crypto/sha512
    const SHA512 = Hash(7)                      // import crypto/sha512
    const MD5SHA1 = Hash(8)                     // no implementation; MD5+SHA1 used for TLS RSA
    const RIPEMD160 = Hash(9)                   // import golang.org/x/crypto/ripemd160
    const SHA3_224 = Hash(10)                    // import golang.org/x/crypto/sha3
    const SHA3_256 = Hash(11)                    // import golang.org/x/crypto/sha3
    const SHA3_384 = Hash(12)                    // import golang.org/x/crypto/sha3
    const SHA3_512 = Hash(13)                    // import golang.org/x/crypto/sha3
    const SHA512_224 = Hash(14)                  // import crypto/sha512
    const SHA512_256 = Hash(15)                  // import crypto/sha512
    const BLAKE2s_256 = Hash(16)                 // import golang.org/x/crypto/blake2s
    const BLAKE2b_256 = Hash(17)                 // import golang.org/x/crypto/blake2b
    const BLAKE2b_384 = Hash(18)                 // import golang.org/x/crypto/blake2b
    const BLAKE2b_512 = Hash(19)                 // import golang.org/x/crypto/blake2b

    interface PrivateKey extends Native { }
    interface PublicKey extends Native { }
    interface Signer extends Native {
        Public(): PublicKey
        Sign(rand: io.Reader, digest: Bytes, opts: SignerOpts): Bytes
    }
    interface SignerOpts extends Native {
        HashFunc(): Hash
    }
    function isDecrypter(v: any): v is Decrypter
    function isDecrypterOpts(v: any): v is DecrypterOpts
    function isHash(v: any): v is Hash
    function isPrivateKey(v: any): v is PrivateKey
    function isPublicKey(v: any): v is PublicKey
    function isSigner(v: any): v is Signer
    function isSignerOpts(v: any): v is SignerOpts
}