declare module "stdgo/mime/quotedprintable" {
    import {
        Float32, Float64,
        Int64, Int32, Int16, Int8, Int,
        Uint64, Uint32, Uint16, Uint8, Uint,
        Number, NumberLike,
        Byte, Bytes, Rune, Runes,
        Float32Slice, Float64Slice,
        Int64Slice, Int32Slice, Int16Slice, Int8Slice, IntSlice,
        Uint64Slice, Uint32Slice, Uint16Slice, Uint8Slice, UintSlice,
        ReadChannel, WriteChannel, Channel,
        Slice, Map,
        Uintptr, Native,
    } from "stdgo/builtin";
    import * as io from "stdgo/io";

    function NewReader(r: io.Reader): ReaderPointer
    interface ReaderPointer extends Native {
        readonly __ReaderPointer: ReaderPointer
        // contains filtered or unexported fields
        Read(p: Bytes): Int
    }
    function NewWriter(w: io.Writer): WriterPointer
    interface WriterPointer extends Native {
        readonly __WriterPointer: WriterPointer
        // Binary mode treats the writer's input as pure binary and processes end of
        // line bytes as binary data.
        Binary: boolean
        // contains filtered or unexported fields

        Close(): void
        Write(p: Bytes): Int
    }

    function isReaderPointer(v: any): v is ReaderPointer
    function isWriterPointer(v: any): v is WriterPointer
}