declare module "stdgo/image/jpeg" {
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
    import * as image from "stdgo/image";
    import * as io from "stdgo/io";

    const DefaultQuality = Int(75)

    function Decode(r: io.Reader): image.Image
    function DecodeConfig(r: io.Reader): image.Config
    function Encode(w: io.Writer, m: image.Image, o: OptionsPointer): void

    interface FormatError extends Error {
        readonly __FormatError: FormatError
    }

    function Options(quality: Int | NumberLike): OptionsPointer
    interface OptionsPointer extends Native {
        Quality: Int
    }

    interface UnsupportedError extends Error {
        readonly __UnsupportedError: UnsupportedError
    }

    function isFormatError(v: any): v is FormatError
    function isOptionsPointer(v: any): v is OptionsPointer
    function isUnsupportedError(v: any): v is UnsupportedError
}