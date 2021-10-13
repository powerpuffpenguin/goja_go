declare module "stdgo/compress/lzw" {
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
    import *as io from "stdgo/io";

    function NewReader(r: io.Reader, order: Order, litWidth: Int | NumberLike): io.ReadCloser
    function NewWriter(w: io.Writer, order: Order, litWidth: Int | NumberLike): io.WriteCloser

    function Order(): Order
    interface Order extends Number {
        readonly __Order: Order
    }
    const LSB: Order
    const MSB: Order

    function isOrder(v: any): v is Order
}