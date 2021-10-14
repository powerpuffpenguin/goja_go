declare module "stdgo/math/cmplx" {
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

    function Abs(x: Complex128): Float64
    function Acos(x: Complex128): Complex128
    function Acosh(x: Complex128): Complex128
    function Asin(x: Complex128): Complex128
    function Asinh(x: Complex128): Complex128
    function Atan(x: Complex128): Complex128
    function Atanh(x: Complex128): Complex128
    function Conj(x: Complex128): Complex128
    function Cos(x: Complex128): Complex128
    function Cosh(x: Complex128): Complex128
    function Cot(x: Complex128): Complex128
    function Exp(x: Complex128): Complex128
    function Inf(): Complex128
    function IsInf(x: Complex128): boolean
    function IsNaN(x: Complex128): boolean
    function Log(x: Complex128): Complex128
    function Log10(x: Complex128): Complex128
    function NaN(): Complex128
    function Phase(x: Complex128): Float64
    /** return (r, θ float64) */
    function Polar(x: Complex128): [Float64, Float64]
    function Pow(x: Complex128, y: Complex128): Complex128
    function Rect(r: Float64 | NumberLike, θ: Float64 | NumberLike): Complex128
    function Sin(x: Complex128): Complex128
    function Sinh(x: Complex128): Complex128
    function Sqrt(x: Complex128): Complex128
    function Tan(x: Complex128): Complex128
    function Tanh(x: Complex128): Complex128
}