declare module "stdgo/encoding/ascii85" {
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

    /** return  (ndst, nsrc int, err error) */
    function Decode(dst: Bytes, src: Bytes, flush: boolean): [Int, Int]
    function Encode(dst: Bytes, src: Bytes): Int
    function MaxEncodedLen(n: Int | NumberLike): Int
    function NewDecoder(r: io.Reader): io.Reader
    function NewEncoder(w: io.Writer): io.WriteCloser

    interface CorruptInputError extends Error {
        readonly __CorruptInputError: CorruptInputError
    }

    function isCorruptInputError(v: any): v is CorruptInputError
}