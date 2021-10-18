declare module "stdgo/crypto/aes" {
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
    import * as cipher from "stdgo/crypto/cipher";

    const BlockSize = 16

    function NewCipher(key: Bytes): cipher.Block

    interface KeySizeError extends Error {
        readonly __KeySizeError: KeySizeError
    }
    function isKeySizeError(v: any): v is KeySizeError
}