declare module "stdgo/encoding/hex" {
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

    const ErrLength: Error

    function Decode(dst: Bytes, src: Bytes): Int
    function DecodeString(s: string): Bytes
    function DecodedLen(x: Int | NumberLike): Int
    function Dump(data: Bytes): string
    function Dumper(w: io.Writer): io.WriteCloser
    function Encode(dst: Bytes, src: Bytes): Int
    function EncodeToString(src: Bytes): string
    function EncodedLen(n: Int | NumberLike): Int
    function NewDecoder(r: io.Reader): io.Reader
    function NewEncoder(w: io.Writer): io.Writer

    interface InvalidByteError extends Error {
        readonly __InvalidByteError: InvalidByteError
    }

    function isInvalidByteError(v: any): v is InvalidByteError
}