declare module "stdgo/encoding/base64" {
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

    const StdPadding: Rune // Standard padding character
    const NoPadding: Rune  // No padding

    const RawStdEncoding: EncodingPointer
    const RawURLEncoding: EncodingPointer
    const StdEncoding: EncodingPointer
    const URLEncoding: EncodingPointer

    function NewDecoder(enc: EncodingPointer, r: io.Reader): io.Reader
    function NewEncoder(enc: EncodingPointer, w: io.Writer): io.WriteCloser

    interface CorruptInputError extends Error {
        readonly __CorruptInputError: CorruptInputError
    }

    function NewEncoding(encoder: string): EncodingPointer
    interface EncodingPointer extends Native {
        readonly __EncodingPointer: EncodingPointer

        Decode(dst: Bytes, src: Bytes): Int
        DecodeString(s: string): Bytes
        DecodedLen(n: Int | NumberLike): Int
        Encode(dst: Bytes, src: Bytes): void
        EncodeToString(src: Bytes): string
        EncodedLen(n: Int | NumberLike): Int
        Strict(): EncodingPointer
        WithPadding(padding: Rune | NumberLike): EncodingPointer
    }

    function isCorruptInputError(v: any): v is CorruptInputError
    function isEncodingPointer(v: any): v is EncodingPointer
}