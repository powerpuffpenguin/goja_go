declare module "stdgo/compress/zlib" {
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
    import * as flate from "stdgo/compress/flate";

    const NoCompression = flate.NoCompression
    const BestSpeed = flate.BestSpeed
    const BestCompression = flate.BestCompression
    const DefaultCompression = flate.DefaultCompression
    const HuffmanOnly = flate.HuffmanOnly

    const ErrChecksum: Error
    const ErrDictionary: Error
    const ErrHeader: Error

    function NewReader(r: io.Reader): io.ReadCloser
    function NewReaderDict(r: io.Reader, dict: Bytes): io.ReadCloser

    interface Resetter extends Native {
        Reset(r: io.Reader, dict: Bytes): void
    }

    function NewWriter(w: io.Writer): WriterPointer
    function NewWriterLevel(w: io.Writer, level: Int | NumberLike): WriterPointer
    function NewWriterLevelDict(w: io.Writer, level: Int | NumberLike, dict: Bytes): WriterPointer
    interface WriterPointer extends Native {
        readonly __WriterPointer: WriterPointer

        Close(): void
        Flush(): void
        Reset(w: io.Writer): void
        Write(p: Bytes): Int
    }

    function isResetter(v: any): v is Resetter
    function isWriterPointer(v: any): v is WriterPointer
}