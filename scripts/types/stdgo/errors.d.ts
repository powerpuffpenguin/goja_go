declare module "stdgo/errors" {
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

    function GoError(err: Error): GoError
    function Wrap(text: string, err: Error): Error
    function New(text: string): Error
    function Unwrap(err: Error): Error
    function Is(err: Error, target: Error): boolean
    function As(err: Error, target: any): boolean

    function isError(v: any): Error
}