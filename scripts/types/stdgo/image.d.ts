declare module "stdgo/image" {
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
    import * as color from "stdgo/image/color";
    import * as io from "stdgo/io";

    // Black is an opaque black uniform image.
    const Black = NewUniform(color.Black)
    // White is an opaque white uniform image.
    const White = NewUniform(color.White)
    // Transparent is a fully transparent uniform image.
    const Transparent = NewUniform(color.Transparent)
    // Opaque is a fully opaque uniform image.
    const Opaque = NewUniform(color.Opaque)

    const ErrFormat: Error

    function RegisterFormat(name: string, magic: string,
        decode: (r: io.Reader) => [Image, Error],
        decodeConfig: (r: io.Reader) => [Config, Error],
    ): void

    function NewAlpha(r: Rectangle): AlphaPointer
    interface AlphaPointer extends Native {
        readonly __AlphaPointer: AlphaPointer

        // Pix holds the image's pixels, as alpha values. The pixel at
        // (x, y) starts at Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*1].
        Pix: Uint8Slice
        // Stride is the Pix stride (in bytes) between vertically adjacent pixels.
        Stride: Int
        // Rect is the image's bounds.
        Rect: Rectangle

        AlphaAt(x: Int | NumberLike, y: Int | NumberLike): color.Alpha
        At(x: Int | NumberLike, y: Int | NumberLike): color.Color
        Bounds(): Rectangle
        ColorModel(): color.Model
        Opaque(): boolean
        PixOffset(x: Int | NumberLike, y: Int | NumberLike): Int
        Set(x: Int | NumberLike, y: Int | NumberLike, c: color.Color): void
        SetAlpha(x: Int | NumberLike, y: Int | NumberLike, c: color.Alpha): void
        SubImage(r: Rectangle): Image
    }

    function NewAlpha16(r: Rectangle): Alpha16Pointer
    interface Alpha16Pointer extends Native {
        readonly __Alpha16Pointer: Alpha16Pointer

        // Pix holds the image's pixels, as alpha values in big-endian format. The pixel at
        // (x, y) starts at Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*2].
        Pix: Uint8Slice
        // Stride is the Pix stride (in bytes) between vertically adjacent pixels.
        Stride: Int
        // Rect is the image's bounds.
        Rect: Rectangle

        Alpha16At(x: Int | NumberLike, y: Int | NumberLike): color.Alpha16
        At(x: Int | NumberLike, y: Int | NumberLike): color.Color
        Bounds(): Rectangle
        ColorModel(): color.Model
        Opaque(): boolean
        PixOffset(x: Int | NumberLike, y: Int | NumberLike): Int
        Set(x: Int | NumberLike, y: Int | NumberLike, c: color.Color): void
        SetAlpha16(x: Int | NumberLike, y: Int | NumberLike, c: color.Alpha16): void
        SubImage(r: Rectangle): Image
    }

    function NewCMYK(r: Rectangle): CMYKPointer
    interface CMYKPointer extends Native {
        readonly __CMYKPointer: CMYKPointer

        // Pix holds the image's pixels, in C, M, Y, K order. The pixel at
        // (x, y) starts at Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*4].
        Pix: Uint8Slice
        // Stride is the Pix stride (in bytes) between vertically adjacent pixels.
        Stride: Int
        // Rect is the image's bounds.
        Rect: Rectangle

        At(x: Int | NumberLike, y: Int | NumberLike): color.Color
        Bounds(): Rectangle
        CMYKAt(x: Int | NumberLike, y: Int | NumberLike): color.CMYK
        ColorModel(): color.Model
        Opaque(): boolean
        PixOffset(x: Int | NumberLike, y: Int | NumberLike): Int
        Set(x: Int | NumberLike, y: Int | NumberLike, c: color.Color): void
        SetCMYK(x: Int | NumberLike, y: Int | NumberLike, c: color.CMYK): void
        SubImage(r: Rectangle): Image
    }

    /** return (Config, string, error) */
    function DecodeConfig(r: io.Reader): [Config, string]
    interface Config extends Native {
        readonly __Config: Config

        ColorModel: color.Model
        Width: Int
        Height: Int
    }

    function NewGray(r: Rectangle): GrayPointer
    interface GrayPointer extends Native {
        readonly __GrayPointer: GrayPointer

        // Pix holds the image's pixels, as gray values. The pixel at
        // (x, y) starts at Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*1].
        Pix: Uint8Slice
        // Stride is the Pix stride (in bytes) between vertically adjacent pixels.
        Stride: Int
        // Rect is the image's bounds.
        Rect: Rectangle

        At(x: Int | NumberLike, y: Int | NumberLike): color.Color
        Bounds(): Rectangle
        ColorModel(): color.Model
        GrayAt(x: Int | NumberLike, y: Int | NumberLike): color.Gray
        Opaque(): boolean
        PixOffset(x: Int | NumberLike, y: Int | NumberLike): Int
        Set(x: Int | NumberLike, y: Int | NumberLike, c: color.Color): void
        SetGray(x: Int | NumberLike, y: Int | NumberLike, c: color.Gray): void
        SubImage(r: Rectangle): Image
    }

    function NewGray16(r: Rectangle): Gray16Pointer
    interface Gray16Pointer extends Native {
        readonly __Gray16Pointer: Gray16Pointer

        // Pix holds the image's pixels, as gray values in big-endian format. The pixel at
        // (x, y) starts at Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*2].
        Pix: Uint8Slice
        // Stride is the Pix stride (in bytes) between vertically adjacent pixels.
        Stride: Int
        // Rect is the image's bounds.
        Rect: Rectangle

        At(x: Int | NumberLike, y: Int | NumberLike): color.Color
        Bounds(): Rectangle
        ColorModel(): color.Model
        Gray16At(x: Int | NumberLike, y: Int | NumberLike): color.Gray16
        Opaque(): boolean
        PixOffset(x: Int | NumberLike, y: Int | NumberLike): Int
        Set(x: Int | NumberLike, y: Int | NumberLike, c: color.Color): void
        SetGray16(x: Int | NumberLike, y: Int | NumberLike, c: color.Gray16): void
        SubImage(r: Rectangle): Image
    }

    /** return (Image, string, error) */
    function Decode(r: io.Reader): [Image, string]
    interface Image extends Native {
        // ColorModel returns the Image's color model.
        ColorModel(): color.Model
        // Bounds returns the domain for which At can return non-zero color.
        // The bounds do not necessarily contain the point (0, 0).
        Bounds(): Rectangle
        // At returns the color of the pixel at (x, y).
        // At(Bounds().Min.X, Bounds().Min.Y) returns the upper-left pixel of the grid.
        // At(Bounds().Max.X-1, Bounds().Max.Y-1) returns the lower-right one.
        At(x: Int | NumberLike, y: Int | NumberLike): color.Color
    }

    function NewNRGBA(r: Rectangle): NRGBAPointer
    interface NRGBAPointer extends Native {
        readonly __NRGBAPointer: NRGBAPointer

        // Pix holds the image's pixels, in R, G, B, A order. The pixel at
        // (x, y) starts at Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*4].
        Pix: Uint8Slice
        // Stride is the Pix stride (in bytes) between vertically adjacent pixels.
        Stride: Int
        // Rect is the image's bounds.
        Rect: Rectangle

        At(x: Int | NumberLike, y: Int | NumberLike): color.Color
        Bounds(): Rectangle
        ColorModel(): color.Model
        NRGBAAt(x: Int | NumberLike, y: Int | NumberLike): color.NRGBA
        Opaque(): boolean
        PixOffset(x: Int | NumberLike, y: Int | NumberLike): Int
        Set(x: Int | NumberLike, y: Int | NumberLike, c: color.Color): void
        SetNRGBA(x: Int | NumberLike, y: Int | NumberLike, c: color.NRGBA): void
        SubImage(r: Rectangle): Image
    }

    function NewNRGBA64(r: Rectangle): NRGBA64Pointer
    interface NRGBA64Pointer extends Native {
        readonly __NRGBA64Pointer: NRGBA64Pointer

        // Pix holds the image's pixels, in R, G, B, A order and big-endian format. The pixel at
        // (x, y) starts at Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*8].
        Pix: Uint8Slice
        // Stride is the Pix stride (in bytes) between vertically adjacent pixels.
        Stride: Int
        // Rect is the image's bounds.
        Rect: Rectangle

        At(x: Int | NumberLike, y: Int | NumberLike): color.Color
        Bounds(): Rectangle
        ColorModel(): color.Model
        NRGBA64At(x: Int | NumberLike, y: Int | NumberLike): color.NRGBA64
        Opaque(): boolean
        PixOffset(x: Int | NumberLike, y: Int | NumberLike): Int
        Set(x: Int | NumberLike, y: Int | NumberLike, c: color.Color): void
        SetNRGBA64(x: Int | NumberLike, y: Int | NumberLike, c: color.NRGBA64): void
        SubImage(r: Rectangle): Image
    }

    function NewNYCbCrA(r: Rectangle, subsampleRatio: YCbCrSubsampleRatio): NYCbCrAPointer
    interface NYCbCrAPointer extends YCbCrPointer {
        readonly __NYCbCrAPointer: NYCbCrAPointer

        YCbCr: YCbCrPointer
        A: Uint8Slice
        AStride: Int

        AOffset(x: Int | NumberLike, y: Int | NumberLike): Int
        At(x: Int | NumberLike, y: Int | NumberLike): color.Color
        ColorModel(): color.Model
        NYCbCrAAt(x: Int | NumberLike, y: Int | NumberLike): color.NYCbCrA
        Opaque(): boolean
        RGBA64At(x: Int | NumberLike, y: Int | NumberLike): color.RGBA64
        SubImage(r: Rectangle): Image
    }

    function NewPaletted(r: Rectangle, p: color.Palette): PalettedPointer
    interface PalettedPointer extends Native {
        readonly __PalettedPointer: PalettedPointer

        // Pix holds the image's pixels, as palette indices. The pixel at
        // (x, y) starts at Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*1].
        Pix: Uint8Slice
        // Stride is the Pix stride (in bytes) between vertically adjacent pixels.
        Stride: Int
        // Rect is the image's bounds.
        Rect: Rectangle
        // Palette is the image's palette.
        Palette: color.Palette

        At(x: Int | NumberLike, y: Int | NumberLike): color.Color
        Bounds(): Rectangle
        ColorIndexAt(x: Int | NumberLike, y: Int | NumberLike): Uint8
        ColorModel(): color.Model
        Opaque(): boolean
        PixOffset(x: Int | NumberLike, y: Int | NumberLike): Int
        Set(x: Int | NumberLike, y: Int | NumberLike, c: color.Color): void
        SetColorIndex(x: Int | NumberLike, y: Int | NumberLike, index: Uint8 | NumberLike): void
        SubImage(r: Rectangle): Image
    }

    interface PalettedImage extends Image {
        // ColorIndexAt returns the palette index of the pixel at (x, y).
        ColorIndexAt(x: Int | NumberLike, y: Int | NumberLike): Uint8
    }

    function Pt(x: Int | NumberLike, y: Int | NumberLike): Point
    interface Point extends Native {
        readonly __Point: Point

        X: Int
        Y: Int

        Add(q: Point): Point
        Div(k: Int | NumberLike): Point
        Eq(q: Point): boolean
        In(r: Rectangle): boolean
        Mod(r: Rectangle): Point
        Mul(k: Int | NumberLike): Point
        String(): string
        Sub(q: Point): Point
    }

    function NewRGBA(r: Rectangle): RGBAPointer
    interface RGBAPointer extends Native {
        readonly __RGBAPointer: RGBAPointer

        // Pix holds the image's pixels, in R, G, B, A order. The pixel at
        // (x, y) starts at Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*4].
        Pix: Uint8Slice
        // Stride is the Pix stride (in bytes) between vertically adjacent pixels.
        Stride: Int
        // Rect is the image's bounds.
        Rect: Rectangle

        At(x: Int | NumberLike, y: Int | NumberLike): color.Color
        Bounds(): Rectangle
        ColorModel(): color.Model
        Opaque(): boolean
        PixOffset(x: Int | NumberLike, y: Int | NumberLike): Int
        RGBAAt(x: Int | NumberLike, y: Int | NumberLike): color.RGBA
        Set(x: Int | NumberLike, y: Int | NumberLike, c: color.Color): void
        SetRGBA(x: Int | NumberLike, y: Int | NumberLike, c: color.RGBA): void
        SubImage(r: Rectangle): Image
    }

    function NewRGBA64(r: Rectangle): RGBA64Pointer
    interface RGBA64Pointer extends Native {
        readonly __RGBA64Pointer: RGBA64Pointer

        // Pix holds the image's pixels, in R, G, B, A order and big-endian format. The pixel at
        // (x, y) starts at Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*8].
        Pix: Uint8Slice
        // Stride is the Pix stride (in bytes) between vertically adjacent pixels.
        Stride: Int
        // Rect is the image's bounds.
        Rect: Rectangle

        At(x: Int | NumberLike, y: Int | NumberLike): color.Color
        Bounds(): Rectangle
        ColorModel(): color.Model
        Opaque(): boolean
        PixOffset(x: Int | NumberLike, y: Int | NumberLike): Int
        RGBA64At(x: Int | NumberLike, y: Int | NumberLike): color.RGBA64
        Set(x: Int | NumberLike, y: Int | NumberLike, c: color.Color): void
        SetRGBA64(x: Int | NumberLike, y: Int | NumberLike, c: color.RGBA64): void
        SubImage(r: Rectangle): Image
    }

    function Rect(x0: Int | NumberLike, y0: Int | NumberLike, x1: Int | NumberLike, y1: Int | NumberLike): Rectangle
    interface Rectangle extends Native {
        readonly __Rectangle: Rectangle

        Min: Point
        Max: Point

        Add(p: Point): Rectangle
        At(x: Int | NumberLike, y: Int | NumberLike): color.Color
        Bounds(): Rectangle
        Canon(): Rectangle
        ColorModel(): color.Model
        Dx(): Int
        Dy(): Int
        Empty(): boolean
        Eq(s: Rectangle): boolean
        In(s: Rectangle): boolean
        Inset(n: Int | NumberLike): Rectangle
        Intersect(s: Rectangle): Rectangle
        Overlaps(s: Rectangle): boolean
        Size(): Point
        String(): string
        Sub(p: Point): Rectangle
        Union(s: Rectangle): Rectangle
    }

    function NewUniform(c: color.Color): UniformPointer
    interface UniformPointer extends color.Color {
        readonly __UniformPointer: UniformPointer

        C: color.Color

        At(x: Int | NumberLike, y: Int | NumberLike): color.Color
        Bounds(): Rectangle
        ColorModel(): color.Model
        Convert(c: color.Color): color.Color
        Opaque(): boolean
    }

    function NewYCbCr(r: Rectangle, subsampleRatio: YCbCrSubsampleRatio): YCbCrPointer
    interface YCbCrPointer extends Native {
        readonly __YCbCrPointer: YCbCrPointer

        Y: Uint8Slice
        Cb: Uint8Slice
        Cr: Uint8Slice
        YStride: Int
        CStride: Int
        SubsampleRatio: YCbCrSubsampleRatio
        Rect: Rectangle

        At(x: Int | NumberLike, y: Int | NumberLike): color.Color
        Bounds(): Rectangle
        COffset(x: Int | NumberLike, y: Int | NumberLike): Int
        ColorModel(): color.Model
        Opaque(): boolean
        SubImage(r: Rectangle): Image
        YCbCrAt(x: Int | NumberLike, y: Int | NumberLike): color.YCbCr
        YOffset(x: Int | NumberLike, y: Int | NumberLike): Int
    }

    function YCbCrSubsampleRatio(v: Int | NumberLike): YCbCrSubsampleRatio
    interface YCbCrSubsampleRatio extends Number {
        String(): string
    }
    const YCbCrSubsampleRatio444 = YCbCrSubsampleRatio(0)
    const YCbCrSubsampleRatio422 = YCbCrSubsampleRatio(1)
    const YCbCrSubsampleRatio420 = YCbCrSubsampleRatio(2)
    const YCbCrSubsampleRatio440 = YCbCrSubsampleRatio(3)
    const YCbCrSubsampleRatio411 = YCbCrSubsampleRatio(4)
    const YCbCrSubsampleRatio410 = YCbCrSubsampleRatio(5)

    function isAlphaPointer(v: any): v is AlphaPointer
    function isAlpha16Pointer(v: any): v is Alpha16Pointer
    function isCMYKPointer(v: any): v is CMYKPointer
    function isConfig(v: any): v is Config
    function isGrayPointer(v: any): v is GrayPointer
    function isGray16Pointer(v: any): v is Gray16Pointer
    function isImage(v: any): v is Image
    function isNRGBAPointer(v: any): v is NRGBAPointer
    function isNRGBA64Pointer(v: any): v is NRGBA64Pointer
    function isNYCbCrAPointer(v: any): v is NYCbCrAPointer
    function isPalettedPointer(v: any): v is PalettedPointer
    function isPalettedImage(v: any): v is PalettedImage
    function isPoint(v: any): v is Point
    function isRGBAPointer(v: any): v is RGBAPointer
    function isRGBA64Pointer(v: any): v is RGBA64Pointer
    function isRectangle(v: any): v is Rectangle
    function isUniformPointer(v: any): v is UniformPointer
    function isYCbCrPointer(v: any): v is YCbCrPointer
    function isYCbCrSubsampleRatio(v: any): v is YCbCrSubsampleRatio
}