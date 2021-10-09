declare module "stdgo/path" {
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

    const ErrBadPattern: Error

    function Base(path: string): string
    function Clean(path: string): string
    function Dir(path: string): string
    function Ext(path: string): string
    function IsAbs(path: string): boolean
    function Join(...elem: Array<string>): string
    function Match(pattern: string, name: string): boolean
    /** return (dir, file string) */
    function Split(path: string): [string, string]

}