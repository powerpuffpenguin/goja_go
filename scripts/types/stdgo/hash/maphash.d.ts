declare module "stdgo/hash/maphash" {
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
    import * as hash from "stdgo/hash";

    interface HashPointer extends Native {
        readonly __HashPointer: HashPointer

        BlockSize(): Int
        Reset(): void
        Seed(): Seed
        SetSeed(seed: Seed): void
        Size(): Int
        Sum(b: Bytes): Bytes
        Sum64(): Uint64
        Write(b: Bytes): Int
        WriteByte(b: Byte): void
        WriteString(s: string): Int
    }

    function MakeSeed(): Seed
    interface Seed extends Native {
        readonly __Seed: Seed
    }

    function isHashPointer(v: any): v is HashPointer
    function isSeed(v: any): v is Seed
}