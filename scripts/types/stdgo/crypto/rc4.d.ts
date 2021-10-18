declare module "stdgo/crypto/rc4" {
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

    function NewCipher(key: Bytes): CipherPointer
    interface CipherPointer extends Native {
        readonly __CipherPointer: CipherPointer

        XORKeyStream(dst: Bytes, src: Bytes): void
    }

    interface KeySizeError extends Error {
        readonly __KeySizeError: KeySizeError
    }
    function isCipherPointer(v: any): v is CipherPointer
    function isKeySizeError(v: any): v is KeySizeError
}