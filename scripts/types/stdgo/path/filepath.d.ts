declare module "stdgo/path/filepath" {
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
    import *as os from "stdgo/os";
    import *as fs from "stdgo/io/fs";
    const Separator = os.PathSeparator
    const ListSeparator = os.PathListSeparator

    const ErrBadPattern: Error
    const SkipDir: Error

    function Abs(path: string): string
    function Base(path: string): string
    function Clean(path: string): string
    function Dir(path: string): string
    function EvalSymlinks(path: string): string
    function Ext(path: string): string
    function FromSlash(path: string): string
    function Glob(pattern: string): Array<string>
    function IsAbs(path: string): boolean
    function Join(...elem: Array<string>): string
    function Match(pattern: string, name: string): boolean
    function Rel(basepath: string, targpath: string): string
    /** return (dir, file string) */
    function Split(path: string): [string, string]
    function SplitList(path: string): Array<string>
    function ToSlash(path: string): string
    function VolumeName(path: string): string
    function Walk(root: string, fn: WalkFunc)
    function WalkDir(root: string, fn: fs.WalkDirFunc)

    type WalkFunc = (path: string, info: fs.FileInfo, err: Error) => Error

    function isWalkFunc(v: any): v is WalkFunc
}