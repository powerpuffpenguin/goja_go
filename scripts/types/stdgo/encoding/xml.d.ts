declare module "stdgo/encoding/xml" {
    import {
        Float32, Float64,
        Int64, Int32, Int16, Int8, Int,
        Uint64, Uint32, Uint16, Uint8, Uint,
        NumberLike,
        Byte, Bytes, Rune, Runes,
        Float32Slice, Float64Slice,
        Int64Slice, Int32Slice, Int16Slice, Int8Slice, IntSlice,
        Uint64Slice, Uint32Slice, Uint16Slice, Uint8Slice, UintSlice,
        Error,
        ReadChannel, WriteChannel, Channel,
        Slice, Map,
        Uintptr, Native,
        Type, Value,
    } from "stdgo/builtin";
    import * as io from "stdgo/io";

    const Header: string
    const HTMLAutoClose: Array<string>
    const HTMLEntity: Map<string, string>

    function Escape(w: io.Writer, s: Bytes): void
    function EscapeText(w: io.Writer, s: Bytes): void
    function Marshal(v: any): Bytes
    function MarshalIndent(v: any, prefix: string, indent: string): Bytes
    function Unmarshal(data: Bytes, v: any): void

    interface Attr extends Native {
        readonly __Attr: Attr
        Name: Name
        Value: string
    }
    interface CharData extends Bytes {
        readonly __CharData: CharData
        Copy(): CharData
    }
    interface Comment extends Bytes {
        readonly __Comment: Comment
        Copy(): Comment
    }

    function NewDecoder(r: io.Reader): DecoderPointer
    function NewTokenDecoder(t: TokenReader): DecoderPointer
    interface DecoderPointer extends Native {
        readonly __DecoderPointer: DecoderPointer
        Strict: boolean
        AutoClose: Array<string>
        Entity: Map<string, string>
        CharsetReader: (charset: string, input: io.Reader) => [io.Reader, Error]
        DefaultSpace: string

        Decode(v: any): void
        DecodeElement(v: any, start: StartElementPointer): void
        InputOffset(): Int64
        RawToken(): Token
        Skip(): void
        Token(): Token
    }
    interface Directive extends Bytes {
        readonly __Directive: Directive
        Copy(): Directive
    }

    function NewEncoder(w: io.Writer): EncoderPointer
    interface EncoderPointer extends Native {
        readonly __EncoderPointer: EncoderPointer

        Encode(v: any): void
        EncodeElement(v: any, start: StartElement): void
        EncodeToken(t: Token): void
        Flush(): void
        Indent(prefix: string, indent: string): void
    }
    interface EndElement extends Native {
        readonly __EndElement: EndElement
        Name: Name
    }
    interface Marshaler extends Native {
        MarshalXML(e: EncoderPointer, start: StartElement): void
    }
    interface MarshalerAttr extends Native {
        MarshalXMLAttr(name: Name): Attr
    }
    interface Name extends Native {
        readonly __Name: Name
        Space: string
        Local: string
    }
    interface ProcInst extends Native {
        readonly __ProcInst: ProcInst
        Target: string
        Inst: Bytes
        Copy(): ProcInst
    }
    interface StartElement extends Native {
        readonly __StartElement: StartElement
        Name: Name
        Attr: Slice<Attr>

        Copy(): StartElement
        End(): EndElement
    }
    interface StartElementPointer extends Native {
        readonly __StartElementPointer: StartElementPointer
    }
    interface SyntaxErrorPointer extends Error {
        Msg: string
        Line: Int
    }
    interface TagPathErrorPointer extends Error {
        Struct: Type
        Field1: string
        Tag1: string
        Field2: string
        Tag2: string
    }

    function CopyToken(t: Token): Token
    interface Token extends Native {
    }

    interface TokenReader extends Native {
        Token(): Token
    }
    interface UnmarshalError extends Error {
    }
    interface Unmarshaler extends Native {
        UnmarshalXML(d: DecoderPointer, start: StartElement): void
    }
    interface UnmarshalerAttr extends Native {
        UnmarshalXMLAttr(attr: Attr): void
    }
    interface UnsupportedTypeErrorPointer extends Error {
        Type: Type
    }
    function isAttr(v: any): v is Attr
    function isCharData(v: any): v is CharData
    function isComment(v: any): v is Comment
    function isDecoderPointer(v: any): v is DecoderPointer
    function isDirective(v: any): v is Directive
    function isEncoderPointer(v: any): v is EncoderPointer
    function isEndElement(v: any): v is EndElement
    function isMarshaler(v: any): v is Marshaler
    function isMarshalerAttr(v: any): v is MarshalerAttr
    function isName(v: any): v is Name
    function isProcInst(v: any): v is ProcInst
    function isStartElement(v: any): v is StartElement
    function isStartElementPointer(v: any): v is StartElementPointer
    function isSyntaxErrorPointer(v: any): v is SyntaxErrorPointer
    function isTagPathErrorPointer(v: any): v is TagPathErrorPointer
    function isToken(v: any): v is Token
    function isTokenReader(v: any): v is TokenReader
    function isUnmarshalError(v: any): v is UnmarshalError
    function isUnmarshaler(v: any): v is Unmarshaler
    function isUnmarshalerAttr(v: any): v is UnmarshalerAttr
    function isUnsupportedTypeErrorPointer(v: any): v is UnsupportedTypeErrorPointer
}