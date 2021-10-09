declare module "stdgo/io" {
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
    } from "stdgo/builtin";

    const SeekStart = 0 // seek relative to the origin of the file
    const SeekCurrent = 1 // seek relative to the current offset
    const SeekEnd = 2 // seek relative to the end

    const EOF: Error
    const ErrClosedPipe: Error
    const ErrNoProgress: Error
    const ErrShortBuffer: Error
    const ErrShortWrite: Error
    const ErrUnexpectedEOF: Error

    function Copy(dst: Writer, src: Reader): Int64
    function CopyBuffer(dst: Writer, src: Reader, buf: Bytes): Int64
    function CopyN(dst: Writer, src: Reader, n: Int64 | NumberLike): Int64
    function Pipe(): [PipeReader, PipeWriter]
    function ReadAll(r: Reader): Bytes
    function ReadAtLeast(r: Reader, buf: Bytes, min: Int | NumberLike): Int
    function ReadFull(r: Reader, buf: Bytes): Int
    function WriteString(w: Writer, s: string): Int

    interface ByteReader extends Native {
        ReadByte(): Byte
    }
    interface ByteScanner extends ByteReader {
        UnreadByte(): void
    }
    interface ByteWriter extends Native {
        WriteByte(c: Byte | NumberLike): void
    }
    interface Closer extends Native {
        Close(): void
    }
    interface LimitedReaderPointer extends Native {
        readonly __LimitedReaderPointer: LimitedReaderPointer

        R: Reader // underlying reader
        N: Int64   // max bytes remaining

        Read(p: Bytes): Int
    }
    interface PipeReaderPointer extends Native {
        readonly __PipeReaderPointer: PipeReaderPointer

        Close(): void
        CloseWithError(err: Error): Error
        Read(data: Bytes): Int
    }
    interface PipeWriterPointer extends Native {
        readonly __PipeWriterPointer: PipeWriterPointer

        Close(): void
        CloseWithError(err: Error): Error
        Write(data: Bytes): Int
    }
    interface ReadCloser extends Reader, Closer { }

    function NopCloser(r: Reader): ReadCloser

    interface ReadSeekCloser extends Reader, Seeker, Closer { }
    interface ReadSeeker extends Reader, Seeker { }
    interface ReadWriteCloser extends Reader, Writer, Closer { }
    interface ReadWriteSeeker extends Reader, Writer, Seeker { }
    interface ReadWriter extends Reader, Writer { }
    interface Reader extends Native {
        Read(data: Bytes): Int
    }

    function LimitReader(r: Reader, n: Int64 | NumberLike): Reader
    function MultiReader(...readers: Array<Reader>): Reader
    function TeeReader(r: Reader, w: Writer): Reader


    interface ReaderAt extends Native {
        ReadAt(p: Bytes, off: Int64 | NumberLike): Int
    }
    interface ReaderFrom extends Native {
        ReadFrom(r: Reader): Int64
    }
    interface RuneReader extends Native {
        ReadRune(): [Rune, Int]
    }
    interface RuneScanner extends RuneReader {
        UnreadRune(): void
    }

    function NewSectionReader(r: ReaderAt, off: Int64 | NumberLike, n: Int64 | NumberLikes): SectionReaderPointer
    interface SectionReaderPointer extends Native {
        readonly __SectionReaderPointer: SectionReaderPointer

        Read(data: Bytes): Int
        ReadAt(p: Bytes, off: Int64 | NumberLike): Int
        Seek(offset: Int64 | NumberLike, whence: Int | NumberLike): Int64
        Size(): Int64
    }

    interface Seeker extends Native {
        Seek(offset: Int64 | NumberLike, whence: Int | NumberLike): Int64
    }
    interface StringWriter extends Native {
        WriteString(s: string): Int
    }
    interface WriteCloser extends Writer, Closer { }
    interface WriteSeeker extends Writer, Seeker { }
    interface Writer extends Native {
        Write(data: Bytes): Int
    }
    const Discard: Writer
    function MultiWriter(...writers: Array<Writer>): Writer

    interface WriterAt extends Native {
        WriteAt(p: Bytes, off: Int64 | NumberLike): Int
    }
    interface WriterTo extends Native {
        WriteTo(w: Writer): Int64
    }

    function isByteReader(v: any): v is ByteReader
    function isByteScanner(v: any): v is ByteScanner
    function isByteWriter(v: any): v is ByteWriter
    function isCloser(v: any): v is Closer
    function isLimitedReaderPointer(v: any): v is LimitedReaderPointer
    function isPipeReaderPointer(v: any): v is PipeReaderPointer
    function isPipeWriterPointer(v: any): v is PipeWriterPointer
    function isReadCloser(v: any): v is ReadCloser
    function isReadSeekCloser(v: any): v is ReadSeekCloser
    function isReadSeeker(v: any): v is ReadSeeker
    function isReadWriteCloser(v: any): v is ReadWriteCloser
    function isReadWriteSeeker(v: any): v is ReadWriteSeeker
    function isReadWriter(v: any): v is ReadWriter
    function isReader(v: any): v is Reader
    function isReaderAt(v: any): v is ReaderAt
    function isReaderFrom(v: any): v is ReaderFrom
    function isRuneReader(v: any): v is RuneReader
    function isRuneScanner(v: any): v is RuneScanner
    function isSectionReaderPointer(v: any): v is SectionReaderPointer
    function isSeeker(v: any): v is Seeker
    function isStringWriter(v: any): v is StringWriter
    function isWriteCloser(v: any): v is WriteCloser
    function isWriteSeeker(v: any): v is WriteSeeker
    function isWriter(v: any): v is Writer
    function isWriterAt(v: any): v is WriterAt
    function isWriterTo(v: any): v is WriterTo
}