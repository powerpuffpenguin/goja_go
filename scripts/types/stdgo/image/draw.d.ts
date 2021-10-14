declare module "stdgo/image/draw" {
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
    import * as image from "stdgo/image";
    import * as color from "stdgo/image/color";

    function Draw(dst: Image, r: image.Rectangle, src: image.Image, sp: image.Point, op: Op | NumberLike): void
    function DrawMask(dst: Image, r: image.Rectangle, src: image.Image, sp: image.Point, mask: image.Image, mp: image.Point, op: Op | NumberLike): void

    interface Drawer extends Native {
        // Draw aligns r.Min in dst with sp in src and then replaces the
        // rectangle r in dst with the result of drawing src on dst.
        Draw(dst: Image, r: image.Rectangle, src: image.Image, sp: image.Point): void
    }
    interface Image extends image.Image {
        Image: image.Image
        Set(x: Int | NumberLike, y: Int | NumberLike, c: color.Color): void
    }

    function Op(v: Int | NumberLike): Op
    interface Op extends Number {
        readonly __Op: Op

        Draw(dst: Image, r: image.Rectangle, src: image.Image, sp: image.Point): void
    }
    // Over specifies “(src in mask) over dst”.
    const Over = Op(0)
    // Src specifies “src in mask”.
    const Src = Op(1)

    interface Quantizer extends Native {
        // Quantize appends up to cap(p) - len(p) colors to p and returns the
        // updated palette suitable for converting m to a paletted image.
        Quantize(p: color.Palette, m: image.Image): color.Palette
    }
    function isDrawer(v: any): v is Drawer
    function isImage(v: any): v is Image
    function isQuantizer(v: any): v is Quantizer
}