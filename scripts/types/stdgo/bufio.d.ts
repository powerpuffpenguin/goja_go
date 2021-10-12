declare module "stdgo/bufio" {
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

    const MaxScanTokenSize = 64 * 1024

    const ErrInvalidUnreadByte: Error
    const ErrInvalidUnreadRune: Error
    const ErrBufferFull: Error
    const ErrNegativeCount: Error

    const ErrTooLong: Error
    const ErrNegativeAdvance: Error
    const ErrAdvanceTooFar: Error
    const ErrBadReadCount: Error

    const ErrFinalToken: Error

    /** return (advance int, token []byte, err error) */
    function ScanBytes(data: Bytes, atEOF: boolean): [Int, Bytes]
    /** return (advance int, token []byte, err error) */
    function ScanLines(data: Bytes, atEOF: boolean): [Int, Bytes]
    /** return (advance int, token []byte, err error) */
    function ScanRunes(data: Bytes, atEOF: boolean): [Int, Bytes]
    /** return (advance int, token []byte, err error) */
    function ScanWords(data: Bytes, atEOF: boolean): [Int, Bytes]

    function NewReadWriter(r: Reader, w: Writer): ReadWriterPointer
    interface ReadWriterPointer extends Native {
        readonly __ReadWriterPointer: ReadWriterPointer

        Reader: ReaderPointer
        Writer: WriterPointer

        Buffered(): Int
        Discard(n: Int | NumberLike): Int
        Peek(n: Int): Bytes
        Read(p: Bytes): Int
        ReadByte(): Byte
        ReadBytes(delim: Byte): Bytes
        /** return (line[]byte, isPrefix bool, err error) */
        ReadLine(): [Bytes, boolean]
        /** return (r rune, size int, err error) */
        ReadRune(): [Rune, Int]
        ReadSlice(delim: Byte): Bytes
        ReadString(delim: Byte): string
        UnreadByte(): void
        UnreadRune(): void
        WriteTo(w: io.Writer): Int64

        Available(): Int
        Buffered(): Int
        Flush(): void
        ReadFrom(r: io.Reader): Int64
        Write(p: Bytes): Int
        WriteByte(c: Byte): void
        WriteRune(r: Rune): Int
        WriteString(s: string): Int
    }

    function NewReader(rd: io.Reader): ReaderPointer
    function NewReaderSize(rd: io.Reader, size: Int | NumberLike): ReaderPointer
    interface ReaderPointer extends Native {
        readonly __ReaderPointer: ReaderPointer

        Buffered(): Int
        Discard(n: Int | NumberLike): Int
        Peek(n: Int): Bytes
        Read(p: Bytes): Int
        ReadByte(): Byte
        ReadBytes(delim: Byte): Bytes
        /** return (line[]byte, isPrefix bool, err error) */
        ReadLine(): [Bytes, boolean]
        /** return (r rune, size int, err error) */
        ReadRune(): [Rune, Int]
        ReadSlice(delim: Byte): Bytes
        ReadString(delim: Byte): string
        Reset(r: io.Reader): void
        Size(): Int
        UnreadByte(): void
        UnreadRune(): void
        WriteTo(w: io.Writer): Int64
    }

    function NewScanner(r: io.Reader): ScannerPointer
    interface ScannerPointer extends Native {
        readonly __ScannerPointer: ScannerPointer

        Buffer(buf: Byte, max: Int | NumberLike): void
        Bytes(): Bytes
        Err(): Error
        Scan(): boolean
        Split(split: SplitFunc): void
        Text(): string
    }
    /** return (advance int, token[]byte, err error) */
    type SplitFunc = (data: Bytes, atEOF: boolean) => [Int, Bytes, Error]

    function NewWriter(w: io.Writer): WriterPointer
    function NewWriterSize(w: io.Writer, size: GoInt): WriterPointer
    class WriterPointer extends Native {
        readonly __WriterPointer: WriterPointer

        Available(): Int
        Buffered(): Int
        Flush(): void
        ReadFrom(r: io.Reader): Int64
        Reset(w: io.Writer): void
        Size(): Int
        Write(p: Bytes): Int
        WriteByte(c: Byte): void
        WriteRune(r: Rune): Int
        WriteString(s: string): Int
    }

    function isReadWriterPointer(v: any): v is ReadWriterPointer
    function isReaderPointer(v: any): v is ReaderPointer
    function isScannerPointer(v: any): v is ScannerPointer
    function isSplitFunc(v: any): v is SplitFunc
    function isWriterPointer(v: any): v is WriterPointer
}