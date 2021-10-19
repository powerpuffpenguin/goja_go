declare module "stdgo/text/template" {
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
        Complex128,
    } from "stdgo/builtin";
    import * as parse from "stdgo/text/template/parse";
    import * as fs from "stdgo/io/fs";
    import * as io from "stdgo/io";

    function HTMLEscape(w: io.Writer, b: Bytes): void
    function HTMLEscapeString(s: string): string
    function HTMLEscaper(...args: Array<any>): string
    /** return (truth, ok bool) */
    function IsTrue(val: any): [boolean, boolean]
    function JSEscape(w: io.Writer, b: Bytes): void
    function JSEscapeString(s: string): string
    function JSEscaper(...args: Array<any>): string
    function URLQueryEscaper(...args: Array<any>): string

    interface ExecError extends Error {
        readonly __ExecError: ExecError
        Name: string // Name of template.
        Err: Error  // Pre-formatted error.

        Unwrap(): error
    }
    function FuncMap(): FuncMap
    interface FuncMap extends Map<string, any> {
        readonly __FuncMap: FuncMap
    }

    function Must(t: TemplatePointer, err: Error): TemplatePointer
    function New(name: string): TemplatePointer
    function ParseFS(fsys: fs.FS, ...patterns: Array<string>): TemplatePointer
    function ParseFiles(...filenames: Array<string>): TemplatePointer
    function ParseGlob(pattern: string): TemplatePointer
    interface TemplatePointer extends parse.TreePointer {
        Tree: parse.TreePointer
        // contains filtered or unexported fields

        AddParseTree(name: string, tree: parse.TreePointer): TemplatePointer
        Clone(): TemplatePointer
        DefinedTemplates(): string
        Delims(left: string, right: string): TemplatePointer
        Execute(wr: io.Writer, data: any): void
        ExecuteTemplate(wr: io.Writer, name: string, data: any): void
        Funcs(funcMap: FuncMap): TemplatePointer
        Lookup(name: string): TemplatePointer
        Name(): string
        New(name: string): TemplatePointer
        Option(...opt: Array<string>): TemplatePointer
        Parse(text: string): TemplatePointer
        ParseFS(fsys: fs.FS, ...patterns: Array<string>): TemplatePointer
        ParseFiles(...filenames: Array<string>): TemplatePointer
        ParseGlob(pattern: string): TemplatePointer
        Templates(): Array<TemplatePointer>
    }
    function isExecError(v: any): v is ExecError
    function isFuncMap(v: any): v is FuncMap
    function isTemplatePointer(v: any): v is TemplatePointer
}