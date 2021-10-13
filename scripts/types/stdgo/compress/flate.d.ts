declare module "stdgo/compress/flate" {
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

    const NoCompression = 0
    const BestSpeed = 1
    const BestCompression = 9
    const DefaultCompression = -1
    const HuffmanOnly = -2

    function NewReader(r: io.Reader): io.ReadCloser
    function NewReaderDict(r: io.Reader, dict: Bytes): io.ReadCloser

    interface CorruptInputError extends Error {
        readonly __CorruptInputError: CorruptInputError
    }
    class InternalError extends Error {
        readonly __InternalError: InternalError
    }
    interface Reader extends io.Reader, io.ByteReader { }
    interface Resetter extends Native {
        Reset(r: io.Reader, dict: Bytes): void
    }

    function NewWriter(w: io.Writer, level: Int | NumberLike): Writer
    function NewWriterDict(w: io.Writer, level: Int | NumberLike, dict: Bytes): Writer
    interface WriterPointer extends Native {
        readonly __WriterPointer: WriterPointer

        Close()
        Flush()
        Reset(dst: io.Writer)
        Write(data: Bytes): Int
    }

    function isCorruptInputError(v: any): v is CorruptInputError
    function isInternalError(v: any): v is InternalError
    function isReader(v: any): v is Reader
    function isResetter(v: any): v is Resetter
    function isWriterPointer(v: any): v is WriterPointer
}