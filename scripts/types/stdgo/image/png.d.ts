declare module "stdgo/image/png" {
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
    import * as io from "stdgo/io";

    function Decode(r: io.Reader): image.Image
    function DecodeConfig(r: io.Reader): image.Config
    function Encode(w: io.Writer, m: image.Image): void

    function CompressionLevel(v: Int | NumberLike): CompressionLevel
    interface CompressionLevel extends Number {
        readonly __CompressionLevel: CompressionLevel
    }
    const DefaultCompression = CompressionLevel(0)
    const NoCompression = CompressionLevel(-1)
    const BestSpeed = CompressionLevel(-2)
    const BestCompression = CompressionLevel(-3)

    function Encoder(level: CompressionLevel | NumberLike, pool: EncoderBufferPool): EncoderPointer
    interface EncoderPointer extends Native {
        readonly __Encoder: Encoder
        CompressionLevel: CompressionLevel

        // BufferPool optionally specifies a buffer pool to get temporary
        // EncoderBuffers when encoding an image.
        BufferPool: EncoderBufferPool

        Encode(w: io.Writer, m: image.Image): void
    }

    interface EncoderBufferPointer extends Native {
        readonly __EncoderBufferPointer: EncoderBufferPointer
    }

    interface EncoderBufferPool extends Native {
        Get(): EncoderBufferPointer
        Put(eb: EncoderBufferPointer): void
    }
    interface FormatError extends Error {
        readonly __FormatError: FormatError
    }
    interface UnsupportedError extends Error {
        readonly __UnsupportedError: UnsupportedError
    }

    function isCompressionLevel(v: any): v is CompressionLevel
    function isEncoderPointer(v: any): v is EncoderPointer
    function isEncoderBufferPointer(v: any): v is EncoderBufferPointer
    function isEncoderBufferPool(v: any): v is EncoderBufferPool
    function isFormatError(v: any): v is FormatError
    function isUnsupportedError(v: any): v is UnsupportedError
}