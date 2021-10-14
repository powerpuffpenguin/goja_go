declare module "stdgo/math/bits" {
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

    const UintSize: number

    /** return (sum, carryOut uint) */
    function Add(x: Uint | NumberLike, y: Uint | NumberLike, carry: Uint | NumberLike): [Uint, Uint]
    /** return (sum, carryOut uint32) */
    function Add32(x: Uint32 | NumberLike, y: Uint32 | NumberLike, carry: Uint32 | NumberLike): [Uint32, Uint32]
    /** return (sum, carryOut uint64) */
    function Add64(x: Uint64 | NumberLike, y: Uint64 | NumberLike, carry: Uint64 | NumberLike): [Uint64, Uint64]
    /** return (quo, rem uint) */
    function Div(hi: Uint | NumberLike, lo: Uint | NumberLike, y: Uint | NumberLike): [Uint, Uint]
    /** return (quo, rem uint32) */
    function Div32(hi: Uint32 | NumberLike, lo: Uint32 | NumberLike, y: Uint32 | NumberLike): [Uint32, Uint32]
    /** return (quo, rem uint64) */
    function Div64(hi: Uint64 | NumberLike, lo: Uint64 | NumberLike, y: Uint64 | NumberLike): [Uint64, Uint64]
    function LeadingZeros(x: Uint | NumberLike): Int
    function LeadingZeros16(x: Uint16 | NumberLike): Int
    function LeadingZeros32(x: Uint32 | NumberLike): Int
    function LeadingZeros64(x: Uint64 | NumberLike): Int
    function LeadingZeros8(x: Uint8 | NumberLike): Int
    function Len(x: Uint | NumberLike): Int
    function Len16(x: Uint16 | NumberLike): Int
    function Len32(x: Uint32 | NumberLike): Int
    function Len64(x: Uint64 | NumberLike): Int
    function Len8(x: Uint8 | NumberLike): Int
    /** return (hi, lo uint) */
    function Mul(x: Uint | NumberLike, y: Uint | NumberLike): [Uint, Uint]
    /** return (hi, lo uint32) */
    function Mul32(x: Uint32 | NumberLike, y: Uint32 | NumberLike): [Uint32, Uint32]
    /** return (hi, lo uint64) */
    function Mul64(x: Uint64 | NumberLike, y: Uint64 | NumberLike): [Uint64, Uint64]
    function OnesCount(x: Uint | NumberLike): Int
    function OnesCount16(x: Uint16 | NumberLike): Int
    function OnesCount32(x: Uint32 | NumberLike): Int
    function OnesCount64(x: Uint64 | NumberLike): Int
    function OnesCount8(x: Uint8 | NumberLike): Int
    function Rem(hi: Uint | NumberLike, lo: Uint | NumberLike, y: Uint | NumberLike): Uint
    function Rem32(hi: Uint32 | NumberLike, lo: Uint32 | NumberLike, y: Uint32 | NumberLike): Uint32
    function Rem64(hi: Uint64 | NumberLike, lo: Uint64 | NumberLike, y: Uint64 | NumberLike): Uint64
    function Reverse(x: Uint | NumberLike): Uint
    function Reverse16(x: Uint16 | NumberLike): Uint16
    function Reverse32(x: Uint32 | NumberLike): Uint32
    function Reverse64(x: Uint64 | NumberLike): Uint64
    function Reverse8(x: Uint8 | NumberLike): Uint8
    function ReverseBytes(x: Uint | NumberLike): Uint
    function ReverseBytes16(x: Uint16 | NumberLike): Uint16
    function ReverseBytes32(x: Uint32 | NumberLike): Uint32
    function ReverseBytes64(x: Uint64 | NumberLike): Uint64
    function RotateLeft(x: Uint | NumberLike, k: Int | NumberLike): Uint
    function RotateLeft16(x: Uint16 | NumberLike, k: Int | NumberLike): Uint16
    function RotateLeft32(x: Uint32 | NumberLike, k: Int | NumberLike): Uint32
    function RotateLeft64(x: Uint64 | NumberLike, k: Int | NumberLike): Uint64
    function RotateLeft8(x: Uint8 | NumberLike, k: Int | NumberLike): Uint8
    /** return (diff, borrowOut uint) */
    function Sub(x: Uint | NumberLike, y: Uint | NumberLike, borrow: Uint | NumberLike): [Uint, Uint]
    /** return (diff, borrowOut uint32) */
    function Sub32(x: Uint32 | NumberLike, y: Uint32 | NumberLike, borrow: Uint32 | NumberLike): [Uint32, Uint32]
    /** return (diff, borrowOut uint64) */
    function Sub64(x: Uint64 | NumberLike, y: Uint64 | NumberLike, borrow: Uint64 | NumberLike): [Uint64, Uint64]
    function TrailingZeros(x: Uint): Int
    function TrailingZeros16(x: Uint16): Int
    function TrailingZeros32(x: Uint32): Int
    function TrailingZeros64(x: Uint64): Int
    function TrailingZeros8(x: Uint8): Int
}