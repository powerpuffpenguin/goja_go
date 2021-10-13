declare module "stdgo/regexp" {
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

    function Match(pattern: string, b: Bytes): boolean
    function MatchReader(pattern: string, r: io.RuneReader): boolean
    function MatchString(pattern: string, s: string): boolean
    function QuoteMeta(s: string): string

    function Compile(expr: string): RegexpPointer
    function CompilePOSIX(expr: string): RegexpPointer
    function MustCompile(str: string): RegexpPointer
    function MustCompilePOSIX(str: string): RegexpPointer
    interface RegexpPointer extends Native {
        readonly __RegexpPointer: RegexpPointer

        Expand(dst: Bytes, template: Bytes, src: Bytes, match: IntSlice): Bytes
        ExpandString(dst: Bytes, template: string, src: string, match: IntSlice): Bytes
        Find(b: Bytes): Bytes
        FindAll(b: Bytes, n: Int | NumberLike): Slice<Bytes>
        FindAllIndex(b: Bytes, n: Int | NumberLike): Slice<IntSlice>
        FindAllString(s: string, n: Int | NumberLike): Array<string>
        FindAllStringIndex(s: string, n: Int | NumberLike): Slice<IntSlice>
        FindAllStringSubmatch(s: string, n: Int | NumberLike): Array<Array<string>>
        FindAllStringSubmatchIndex(s: string, n: Int | NumberLike): Slice<IntSlice>
        FindAllSubmatch(b: Bytes, n: Int | NumberLike): Slice<Slice<Bytes>>
        FindAllSubmatchIndex(b: Bytes, n: Int | NumberLike): Slice<IntSlice>
        FindIndex(b: Bytes): IntSlice
        FindReaderIndex(r: io.RuneReader): IntSlice
        FindReaderSubmatchIndex(r: io.RuneReader): IntSlice
        FindString(s: string): string
        FindStringIndex(s: string): IntSlice
        FindStringSubmatch(s: string): Array<string>
        FindStringSubmatchIndex(s: string): IntSlice
        FindSubmatch(b: Bytes): Slice<Bytes>
        FindSubmatchIndex(b: Bytes): IntSlice
        /** return (prefix string, complete bool) */
        LiteralPrefix(): [string, boolean]
        Longest(): void
        Match(b: Bytes): boolean
        MatchReader(r: io.RuneReader): boolean
        MatchString(s: string): boolean
        NumSubexp(): Int
        ReplaceAll(src: Bytes, repl: Bytes): Bytes
        ReplaceAllFunc(src: Bytes, repl: (b: Bytes) => Bytes): Bytes
        ReplaceAllLiteral(src: Bytes, repl: Bytes): Bytes
        ReplaceAllLiteralString(src: string, repl: string): string
        ReplaceAllString(src: string, repl: string): string
        ReplaceAllStringFunc(src: string, repl: (s: string) => string): string
        Split(s: string, n: Int | NumberLike): Array<string>
        String(): string
        SubexpIndex(name: string): Int
        SubexpNames(): Array<string>
    }

    function isRegexpPointer(v: any): v is RegexpPointer
}