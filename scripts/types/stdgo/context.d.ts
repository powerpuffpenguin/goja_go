declare module "stdgo/context" {
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
    import * as time from "stdgo/time";

    const Canceled: Error
    const DeadlineExceeded: Error

    function WithCancel(parent: Context): [Context, CancelFunc]
    function WithDeadline(parent: Context, d: time.Time): [Context, CancelFunc]
    function WithTimeout(parent: Context, timeout: time.Duration): [Context, CancelFunc]

    type CancelFunc = () => void

    function Background(): Context
    function TODO(): Context
    function WithValue(parent: Context, key: any, val: any): Context

    interface Context extends Native {
        /** return (deadline time.Time, ok bool) */
        Deadline(): [time.Time, boolean]
        Done(): ReadChannel<null>
        Err(): Error
        Value(key: any): any
    }

    function isCancelFunc(v: any): v is CancelFunc
    function isContext(v: any): v is Context
}