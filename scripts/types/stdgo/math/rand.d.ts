declare module "stdgo/math/rand" {
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

    function ExpFloat64(): Float64
    function Float32(): Float32
    function Float64(): Float64
    function Int(): Int
    function Int31(): Int32
    function Int31n(n: Int32 | NumberLike): Int32
    function Int63(): Int64
    function Int63n(n: Int64 | NumberLike): Int64
    function Intn(n: Int | NumberLike): Int
    function NormFloat64(): Float64
    function Perm(n: Int | NumberLike): IntSlice
    function Read(p: Bytes): Int
    function Seed(seed: Int64 | NumberLike): void
    function Shuffle(n: Int | NumberLike, swap: (i: Int | NumberLike, j: Int | NumberLike) => void): void
    function Uint32(): Uint32
    function Uint64(): Uint64

    function New(src: Source): RandPointer
    interface RandPointer extends Native {
        readonly __RandPointer: RandPointer

        ExpFloat64(): Float64
        Float32(): Float32
        Float64(): Float64
        Int(): Int
        Int31(): Int32
        Int31n(n: Int32 | NumberLike): Int32
        Int63(): Int64
        Int63n(n: Int64 | NumberLike): Int64
        Intn(n: Int | NumberLike): Int
        NormFloat64(): Float64
        Perm(n: Int | NumberLike): IntSlice
        Read(p: Bytes): Int
        Seed(seed: Int64 | NumberLike): void
        Shuffle(n: Int | NumberLike, swap: (i: Int | NumberLike, j: Int | NumberLike) => void): void
        Uint32(): Uint32
        Uint64(): Uint64
    }

    function NewSource(seed: Int64 | NumberLike): Source
    interface Source extends Native {
        Int63(): Int64
        Seed(seed: Int64 | NumberLike): void
    }
    interface Source64 extends Source {
        Uint64(): Uint64
    }

    function NewZipf(r: RandPointer, s: Float64 | NumberLike, v: Float64 | NumberLike, imax: Uint64 | NumberLike): ZipfPointer
    interface ZipfPointer extends Native {
        readonly __ZipfPointer: ZipfPointer

        Uint64(): Uint64
    }
    
    function isRandPointer(v: any): v is RandPointer
    function isSource(v: any): v is Source
    function isSource64(v: any): v is Source64
    function isZipfPointer(v: any): v is ZipfPointer
}