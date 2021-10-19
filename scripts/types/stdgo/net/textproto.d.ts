declare module "stdgo/net/textproto" {
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
    import * as bufio from "stdgo/bufio";

    function CanonicalMIMEHeaderKey(s: string): string
    function TrimBytes(b: Bytes): Bytes
    function TrimString(s: string): string

    function Dial(network: string, addr: string): ConnPointer
    function NewConn(conn: io.ReadWriteCloser): ConnPointer
    interface ConnPointer extends ReaderPointer, WriterPointer, PipelinePointer {
        readonly __ConnPointer: ConnPointer
        // contains filtered or unexported fields
        Close(): void
        Cmd(format: string, ...args: Array<any>): Uint
    }
    interface ErrorPointer extends Error {
        readonly __ErrorPointer: ErrorPointer
        Code: Int
        Msg: string
    }
    interface MIMEHeader extends Map<string, Array<string>> {
        readonly __MIMEHeader: MIMEHeader
        Add(key: string, value: string): void
        Del(key: string): void
        Get(key: string): string
        Set(key: string, value: string): void
        Values(key: string): Array<string>
    }
    interface PipelinePointer extends Native {
        readonly __PipelinePointer: PipelinePointer
        // contains filtered or unexported fields
        EndRequest(id: Uint | NumberLike): void
        EndResponse(id: Uint | NumberLike): void
        Next(): Uint
        StartRequest(id: Uint | NumberLike): void
        StartResponse(id: Uint | NumberLike): void
    }
    interface ProtocolError extends Error {
        readonly __ProtocolError: ProtocolError
    }

    function NewReader(r: bufio.ReaderPointer): ReaderPointer
    interface ReaderPointer extends Native {
        R: bufio.ReaderPointer
        // contains filtered or unexported fields

        DotReader(): io.Reader
        /** return (code int, message string, err error) */
        ReadCodeLine(expectCode: Int | NumberLike): [Int, string]
        ReadContinuedLine(): string
        ReadContinuedLineBytes(): Bytes
        ReadDotBytes(): Bytes
        ReadDotLines(): Array<string>
        ReadLine(): string
        ReadLineBytes(): Bytes
        ReadMIMEHeader(): MIMEHeader
        /** (code int, message string, err error) */
        ReadResponse(expectCode: Int | NumberLike): [Int, string]
    }
    function NewWriter(w: bufio.WriterPointer): WriterPointer
    interface WriterPointer extends Native {
        W: bufio.WriterPointer
        // contains filtered or unexported fields
        DotWriter(): io.WriteCloser
        PrintfLine(format: string, ...args: Array<any>): void
    }
    function isConnPointer(v: any): v is ConnPointer
    function isErrorPointer(v: any): v is ErrorPointer
    function isMIMEHeader(v: any): v is MIMEHeader
    function isPipelinePointer(v: any): v is PipelinePointer
    function isProtocolError(v: any): v is ProtocolError
    function isReaderPointer(v: any): v is ReaderPointer
    function isWriterPointer(v: any): v is WriterPointer
}