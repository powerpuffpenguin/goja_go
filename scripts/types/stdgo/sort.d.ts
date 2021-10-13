declare module "stdgo/sort" {
    import {
        Float32, Float64,
        Int64, Int32, Int16, Int8, Int,
        Uint64, Uint32, Uint16, Uint8, Uint,
        Number, NumberLike,
        Byte, Bytes, Rune, Runes,
        Error,
        ReadChannel, WriteChannel, Channel,
        Slice, Map,
        Uintptr, Native,
    } from "stdgo/builtin";
    import * as builtin from "stdgo/builtin";

    function Float64s(x: builtin.Float64Slice): void
    function Float64sAreSorted(x: builtin.Float64Slice): boolean
    function Ints(x: builtin.IntSlice): void
    function IntsAreSorted(x: builtin.IntSlice): boolean
    function IsSorted(data: Interface): boolean
    function Search(n: int, f: (index: Int) => boolea): Int
    function SearchFloat64s(a: builtin.Float64Slice, x: Float64 | NumberLike): Int
    function SearchInts(a: builtin.IntSlice, x: Int | NumberLike): Int
    function SearchStrings(a: Array<string>, x: string): Int
    function Slice(x: Slice, less: (i: Int, j: Int) => boolean): void
    function SliceIsSorted(x: Slice, less: (i: Int, j: Int) => boolean): boolean
    function SliceStable(x: Slice, less: (i: Int, j: Int) => boolean): void
    function Sort(data: Interface): void
    function Stable(data: Interface): void
    function Strings(x: Array<string>): void
    function StringsAreSorted(x: Array<string>): boolean


    function Float64Slice(v: Slice<Float64>): Float64Slice
    interface Float64Slice extends Slice<Float64> {
        readonly __Float64Slice: Float64Slice
        Len(): Int
        Less(i: Int | NumberLike, j: Int | NumberLike): boolean
        Search(x: NumberLike): Int
        Sort(): void
        Swap(i: Int | NumberLike, j: Int | NumberLike): void
    }
    function IntSlice(v: Slice<Int64>): IntSlice
    interface IntSlice extends Slice<Int64> {
        readonly __IntSlice: IntSlice
        Len(): Int
        Less(i: Int | NumberLike, j: Int | NumberLike): boolean
        Search(x: Int | NumberLike): Int
        Sort(): void
        Swap(i: Int | NumberLike, j: Int | NumberLike): void
    }

    function Reverse(data: Interface): Interface
    interface Interface extends Native {
        Len(): Int
        Less(i: Int, j: Int): boolean
        Swap(i: Int, j: Int): void
    }

    function StringSlice(v: Array<string>): StringSlice
    interface StringSlice extends Slice<string> {
        readonly __StringSlice: StringSlice
        Len(): Int
        Less(i: Int | NumberLike, j: Int | NumberLike): boolean
        Search(x: string): Int
        Sort(): void
        Swap(i: Int | NumberLike, j: Int | NumberLike): void
    }

    function isFloat64Slice(v: any): v is Float64Slice
    function isInterface(v: any): v is Interface
    function isIntSlice(v: any): v is IntSlice
    function isStringSlice(v: any): v is StringSlice
}