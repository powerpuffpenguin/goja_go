declare module "stdgo/strconv" {
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
        Complex128,
    } from "stdgo/builtin";

    const IntSize: number
    const ErrRange: Error
    const ErrSyntax: Error

    function AppendBool(dst: Bytes, b: boolean): Bytes
    function AppendFloat(dst: Bytes, f: Float64 | NumberLike, fmt: Byte | NumberLike, prec: Int | NumberLike, bitSize: Int | NumberLike): Bytes
    function AppendInt(dst: Bytes, i: Int64 | NumberLike, base: Int | NumberLike): Bytes
    function AppendQuote(dst: Bytes, s: string): Bytes
    function AppendQuoteRune(dst: Bytes, r: Rune | NumberLike): Bytes
    function AppendQuoteRuneToASCII(dst: Bytes, r: Rune | NumberLike): Bytes
    function AppendQuoteRuneToGraphic(dst: Bytes, r: Rune | NumberLike): Bytes
    function AppendQuoteToASCII(dst: Bytes, s: string): Bytes
    function AppendQuoteToGraphic(dst: Bytes, s: string): Bytes
    function AppendUint(dst: Bytes, i: Uint64 | NumberLike, base: Int | NumberLike): Bytes
    function Atoi(s: string): Int
    function CanBackquote(s: string): boolean
    function FormatBool(b: boolean): string
    function FormatComplex(c: Complex128, fmt: Byte | NumberLike, prec: Int | NumberLike, bitSize: Int | NumberLike): string
    function FormatFloat(f: Float64 | NumberLike, fmt: Byte | NumberLike, prec: Int | NumberLike, bitSize: Int | NumberLike): string
    function FormatInt(i: Int64 | NumberLike, base: Int | NumberLike): string
    function FormatUint(i: Uint64 | NumberLike, base: Int | NumberLike): string
    function IsGraphic(r: Rune): boolean
    function IsPrint(r: Rune): boolean
    function Itoa(i: Int | NumberLike): string
    function ParseBool(str: string): boolean
    function ParseComplex(s: string, bitSize: Int | NumberLike): Complex128
    function ParseFloat(s: string, bitSize: Int | NumberLike): Float64
    function ParseInt(s: string, base: Int | NumberLike, bitSize: Int | NumberLike): Int64
    function ParseUint(s: string, base: Int | NumberLike, bitSize: Int | NumberLike): Uint64
    function Quote(s: string): string
    function QuoteRune(r: Rune | NumberLike): string
    function QuoteRuneToASCII(r: Rune | NumberLike): string
    function QuoteRuneToGraphic(r: Rune | NumberLike): string
    function QuoteToASCII(s: string): string
    function QuoteToGraphic(s: string): string
    function Unquote(s: string): string
    /** return (value rune, multibyte bool, tail string, err error) */
    function UnquoteChar(s: string, quote: Byte | NumberLike): [Rune, boolean, string]

    interface NumErrorPointer extends Error {
        Func: string // the failing function (ParseBool, ParseInt, ParseUint, ParseFloat, ParseComplex)
        Num: string // the input
        Err: Error  // the reason the conversion failed (e.g. ErrRange, ErrSyntax, etc.)

        Unwrap(): Error
    }
    function isNumErrorPointer(v: any): v is NumErrorPointer
}