declare module "stdgo/html/template" {
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

    interface CSS extends Native {
        readonly __CSS: CSS
    }
    interface ErrorPointer extends Error {
        readonly __ErrorPointer: ErrorPointer
        // ErrorCode describes the kind of error.
        ErrorCode: ErrorCode
        // Node is the node that caused the problem, if known.
        // If not nil, it overrides Name and Line.
        Node: parse.Node
        // Name is the name of the template in which the error was encountered.
        Name: string
        // Line is the line number of the error in the template source or 0.
        Line: Int
        // Description is a human-readable description of the problem.
        Description: string
    }
    function ErrorCode(v: Int | NumberLike): ErrorCode
    interface ErrorCode extends Number {
        readonly __ErrorCode: ErrorCode
    }
    const OK = ErrorCode(0)
    const ErrAmbigContext = ErrorCode(1)
    const ErrBadHTML = ErrorCode(2)
    const ErrBranchEnd = ErrorCode(3)
    const ErrEndContext = ErrorCode(4)
    const ErrNoSuchTemplate = ErrorCode(5)
    const ErrOutputContext = ErrorCode(6)
    const ErrPartialCharset = ErrorCode(7)
    const ErrPartialEscape = ErrorCode(8)
    const ErrRangeLoopReentry = ErrorCode(9)
    const ErrSlashAmbig = ErrorCode(10)
    const ErrPredefinedEscaper = ErrorCode(11)

    function FuncMap(): FuncMap
    interface FuncMap extends Map<string, any> {
        readonly __FuncMap: FuncMap
    }
    interface HTML extends Native {
        readonly __HTML: HTML
    }
    interface HTMLAttr extends Native {
        readonly __HTMLAttr: HTMLAttr
    }
    interface JS extends Native {
        readonly __JS: JS
    }
    interface JSStr extends Native {
        readonly __JSStr: JSStr
    }
    interface Srcset extends Native {
        readonly __Srcset: Srcset
    }

    function Must(t: TemplatePointer, err: Error): TemplatePointer
    function New(name: string): TemplatePointer
    function ParseFS(fs: fs.FS, ...patterns: Array<string>): TemplatePointer
    function ParseFiles(...filenames: Array<string>): TemplatePointer
    function ParseGlob(pattern: string): TemplatePointer
    interface TemplatePointer extends Native {
        readonly __TemplatePointer: TemplatePointer
        // The underlying template's parse tree, updated to be HTML-safe.
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
        ParseFS(fs: fs.FS, ...patterns: Array<string>): TemplatePointer
        ParseFiles(...filenames: Array<string>): TemplatePointer
        ParseGlob(pattern: string): TemplatePointer
        Templates(): Array<TemplatePointer>
    }
    interface URL extends Native {
        readonly __URL: URL
    }
    function isCSS(v: any): v is CSS
    function isErrorPointer(v: any): v is ErrorPointer
    function isErrorCode(v: any): v is ErrorCode
    function isFuncMap(v: any): v is FuncMap
    function isHTML(v: any): v is HTML
    function isHTMLAttr(v: any): v is HTMLAttr
    function isJS(v: any): v is JS
    function isJSStr(v: any): v is JSStr
    function isSrcset(v: any): v is Srcset
    function isTemplatePointer(v: any): v is TemplatePointer
    function isURL(v: any): v is URL
}