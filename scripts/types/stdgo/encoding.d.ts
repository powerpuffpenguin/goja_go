declare module "stdgo/encoding" {
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

    interface BinaryMarshaler extends Native {
        MarshalBinary(): Bytes
    }
    interface BinaryUnmarshaler extends Native {
        UnmarshalBinary(data: Bytes): void
    }
    interface TextMarshaler extends Native {
        MarshalText(): Bytes
    }
    interface TextUnmarshaler extends Native {
        UnmarshalText(text: Bytes): void
    }

    function isBinaryMarshaler(v: any): v is BinaryMarshaler
    function isBinaryUnmarshaler(v: any): v is BinaryUnmarshaler
    function isTextMarshaler(v: any): v is TextMarshaler
    function isTextUnmarshaler(v: any): v is TextUnmarshaler
}