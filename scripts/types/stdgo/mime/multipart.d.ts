declare module "stdgo/mime/multipart" {
    import {
        Float32, Float64,
        Int64, Int32, Int16, Int8, Int,
        Uint64, Uint32, Uint16, Uint8, Uint,
        Number, NumberLike,
        Byte, Bytes, Rune, Runes,
        Float32Slice, Float64Slice,
        Int64Slice, Int32Slice, Int16Slice, Int8Slice, IntSlice,
        Uint64Slice, Uint32Slice, Uint16Slice, Uint8Slice, UintSlice,
        ReadChannel, WriteChannel, Channel,
        Slice, Map,
        Uintptr, Native,
    } from "stdgo/builtin";
    import * as io from "stdgo/io";
    import * as textproto from "stdgo/net/textproto";

    const ErrMessageTooLarge: Error

    interface File extends io.Reader, io.ReaderAt, io.Seeker, io.Closer {
    }
    interface FileHeaderPointer extends Native {
        readonly __FileHeaderPointer: FileHeaderPointer
        Filename: string
        Header: textproto.MIMEHeader
        Size: Int64
        // contains filtered or unexported fields
        Open(): File
    }
    interface FormPointer extends Native {
        readonly __FormPointer: FormPointer
        Value: Map<string, Array<string>>
        File: Map<string, Array<FileHeaderPointer>>
        RemoveAll(): void
    }
    interface PartPointer extends Native {
        readonly __PartPointer: PartPointer
        Header: textproto.MIMEHeader
        Close(): void
        FileName(): string
        FormName(): string
        Read(d: Bytes): Int
    }

    function NewReader(r: io.Reader, boundary: string): ReaderPointer
    interface ReaderPointer extends Native {
        readonly __ReaderPointer: ReaderPointer
        // contains filtered or unexported fields
        NextPart(): PartPointer
        NextRawPart(): PartPointer
        ReadForm(maxMemory: Int64 | NumberLike): FormPointer
    }
    function NewWriter(w: io.Writer): WriterPointer
    interface WriterPointer extends Native {
        readonly __WriterPointer: WriterPointer
        // contains filtered or unexported fields

        Boundary(): string
        Close(): void
        CreateFormField(fieldname: string): io.Writer
        CreateFormFile(fieldname: string, filename: string): io.Writer
        CreatePart(header: textproto.MIMEHeader): io.Writer
        FormDataContentType(): string
        SetBoundary(boundary: string): void
        WriteField(fieldname: string, value: string): void
    }
    function isFile(v: any): v is File
    function isFileHeaderPointer(v: any): v is FileHeaderPointer
    function isFormPointer(v: any): v is FormPointer
    function isPartPointer(v: any): v is PartPointer
    function isReaderPointer(v: any): v is ReaderPointer
    function isWriterPointer(v: any): v is WriterPointer
}