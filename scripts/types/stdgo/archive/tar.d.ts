declare module "stdgo/archive/tar" {
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
    import *as time from "stdgo/time";
    import *as fs from "stdgo/io/fs";
    import *as io from "stdgo/io";

    const TypeReg: Byte
    const TypeRegA: Byte
    const TypeLink: Byte
    const TypeSymlink: Byte
    const TypeChar: Byte
    const TypeBlock: Byte
    const TypeDir: Byte
    const TypeFifo: Byte
    const TypeCont: Byte
    const TypeXHeader: Byte
    const TypeXGlobalHeader: Byte
    const TypeGNUSparse: Byte
    const TypeGNULongName: Byte
    const TypeGNULongLink: Byte

    const ErrHeader: Error
    const ErrWriteTooLong: Error
    const ErrFieldTooLong: Error
    const ErrWriteAfterClose: Error

    interface Format extends Number {
        readonly __Format: Format
        String(): string
    }

    function FileInfoHeader(fi: fs.FileInfo, link: string): HeaderPointer
    interface HeaderPointer extends Native {
        readonly __Header: Header
        // Typeflag is the type of header entry.
        // The zero value is automatically promoted to either TypeReg or TypeDir
        // depending on the presence of a trailing slash in Name.
        Typeflag: Byte

        Name: string // Name of file entry
        Linkname: string // Target name of link (valid for TypeLink or TypeSymlink)

        Size: Int64  // Logical file size in bytes
        Mode: Int64  // Permission and mode bits
        Uid: Int    // User ID of owner
        Gid: Int    // Group ID of owner
        Uname: string // User name of owner
        Gname: string // Group name of owner

        // If the Format is unspecified, then Writer.WriteHeader rounds ModTime
        // to the nearest second and ignores the AccessTime and ChangeTime fields.
        //
        // To use AccessTime or ChangeTime, specify the Format as PAX or GNU.
        // To use sub-second resolution, specify the Format as PAX.
        ModTime: time.Time // Modification time
        AccessTime: time.Time // Access time (requires either PAX or GNU support)
        ChangeTime: time.Time // Change time (requires either PAX or GNU support)

        Devmajor: Int64 // Major device number (valid for TypeChar or TypeBlock)
        Devminor: Int64 // Minor device number (valid for TypeChar or TypeBlock)

        // Xattrs stores extended attributes as PAX records under the
        // "SCHILY.xattr." namespace.
        //
        // The following are semantically equivalent:
        //  h.Xattrs[key] = value
        //  h.PAXRecords["SCHILY.xattr."+key] = value
        //
        // When Writer.WriteHeader is called, the contents of Xattrs will take
        // precedence over those in PAXRecords.
        //
        // Deprecated: Use PAXRecords instead.
        Xattrs: Map<string, string>

        // PAXRecords is a map of PAX extended header records.
        //
        // User-defined records should have keys of the following form:
        //	VENDOR.keyword
        // Where VENDOR is some namespace in all uppercase, and keyword may
        // not contain the '=' character (e.g., "GOLANG.pkg.version").
        // The key and value should be non-empty UTF-8 strings.
        //
        // When Writer.WriteHeader is called, PAX records derived from the
        // other fields in Header take precedence over PAXRecords.
        PAXRecords: Map<string, string>

        // Format specifies the format of the tar header.
        //
        // This is set by Reader.Next as a best-effort guess at the format.
        // Since the Reader liberally reads some non-compliant files,
        // it is possible for this to be FormatUnknown.
        //
        // If the format is unspecified when Writer.WriteHeader is called,
        // then it uses the first format (in the order of USTAR, PAX, GNU)
        // capable of encoding this Header (see Format).
        Format: Format

        FileInfo(): fs.FileInfo
    }

    function NewReader(r: io.Reader): ReaderPointer
    interface ReaderPointer extends Native {
        readonly __ReaderPointer: ReaderPointer

        Next(): HeaderPointer
        Read(b: Bytes): Int
    }

    function NewWriter(w: io.Writer): WriterPointer
    class WriterPointer extends Native {
        readonly __WriterPointer: WriterPointer

        Close(): void
        Flush(): void
        Write(b: Bytes): Int
        WriteHeader(hdr: HeaderPointer)
    }

    function isFormat(v: any): v is Format
    function isHeaderPointer(v: any): v is HeaderPointer
    function isReaderPointer(v: any): v is ReaderPointer
    function isWriterPointer(v: any): v is WriterPointer
}