declare module "stdgo/os/exec" {
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
    import * as context from "stdgo/context";
    import * as io from "stdgo/io";
    import * as os from "stdgo/os";
    import * as syscall from "stdgo/syscall";

    const ErrNotFound: Error

    function LookPath(file: string): string

    function Command(name: string, ...arg: Array<string>): CmdPointer
    function CommandContext(ctx: context.Context, name: string, ...arg: Array<string>): CmdPointer
    interface CmdPointer extends Native {
        readonly __CmdPointer: CmdPointer
        Path: string
        Args: Array<string>
        Env: Array<string>
        Stdin: io.Reader
        Stdout: io.Writer
        Stderr: io.Writer
        ExtraFiles: Array<os.FilePointer>
        SysProcAttr: syscall.SysProcAttrPointer
        Process: os.ProcessPointer
        ProcessState: os.ProcessStatePointer

        CombinedOutput(): Bytes
        Output(): Bytes
        Run(): void
        Start(): void
        StderrPipe(): io.ReadCloser
        StdinPipe(): io.WriteCloser
        StdoutPipe(): io.ReadCloser
        String(): string
        Wait(): void
    }

    interface ErrorPointer extends Error {
        readonly __ErrorPointer: ErrorPointer
        // Name is the file name for which the error occurred.
        Name: string
        // Err is the underlying error.
        Err: Error

        Unwrap(): Error
    }
    interface ExitErrorPointer extends os.ProcessStatePointer, Error {
        ProcessState: os.ProcessStatePointer
        Stderr: Bytes
    }

    function isCmdPointer(v: any): v is CmdPointer
    function isErrorPointer(v: any): v is ErrorPointer
    function isExitErrorPointer(v: any): v is ExitErrorPointer
}