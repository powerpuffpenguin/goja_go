declare module "stdgo/crypto/tls" {
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
    import * as crypto from "stdgo/crypto";
    import * as x509 from "stdgo/crypto/x509";
    import * as net from "stdgo/net";
    import * as io from "stdgo/io";
    import * as time from "stdgo/time";

    // TLS 1.0 - 1.2 cipher suites.
    const TLS_RSA_WITH_RC4_128_SHA = Uint16(0x0005)
    const TLS_RSA_WITH_3DES_EDE_CBC_SHA = Uint16(0x000a)
    const TLS_RSA_WITH_AES_128_CBC_SHA = Uint16(0x002f)
    const TLS_RSA_WITH_AES_256_CBC_SHA = Uint16(0x0035)
    const TLS_RSA_WITH_AES_128_CBC_SHA256 = Uint16(0x003c)
    const TLS_RSA_WITH_AES_128_GCM_SHA256 = Uint16(0x009c)
    const TLS_RSA_WITH_AES_256_GCM_SHA384 = Uint16(0x009d)
    const TLS_ECDHE_ECDSA_WITH_RC4_128_SHA = Uint16(0xc007)
    const TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA = Uint16(0xc009)
    const TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA = Uint16(0xc00a)
    const TLS_ECDHE_RSA_WITH_RC4_128_SHA = Uint16(0xc011)
    const TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA = Uint16(0xc012)
    const TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA = Uint16(0xc013)
    const TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA = Uint16(0xc014)
    const TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256 = Uint16(0xc023)
    const TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256 = Uint16(0xc027)
    const TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256 = Uint16(0xc02f)
    const TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256 = Uint16(0xc02b)
    const TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384 = Uint16(0xc030)
    const TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384 = Uint16(0xc02c)
    const TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256 = Uint16(0xcca8)
    const TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256 = Uint16(0xcca9)

    // TLS 1.3 cipher suites.
    const TLS_AES_128_GCM_SHA256 = Uint16(0x1301)
    const TLS_AES_256_GCM_SHA384 = Uint16(0x1302)
    const TLS_CHACHA20_POLY1305_SHA256 = Uint16(0x1303)

    // TLS_FALLBACK_SCSV isn't a standard cipher suite but an indicator
    // that the client is doing version fallback. See RFC 7507.
    const TLS_FALLBACK_SCSV = Uint16(0x5600)

    // Legacy names for the corresponding cipher suites with the correct _SHA256
    // suffix, retained for backward compatibility.
    const TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305 = TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256
    const TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305 = TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256


    const VersionTLS10 = Int(0x0301)
    const VersionTLS11 = Int(0x0302)
    const VersionTLS12 = Int(0x0303)
    const VersionTLS13 = Int(0x0304)

    // Deprecated: SSLv3 is cryptographically broken, and is no longer
    // supported by this package. See golang.org/issue/32716.
    const VersionSSL30 = Int(0x0300)

    function CipherSuiteName(id: Uint16 | NumberLike): string
    function Listen(network: string, laddr: string, config: ConfigPointer): net.Listener
    function NewListener(inner: net.Listener, config: ConfigPointer): net.Listener

    function LoadX509KeyPair(certFile: string, keyFile: string): Certificate
    function X509KeyPair(certPEMBlock: Bytes, keyPEMBlock: Bytes): Certificate
    interface Certificate extends Native {
        readonly __Certificate: Certificate

        Certificate: Slice<Bytes>
        // PrivateKey contains the private key corresponding to the public key in
        // Leaf. This must implement crypto.Signer with an RSA, ECDSA or Ed25519 PublicKey.
        // For a server up to TLS 1.2, it can also implement crypto.Decrypter with
        // an RSA PublicKey.
        PrivateKey: crypto.PrivateKey
        // SupportedSignatureAlgorithms is an optional list restricting what
        // signature algorithms the PrivateKey can be used for.
        SupportedSignatureAlgorithms: Array<SignatureScheme>
        // OCSPStaple contains an optional OCSP response which will be served
        // to clients that request it.
        OCSPStaple: Bytes
        // SignedCertificateTimestamps contains an optional list of Signed
        // Certificate Timestamps which will be served to clients that request it.
        SignedCertificateTimestamps: Slice<Bytes>
        // Leaf is the parsed form of the leaf certificate, which may be initialized
        // using x509.ParseCertificate to reduce per-handshake processing. If nil,
        // the leaf certificate will be parsed as needed.
        Leaf: x509.CertificatePointer
    }
    interface CertificatePointer extends Certificate {
        readonly __CertificatePointer: CertificatePointer
    }

    interface CertificateRequestInfo extends Native {
        readonly __CertificateRequestInfoPointer: CertificateRequestInfoPointer
        // AcceptableCAs contains zero or more, DER-encoded, X.501
        // Distinguished Names. These are the names of root or intermediate CAs
        // that the server wishes the returned certificate to be signed by. An
        // empty slice indicates that the server has no preference.
        AcceptableCAs: Slice<Bytes>

        // SignatureSchemes lists the signature schemes that the server is
        // willing to verify.
        SignatureSchemes: Array<SignatureScheme>

        // Version is the TLS version that was negotiated for this connection.
        Version: Uint16
        // contains filtered or unexported fields

        Context(): context.Context
        SupportsCertificate(c: x509.CertificatePointer): void
    }
    interface CertificateRequestInfoPointer extends CertificateRequestInfo {
        readonly __CertificateRequestInfoPointer: CertificateRequestInfoPointer
    }

    function CipherSuites(): Slice<CipherSuitePointer>
    function InsecureCipherSuites(): Slice<CipherSuitePointer>
    interface CipherSuitePointer extends Native {
        readonly __CipherSuitePointer: CipherSuitePointer
        ID: Uint16
        Name: string

        // Supported versions is the list of TLS protocol versions that can
        // negotiate this cipher suite.
        SupportedVersions: Uint16Slice

        // Insecure is true if the cipher suite has known security issues
        // due to its primitives, design, or implementation.
        Insecure: boolean
    }

    function ClientAuthType(v: any): ClientAuthType
    interface ClientAuthType extends Number {
        readonly __ClientAuthType: ClientAuthType

        String(): string
    }
    const NoClientCert = ClientAuthType(0)
    const RequestClientCert = ClientAuthType(1)
    const RequireAnyClientCert = ClientAuthType(2)
    const VerifyClientCertIfGiven = ClientAuthType(3)
    const RequireAndVerifyClientCert = ClientAuthType(4)

    interface ClientHelloInfoPointer extends Native {
        readonly __ClientHelloInfoPointer: ClientHelloInfoPointer
        CipherSuites: Uint16Slice
        ServerName: string
        SupportedCurves: Array<CurveID>
        SupportedPoints: Uint8Slice
        SignatureSchemes: Array<SignatureScheme>
        SupportedProtos: Array<string>
        SupportedVersions: Uint16Slice
        Conn: net.Conn

        Context(): context.Context
        SupportsCertificate(c: CertificatePointer): void
    }

    function NewLRUClientSessionCache(capacity: Int | NumberLike): ClientSessionCache
    interface ClientSessionCache extends Native {
        /** return (session :ClientSessionState, ok bool) */
        Get(sessionKey: string): [ClientSessionStatePointer, boolean]
        Put(sessionKey: string, cs: ClientSessionStatePointer): void
    }
    interface ClientSessionStatePointer extends Native {
        readonly __ClientSessionStatePointer: ClientSessionStatePointer
    }

    function Config(): ConfigPointer
    interface ConfigPointer extends Native {
        readonly __ConfigPointer: ConfigPointer
        // Rand provides the source of entropy for nonces and RSA blinding.
        // If Rand is nil, TLS uses the cryptographic random reader in package
        // crypto/rand.
        // The Reader must be safe for use by multiple goroutines.
        Rand: io.Reader

        // Time returns the current time as the number of seconds since the epoch.
        // If Time is nil, TLS uses time.Now.
        Time: () => time.Time

        // Certificates contains one or more certificate chains to present to the
        // other side of the connection. The first certificate compatible with the
        // peer's requirements is selected automatically.
        //
        // Server configurations must set one of Certificates, GetCertificate or
        // GetConfigForClient. Clients doing client-authentication may set either
        // Certificates or GetClientCertificate.
        //
        // Note: if there are multiple Certificates, and they don't have the
        // optional field Leaf set, certificate selection will incur a significant
        // per-handshake performance cost.
        Certificates: Slice<Certificate>

        // NameToCertificate maps from a certificate name to an element of
        // Certificates. Note that a certificate name can be of the form
        // '*.example.com' and so doesn't have to be a domain name as such.
        //
        // Deprecated: NameToCertificate only allows associating a single
        // certificate with a given name. Leave this field nil to let the library
        // select the first compatible chain from Certificates.
        NameToCertificate: Map<string, CertificatePointer>

        // GetCertificate returns a Certificate based on the given
        // ClientHelloInfo. It will only be called if the client supplies SNI
        // information or if Certificates is empty.
        //
        // If GetCertificate is nil or returns nil, then the certificate is
        // retrieved from NameToCertificate. If NameToCertificate is nil, the
        // best element of Certificates will be used.
        GetCertificate: (c: ClientHelloInfoPointer) => [CertificatePointer, Error]

        // GetClientCertificate, if not nil, is called when a server requests a
        // certificate from a client. If set, the contents of Certificates will
        // be ignored.
        //
        // If GetClientCertificate returns an error, the handshake will be
        // aborted and that error will be returned. Otherwise
        // GetClientCertificate must return a non-nil Certificate. If
        // Certificate.Certificate is empty then no certificate will be sent to
        // the server. If this is unacceptable to the server then it may abort
        // the handshake.
        //
        // GetClientCertificate may be called multiple times for the same
        // connection if renegotiation occurs or if TLS 1.3 is in use.
        GetClientCertificate: (c: CertificateRequestInfo) => [CertificatePointer, Error]

        // GetConfigForClient, if not nil, is called after a ClientHello is
        // received from a client. It may return a non-nil Config in order to
        // change the Config that will be used to handle this connection. If
        // the returned Config is nil, the original Config will be used. The
        // Config returned by this callback may not be subsequently modified.
        //
        // If GetConfigForClient is nil, the Config passed to Server() will be
        // used for all connections.
        //
        // If SessionTicketKey was explicitly set on the returned Config, or if
        // SetSessionTicketKeys was called on the returned Config, those keys will
        // be used. Otherwise, the original Config keys will be used (and possibly
        // rotated if they are automatically managed).
        GetConfigForClient: (c: ClientHelloInfoPointer) => [ConfigPointer, Error]

        // VerifyPeerCertificate, if not nil, is called after normal
        // certificate verification by either a TLS client or server. It
        // receives the raw ASN.1 certificates provided by the peer and also
        // any verified chains that normal processing found. If it returns a
        // non-nil error, the handshake is aborted and that error results.
        //
        // If normal verification fails then the handshake will abort before
        // considering this callback. If normal verification is disabled by
        // setting InsecureSkipVerify, or (for a server) when ClientAuth is
        // RequestClientCert or RequireAnyClientCert, then this callback will
        // be considered but the verifiedChains argument will always be nil.
        VerifyPeerCertificate: (rawCerts: Slice<Bytes>, verifiedChains: Slice<Slice<x509.CertificatePointer>>) => Error

        // VerifyConnection, if not nil, is called after normal certificate
        // verification and after VerifyPeerCertificate by either a TLS client
        // or server. If it returns a non-nil error, the handshake is aborted
        // and that error results.
        //
        // If normal verification fails then the handshake will abort before
        // considering this callback. This callback will run for all connections
        // regardless of InsecureSkipVerify or ClientAuth settings.
        VerifyConnection: (c: ConnectionState) => Error

        // RootCAs defines the set of root certificate authorities
        // that clients use when verifying server certificates.
        // If RootCAs is nil, TLS uses the host's root CA set.
        RootCAs: x509.CertPoolPointer

        // NextProtos is a list of supported application level protocols, in
        // order of preference. If both peers support ALPN, the selected
        // protocol will be one from this list, and the connection will fail
        // if there is no mutually supported protocol. If NextProtos is empty
        // or the peer doesn't support ALPN, the connection will succeed and
        // ConnectionState.NegotiatedProtocol will be empty.
        NextProtos: Array<string>

        // ServerName is used to verify the hostname on the returned
        // certificates unless InsecureSkipVerify is given. It is also included
        // in the client's handshake to support virtual hosting unless it is
        // an IP address.
        ServerName: string

        // ClientAuth determines the server's policy for
        // TLS Client Authentication. The default is NoClientCert.
        ClientAuth: ClientAuthType

        // ClientCAs defines the set of root certificate authorities
        // that servers use if required to verify a client certificate
        // by the policy in ClientAuth.
        ClientCAs: x509.CertPoolPointer

        // InsecureSkipVerify controls whether a client verifies the server's
        // certificate chain and host name. If InsecureSkipVerify is true, crypto/tls
        // accepts any certificate presented by the server and any host name in that
        // certificate. In this mode, TLS is susceptible to machine-in-the-middle
        // attacks unless custom verification is used. This should be used only for
        // testing or in combination with VerifyConnection or VerifyPeerCertificate.
        InsecureSkipVerify: boolean

        // CipherSuites is a list of enabled TLS 1.0–1.2 cipher suites. The order of
        // the list is ignored. Note that TLS 1.3 ciphersuites are not configurable.
        //
        // If CipherSuites is nil, a safe default list is used. The default cipher
        // suites might change over time.
        CipherSuites: Uint16Slice

        // PreferServerCipherSuites is a legacy field and has no effect.
        //
        // It used to control whether the server would follow the client's or the
        // server's preference. Servers now select the best mutually supported
        // cipher suite based on logic that takes into account inferred client
        // hardware, server hardware, and security.
        //
        // Deprected: PreferServerCipherSuites is ignored.
        PreferServerCipherSuites: boolean

        // SessionTicketsDisabled may be set to true to disable session ticket and
        // PSK (resumption) support. Note that on clients, session ticket support is
        // also disabled if ClientSessionCache is nil.
        SessionTicketsDisabled: boolean

        // SessionTicketKey is used by TLS servers to provide session resumption.
        // See RFC 5077 and the PSK mode of RFC 8446. If zero, it will be filled
        // with random data before the first server handshake.
        //
        // Deprecated: if this field is left at zero, session ticket keys will be
        // automatically rotated every day and dropped after seven days. For
        // customizing the rotation schedule or synchronizing servers that are
        // terminating connections for the same host, use SetSessionTicketKeys.
        SessionTicketKey: Bytes

        // ClientSessionCache is a cache of ClientSessionState entries for TLS
        // session resumption. It is only used by clients.
        ClientSessionCache: ClientSessionCache

        // MinVersion contains the minimum TLS version that is acceptable.
        // If zero, TLS 1.0 is currently taken as the minimum.
        MinVersion: Uint16

        // MaxVersion contains the maximum TLS version that is acceptable.
        // If zero, the maximum version supported by this package is used,
        // which is currently TLS 1.3.
        MaxVersion: Uint16

        // CurvePreferences contains the elliptic curves that will be used in
        // an ECDHE handshake, in preference order. If empty, the default will
        // be used. The client will use the first preference as the type for
        // its key share in TLS 1.3. This may change in the future.
        CurvePreferences: Slice<CurveID>

        // DynamicRecordSizingDisabled disables adaptive sizing of TLS records.
        // When true, the largest possible TLS record size is always used. When
        // false, the size of TLS records may be adjusted in an attempt to
        // improve latency.
        DynamicRecordSizingDisabled: boolean

        // Renegotiation controls what types of renegotiation are supported.
        // The default, none, is correct for the vast majority of applications.
        Renegotiation: RenegotiationSupport

        // KeyLogWriter optionally specifies a destination for TLS master secrets
        // in NSS key log format that can be used to allow external programs
        // such as Wireshark to decrypt TLS connections.
        // See https://developer.mozilla.org/en-US/docs/Mozilla/Projects/NSS/Key_Log_Format.
        // Use of KeyLogWriter compromises security and should only be
        // used for debugging.
        KeyLogWriter: io.Writer
        // contains filtered or unexported fields

        Clone(): Config
        SetSessionTicketKeys(keys: Slice<Slice<Byte>>): void
    }

    function Client(conn: net.Conn, config: ConfigPointer): ConnPointer
    function Dial(network: string, addr: string, config: ConfigPointer): ConnPointer
    function DialWithDialer(dialer: net.DialerPointer, network: string, addr: string, config: ConfigPointer): ConnPointer
    function Server(conn: net.Conn, config: ConfigPointer): ConnPointer
    interface ConnPointer extends Native {
        // contains filtered or unexported fields
    }
    interface ConnectionState extends Native {
        readonly __ConnectionState: ConnectionState
        // Version is the TLS version used by the connection (e.g. VersionTLS12).
        Version: Uint16

        // HandshakeComplete is true if the handshake has concluded.
        HandshakeComplete: boolean

        // DidResume is true if this connection was successfully resumed from a
        // previous session with a session ticket or similar mechanism.
        DidResume: boolean

        // CipherSuite is the cipher suite negotiated for the connection (e.g.
        // TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256, TLS_AES_128_GCM_SHA256).
        CipherSuite: Uint16

        // NegotiatedProtocol is the application protocol negotiated with ALPN.
        NegotiatedProtocol: string

        // NegotiatedProtocolIsMutual used to indicate a mutual NPN negotiation.
        //
        // Deprecated: this value is always true.
        NegotiatedProtocolIsMutual: boolean

        // ServerName is the value of the Server Name Indication extension sent by
        // the client. It's available both on the server and on the client side.
        ServerName: string

        // PeerCertificates are the parsed certificates sent by the peer, in the
        // order in which they were sent. The first element is the leaf certificate
        // that the connection is verified against.
        //
        // On the client side, it can't be empty. On the server side, it can be
        // empty if Config.ClientAuth is not RequireAnyClientCert or
        // RequireAndVerifyClientCert.
        PeerCertificates: Array<x509.CertificatePointer>

        // VerifiedChains is a list of one or more chains where the first element is
        // PeerCertificates[0] and the last element is from Config.RootCAs (on the
        // client side) or Config.ClientCAs (on the server side).
        //
        // On the client side, it's set if Config.InsecureSkipVerify is false. On
        // the server side, it's set if Config.ClientAuth is VerifyClientCertIfGiven
        // (and the peer provided a certificate) or RequireAndVerifyClientCert.
        VerifiedChains: Array<Array<x509.CertificatePointer>>

        // SignedCertificateTimestamps is a list of SCTs provided by the peer
        // through the TLS handshake for the leaf certificate, if any.
        SignedCertificateTimestamps: Slice<Bytes>

        // OCSPResponse is a stapled Online Certificate Status Protocol (OCSP)
        // response provided by the peer for the leaf certificate, if any.
        OCSPResponse: Bytes

        // TLSUnique contains the "tls-unique" channel binding value (see RFC 5929,
        // Section 3). This value will be nil for TLS 1.3 connections and for all
        // resumed connections.
        //
        // Deprecated: there are conditions in which this value might not be unique
        // to a connection. See the Security Considerations sections of RFC 5705 and
        // RFC 7627, and https://mitls.org/pages/attacks/3SHAKE#channelbindings.
        TLSUnique: Bytes
        // contains filtered or unexported fields
        ExportKeyingMaterial(label: string, context: Bytes, length: Int | NumberLike): Bytes
    }
    function CurveID(v: Uint16 | NumberLike): CurveID
    interface CurveID extends Number {
        readonly __CurveID: CurveID

        String(): string
    }
    const CurveP256 = CurveID(23)
    const CurveP384 = CurveID(24)
    const CurveP521 = CurveID(25)
    const X25519 = CurveID(29)

    function Dialer(dialer: net.DialerPointer, config: ConfigPointer): DialerPointer
    interface DialerPointer extends Native {
        readonly __DialerPointer: DialerPointer
        // NetDialer is the optional dialer to use for the TLS connections'
        // underlying TCP connections.
        // A nil NetDialer is equivalent to the net.Dialer zero value.
        NetDialer: net.DialerPointer

        // Config is the TLS configuration to use for new connections.
        // A nil configuration is equivalent to the zero
        // configuration; see the documentation of Config for the
        // defaults.
        Config: ConfigPointer

        Dial(network: string, addr: string): net.Conn
        DialContext(ctx: context.Context, network: string, addr: string): net.Conn
    }
    interface RecordHeaderError extends Error {
        readonly __RecordHeaderError: RecordHeaderError
        // Msg contains a human readable string that describes the error.
        Msg: string
        // RecordHeader contains the five bytes of TLS record header that
        // triggered the error.
        RecordHeader: Bytes
        // Conn provides the underlying net.Conn in the case that a client
        // sent an initial handshake that didn't look like TLS.
        // It is nil if there's already been a handshake or a TLS alert has
        // been written to the connection.
        Conn: net.Conn
    }
    function RenegotiationSupport(v: Int | NumberLike): RenegotiationSupport
    interface RenegotiationSupport extends Number {
        readonly __RenegotiationSupport: RenegotiationSupport
    }
    const RenegotiateNever = RenegotiationSupport(0)
    const RenegotiateOnceAsClient = RenegotiationSupport(1)
    const RenegotiateFreelyAsClient = RenegotiationSupport(2)

    function SignatureScheme(v: Uint16 | NumberLike): SignatureScheme
    interface SignatureScheme extends Number {
        readonly __SignatureScheme: SignatureScheme
        String(): string
    }
    const PKCS1WithSHA256 = SignatureScheme(0x0401)
    const PKCS1WithSHA384 = SignatureScheme(0x0501)
    const PKCS1WithSHA512 = SignatureScheme(0x0601)

    // RSASSA-PSS algorithms with public key OID rsaEncryption.
    const PSSWithSHA256 = SignatureScheme(0x0804)
    const PSSWithSHA384 = SignatureScheme(0x0805)
    const PSSWithSHA512 = SignatureScheme(0x0806)

    // ECDSA algorithms. Only constrained to a specific curve in TLS 1.3.
    const ECDSAWithP256AndSHA256 = SignatureScheme(0x0403)
    const ECDSAWithP384AndSHA384 = SignatureScheme(0x0503)
    const ECDSAWithP521AndSHA512 = SignatureScheme(0x0603)

    // EdDSA algorithms.
    const Ed25519 = SignatureScheme(0x0807)

    // Legacy signature and hash algorithms for TLS 1.2.
    const PKCS1WithSHA1 = SignatureScheme(0x0201)
    const ECDSAWithSHA1 = SignatureScheme(0x0203)

    function isCertificate(v: any): v is Certificate
    function isCertificatePointer(v: any): v is CertificatePointer
    function isCertificateRequestInfo(v: any): v is CertificateRequestInfo
    function isCertificateRequestInfoPointer(v: any): v is CertificateRequestInfoPointer
    function isCipherSuitePointer(v: any): v is CipherSuitePointer
    function isClientAuthType(v: any): v is ClientAuthType
    function isClientHelloInfoPointer(v: any): v is ClientHelloInfoPointer
    function isClientSessionStatePointer(v: any): v is ClientSessionStatePointer
    function isConfigPointer(v: any): v is ConfigPointer
    function isConnPointer(v: any): v is ConnPointer
    function isConnectionState(v: any): v is ConnectionState
    function isCurveID(v: any): v is CurveID
    function isDialerPointer(v: any): v is DialerPointer
    function isRecordHeaderError(v: any): v is RecordHeaderError
    function isRenegotiationSupport(v: any): v is RenegotiationSupport
    function isSignatureScheme(v: any): v is SignatureScheme
}