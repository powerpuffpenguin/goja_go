declare module "stdgo/strings" {
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
    import * as unicode from "stdgo/unicode";
    import * as io from "stdgo/io";

    function At(s: string, index: Int | NumberLike): Rune
    /** if f return false break foreach */
    function Foreach(s: string, f: (r: Rune, index?: Int) => boolean): void
    function String(r: Rune): string

    function Compare(a: string, b: string): Int
    function Contains(s: string, substr: string): boolean
    function ContainsAny(s: string, chars: string): boolean
    function ContainsRune(s: string, r: Rune): boolean
    function Count(s: string, substr: string): Int
    function EqualFold(s: string, t: string): boolean
    function Fields(s: string): Array<string>
    function FieldsFunc(s: string, f: (r: Rune) => boolean): Array<string>
    function HasPrefix(s: string, prefix: string): boolean
    function HasSuffix(s: string, suffix: string): boolean
    function Index(s: string, substr: string): Int
    function IndexAny(s: string, chars: string): Int
    function IndexByte(s: string, c: Byte): Int
    function IndexFunc(s: string, f: (r: Rune) => boolean): Int
    function IndexRune(s: string, r: Rune): Int
    function Join(elems: Array<string>, sep: string): string
    function LastIndex(s: string, substr: string): Int
    function LastIndexAny(s: string, chars: string): Int
    function LastIndexByte(s: string, c: Byte): Int
    function LastIndexFunc(s: string, f: (r: Rune) => boolean): Int
    function Map(mapping: (r: Rune) => Rune, s: string): string
    function Repeat(s: string, count: Int | NumberLike): string
    function Replace(s: string, old: string, newstr: string, n: Int | NumberLike): string
    function ReplaceAll(s: string, old: string, newstr: string): string
    function Split(s: string, sep: string): Array<string>
    function SplitAfter(s: string, sep: string): Array<string>
    function SplitAfterN(s: string, sep: string, n: Int | NumberLike): Array<string>
    function SplitN(s: string, sep: string, n: Int | NumberLike): Array<string>
    function Title(s: string): string
    function ToLower(s: string): string
    function ToLowerSpecial(c: unicode.SpecialCase, s: string): string
    function ToTitle(s: string): string
    function ToTitleSpecial(c: unicode.SpecialCase, s: string): string
    function ToUpper(s: string): string
    function ToUpperSpecial(c: unicode.SpecialCase, s: string): string
    function ToValidUTF8(s: string, replacement: string): string
    function Trim(s: string, cutset: string): string
    function TrimFunc(s: string, f: (r: Rune) => boolean): string
    function TrimLeft(s: string, cutset: string): string
    function TrimLeftFunc(s: string, f: (r: Rune) => boolean): string
    function TrimPrefix(s: string, prefix: string): string
    function TrimRight(s: string, cutset: string): string
    function TrimRightFunc(s: string, f: (r: Rune) => boolean): string
    function TrimSpace(s: string): string
    function TrimSuffix(s: string, suffix: string): string

    function Builder(): BuilderPointer
    interface BuilderPointer extends Native {
        readonly __BuilderPointer: BuilderPointer

        Cap(): Int
        Grow(n: Int | NumberLike): void
        Len(): Int
        Reset(): void
        String(): string
        Write(p: Bytes): Int
        WriteByte(c: Byte): void
        WriteRune(r: Rune): Int
        WriteString(s: string): Int
    }

    function NewReader(s: string): ReaderPointer
    interface ReaderPointer extends Native {
        readonly __ReaderPointer: ReaderPointer

        Len(): Int
        Read(b: Bytes): Int
        ReadAt(b: Bytes, off: Int64 | NumberLike): Int
        ReadByte(): Byte
        /** return (ch rune, size int, err error) */
        ReadRune(): [Rune, Int]
        Reset(s: string): void
        Seek(offset: Int64 | NumberLike, whence: Int | NumberLike): Int64
        Size(): Int64
        UnreadByte(): void
        UnreadRune(): void
        WriteTo(w: io.Writer): Int64
    }

    function NewReplacer(...oldnew: Array<string>): ReplacerPointer
    interface ReplacerPointer extends Native {
        readonly __ReplacerPointer: ReplacerPointer

        Replace(s: string): string
        WriteString(w: io.Writer, s: string): Int
    }

    function isBuilderPointer(v: any): v is BuilderPointer
    function isReaderPointer(v: any): v is ReaderPointer
    function isReplacerPointer(v: any): v is ReplacerPointer
}