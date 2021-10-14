declare module "stdgo/index/suffixarray" {
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
    import * as regexp from "stdgo/regexp";
    import * as io from "stdgo/io";

    function New(data: Bytes): IndexPointer
    interface IndexPointer extends Navigator {
        readonly __IndexPointer: IndexPointer

        Bytes(): Bytes
        FindAllIndex(r: regexp.RegexpPointer, n: Int | NumberLike): Slice<IntSlice>
        Lookup(s: Bytes, n: Int | NumberLike): IntSlice
        Read(r: io.Reader): void
        Write(w: io.Writer): void
    }

    function isIndexPointer(v: any): v is IndexPointer
}