declare module "stdgo/os/signal" {
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
    import * as context from "stdgo/context";
    import * as os from "stdgo/os";

    function Ignore(...sig: Array<os.Signal>): void
    function Ignored(sig: os.Signal): boolean
    function Notify(c: WriteChannel<os.Signal>, ...sig: Array<os.Signal>): void
    /** return (ctx context.Context, stop context.CancelFunc) */
    function NotifyContext(parent: context.Context, ...signals: Array<os.Signal>): [context.Context, context.CancelFunc]
    function Reset(...sig: Array<os.Signal>): void
    function Stop(c: WriteChannel<os.Signal>): void
}