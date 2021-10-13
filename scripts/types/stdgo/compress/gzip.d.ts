declare module "stdgo/compress/gzip" {
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
    import * as flate from "stdgo/compress/flate";
    import *as io from "stdgo/io";
    import *as time from "stdgo/time";

    const NoCompression = flate.NoCompression
    const BestSpeed = flate.BestSpeed
    const BestCompression = flate.BestCompression
    const DefaultCompression = flate.DefaultCompression
    const HuffmanOnly = flate.HuffmanOnly

    const ErrChecksum: Error
    const ErrHeader: Error

    interface Header extends Native {
        readonly __Header: Header

        Comment: string    // comment
        Extra: Bytes    // "extra data"
        ModTime: time.Time // modification time
        Name: string    // file name
        OS: Byte      // operating system type
    }

    function NewReader(r: io.Reader): Reader
    interface ReaderPointer extends Header {
        readonly __ReaderPointer: ReaderPointer
        Header: Header

        Close(): void
        Multistream(ok: boolean): void
        Read(p: Bytes): Int
        Reset(r: io.Reader): void
    }

    function NewWriter(w: io.Writer): Writer
    function NewWriterLevel(w: io.Writer, level: Int | NumberLike): Writer
    interface WriterPointer extends Header {
        readonly __WriterPointer: WriterPointer
        Header: Header

        Close(): void
        Flush(): void
        Reset(w: io.Writer): void
        Write(p: Bytes): Int
    }

    function isHeader(v: any): v is Header
    function isReaderPointer(v: any): v is ReaderPointer
    function isWriterPointer(v: any): v is WriterPointer
}