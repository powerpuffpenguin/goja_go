declare module "stdgo/net/mail" {
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
    import * as mime from "stdgo/mime";
    import * as time from "stdgo/time";

    const ErrHeaderNotPresent: Error

    function ParseDate(date: string): time.Time

    function ParseAddress(address: string): AddressPointer
    function ParseAddressList(list: string): Array<AddressPointer>
    interface AddressPointer extends Native {
        readonly __AddressPointer: AddressPointer
        Name: string // Proper name; may be empty.
        Address: string // user@domain

        String(): string
    }

    function AddressParser(wordDecoder: mime.WordDecoderPointer): AddressParserPointer
    interface AddressParserPointer extends Native {
        readonly __AddressParserPointer: AddressParserPointer
        // WordDecoder optionally specifies a decoder for RFC 2047 encoded-words.
        WordDecoder: mime.WordDecoderPointer

        Parse(address: string): AddressPointer
        ParseList(list: string): Array<AddressPointer>
    }

    function Header(): Header
    interface Header extends Map<string, Array<string>> {
        readonly __Header: Header

        AddressList(key: string): Array<AddressPointer>
        Date(): time.Time
        Get(key: string): string
    }

    function ReadMessage(r: io.Reader): MessagePointer
    interface MessagePointer extends Native {
        Header: Header
        Body: io.Reader
    }
    function isAddressPointer(v: any): v is AddressPointer
    function isAddressParserPointer(v: any): v is AddressParserPointer
    function isHeader(v: any): v is Header
    function isMessagePointer(v: any): v is MessagePointer
}