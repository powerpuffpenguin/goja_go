declare module "stdgo/crypto/rand" {
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
    import * as big from "stdgo/math/big";

    const Reader: io.Reader

    function Int(rand: io.Reader, max: big.IntPointer): big.IntPointer
    function Prime(rand: io.Reader, bits: Int | NumberLike): big.IntPointer
    function Read(b: Bytes): Int
}