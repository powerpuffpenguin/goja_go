declare module "stdgo/net" {
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
    import {
        Error as Error0
    } from "stdgo/builtin";
    import * as io from "stdgo/io";
    import * as os from "stdgo/os";
    import * as time from "stdgo/time";
    import * as syscall from "stdgo/syscall";
    import * as context from "stdgo/context";

    const IPv4len = 4
    const IPv6len = 16

    const IPv4bcast: IP// IPv4(255, 255, 255, 255) // limited broadcast
    const IPv4allsys: IP// IPv4(224, 0, 0, 1)       // all systems
    const IPv4allrouter: IP// IPv4(224, 0, 0, 2)       // all routers
    const IPv4zero: IP// IPv4(0, 0, 0, 0)         // all zeros

    const IPv6zero: IP // IP{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
    const IPv6unspecified: IP // IP{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
    const IPv6loopback: IP // IP{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
    const IPv6interfacelocalallnodes: IP // IP{0xff, 0x01, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x01}
    const IPv6linklocalallnodes: IP // IP{0xff, 0x02, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x01}
    const IPv6linklocalallrouters: IP // IP{0xff, 0x02, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x02}

    const DefaultResolver: ResolverPointer
    const ErrClosed: Error
    const ErrWriteToConnected: Error

    function JoinHostPort(host: string, port: string): string
    function LookupAddr(addr: string): Array<string>
    function LookupCNAME(host: string): string
    function LookupHost(host: string): Array<string>
    function LookupPort(network: string, service: string): Int
    function LookupTXT(name: string): Array<string>
    function ParseCIDR(s: string): [IP, IPNetPointer]
    function Pipe(): [Conn, Conn]
    /** return (host, port string, err error) */
    function SplitHostPort(hostport: string): [string, string]

    function InterfaceAddrs(): Array<Addr>
    interface Addr extends Native {
        Network(): string // name of the network (for example, "tcp", "udp")
        String(): string  // string form of address (for example, "192.0.2.1:25", "[2001:db8::1]:80")
    }

    interface AddrError extends Error {
        Err: string
        Addr: string

        Temporary(): boolean
        Timeout(): boolean
    }

    interface Buffers extends Native {
        readonly __Buffers: Buffers

        Read(p: Bytes): Int
        WriteTo(w: io.Writer): Int64
    }

    function Dial(network: string, address: string): Conn
    function DialTimeout(network: string, address: string, timeout: time.Duration): Conn
    function FileConn(f: os.FilePointer): Conn
    interface Conn extends Native {
        Read(b: Bytes): Int
        Write(b: Bytes): Int
        Close(): void
        LocalAddr(): Addr
        RemoteAddr(): Addr
        SetDeadline(t: time.Time)
        SetReadDeadline(t: time.Time)
        SetWriteDeadline(t: time.Time)
    }

    interface DNSConfigErrorPointer extends Error {
        readonly __DNSConfigErrorPointer: DNSConfigErrorPointer

        Err: Error

        Temporary(): boolean
        Timeout(): boolean
        Unwrap(): Error
    }
    interface DNSErrorPointer extends Error {
        readonly __DNSErrorPointer: DNSErrorPointer

        Err: string // description of the error
        Name: string // name looked for
        Server: string // server used
        IsTimeout: boolean   // if true, timed out; not all timeouts set this
        IsTemporary: boolean   // if true, error is temporary; not all errors set this
        IsNotFound: boolean   // if true, host could not be found

        Temporary(): boolean
        Timeout(): boolean
    }

    function Dialer(): DialerPointer
    interface DialerPointer extends Native {
        readonly __DialerPointer: DialerPointer

        Timeout: time.Duration
        Deadline: time.Time
        LocalAddr: Addr
        DualStack: boolean
        FallbackDelay: time.Duration
        KeepAlive: time.Duration
        Resolver: ResolverPointer
        Cancel: ReadChannel<null>
        Control: (network: string, address: string, c: syscall.RawConn) => void

        Dial(network: string, address: string): Conn
        DialContext(ctx: context.Context, network: string, address: string): Conn
    }

    interface Error extends Error {
        Timeout(): boolean   // Is the error a timeout?
        Temporary(): boolean // Is the error temporary?
    }

    function Flags(val: Uint | NumberLike): Flags
    interface Flags extends Number {
        readonly __Flags: Flags
        String(): string
    }
    const FlagUp: Flags            // interface is up
    const FlagBroadcast: Flags                      // interface supports broadcast access capability
    const FlagLoopback: Flags                       // interface is a loopback interface
    const FlagPointToPoint: Flags                   // interface belongs to a point-to-point link
    const FlagMulticast: Flags                      // interface supports multicast access capability

    function ParseMAC(s: string): HardwareAddr
    interface HardwareAddr extends Native {
        readonly __HardwareAddr: HardwareAddr
        String(): string
    }

    function IPv4(a: Byte | NumberLike, b: Byte | NumberLike, c: Byte | NumberLike, d: Byte | NumberLike): IP
    function LookupIP(host: string): Array<IP>
    function ParseIP(s: string): Array<IP>
    interface IP extends Native {
        readonly __IP: IP

        DefaultMask(): IPMask
        Equal(x: IP): boolean
        IsGlobalUnicast(): boolean
        IsInterfaceLocalMulticast(): boolean
        IsLinkLocalMulticast(): boolean
        IsLinkLocalUnicast(): boolean
        IsLoopback(): boolean
        IsMulticast(): boolean
        IsUnspecified(): boolean
        MarshalText(): GoBytes
        Mask(mask: IPMask): IP
        String(): string
        To16(): IP
        To4(): IP
    }

    function ResolveIPAddr(network: string, address: string): IPAddrPointer
    interface IPAddrPointer extends Native {
        readonly __IPAddrPointer: IPAddrPointer

        IP: IP
        Zone: string // IPv6 scoped addressing zone

        Network(): string
        String(): string
    }

    function DialIP(network: string, laddr: IPAddrPointer, raddr: IPAddrPointer): IPConnPointer
    function ListenIP(network: string, laddr: IPAddrPointer): IPConnPointer
    interface IPConnPointer extends Native {
        readonly __IPConnPointer: IPConnPointer

        Close(): void
        File(): os.FilePointer
        LocalAddr(): Addr
        Read(b: Bytes): Int
        /** return (int, Addr, error) */
        ReadFrom(b: Bytes): [Int, Addr]
        /** return (int, * IPAddr, error) */
        ReadFromIP(b: Bytes): [Int, IPAddrPointer]
        /** return (n, oobn, flags int, addr * IPAddr, err error)*/
        ReadMsgIP(b: Bytes, oob: Bytes): [Int, Int, Int, IPAddrPointer]
        RemoteAddr(): Addr
        SetDeadline(t: time.Time): void
        SetReadBuffer(bytes: Int): void
        SetReadDeadline(t: time.Time): void
        SetWriteBuffer(bytes: Int): void
        SetWriteDeadline(t: time.Time): void
        SyscallConn(): syscall.RawConn
        Write(b: Bytes): Int
        /** return (n, oobn int, err error) */
        WriteMsgIP(b: Bytes, oob: Bytes, addr: IPAddrPointer): [Int, Int]
        WriteTo(b: Bytes, addr: Addr): Int
        WriteToIP(b: Bytes, addr: IPAddrPointer): Int
    }

    function CIDRMask(ones: Int | NumberLike, bits: Int | NumberLike): IPMask
    function IPv4Mask(a: Byte | NumberLike, b: Byte | NumberLike, c: Byte | NumberLike, d: Byte | NumberLike): IPMask
    interface IPMask extends Native {
        readonly __IPMask: IPMask

        /** return (ones, bits int) */
        Size(): [Int, Int]
        String(): string
    }

    interface IPNetPointer extends Native {
        readonly __IPNetPointer: IPNetPointer

        IP: IP     // network number
        Mask: IPMask // network mask

        Contains(ip: IP): boolean
        Network(): string
        String(): string
    }

    function InterfaceByIndex(index: Int | NumberLike): InterfacePointer
    function InterfaceByName(name: string): InterfacePointer
    function Interfaces(): Array<InterfacePointer>
    interface InterfacePointer extends Native {
        readonly __InterfacePointer: InterfacePointer

        Index: Int | number          // positive integer that starts at one, zero is never used
        MTU: Int | number         // maximum transmission unit
        Name: string       // e.g., "en0", "lo0", "eth0.100"
        HardwareAddr: HardwareAddr // IEEE MAC-48, EUI-48 and EUI-64 form
        Flags: Flags        // e.g., FlagUp, FlagLoopback, FlagMulticast

        Addrs(): Array<Addr>
        MulticastAddrs(): Array<Addr>
    }

    interface InvalidAddrError extends Error0 {
        Temporary(): boolean
        Timeout(): boolean
    }

    interface ListenConfigPointer extends Native {
        readonly __ListenConfigPointer: ListenConfigPointer

        Listen(ctx: context.Context, network: string, address: string): Listener
        ListenPacket(ctx: context.Context, network: string, address: string): PacketConn
    }

    function FileListener(f: os.FilePointer): Listener
    function Listen(network: string, address: string): Listener
    interface Listener extends Native {
        Accept(): Conn
        Close(): void
        Addr(): Addr
    }

    function LookupMX(name: string): Array<MXPointer>
    interface MXPointer extends Native {
        readonly __MXPointer: MXPointer

        Host: string
        Pref: GoUint16 | number
    }

    function LookupNS(name: string): Array<NSPointer>
    interface NSPointer extends Native {
        readonly __NSPointer: NSPointer
        Host: string
    }

    interface OpErrorPointer extends Error0 {
        readonly __OpErrorPointer: OpErrorPointer

        Op: string
        Net: string
        Source: Addr
        Addr: Addr
        Err: Error0

        Temporary(): boolean
        Timeout(): boolean
        Unwrap(): Error0
    }

    function FilePacketConn(f: os.FilePointer): PacketConn
    function ListenPacket(network: string, address: string): PacketConn
    interface PacketConn extends Native {
        /** return (n int, addr Addr, err error) */
        ReadFrom(p: Bytes): [Int, Addr]
        WriteTo(p: Bytes, addr: Addr): Int
        Close(): void
        LocalAddr(): Addr
        SetDeadline(t: time.Time)
        SetReadDeadline(t: time.Time)
        SetWriteDeadline(t: time.Time)
    }

    interface ParseErrorPointer extends Error0 {
        readonly __ParseErrorPointer: ParseErrorPointer
        Type: string
        Text: string

        Temporary(): boolean
        Timeout(): boolean
    }

    function Resolver(): ResolverPointer
    interface ResolverPointer extends Native {
        readonly __ResolverPointer: ResolverPointer

        PreferGo: boolean
        StrictErrors: boolean
        Dial: (ctx: context.Context, network: string, address: string) => Conn

        LookupAddr(ctx: context.Context, addr: string): Array<string>
        LookupCNAME(ctx: context.Context, host: string): string
        LookupHost(ctx: context.Context, host: string): Array<string>
        LookupIP(ctx: context.Context, network: string, host: string): Array<IP>
        LookupIPAddr(ctx: context.Context, host: string): Array<IPAddrPointer>
        LookupMX(ctx: context.Context, name: string): Array<MXPointer>
        LookupNS(ctx: context.Context, name: string): Array<NSPointer>
        LookupPort(ctx: context.Context, network: string, service: string): Int
        /** return (string, []*SRV, error) */
        LookupSRV(ctx: context.Context, service: string, proto: string, name: string): [string, SRVPointer]
        LookupTXT(ctx: context.Context, name: string): Array<string>
    }

    /** return (cname string, addrs[] * SRV, err error) */
    function LookupSRV(service: string, proto: string, name: string): [string, Array<SRVPointer>]
    interface SRVPointer extends Native {
        readonly __SRVPointer: SRVPointer
        Target: string
        Port: Uint16
        Priority: Uint16
        Weight: Uint16
    }

    function ResolveTCPAddr(network: string, address: string): TCPAddrPointer
    interface TCPAddrPointer extends Native {
        readonly __TCPAddrPointer: TCPAddrPointer
        IP: IP
        Port: Int
        Zone: string // IPv6 scoped addressing zone

        Network(): string
        String(): string
    }

    function DialTCP(network: string, laddr: TCPAddrPointer, raddr: TCPAddrPointer): TCPConnPointer
    interface TCPConnPointer extends Native {
        readonly __TCPConnPointer: TCPConnPointer

        Close(): void
        CloseRead(): void
        CloseWrite(): void
        File(): os.FilePointer
        LocalAddr(): Addr
        Read(b: Bytes): Int
        ReadFrom(r: io.Reader): Int64
        RemoteAddr(): Addr
        SetDeadline(t: time.Time): void
        SetKeepAlive(keepalive: boolean): void
        SetKeepAlivePeriod(d: time.Duration): void
        SetLinger(sec: Int | NumberLike): void
        SetNoDelay(noDelay: boolean): void
        SetReadBuffer(bytes: Int | NumberLike): void
        SetReadDeadline(t: time.Time): void
        SetWriteBuffer(bytes: Int | NumberLike): void
        SetWriteDeadline(t: time.Time): void
        SyscallConn(): syscall.RawConn
        Write(b: Bytes): Int
    }

    function ListenTCP(network: string, laddr: TCPAddr): TCPListenerPointer
    interface TCPListenerPointer extends Native {
        readonly __TCPListenerPointer: TCPListenerPointer

        Accept(): Conn
        AcceptTCP(): TCPConnPointer
        Addr(): Addr
        Close(): void
        File(): os.FilePointer
        SetDeadline(t: time.Time): void
        SyscallConn(): syscall.RawConn
    }

    function ResolveUDPAddr(network: string, address: string): UDPAddrPointer
    interface UDPAddrPointer extends Native {
        readonly __UDPAddrPointer: UDPAddrPointer
        IP: IP
        Port: Int
        Zone: string // IPv6 scoped addressing zone

        Network(): string
        String(): string
    }

    function DialUDP(network: string, laddr: UDPAddrPointer, raddr: UDPAddrPointer): UDPConnPointer
    function ListenMulticastUDP(network: string, ifi: InterfacePointer, gaddr: UDPAddrPointer): UDPConnPointer
    function ListenUDP(network: string, laddr: UDPAddrPointer): UDPConnPointer
    interface UDPConnPointer extends Native {
        readonly __UDPConnPointer: UDPConnPointer

        Close(): void
        File(): os.FilePointer
        LocalAddr(): Addr
        Read(b: Bytes): Int
        /** return (int, Addr, error) */
        ReadFrom(b: Bytes): [Int, Addr]
        /** return (n int, addr *UDPAddr, err error) */
        ReadFromUDP(b: Bytes): [Int, UDPAddrPointer]
        /** return (n, oobn, flags int, addr *UDPAddr, err error) */
        ReadMsgUDP(b: Bytes, oob: Bytes): [Int, Int, Int, UDPAddrPointer]
        RemoteAddr(): Addr
        SetDeadline(t: time.Time): void
        SetReadBuffer(bytes: Int | NumberLike): void
        SetReadDeadline(t: time.Time): void
        SetWriteBuffer(bytes: Int | NumberLike): void
        SetWriteDeadline(t: time.Time): void
        SyscallConn(): syscall.RawConn
        Write(b: Bytes): Int
        /** return (n, oobn int, err error) */
        WriteMsgUDP(b: Bytes, oob: Bytes, addr: UDPAddrPointer): [Int, Int]
        WriteTo(b: Bytes, addr: Addr): Int
        WriteToUDP(b: Bytes, addr: UDPAddrPointer): Int
    }

    function ResolveUnixAddr(network: string, address: string): UnixAddrPointer
    interface UnixAddrPointer extends Native {
        readonly __UnixAddrPointer: UnixAddrPointer
        Name: string
        Net: string

        Network(): string
        String(): string
    }

    function DialUnix(network: string, laddr: UnixAddrPointer, raddr: UnixAddrPointer): UnixConnPointer
    function ListenUnixgram(network: string, laddr: UnixAddrPointer): UnixConnPointer
    interface UnixConnPointer extends Native {
        readonly __UnixConnPointer: UnixConnPointer

        Close(): void
        CloseRead(): void
        CloseWrite(): void
        File(): os.FilePointer
        LocalAddr(): Addr
        Read(b: Bytes): Int
        /** return (int, Addr, error) */
        ReadFrom(b: Bytes): [Int, Addr]
        /** return  (int, *UnixAddr, error) */
        ReadFromUnix(b: Bytes): [Int, UnixAddrPointer]
        /** return (n, oobn, flags int, addr *UnixAddr, err error) */
        ReadMsgUnix(b: Bytes, oob: Bytes): [Int, Int, Int, UnixAddrPointer]
        RemoteAddr(): Addr
        SetDeadline(t: time.Time): void
        SetReadBuffer(bytes: Int | NumberLike): void
        SetReadDeadline(t: time.Time): void
        SetWriteBuffer(bytes: Int | NumberLike): void
        SetWriteDeadline(t: time.Time): void
        SyscallConn(): syscall.RawConn
        Write(b: Bytes): Int
        /** return (n, oobn int, err error) */
        WriteMsgUnix(b: Bytes, oob: Bytes, addr: UnixAddrPointer): [Int, Int]
        WriteTo(b: Bytes, addr: Addr): Int
        WriteToUnix(b: Bytes, addr: UnixAddrPointer): Int
    }

    function ListenUnix(network: string, laddr: UnixAddrPointer): UnixListenerPointer
    interface UnixListenerPointer extends Native {
        readonly __UnixListenerPointer: UnixListenerPointer

        Accept(): Conn
        AcceptUnix(): UnixConnPointer
        Addr(): Addr
        Close(): void
        File(): os.FilePointer
        SetDeadline(t: time.Time): void
        SetUnlinkOnClose(unlink: boolean): void
        SyscallConn(): syscall.RawConn
    }

    interface UnknownNetworkError extends Error0 {
        readonly __UnknownNetworkError: UnknownNetworkError

        Temporary(): boolean
        Timeout(): boolean
    }

    function isAddr(v: any): v is Addr
    function isAddrError(v: any): v is AddrError
    function isBuffers(v: any): v is Buffers
    function isConn(v: any): v is Conn
    function isDNSConfigErrorPointer(v: any): v is DNSConfigErrorPointer
    function isDNSErrorPointer(v: any): v is DNSErrorPointer
    function isDialerPointer(v: any): v is DialerPointer
    function isError(v: any): v is Error
    function isFlags(v: any): v is Flags
    function isHardwareAddr(v: any): v is HardwareAddr
    function isIP(v: any): v is IP
    function isIPAddrPointer(v: any): v is IPAddrPointer
    function isIPConnPointer(v: any): v is IPConnPointer
    function isIPMask(v: any): v is IPMask
    function isIPNetPointer(v: any): v is IPNetPointer
    function isInterfacePointer(v: any): v is InterfacePointer
    function isInvalidAddrError(v: any): v is InvalidAddrError
    function isListenConfigPointer(v: any): v is ListenConfigPointer
    function isListener(v: any): v is Listener
    function isMXPointer(v: any): v is MXPointer
    function isNSPointer(v: any): v is NSPointer
    function isOpErrorPointer(v: any): v is OpErrorPointer
    function isPacketConn(v: any): v is PacketConn
    function isParseErrorPointer(v: any): v is ParseErrorPointer
    function isResolverPointer(v: any): v is ResolverPointer
    function isSRVPointer(v: any): v is SRVPointer
    function isTCPAddrPointer(v: any): v is TCPAddrPointer
    function isTCPConnPointer(v: any): v is TCPConnPointer
    function isTCPListenerPointer(v: any): v is TCPListenerPointer
    function isUDPAddrPointer(v: any): v is UDPAddrPointer
    function isUDPConnPointer(v: any): v is UDPConnPointer
    function isUnixAddrPointer(v: any): v is UnixAddrPointer
    function isUnixConnPointer(v: any): v is UnixConnPointer
    function isUnixListenerPointer(v: any): v is UnixListenerPointer
    function isUnknownNetworkError(v: any): v is UnknownNetworkError
}