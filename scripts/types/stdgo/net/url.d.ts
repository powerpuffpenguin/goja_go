declare module "stdgo/net/url" {
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

    function PathEscape(s: string): string
    function PathUnescape(s: string): string
    function QueryEscape(s: string): string
    function QueryUnescape(s: string): string

    interface ErrorPointer extends Error {
        readonly __ErrorPointer: ErrorPointer

        Op: string
        URL: string
        Err: Error

        Temporary(): boolean
        Timeout(): boolean
        Unwrap(): Error
    }

    interface EscapeError extends Error {
        readonly __EscapeError: EscapeError
    }
    interface InvalidHostError extends Error {
        readonly __InvalidHostError: InvalidHostError
    }

    function Parse(rawURL: string): URLPointer
    function ParseRequestURI(rawURL: string): URLPointer
    interface URLPointer extends Native {
        readonly __URLPointer: URLPointer
        String()

        Scheme: string
        Opaque: string
        User: UserinfoPointer | null
        Host: string
        Path: string
        RawPath: string
        ForceQuery: boolean
        RawQuery: string
        Fragment: string
        RawFragment: string

        EscapedFragment(): string
        EscapedPath(): string
        Hostname(): boolean
        IsAbs(): boolean
        MarshalBinary(): Bytes
        Parse(ref: string): URLPointer
        Port(): string
        Query(): Values
        Redacted(): string
        RequestURI(): string
        ResolveReference(url: URLPointer): URLPointer
        UnmarshalBinary(text: Bytes): URLPointer
    }

    function User(username: string): UserinfoPointer
    function UserPassword(username: string, password: string): UserinfoPointer
    interface UserinfoPointer extends Native {
        readonly __UserinfoPointer: UserinfoPointer
        String(): string
        Username(): string
        Password(): [string, boolean]
    }

    function ParseQuery(query: string): Values
    function Values(): Values
    interface Values extends Map<string, Slice<string>> {
        readonly __Values: Values

        Add(key: string, value: string): void
        Del(key: string): void
        Encode(): string
        Get(key: string): string
        Set(key: string, value: string): void
    }

    function isErrorPointer(v: any): v is ErrorPointer
    function isEscapeError(v: any): v is EscapeError
    function isInvalidHostError(v: any): v is InvalidHostError
    function isURLPointer(v: any): v is URLPointer
    function isUserinfoPointer(v: any): v is UserinfoPointer
    function isValues(v: any): v is Values
}