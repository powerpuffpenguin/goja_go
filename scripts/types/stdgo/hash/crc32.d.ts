declare module "stdgo/hash/crc32" {
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

    const IEEE: Uint32
    const Castagnoli: Uint32
    const Koopman: Uint32

    const Size = 4
    const IEEETable: TablePointer

    function Checksum(data: Bytes, tab: TablePointer): Uint32
    function ChecksumIEEE(data: Bytes): Uint32
    function New(tab: TablePointer): hash.Hash32
    function NewIEEE(): hash.Hash32
    function Update(crc: Uint32 | NumberLike, tab: TablePointer, p: Bytes): Uint32

    function MakeTable(poly: Uint32 | NumberLike): TablePointer
    interface TablePointer extends Native {
        readonly __TablePointer: TablePointer
    }

    function isTablePointer(v: any): v is TablePointer
}