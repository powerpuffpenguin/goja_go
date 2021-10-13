declare module "stdgo/compress/bzip2" {
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
    import *as io from "stdgo/io";

    function NewReader(r: io.Reader): io.Reader

    interface StructuralError extends Error {
        readonly __StructuralError: StructuralError
    }

    function isStructuralError(v: any): v is StructuralError
}