declare module "stdgo/mime" {
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


    // BEncoding represents Base64 encoding scheme as defined by RFC 2045.
    const BEncoding = WordEncoder('b')
    // QEncoding represents the Q-encoding scheme as defined by RFC 2047.
    const QEncoding = WordEncoder('q')

    const ErrInvalidMediaParameter: Error

    function AddExtensionType(ext: string, typ: string): void
    function ExtensionsByType(typ: string): Array<string>
    function FormatMediaType(t: string, param: Map<string, string>): string
    /** return (mediatype string, params map[string]string, err error) */
    function ParseMediaType(v: string): [string, Map<string, string>]
    function TypeByExtension(ext: string): string

    interface WordDecoderPointer extends Native {
        readonly __WordDecoderPointer: WordDecoderPointer
        // CharsetReader, if non-nil, defines a function to generate
        // charset-conversion readers, converting from the provided
        // charset into UTF-8.
        // Charsets are always lower-case. utf-8, iso-8859-1 and us-ascii charsets
        // are handled by default.
        // One of the CharsetReader's result values must be non-nil.
        CharsetReader: (charset: string, input: io.Reader) => [io.Reader, Error]

        Decode(word: string): string
        DecodeHeader(header: string): string
    }

    function WordEncoder(v: Byte | Number): WordEncoder
    interface WordEncoder extends Number {
        readonly __WordEncoder: WordEncoder
        Encode(charset: string, s: string): string
    }
    function isWordDecoderPointer(v: any): v is WordDecoderPointer
    function isWordEncoder(v: any): v is WordEncoder
}