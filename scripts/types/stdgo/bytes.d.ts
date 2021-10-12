declare module "stdgo/bytes" {
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
    import * as unicode from 'stdgo/unicode';
    import * as io from 'stdgo/io';

    function Bytes(str: string): Bytes
    function String(b: Bytes): string

    const MinRead = 512
    const ErrTooLarge: Error

    function Compare(a: Bytes, b: Bytes): Int
    function Contains(b: Bytes, subslice: Bytes): boolean
    function ContainsAny(b: Bytes, chars: string): boolean
    function ContainsRune(b: Bytes, r: Rune | NumberLike): boolean
    function Count(s: Bytes, sep: Bytes): Int
    function Equal(a: Bytes, b: Bytes): boolean
    function EqualFold(s: Bytes, t: Bytes): boolean
    function Fields(s: Bytes): Array<Bytes>
    function FieldsFunc(s: Bytes, f: (rune: Rune) => boolean): Array<Bytes>
    function HasPrefix(s: Bytes, prefix: Bytes): boolean
    function HasSuffix(s: Bytes, suffix: Bytes): boolean
    function Index(s: Bytes, sep: Bytes): Int
    function IndexAny(s: Bytes, chars: string): Int
    function IndexByte(b: Bytes, c: Byte): Int
    function IndexFunc(s: Bytes, f: (r: Rune) => boolean): Int
    function IndexRune(s: Bytes, r: Rune | NumberLike): Int
    function Join(s: Array<Bytes>, sep: Bytes): Bytes
    function LastIndex(s: Bytes, sep: Bytes): Int
    function LastIndexAny(s: Bytes, chars: string): Int
    function LastIndexByte(s: Bytes, c: Byte): Int
    function LastIndexFunc(s: Bytes, f: (r: Rune) => boolean): Int
    function Map(mapping: (r: Rune) => Rune, s: Bytes): Bytes
    function Repeat(b: Bytes, count: Int | NumberLike): Bytes
    function Replace(s: Bytes, old: Bytes, newb: Bytes, n: Int | NumberLike): Bytes
    function ReplaceAll(s: Bytes, old: Bytes, newb: Bytes): Bytes
    function Runes(s: Bytes): Runes
    function Split(s: Bytes, sep: Bytes): Array<Bytes>
    function SplitAfter(s: Bytes, sep: Bytes): Array<Bytes>
    function SplitAfterN(s: Bytes, sep: Bytes, n: Int | NumberLike): Array<Bytes>
    function SplitN(s: Bytes, sep: Bytes, n: Int | NumberLike): Array<Bytes>
    function Title(s: Bytes): Bytes
    function ToLower(s: Bytes): Bytes
    function ToLowerSpecial(c: unicode.SpecialCase, s: Bytes): Bytes
    function ToTitle(s: Bytes): Bytes
    function ToTitleSpecial(c: unicode.SpecialCase, s: Bytes): Bytes
    function ToUpper(s: Bytes): Bytes
    function ToUpperSpecial(c: unicode.SpecialCase, s: Bytes): Bytes
    function ToValidUTF8(s: Bytes, replacement: Bytes): Bytes
    function Trim(s: Bytes, cutset: string): Bytes
    function TrimFunc(s: Bytes, f: (r: Rune) => boolean): Bytes
    function TrimLeft(s: Bytes, cutset: string): Bytes
    function TrimLeftFunc(s: Bytes, f: (r: Rune) => boolean): Bytes
    function TrimPrefix(s: Bytes, prefix: Bytes): Bytes
    function TrimRight(s: Bytes, cutset: string): Bytes
    function TrimRightFunc(s: Bytes, f: (r: Rune) => boolean): Bytes
    function TrimSpace(s: Bytes): Bytes
    function TrimSuffix(s: Bytes, suffix: Bytes): Bytes

    function NewBuffer(buf: Bytes | null | undefined): BufferPointer
    function NewBufferString(s: string): BufferPointer
    interface BufferPointer extends Native {
        readonly __BufferPointer: BufferPointer

        Bytes(): Bytes
        Cap(): Int
        Grow(n: Int | NumberLike): void
        Len(): Int
        Next(n: Int | NumberLike): Bytes
        Read(p: Bytes): Int
        ReadByte(): Byte
        ReadBytes(delim: Byte): Bytes
        ReadFrom(r: io.Reader): Int64
        /** return (r rune, size int, err error) */
        ReadRune(): [Rune, Int]
        ReadString(delim: Byte): string
        Reset(): void
        String(): string
        Truncate(n: Int | NumberLike)
        UnreadByte(): void
        UnreadRune(): void
        Write(p: Bytes): Int
        WriteByte(c: Byte): void
        WriteRune(r: Rune): Int
        WriteString(s: string): Int
        WriteTo(w: io.Writer): Int64
    }

    function NewReader(b: GoBytes): ReaderPointer
    interface ReaderPointer extends Native {
        readonly __ReaderPointer: ReaderPointer

        Len(): Int
        Read(b: Bytes): Int
        ReadAt(b: Bytes, off: Int64 | NumberLike): Int
        ReadByte(): Byte
        /** return (ch rune, size int, err error) */
        ReadRune(): [Rune, Int]
        Reset(b: Bytes)
        Seek(offset: Int64 | NumberLike, whence: Int | NumberLike): Int64
        Size(): Int64
        UnreadByte(): void
        UnreadRune(): void
        WriteTo(w: io.Writer): Int64
    }

    function isBufferPointer(v: any, isptr?: boolean): v is BufferPointer
    function isReaderPointer(v: any, isptr?: boolean): v is ReaderPointer
}