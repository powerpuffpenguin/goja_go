declare module "stdgo/net/smtp" {
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
    import * as textproto from "stdgo/net/textproto";
    import * as net from "stdgo/net";
    import * as io from "stdgo/io";
    import * as tls from "stdgo/crypto/tls";

    function SendMail(addr: string, a: Auth, from: string, to: Array<string>, msg: Bytes): void

    function CRAMMD5Auth(username: string, secret: string): Auth
    function PlainAuth(identity: string, username: string, password: string, host: string): Auth
    interface Auth extends Native {
        /** return (proto string, toServer []byte, err error) */
        Start(server: ServerInfoPointer): [string, Bytes]
        Next(fromServer: Bytes, more: boolean): Bytes
    }

    function Dial(addr: string): ClientPointer
    function NewClient(conn: net.Conn, host: string): ClientPointer
    interface ClientPointer extends Native {
        readonly __ClientPointer: ClientPointer
        // Text is the textproto.Conn used by the Client. It is exported to allow for
        // clients to add extensions.
        Text: textproto.ConnPointer
        // contains filtered or unexported fields
        Auth(a: Auth): void
        Close(): void
        Data(): io.WriteCloser
        /** return (bool, string) */
        Extension(ext: string): [boolean, string]
        Hello(localName: string): void
        Mail(from: string): void
        Noop(): void
        Quit(): void
        Rcpt(to: string): void
        Reset(): void
        StartTLS(config: tls.ConfigPointer): void
        /** return (state tls.ConnectionState, ok bool) */
        TLSConnectionState(): [tls.ConnectionState, bool]
        Verify(addr: string): void
    }
    interface ServerInfoPointer extends Native {
        readonly __ServerInfoPointer: ServerInfoPointer
        Name: string   // SMTP server name
        TLS: boolean     // using TLS, with valid certificate for Name
        Auth: Array<string> // advertised authentication mechanisms
    }
    function isAuth(v: any): v is Auth
    function isClientPointer(v: any): v is ClientPointer
    function isServerInfoPointer(v: any): v is ServerInfoPointer
}