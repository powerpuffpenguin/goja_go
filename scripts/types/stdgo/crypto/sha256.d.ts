declare module "stdgo/crypto/sha256" {
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
    import * as hash from "stdgo/hash";

    const BlockSize = 64
    const Size = 32
    const Size224 = 28

    function New(): hash.Hash
    function New224(): hash.Hash
    function Sum224(data: Bytes): Bytes
    function Sum256(data: Bytes): Bytes
}