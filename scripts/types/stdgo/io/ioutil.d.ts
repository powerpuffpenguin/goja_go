declare module "stdgo/io/ioutil" {
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

    import * as io from 'stdgo/io'
    import * as os from 'stdgo/os'
    import * as fs from 'stdgo/io/fs'
    const Discard: io.Writer

    function NopCloser(r: io.Reader): io.ReadCloser
    function ReadAll(r: io.Reader): Bytes
    function ReadDir(dirname: string): Array<fs.FileInfo>
    function ReadFile(filename: string): Bytes
    function TempDir(dir: sring, pattern: string): string
    function TempFile(dir: string, pattern: string): os.File
    function WriteFile(filename: string, data: Bytes, perm: fs.FileMode)
}