declare module "stdgo/archive/zip" {
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
    import * as fs from "stdgo/io/fs";
    import * as time from "stdgo/time";

    const Store = 0 // no compression
    const Deflate = 8 // DEFLATE compressed

    const ErrFormat: Error
    const ErrAlgorithm: Error
    const ErrChecksum: Error

    function RegisterCompressor(method: Uint16 | NumberLike, comp: Compressor): void
    function RegisterDecompressor(method: Uint16 | NumberLike, dcomp: Decompressor): void

    type Compressor = (w: io.Writer) => [io.WriteCloser, Error]
    type Decompressor = (r: io.Reader) => io.ReadCloser

    interface FilePointer extends Native {
        readonly __FilePointer: FilePointer

        FileHeader: FileHeaderPointer

        DataOffset(): Int64
        Open(): io.ReadCloser
        OpenRaw(): io.Reader

        // Name is the name of the file.
        //
        // It must be a relative path, not start with a drive letter (such as "C:"),
        // and must use forward slashes instead of back slashes. A trailing slash
        // indicates that this file is a directory and should have no data.
        //
        // When reading zip files, the Name field is populated from
        // the zip file directly and is not validated for correctness.
        // It is the caller's responsibility to sanitize it as
        // appropriate, including canonicalizing slash directions,
        // validating that paths are relative, and preventing path
        // traversal through filenames ("../../../").
        Name: string

        // Comment is any arbitrary user-defined string shorter than 64KiB.
        Comment: string

        // NonUTF8 indicates that Name and Comment are not encoded in UTF-8.
        //
        // By specification, the only other encoding permitted should be CP-437,
        // but historically many ZIP readers interpret Name and Comment as whatever
        // the system's local character encoding happens to be.
        //
        // This flag should only be set if the user intends to encode a non-portable
        // ZIP file for a specific localized region. Otherwise, the Writer
        // automatically sets the ZIP format's UTF-8 flag for valid UTF-8 strings.
        NonUTF8: boolean

        CreatorVersion: Uint16
        ReaderVersion: Uint16
        Flags: Uint16

        // Method is the compression method. If zero, Store is used.
        Method: Uint16

        // Modified is the modified time of the file.
        //
        // When reading, an extended timestamp is preferred over the legacy MS-DOS
        // date field, and the offset between the times is used as the timezone.
        // If only the MS-DOS date is present, the timezone is assumed to be UTC.
        //
        // When writing, an extended timestamp (which is timezone-agnostic) is
        // always emitted. The legacy MS-DOS date field is encoded according to the
        // location of the Modified time.
        Modified: time.Time
        ModifiedTime: Uint16 // Deprecated: Legacy MS-DOS date; use Modified instead.
        ModifiedDate: Uint16 // Deprecated: Legacy MS-DOS time; use Modified instead.

        CRC32: Uint32
        CompressedSize: Uint32 // Deprecated: Use CompressedSize64 instead.
        UncompressedSize: Uint32 // Deprecated: Use UncompressedSize64 instead.
        CompressedSize64: Uint64
        UncompressedSize64: Uint64
        Extra: Bytes
        ExternalAttrs: Uint32 // Meaning depends on CreatorVersion

        FileInfo(): fs.FileInfo
        Mode(): fs.FileMode
        SetMode(mode: fs.FileMode)
    }

    function FileInfoHeader(fi: fs.FileInfo): FileHeaderPointer
    interface FileHeaderPointer extends Native {
        readonly __FileHeaderPointer: FileHeaderPointer
        // Name is the name of the file.
        //
        // It must be a relative path, not start with a drive letter (such as "C:"),
        // and must use forward slashes instead of back slashes. A trailing slash
        // indicates that this file is a directory and should have no data.
        //
        // When reading zip files, the Name field is populated from
        // the zip file directly and is not validated for correctness.
        // It is the caller's responsibility to sanitize it as
        // appropriate, including canonicalizing slash directions,
        // validating that paths are relative, and preventing path
        // traversal through filenames ("../../../").
        Name: string

        // Comment is any arbitrary user-defined string shorter than 64KiB.
        Comment: string

        // NonUTF8 indicates that Name and Comment are not encoded in UTF-8.
        //
        // By specification, the only other encoding permitted should be CP-437,
        // but historically many ZIP readers interpret Name and Comment as whatever
        // the system's local character encoding happens to be.
        //
        // This flag should only be set if the user intends to encode a non-portable
        // ZIP file for a specific localized region. Otherwise, the Writer
        // automatically sets the ZIP format's UTF-8 flag for valid UTF-8 strings.
        NonUTF8: boolean

        CreatorVersion: Uint16
        ReaderVersion: Uint16
        Flags: Uint16

        // Method is the compression method. If zero, Store is used.
        Method: Uint16

        // Modified is the modified time of the file.
        //
        // When reading, an extended timestamp is preferred over the legacy MS-DOS
        // date field, and the offset between the times is used as the timezone.
        // If only the MS-DOS date is present, the timezone is assumed to be UTC.
        //
        // When writing, an extended timestamp (which is timezone-agnostic) is
        // always emitted. The legacy MS-DOS date field is encoded according to the
        // location of the Modified time.
        Modified: time.Time
        ModifiedTime: Uint16 // Deprecated: Legacy MS-DOS date; use Modified instead.
        ModifiedDate: Uint16 // Deprecated: Legacy MS-DOS time; use Modified instead.

        CRC32: Uint32
        CompressedSize: Uint32 // Deprecated: Use CompressedSize64 instead.
        UncompressedSize: Uint32 // Deprecated: Use UncompressedSize64 instead.
        CompressedSize64: Uint64
        UncompressedSize64: Uint64
        Extra: Bytes
        ExternalAttrs: Uint32 // Meaning depends on CreatorVersion

        FileInfo(): fs.FileInfo
        Mode(): fs.FileMode
        SetMode(mode: fs.FileMode)
    }

    function OpenReader(name: string): ReadCloserPointer
    class ReadCloserPointer extends Native {
        readonly __ReadCloserPointer: ReadCloserPointer
        Reader: ReaderPointer

        Close()

        File: Array<File>
        Comment: string

        Open(name: string): fs.File
        RegisterDecompressor(method: Uint16 | NumberLike, dcomp: Decompressor): void
    }

    function NewReader(r: io.ReaderAt, size: Int64 | NumberLike): ReaderPointer
    class ReaderPointer extends Native {
        readonly __ReaderPointer: ReaderPointer

        File: Array<File>
        Comment: string

        Open(name: string): fs.File
        RegisterDecompressor(method: Uint16 | NumberLike, dcomp: Decompressor): void
    }

    function NewWriter(w: io.Writer): WriterPointer
    class WriterPointer extends Native {
        readonly __WriterPointer: WriterPointer

        Close(): void
        Copy(f: File): void
        Create(name: string): io.Writer
        CreateHeader(fh: FileHeaderPointer): io.Writer
        CreateRaw(fh: FileHeaderPointer): io.Writer
        Flush(): void
        RegisterCompressor(method: Uint16 | NumberLike, comp: Compressor): void
        SetComment(comment: string): void
        SetOffset(n: Int64 | NumberLike): void
    }
    function isCompressor(v: any): v is Compressor
    function isDecompressor(v: any): v is Decompressor
    function isFilePointer(v: any): v is FilePointer
    function isFileHeaderPointer(v: any): v is FileHeaderPointer
    function isReadCloserPointer(v: any): v is ReadCloserPointer
    function isReaderPointer(v: any): v is ReaderPointer
    function isWriterPointer(v: any): v is WriterPointer
}