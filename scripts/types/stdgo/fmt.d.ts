declare module "stdgo/fmt" {
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

    function Errorf(format: string, ...a: Array<any>)
    function Fprint(w: io.Writer, ...a: Array<any>): Int
    function Fprintf(w: io.Writer, format: string, ...a: Array<any>): Int
    function Fprintln(w: io.Writer, ...a: Array<any>): Int
    function Fscan(r: io.Reader, ...a: Array<any>): Int
    function Fscanf(r: io.Reader, format: string, ...a: Array<any>): Int
    function Fscanln(r: io.Reader, ...a: Array<any>): Int
    function Print(...a: Array<any>): Int
    function Printf(format: string, ...a: Array<any>): Int
    function Println(...a: Array<any>): Int
    function Scan(...a: Array<any>): Int
    function Scanf(format: string, ...a: Array<any>): Int
    function Scanln(...a: Array<any>): Int
    function Sprint(...a: Array<any>): string
    function Sprintf(format: string, ...a: Array<any>): string
    function Sprintln(...a: Array<any>): string
    function Sscan(str: string, ...a: Array<any>): Int
    function Sscanf(str: string, format: string, ...a: Array<any>): Int
    function Sscanln(str: string, ...a: Array<any>): Int

    interface Formatter extends Native {
        Format(f: State, verb: Rune): void
    }
    interface GoStringer extends Native {
        GoString(): string
    }
    interface ScanState extends Native {
        /** return  (r rune, size int, err error) */
        ReadRune(): [Rune, Int]
        UnreadRune()
        SkipSpace()
        /** return (token []byte, err error) */
        Token(skipSpace: boolean, f: (rune: Rune) => boolean): Bytes
        /** return (wid int, ok bool) */
        Width(): [Int, boolean]
        Read(buf: Bytes): Int
    }
    interface Scanner extends Native {
        Scan(state: ScanState, verb: Rune): void
    }
    interface State extends Native {
        Write(b: Bytes): Int
        /** return (wid int, ok bool) */
        Width(): [Int, boolean]
        /** return (prec int, ok bool) */
        Precision(): [Int, boolean]
        Flag(c: Int): boolean
    }
    interface Stringer extends Native {
        String(): string
    }

    function isFormatter(v: any): v is Formatter
    function isGoStringer(v: any): v is GoStringer
    function isScanState(v: any): v is ScanState
    function isScanner(v: any): v is Scanner
    function isState(v: any): v is State
    function isStringer(v: any): v is Stringer
}