declare module "stdgo/encoding/csv" {
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
    
    const ErrTrailingComma:Error
    const ErrBareQuote:Error
    const ErrQuote:Error
    const ErrFieldCount:Error

    interface ParseErrorPointer extends Error {
        readonly __ParseErrorPointer: ParseErrorPointer

        StartLine: Int   // Line where the record starts
        Line: Int   // Line where the error occurred
        Column: Int   // Column (1-based byte index) where the error occurred
        Err: Error // The actual error

        Unwrap(): Error
    }

    function NewReader(r: io.Reader): ReaderPointer
    interface ReaderPointer extends Native {
        readonly __ReaderPointer: ReaderPointer
        // Comma is the field delimiter.
        // It is set to comma (',') by NewReader.
        // Comma must be a valid rune and must not be \r, \n,
        // or the Unicode replacement character (0xFFFD).
        Comma: Rune

        // Comment, if not 0, is the comment character. Lines beginning with the
        // Comment character without preceding whitespace are ignored.
        // With leading whitespace the Comment character becomes part of the
        // field, even if TrimLeadingSpace is true.
        // Comment must be a valid rune and must not be \r, \n,
        // or the Unicode replacement character (0xFFFD).
        // It must also not be equal to Comma.
        Comment: Rune

        // FieldsPerRecord is the number of expected fields per record.
        // If FieldsPerRecord is positive, Read requires each record to
        // have the given number of fields. If FieldsPerRecord is 0, Read sets it to
        // the number of fields in the first record, so that future records must
        // have the same field count. If FieldsPerRecord is negative, no check is
        // made and records may have a variable number of fields.
        FieldsPerRecord: Int

        // If LazyQuotes is true, a quote may appear in an unquoted field and a
        // non-doubled quote may appear in a quoted field.
        LazyQuotes: boolean

        // If TrimLeadingSpace is true, leading white space in a field is ignored.
        // This is done even if the field delimiter, Comma, is white space.
        TrimLeadingSpace: boolean

        // ReuseRecord controls whether calls to Read may return a slice sharing
        // the backing array of the previous call's returned slice for performance.
        // By default, each call to Read returns newly allocated memory owned by the caller.
        ReuseRecord: boolean

        TrailingComma: boolean // Deprecated: No longer used.
        // contains filtered or unexported fields

        Read(): Array<string>
        ReadAll(): Array<Array<string>>
    }

    function NewWriter(w: io.Writer): WriterPointer
    interface WriterPointer extends Native {
        readonly __WriterPointer: WriterPointer

        Comma: Rune // Field delimiter (set to ',' by NewWriter)
        UseCRLF: boolean // True to use \r\n as the line terminator

        Error(): Error
        Flush(): void
        Write(record: Array<string>): void
        WriteAll(records: Array<Array<string>>): void
    }

    function isParseErrorPointer(v: any): v is ParseErrorPointer
    function isReaderPointer(v: any): v is ReaderPointer
    function isWriterPointer(v: any): v is WriterPointer
}