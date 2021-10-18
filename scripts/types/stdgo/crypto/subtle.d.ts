declare module "stdgo/crypto/subtle" {
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

    function ConstantTimeByteEq(x: Uint8 | NumberLike, y: Uint8 | NumberLike): Int
    function ConstantTimeCompare(x: Bytes, y: Bytes): Int
    function ConstantTimeCopy(v: Int | NumberLike, x: Bytes, y: Bytes): void
    function ConstantTimeEq(x: Int32 | NumberLike, y: Int32 | NumberLike): Int
    function ConstantTimeLessOrEq(x: Int | NumberLike, y: Int | NumberLike): Int
    function ConstantTimeSelect(v: Int | NumberLike, x: Int | NumberLike, y: Int | NumberLike): Int
}