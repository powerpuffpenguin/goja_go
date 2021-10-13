declare module "stdgo/hash/crc64" {
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

    const ISO: Uint64
    const ECMA: Uint64

    const Size = 8

    function Checksum(data: Bytes, tab: TablePointer): Uint64
    function New(tab: TablePointer): hash.Hash64
    function Update(crc: Uint64 | NumberLike, tab: TablePointer, p: Bytes): Uint64

    function MakeTable(poly: Uint64 | NumberLike): TablePointer
    interface TablePointer extends Native {
        readonly __TablePointer: TablePointer
    }

    function isTablePointer(v: any, isptr?: boolean): v is TablePointer
}