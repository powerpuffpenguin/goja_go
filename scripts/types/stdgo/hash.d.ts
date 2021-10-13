declare module "stdgo/hash" {
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

    interface Hash extends io.Writer {
        Sum(b: Bytes): Bytes
        Reset(): void
        Size(): Int
        BlockSize(): Int
    }

    interface Hash32 extends Hash {
        Sum32(): Uint32
    }

    interface Hash64 extends Hash {
        Sum64(): Uint64
    }

    function isHash(v: any): v is Hash
    function isHash32(v: any): v is Hash32
    function isHash64(v: any): v is Hash64
}