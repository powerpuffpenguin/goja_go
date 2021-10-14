declare module "stdgo/math/big" {
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
    import * as math from "stdgo/math";
    import * as rand from "stdgo/math/rand";
    import * as fmt from "stdgo/fmt";

    const MaxExp = math.MaxInt32  // largest supported exponent
    const MinExp = math.MinInt32  // smallest supported exponent
    const MaxPrec = math.MaxUint32 // largest (theoretically) supported precision; likely memory-limited
    const MaxBase: Byte

    function Jacobi(x: IntPointer, y: IntPointer): Int

    function Accuracy(v: Int8 | NumberLike): Accuracy
    interface Accuracy extends Number {
        readonly __Accuracy: Accuracy
        String(): string
    }
    const Below = Accuracy(-1)
    const Exact = Accuracy(0)
    const Above = Accuracy(+1)

    interface ErrNaN extends Error { }

    function NewFloat(x: Float64 | NumberLike): FloatPointer
    /** return (f * Float, b int, err error) */
    function ParseFloat(s: string, base: Int | NumberLike, prec: Uint | NumberLike, mode: RoundingMode): [FloatPointer, Int]
    interface FloatPointer extends Native {
        readonly __FloatPointer: FloatPointer

        Abs(x: FloatPointer): FloatPointer
        Acc(): Accuracy
        Add(x: FloatPointer, y: FloatPointer): FloatPointer
        Append(buf: Bytes, fmt: Byte, prec: Int | NumberLike): Bytes
        Cmp(y: FloatPointer): Int
        Copy(x: FloatPointer): FloatPointer
        /** return (float32, Accuracy) */
        Float32(): [Float32, Accuracy]
        /** return (float64, Accuracy) */
        Float64(): [Float64, Accuracy]
        Format(s: fmt.State, format: Rune | NumberLike): void
        GobDecode(buf: Bytes): void
        GobEncode(): Bytes
        /** return (* Int, Accuracy) */
        Int(z: IntPointer): [IntPointer, Accuracy]
        /** return (int64, Accuracy) */
        Int64(): [Int64, Accuracy]
        IsInf(): boolean
        IsInt(): boolean
        MantExp(mant: FloatPointer): Int
        MarshalText(): Bytes
        MinPrec(): Uint
        Mode(): RoundingMode
        Mul(x: FloatPointer, y: FloatPointer): FloatPointer
        Neg(x: FloatPointer): FloatPointer
        /** return (f * Float, b int, err error) */
        Parse(s: string, base: Int | NumberLike): [FloatPointer, Int]
        Prec(): Uint
        Quo(x: FloatPointer, y: FloatPointer): FloatPointer
        /** return (* Rat, Accuracy) */
        Rat(z: RatPointer): [RatPointer, Accuracy]
        Scan(s: fmt.ScanState, ch: Rune | NumberLike): void
        Set(x: FloatPointer): FloatPointer
        SetFloat64(x: Float64 | NumberLike): FloatPointer
        SetInf(signbit: boolean): FloatPointer
        SetInt(x: IntPointer): FloatPointer
        SetInt64(x: Int64 | NumberLike): FloatPointer
        SetMantExp(mant: FloatPointer, exp: Int | NumberLike): FloatPointer
        SetMode(mode: RoundingMode): FloatPointer
        SetPrec(prec: Uint | NumberLike): FloatPointer
        SetRat(x: RatPointer): FloatPointer
        SetString(s: string): [FloatPointer, boolean]
        SetUint64(x: Uint64 | NumberLike): FloatPointer
        Sign(): Int
        Signbit(): boolean
        Sqrt(x: FloatPointer): FloatPointer
        String(): string
        Sub(x: FloatPointer, y: FloatPointer): FloatPointer
        Text(format: Byte | NumberLike, prec: Int | NumberLike): string
        Uint64(): [Uint64, Accuracy]
        UnmarshalText(text: Bytes): void
    }

    function NewInt(x: Int64 | NumberLike): IntPointer
    interface IntPointer extends Native {
        readonly __IntPointer: IntPointer

        Abs(x: IntPointer): IntPointer
        Add(x: IntPointer, y: IntPointer): IntPointer
        And(x: IntPointer, y: IntPointer): IntPointer
        AndNot(x: IntPointer, y: IntPointer): IntPointer
        Append(buf: Bytes, base: Int | NumberLike): Bytes
        Binomial(n: Int64 | NumberLike, k: Int64 | NumberLike): IntPointer
        Bit(i: Int | NumberLike): Uint
        BitLen(): Int
        Bits(): Array<Word>
        Bytes(): Bytes
        Cmp(y: IntPointer): Int
        CmpAbs(y: IntPointer): Int
        Div(x: IntPointer, y: IntPointer): IntPointer
        DivMod(x: IntPointer, y: IntPointer, m: IntPointer): [IntPointer, IntPointer]
        Exp(x: IntPointer, y: IntPointer, m: IntPointer): IntPointer
        FillBytes(buf: Bytes): Bytes
        Format(s: fmt.State, ch: Rune | NumberLike): void
        GCD(x: IntPointer, y: IntPointer, a: IntPointer, b: IntPointer): IntPointer
        GobDecode(buf: Bytes): void
        GobEncode(): Bytes
        Int64(): Int64
        IsInt64(): boolean
        IsUint64(): boolean
        Lsh(x: IntPointer, n: Uint | NumberLike): IntPointer
        MarshalJSON(): Bytes
        MarshalText(): Bytes
        Mod(x: IntPointer, y: IntPointer): IntPointer
        ModInverse(g: IntPointer, n: IntPointer): IntPointer
        ModSqrt(x: IntPointer, p: IntPointer): IntPointer
        Mul(x: IntPointer, y: IntPointer): IntPointer
        MulRange(a: Int64 | NumberLike, b: Int64 | NumberLike): IntPointer
        Neg(x: IntPointer): IntPointer
        Not(x: IntPointer): IntPointer
        Or(x: IntPointer, y: IntPointer): IntPointer
        ProbablyPrime(n: Int | NumberLike): boolean
        Quo(x: IntPointer, y: IntPointer): IntPointer
        QuoRem(x: IntPointer, y: IntPointer, r: IntPointer): [IntPointer, IntPointer]
        Rand(rnd: rand.RandPointer, n: IntPointer): IntPointer
        Rem(x: IntPointer, y: IntPointer): IntPointer
        Rsh(x: IntPointer, n: Uint | NumberLike): IntPointer
        Scan(s: fmt.ScanState, ch: Rune | NumberLike): void
        Set(x: IntPointer): IntPointer
        SetBit(x: IntPointer, i: Int | NumberLike, b: Uint | NumberLike): IntPointer
        SetBits(abs: Array<Word>): IntPointer
        SetBytes(buf: Bytes): IntPointer
        SetInt64(x: Int64 | NumberLike): IntPointer
        SetString(s: string, base: Int | NumberLike): [IntPointer, boolean]
        SetUint64(x: Uint64 | NumberLike): IntPointer
        Sign(): Int
        Sqrt(x: IntPointer): IntPointer
        String(): string
        Sub(x: IntPointer, y: IntPointer): IntPointer
        Text(base: Int | NumberLike): string
        TrailingZeroBits(): Uint
        Uint64(): Uint64
        UnmarshalJSON(text: Bytes): void
        UnmarshalText(text: Bytes): void
        Xor(x: IntPointer, y: IntPointer): IntPointer
    }

    function NewRat(a: Int64 | NumberLike, b: Int64 | NumberLike): RatPointer
    interface RatPointer extends Native {
        readonly __RatPointer: RatPointer

        Abs(x: RatPointer): RatPointer
        Add(x: RatPointer, y: RatPointer): RatPointer
        Cmp(y: RatPointer): Int
        Denom(): IntPointer
        /** return (f float32, exact bool) */
        Float32(): [Float32, boolean]
        /** return (f float64, exact bool) */
        Float64(): [Float64, boolean]
        FloatString(prec: Int | NumberLike): string
        GobDecode(buf: Bytes): void
        GobEncode(): Bytes
        Inv(x: RatPointer): RatPointer
        IsInt(): boolean
        MarshalText(): Bytes
        Mul(x: RatPointer, y: RatPointer): RatPointer
        Neg(x: RatPointer): RatPointer
        Num(): IntPointer
        Quo(x: RatPointer, y: RatPointer): RatPointer
        RatString(): string
        Scan(s: fmt.ScanState, ch: Rune | NumberLike): void
        Set(x: RatPointer): RatPointer
        SetFloat64(f: Float64 | NumberLike): RatPointer
        SetFrac(a: IntPointer, b: IntPointer): RatPointer
        SetFrac64(a: Int64 | NumberLike, b: Int64 | NumberLike): RatPointer
        SetInt(x: IntPointer): RatPointer
        SetInt64(x: Int64 | NumberLike): RatPointer
        SetString(s: string): [RatPointer, bool]
        SetUint64(x: Uint64 | NumberLike): RatPointer
        Sign(): Int
        String(): string
        Sub(x: RatPointer, y: RatPointer): RatPointer
        UnmarshalText(text: Bytes): void
    }

    function RoundingMode(v: Byte | NumberLike): RoundingMode
    interface RoundingMode extends Number {
        readonly __RoundingMode: RoundingMode

        String(): string
    }
    const ToNearestEven = RoundingMode(0) // == IEEE 754-2008 roundTiesToEven
    const ToNearestAway = RoundingMode(1)                    // == IEEE 754-2008 roundTiesToAway
    const ToZero = RoundingMode(2)                       // == IEEE 754-2008 roundTowardZero
    const AwayFromZero = RoundingMode(3)                     // no IEEE 754-2008 equivalent
    const ToNegativeInf = RoundingMode(4)                  // == IEEE 754-2008 roundTowardNegative
    const ToPositiveInf = RoundingMode(5)                   // == IEEE 754-2008 roundTowardPositive

    function Word(v: Uint | NumberLike): Word
    interface Word extends Number {
        readonly __Word: Word
    }

    function isAccuracy(v: any): v is Accuracy
    function isErrNaN(v: any): v is ErrNaN
    function isFloatPointer(v: any): v is FloatPointer
    function isIntPointer(v: any): v is IntPointer
    function isRatPointer(v: any): v is RatPointer
    function isRoundingMode(v: any): v is RoundingMode
    function isWord(v: any): v is Word
}