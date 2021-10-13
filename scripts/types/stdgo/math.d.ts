declare module "stdgo/math" {
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

    const E: Float64 // https://oeis.org/A001113
    const Pi: Float64 // https://oeis.org/A000796
    const Phi: Float64 // https://oeis.org/A001622
    const Sqrt2: Float64 // https://oeis.org/A002193
    const SqrtE: Float64 // https://oeis.org/A019774
    const SqrtPi: Float64 // https://oeis.org/A002161
    const SqrtPhi: Float64 // https://oeis.org/A139339
    const Ln2: Float64 // https://oeis.org/A002162
    const Log2E: Float64
    const Ln10: Float64 // https://oeis.org/A002392
    const Log10E: Float64

    const MaxFloat32: Float32 // 2**127 * (2**24 - 1) / 2**23
    const SmallestNonzeroFloat32: Float32 // 1 / 2**(127 - 1 + 23)

    const MaxFloat64: Float64// 2**1023 * (2**53 - 1) / 2**52
    const SmallestNonzeroFloat64: Float64 // 1 / 2**(1023 - 1 + 52)

    const MaxInt64: Int64
    const MaxInt32: Int32
    const MaxInt16: Int16
    const MaxInt8: Int8
    const MaxUint64: Uint64
    const MaxUint32: Uint32
    const MaxUint16: Uint16
    const MaxUint8: Uint8
    const MinInt64: Int64
    const MinInt32: Int32
    const MinInt16: Int16
    const MinInt8: Int8

    function Abs(x: Float64 | NumberLike): Float64
    function Acos(x: Float64 | NumberLike): Float64
    function Acosh(x: Float64 | NumberLike): Float64
    function Asin(x: Float64 | NumberLike): Float64
    function Asinh(x: Float64 | NumberLike): Float64
    function Atan(x: Float64 | NumberLike): Float64
    function Atan2(y: Float64 | NumberLike, x: Float64 | NumberLike): Float64
    function Atanh(x: Float64 | NumberLike): Float64
    function Cbrt(x: Float64 | NumberLike): Float64
    function Ceil(x: Float64 | NumberLike): Float64
    function Copysign(x, y: Float64 | NumberLike): Float64
    function Cos(x: Float64 | NumberLike): Float64
    function Cosh(x: Float64 | NumberLike): Float64
    function Dim(x: Float64 | NumberLike, y: Float64 | NumberLike): Float64
    function Erf(x: Float64 | NumberLike): Float64
    function Erfc(x: Float64 | NumberLike): Float64
    function Erfcinv(x: Float64 | NumberLike): Float64
    function Erfinv(x: Float64 | NumberLike): Float64
    function Exp(x: Float64 | NumberLike): Float64
    function Exp2(x: Float64 | NumberLike): Float64
    function Expm1(x: Float64 | NumberLike): Float64
    function FMA(x: Float64 | NumberLike, y: Float64 | NumberLike, z: Float64 | NumberLike): Float64
    function Float32bits(f: Float32 | NumberLike): Uint32
    function Float32frombits(b: Uint32 | NumberLike): Float32
    function Float64bits(f: Float64 | NumberLike): Uint64
    function Float64frombits(b: Uint64 | NumberLike): Float64
    function Floor(x: Float64 | NumberLike): Float64
    /** return (frac float64, exp int) */
    function Frexp(f: Float64 | NumberLike): [Float64, Int]
    function Gamma(x: Float64 | NumberLike): Float64
    function Hypot(p: Float64 | NumberLike, q: Float64 | NumberLike): Float64
    function Ilogb(x: Float64 | NumberLike): Int
    function Inf(sign: Int | NumberLike): Float64
    function IsInf(f: Float64 | NumberLike, sign: Int | NumberLike): boolean
    function IsNaN(f: Float64 | NumberLike): boolean
    function J0(x: Float64 | NumberLike): Float64
    function J1(x: Float64 | NumberLike): Float64
    function Jn(n: Int | NumberLike, x: Float64 | NumberLike): Float64
    function Ldexp(frac: Float64 | NumberLike, exp: Int | NumberLike): Float64
    /** return (lgamma float64, sign int) */
    function Lgamma(x: Float64 | NumberLike): [Float64, Int]
    function Log(x: Float64 | NumberLike): Float64
    function Log10(x: Float64 | NumberLike): Float64
    function Log1p(x: Float64 | NumberLike): Float64
    function Log2(x: Float64 | NumberLike): Float64
    function Logb(x: Float64 | NumberLike): Float64
    function Max(x: Float64 | NumberLike, y: Float64 | NumberLike): Float64
    function Min(x: Float64 | NumberLike, y: Float64 | NumberLike): Float64
    function Mod(x: Float64 | NumberLike, y: Float64 | NumberLike): Float64
    /** reutrn (int float64, frac float64) */
    function Modf(f: Float64 | NumberLike): [Float64, Float64]
    function NaN(): Float64
    function Nextafter(x: Float64 | NumberLike, y: Float64 | NumberLike): Float64
    function Nextafter32(x: Float32 | NumberLike, y: Float32 | NumberLike): Float32
    function Pow(x: Float64 | NumberLike, y: Float64 | NumberLike): Float64
    function Pow10(n: Int | NumberLike): Float64
    function Remainder(x: Float64 | NumberLike, y: Float64 | NumberLike): Float64
    function Round(x: Float64 | NumberLike): Float64
    function RoundToEven(x: Float64 | NumberLike): Float64
    function Signbit(x: Float64 | NumberLike): boolean
    function Sin(x: Float64 | NumberLike): Float64
    /** return (sin, cos float64) */
    function Sincos(x: Float64 | NumberLike): [Float64, Float64]
    function Sinh(x: Float64 | NumberLike): Float64
    function Sqrt(x: Float64 | NumberLike): Float64
    function Tan(x: Float64 | NumberLike): Float64
    function Tanh(x: Float64 | NumberLike): Float64
    function Trunc(x: Float64 | NumberLike): Float64
    function Y0(x: Float64 | NumberLike): Float64
    function Y1(x: Float64 | NumberLike): Float64
    function Yn(n: Int | NumberLike, x: Float64 | NumberLike): Float64
}