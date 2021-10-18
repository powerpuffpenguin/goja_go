declare module "stdgo/crypto/sha512" {
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

    const Size = 64
    const Size224 = 28
    const Size256 = 32
    const Size384 = 48
    const BlockSize = 128

    function New(): hash.Hash
    function New384(): hash.Hash
    function New512_224(): hash.Hash
    function New512_256(): hash.Hash
    function Sum384(data: Bytes): Bytes
    function Sum512(data: Bytes): Bytes
    function Sum512_224(data: Bytes): Bytes
    function Sum512_256(data: Bytes): Bytes
}