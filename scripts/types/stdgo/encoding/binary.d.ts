declare module "stdgo/encoding/binary" {
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

    const MaxVarintLen16 = 3
    const MaxVarintLen32 = 5
    const MaxVarintLen64 = 10

    const BigEndian: ByteOrder
    const LittleEndian: ByteOrder

    function PutUvarint(buf: Bytes, x: Uint64 | NumberLike): Int
    function PutVarint(buf: Bytes, x: Int64 | NumberLike): Int
    function Read(r: io.Reader, order: ByteOrder, data: any): void
    function ReadUvarint(r: io.ByteReader): Uint64
    function ReadVarint(r: io.ByteReader): Int64
    function Size(v: any): Int
    function Uvarint(buf: Bytes): Uint64
    /** return (int64, int) */
    function Varint(buf: Bytes): [Int64, Int]
    function Write(w: io.Writer, order: ByteOrder, data: any): void

    interface ByteOrder extends Native {
        Uint16(b: Bytes): Uint16
        Uint32(b: Bytes): Uint32
        Uint64(b: Bytes): Uint64
        PutUint16(b: Bytes, v: Uint16 | NumberLike): void
        PutUint32(b: Bytes, v: Uint32 | NumberLike): void
        PutUint64(b: Bytes, v: Uint64 | NumberLike): void
        String(): string
    }

    function isByteOrder(v: any): v is ByteOrde
}