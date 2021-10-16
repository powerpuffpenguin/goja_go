declare module "stdgo/encoding/gob" {
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

    function Register(value: any): void
    function RegisterName(name: string, value: any): void

    interface CommonType extends Native {
        readonly __CommonType: CommonType
        Name: string
    }

    function NewDecoder(r: io.Reader): DecoderPointer
    interface DecoderPointer extends Native {
        readonly __DecoderPointer: DecoderPointer

        Decode(e: any): void
        DecodeValue(v: Value): void
    }

    function NewEncoder(w: io.Writer): EncoderPointer
    interface EncoderPointer extends Native {
        readonly __EncoderPointer: EncoderPointer

        Encode(e: any): void
        EncodeValue(value: Value): void
    }

    interface GobDecoder extends Native {
        GobDecode(b: Bytes): void
    }
    interface GobEncoder extends Native {
        GobEncode(): Bytes
    }
    function isCommonType(v: any): v is CommonType
    function isDecoderPointer(v: any): v is DecoderPointer
    function isEncoderPointer(v: any): v is EncoderPointer
    function isGobDecoder(v: any): v is GobDecoder
    function isGobEncoder(v: any): v is GobEncoder
}