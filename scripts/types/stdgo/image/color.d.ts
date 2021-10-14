declare module "stdgo/image/color" {
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

    const Black = Gray16(0)
    const White = Gray16(0xffff)
    const Transparent = Alpha16(0)
    const Opaque = Alpha16(0xffff)

    function CMYKToRGB(c: Uint8 | NumberLike, m: Uint8 | NumberLike, y: Uint8 | NumberLike, k: Uint8 | NumberLike): [Uint8, Uint8, Uint8]
    function RGBToCMYK(r: Uint8 | NumberLike, g: Uint8 | NumberLike, b: Uint8 | NumberLike): [Uint8, Uint8, Uint8, Uint8]
    function RGBToYCbCr(r: Uint8 | NumberLike, g: Uint8 | NumberLike, b: Uint8 | NumberLike): [Uint8, Uint8, Uint8]
    function YCbCrToRGB(y: Uint8 | NumberLike, cb: Uint8 | NumberLike, cr: Uint8 | NumberLike): [Uint8, Uint8, Uint8]

    function Alpha(a: Uint8 | NumberLike): Alpha
    interface Alpha extends Color {
        readonly __Alpha: Alpha

        A: Uint8
    }

    function Alpha16(a: Uint16 | NumberLike): Alpha16
    interface Alpha16 extends Color {
        readonly __Alpha16: Alpha16

        A: Uint16
    }

    function CMYK(c: Uint8 | NumberLike, m: Uint8 | NumberLike, y: Uint8 | NumberLike, k: Uint8 | NumberLike): CMYK
    interface CMYK extends Color {
        readonly __CMYK: CMYK

        C: Uint8
        M: Uint8
        Y: Uint8
        K: Uint8
    }

    interface Color extends Native {
        /** return (r, g, b, a uint32) */
        RGBA(): [Uint32, Uint32, Uint32, Uint32]
    }

    function Gray(y: Uint8 | NumberLike): Gray
    interface Gray extends Color {
        readonly __Gray: Gray

        Y: Uint8
    }

    function Gray16(y: Uint16 | NumberLike): Gray16
    interface Gray16 extends Color {
        readonly __Gray16: Gray16

        Y: Uint16
    }

    function ModelFunc(f: (c: Color) => Color): Model
    interface Model extends Native {
        Convert(c: Color): Color
    }
    const RGBAModel: Model
    const RGBA64Model: Model
    const NRGBAModel: Model
    const NRGBA64Model: Model
    const AlphaModel: Model
    const Alpha16Model: Model
    const GrayModel: Model
    const Gray16Model: Model
    const CMYKModel: Model
    const NYCbCrAModel: Model
    const YCbCrModel: Model

    function NRGBA(r: Uint8 | NumberLike, g: Uint8 | NumberLike, b: Uint8 | NumberLike, a: Uint8 | NumberLike): NRGBA
    interface NRGBA extends Color {
        readonly __NRGBA: NRGBA

        R: Uint8
        G: Uint8
        B: Uint8
        A: Uint8
    }

    function NRGBA64(r: Uint16 | NumberLike, g: Uint16 | NumberLike, b: Uint16 | NumberLike, a: Uint16 | NumberLike): NRGBA
    interface NRGBA64 extends Color {
        readonly __NRGBA64: NRGBA64

        R: Uint16
        G: Uint16
        B: Uint16
        A: Uint16
    }

    function NYCbCrA(y: Uint8 | NumberLike, cb: Uint8 | NumberLike, cr: Uint8 | NumberLike, a: Uint8 | NumberLike): NYCbCrA
    interface NYCbCrA extends YCbCr {
        readonly __NYCbCrA: NYCbCrA
        YCbCr: YCbCr
        A: Uint8
    }

    function Palette(v: Slice<Color> | Array<Color>): Palette
    interface Palette extends Slice<Color> {
        Convert(c: Color): Color
        Index(c: Color): Int
    }

    function RGBA(r: Uint8 | NumberLike, g: Uint8 | NumberLike, b: Uint8 | NumberLike, a: Uint8 | NumberLike): RGBA
    interface RGBA extends Color {
        readonly __RGBA: RGBA

        R: Uint8
        G: Uint8
        B: Uint8
        A: Uint8
    }

    function RGBA64(r: Uint16 | NumberLike, g: Uint16 | NumberLike, b: Uint16 | NumberLike, a: Uint16 | NumberLike): RGBA64
    interface RGBA64 extends Color {
        readonly __RGBA64: RGBA64

        R: Uint16
        G: Uint16
        B: Uint16
        A: Uint16
    }

    function YCbCr(y: Uint8 | NumberLike, cb: Uint8 | NumberLike, cr: Uint8 | NumberLike): YCbCr
    interface YCbCr extends Color {
        readonly __YCbCr: YCbCr

        Y: Uint8
        Cb: Uint8
        Cr: Uint8
    }

    function isAlpha(v: any): v is Alpha
    function isAlpha16(v: any): v is Alpha16
    function isCMYK(v: any): v is CMYK
    function isColor(v: any): v is Color
    function isGray(v: any): v is Gray
    function isGray16(v: any): v is Gray16
    function isModel(v: any): v is Model
    function isNRGBA(v: any): v is NRGBA
    function isNRGBA64(v: any): v is NRGBA64
    function isNYCbCrA(v: any): v is NYCbCrA
    function isPalette(v: any): v is Palette
    function isRGBA(v: any): v is RGBA
    function isRGBA64(v: any): v is RGBA64
    function isYCbCr(v: any): v is YCbCr
}