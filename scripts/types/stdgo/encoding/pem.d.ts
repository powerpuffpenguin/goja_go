declare module "stdgo/encoding/pem" {
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
    import * as io from "stdgo/io";

    function Encode(out: io.Writer, b: BlockPointer): void
    function EncodeToMemory(b: BlockPointer): Bytes

    /** return (p *Block, rest []byte) */
    function Decode(data: Bytes): [BlockPointer, Bytes]
    interface BlockPointer extends Native {
        readonly __BlockPointer: BlockPointer
        Type: string            // The type, taken from the preamble (i.e. "RSA PRIVATE KEY").
        Headers: Map<string, string> // Optional headers.
        Bytes: Bytes            // The decoded bytes of the contents. Typically a DER encoded ASN.1 structure.
    }

    function isBlockPointer(v: any): v is BlockPointer
}