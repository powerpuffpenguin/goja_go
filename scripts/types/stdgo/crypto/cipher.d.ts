declare module "stdgo/crypto/cipher" {
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

    function NewGCM(cipher: Block): AEAD
    function NewGCMWithNonceSize(cipher: Block, size: Int | NumberLike): AEAD
    function NewGCMWithTagSize(cipher: Block, tagSize: Int | NumberLike): AEAD
    interface AEAD extends Native {
        NonceSize(): Int
        Overhead(): Int
        Seal(dst: Bytes, nonce: Bytes, plaintext: Bytes, additionalData: Bytes): Bytes
        Open(dst: Bytes, nonce: Bytes, ciphertext: Bytes, additionalData: Bytes): Bytes
    }
    interface Block extends Native {
        BlockSize(): Int
        Encrypt(dst: Bytes, src: Bytes): void
        Decrypt(dst: Bytes, src: Bytes): void
    }
    function NewCBCDecrypter(b: Block, iv: Bytes): BlockMode
    function NewCBCEncrypter(b: Block, iv: Bytes): BlockMode
    interface BlockMode extends Native {
        BlockSize(): Int
        CryptBlocks(dst: Bytes, src: Bytes): void
    }

    function NewCFBDecrypter(block: Block, iv: Bytes): Stream
    function NewCFBEncrypter(block: Block, iv: Bytes): Stream
    function NewCTR(block: Block, iv: Bytes): Stream
    function NewOFB(b: Block, iv: Bytes): Stream
    interface Stream extends Native {
        XORKeyStream(dst: Bytes, src: Bytes): void
    }

    interface StreamReader extends Native {
        readonly __StreamReader: StreamReader
        S: Stream
        R: io.Reader

        Read(dst: Bytes): Int
    }

    interface StreamWriter extends Native {
        readonly __StreamWriter: StreamWriter
        S: Stream
        W: io.Writer
        Err: Error // unused

        Close(): void
        Write(src: Bytes): Int
    }

    function isAEAD(v: any): v is AEAD
    function isBlock(v: any): v is Block
    function isBlockMode(v: any): v is BlockMode
    function isStream(v: any): v is Stream
    function isStreamReader(v: any): v is StreamReader
    function isStreamWriter(v: any): v is StreamWriter
}