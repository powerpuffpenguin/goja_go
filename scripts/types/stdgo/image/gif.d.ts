declare module "stdgo/image/gif" {
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
    import * as draw from "stdgo/image/draw";
    import * as io from "stdgo/io";

    const DisposalNone = Uint8(0x01)
    const DisposalBackground = Uint8(0x02)
    const DisposalPrevious = Uint8(0x03)

    function Decode(r: io.Reader): image.Image
    function DecodeConfig(r: io.Reader): image.Config
    function Encode(w: io.Writer, m: image.Image, o: OptionsPointer): void
    function EncodeAll(w: io.Writer, g: GIFPointer): void

    function DecodeAll(r: io.Reader): GIFPointer
    interface GIFPointer extends Native {
        readonly __GIFPointer: GIFPointer

        Image: Slice<image.PalettedPointer> // The successive images.
        Delay: IntSlice             // The successive delay times, one per frame, in 100ths of a second.
        // LoopCount controls the number of times an animation will be
        // restarted during display.
        // A LoopCount of 0 means to loop forever.
        // A LoopCount of -1 means to show each frame only once.
        // Otherwise, the animation is looped LoopCount+1 times.
        LoopCount: Int
        // Disposal is the successive disposal methods, one per frame. For
        // backwards compatibility, a nil Disposal is valid to pass to EncodeAll,
        // and implies that each frame's disposal method is 0 (no disposal
        // specified).
        Disposal: Byte
        // Config is the global color table (palette), width and height. A nil or
        // empty-color.Palette Config.ColorModel means that each frame has its own
        // color table and there is no global color table. Each frame's bounds must
        // be within the rectangle defined by the two points (0, 0) and
        // (Config.Width, Config.Height).
        //
        // For backwards compatibility, a zero-valued Config is valid to pass to
        // EncodeAll, and implies that the overall GIF's width and height equals
        // the first frame's bounds' Rectangle.Max point.
        Config: image.Config
        // BackgroundIndex is the background index in the global color table, for
        // use with the DisposalBackground disposal method.
        BackgroundIndex: Byte
    }

    function Options(numColors: Int | NumberLike, quantizer: draw.Quantizer, drawer: draw.Drawer): OptionsPointer
    interface OptionsPointer extends Native {
        readonly __OptionsPointer: OptionsPointer
        // NumColors is the maximum number of colors used in the image.
        // It ranges from 1 to 256.
        NumColors: Int

        // Quantizer is used to produce a palette with size NumColors.
        // palette.Plan9 is used in place of a nil Quantizer.
        Quantizer: draw.Quantizer

        // Drawer is used to convert the source image to the desired palette.
        // draw.FloydSteinberg is used in place of a nil Drawer.
        Drawer: draw.Drawer
    }

    function isGIFPointer(v: any): v is GIFPointer
    function isOptionsPointer(v: any): v is OptionsPointer
}