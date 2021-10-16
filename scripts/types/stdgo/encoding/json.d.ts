declare module "stdgo/encoding/json" {
    import {
        Float32, Float64,
        Int64, Int32, Int16, Int8, Int,
        Uint64, Uint32, Uint16, Uint8, Uint,
        NumberLike,
        Byte, Bytes, Rune, Runes,
        Float32Slice, Float64Slice,
        Int64Slice, Int32Slice, Int16Slice, Int8Slice, IntSlice,
        Uint64Slice, Uint32Slice, Uint16Slice, Uint8Slice, UintSlice,
        Error,
        ReadChannel, WriteChannel, Channel,
        Slice, Map,
        Uintptr, Native,
        Type, Value,
    } from "stdgo/builtin";
    import * as io from "stdgo/io";
    import * as bytes from "stdgo/bytes";

    function Compact(dst: bytes.BufferPointer, src: Bytes): void
    function HTMLEscape(dst: bytes.BufferPointer, src: Bytes): void
    function Indent(dst: bytes.BufferPointer, src: Bytes, prefix: string, indent: string): void
    function Marshal(v: any): Bytes
    function MarshalIndent(v: any, prefix: string, indent: string): Bytes
    function Unmarshal(data: Bytes, v: any): void
    function Valid(data: Bytes): boolean

    function NewDecoder(r: io.Reader): DecoderPointer
    interface DecoderPointer extends Native {
        readonly __DecoderPointer: DecoderPointer

        Buffered(): io.Reader
        Decode(v: any): void
        DisallowUnknownFields(): void
        InputOffset(): Int64
        More(): boolean
        Token(): Token
        UseNumber(): void
    }

    interface Delim extends Number {
        readonly __Delim: Delim

        String(): string
    }

    function NewEncoder(w: io.Writer): EncoderPointer
    interface EncoderPointer extends Native {
        readonly __EncoderPointer: EncoderPointer

        Encode(v: any): void
        SetEscapeHTML(on: boolean): void
        SetIndent(prefix: string, indent: string): void
    }

    interface InvalidUnmarshalErrorPointer extends Error {
        readonly __InvalidUnmarshalErrorPointer: InvalidUnmarshalErrorPointer

        Type: Type
    }

    interface Marshaler extends Native {
        MarshalJSON(): Bytes
    }
    interface MarshalerErrorPointer extends Error {
        readonly __MarshalerErrorPointer: MarshalerErrorPointer

        Type: Type
        Err: Error

        Unwrap(): Error
    }
    interface Number extends Native {
        readonly __Number: Number

        Float64(): Float64
        Int64(): Int64
        String(): string
    }

    interface RawMessage extends Uint8Slice {
        readonly __RawMessage: RawMessage

        MarshalJSON(): Bytes
        UnmarshalJSON(data: Bytes): void
    }

    interface SyntaxErrorPointer extends Error {
        readonly SyntaxErrorPointer: SyntaxErrorPointer
        Offset: Int64 // error occurred after reading Offset bytes
    }
    interface Token extends Native {
    }
    interface UnmarshalTypeErrorPointer extends Error {
        readonly __UnmarshalTypeErrorPointer: UnmarshalTypeErrorPointer
        Value: string       // description of JSON value - "bool", "array", "number -5"
        Type: Type // type of Go value it could not be assigned to
        Offset: Int64        // error occurred after reading Offset bytes
        Struct: string       // name of the struct type containing the field
        Field: string       // the full path from root node to the field
    }
    interface Unmarshaler extends Native {
        UnmarshalJSON(b: Bytes): void
    }
    interface UnsupportedTypeErrorPointer extends Error {
        readonly __UnsupportedTypeErrorPointer: UnsupportedTypeErrorPointer
        Type: Type
    }
    interface UnsupportedValueErrorPointer extends Error {
        readonly __UnsupportedValueErrorPointer: UnsupportedValueErrorPointer
        Value: Value
        Str: string
    }

    function isDecoderPointer(v: any): v is DecoderPointer
    function isDelim(v: any): v is Delim
    function isEncoderPointer(v: any): v is EncoderPointer
    function isInvalidUnmarshalErrorPointer(v: any): v is InvalidUnmarshalErrorPointer
    function isMarshaler(v: any): v is Marshaler
    function isMarshalerErrorPointer(v: any): v is MarshalerErrorPointer
    function isNumber(v: any): v is Number
    function isRawMessage(v: any): v is RawMessage
    function isSyntaxErrorPointer(v: any): v is SyntaxErrorPointer
    function isToken(v: any): v is Token
    function isUnmarshalTypeErrorPointer(v: any): v is UnmarshalTypeErrorPointer
    function isUnmarshaler(v: any): v is Unmarshaler
    function isUnsupportedTypeErrorPointer(v: any): v is UnsupportedTypeErrorPointer
    function isUnsupportedValueErrorPointer(v: any): v is UnsupportedValueErrorPointer
}