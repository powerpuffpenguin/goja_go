declare module "stdgo/net/http" {
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
    import * as time from "stdgo/time";
    import * as fmt from "stdgo/fmt";
    import * as bufio from "stdgo/bufio";
    import * as io from "stdgo/io";
    import * as fs from "stdgo/io/fs";
    import * as url from "stdgo/net/url";
    import * as tls from "stdgo/crypto/tls";
    import * as multipart from "stdgo/mime/multipart"

    const MethodGet = "GET"
    const MethodHead = "HEAD"
    const MethodPost = "POST"
    const MethodPut = "PUT"
    const MethodPatch = "PATCH" // RFC 5789
    const MethodDelete = "DELETE"
    const MethodConnect = "CONNECT"
    const MethodOptions = "OPTIONS"
    const MethodTrace = "TRACE"

    const StatusContinue = Int(100) // RFC 7231, 6.2.1
    const StatusSwitchingProtocols = Int(101) // RFC 7231, 6.2.2
    const StatusProcessing = Int(102) // RFC 2518, 10.1
    const StatusEarlyHints = Int(103) // RFC 8297
    const StatusOK = Int(200) // RFC 7231, 6.3.1
    const StatusCreated = Int(201) // RFC 7231, 6.3.2
    const StatusAccepted = Int(202) // RFC 7231, 6.3.3
    const StatusNonAuthoritativeInfo = Int(203) // RFC 7231, 6.3.4
    const StatusNoContent = Int(204) // RFC 7231, 6.3.5
    const StatusResetContent = Int(205) // RFC 7231, 6.3.6
    const StatusPartialContent = Int(206) // RFC 7233, 4.1
    const StatusMultiStatus = Int(207) // RFC 4918, 11.1
    const StatusAlreadyReported = Int(208) // RFC 5842, 7.1
    const StatusIMUsed = Int(226) // RFC 3229, 10.4.1
    const StatusMultipleChoices = Int(300) // RFC 7231, 6.4.1
    const StatusMovedPermanently = Int(301) // RFC 7231, 6.4.2
    const StatusFound = Int(302) // RFC 7231, 6.4.3
    const StatusSeeOther = Int(303) // RFC 7231, 6.4.4
    const StatusNotModified = Int(304) // RFC 7232, 4.1
    const StatusUseProxy = Int(305) // RFC 7231, 6.4.5
    const StatusTemporaryRedirect = Int(307) // RFC 7231, 6.4.7
    const StatusPermanentRedirect = Int(308) // RFC 7538, 3
    const StatusBadRequest = Int(400) // RFC 7231, 6.5.1
    const StatusUnauthorized = Int(401) // RFC 7235, 3.1
    const StatusPaymentRequired = Int(402) // RFC 7231, 6.5.2
    const StatusForbidden = Int(403) // RFC 7231, 6.5.3
    const StatusNotFound = Int(404) // RFC 7231, 6.5.4
    const StatusMethodNotAllowed = Int(405) // RFC 7231, 6.5.5
    const StatusNotAcceptable = Int(406) // RFC 7231, 6.5.6
    const StatusProxyAuthRequired = Int(407) // RFC 7235, 3.2
    const StatusRequestTimeout = Int(408) // RFC 7231, 6.5.7
    const StatusConflict = Int(409) // RFC 7231, 6.5.8
    const StatusGone = Int(410) // RFC 7231, 6.5.9
    const StatusLengthRequired = Int(411) // RFC 7231, 6.5.10
    const StatusPreconditionFailed = Int(412) // RFC 7232, 4.2
    const StatusRequestEntityTooLarge = Int(413) // RFC 7231, 6.5.11
    const StatusRequestURITooLong = Int(414) // RFC 7231, 6.5.12
    const StatusUnsupportedMediaType = Int(415) // RFC 7231, 6.5.13
    const StatusRequestedRangeNotSatisfiable = Int(416) // RFC 7233, 4.4
    const StatusExpectationFailed = Int(417) // RFC 7231, 6.5.14
    const StatusTeapot = Int(418) // RFC 7168, 2.3.3
    const StatusMisdirectedRequest = Int(421) // RFC 7540, 9.1.2
    const StatusUnprocessableEntity = Int(422) // RFC 4918, 11.2
    const StatusLocked = Int(423) // RFC 4918, 11.3
    const StatusFailedDependency = Int(424) // RFC 4918, 11.4
    const StatusTooEarly = Int(425) // RFC 8470, 5.2.
    const StatusUpgradeRequired = Int(426) // RFC 7231, 6.5.15
    const StatusPreconditionRequired = Int(428) // RFC 6585, 3
    const StatusTooManyRequests = Int(429) // RFC 6585, 4
    const StatusRequestHeaderFieldsTooLarge = Int(431) // RFC 6585, 5
    const StatusUnavailableForLegalReasons = Int(451) // RFC 7725, 3
    const StatusInternalServerError = Int(500) // RFC 7231, 6.6.1
    const StatusNotImplemented = Int(501) // RFC 7231, 6.6.2
    const StatusBadGateway = Int(502) // RFC 7231, 6.6.3
    const StatusServiceUnavailable = Int(503) // RFC 7231, 6.6.4
    const StatusGatewayTimeout = Int(504) // RFC 7231, 6.6.5
    const StatusHTTPVersionNotSupported = Int(505) // RFC 7231, 6.6.6
    const StatusVariantAlsoNegotiates = Int(506) // RFC 2295, 8.1
    const StatusInsufficientStorage = Int(507) // RFC 4918, 11.5
    const StatusLoopDetected = Int(508) // RFC 5842, 7.2
    const StatusNotExtended = Int(510) // RFC 2774, 7
    const StatusNetworkAuthenticationRequired = Int(511) // RFC 6585, 6

    const DefaultMaxHeaderBytes = Int(1 << 20) // 1 MB
    const DefaultMaxIdleConnsPerHost = Int(2)
    const TimeFormat = "Mon, 02 Jan 2006 15:04:05 GMT"
    const TrailerPrefix = "Trailer:"

    const ErrNotSupported: Error
    const ErrMissingBoundary: Error
    const ErrNotMultipart: Error
    const ErrHeaderTooLong: Error
    const ErrShortBody: Error
    const ErrMissingContentLength: Error
    const ErrBodyNotAllowed: Error
    const ErrHijacked: Error
    const ErrContentLength: Error
    const ErrWriteAfterFlush: Error

    const ServerContextKey: fmt.Stringer
    const LocalAddrContextKey: fmt.Stringer

    const DefaultClient: ClientPointer
    const DefaultServeMux: ServeMuxPointer
    const ErrAbortHandler: Error
    const ErrBodyReadAfterClose: Error
    const ErrHandlerTimeout: Error
    const ErrLineTooLong: Error
    const ErrMissingFile: Error
    const ErrNoCookie: Error
    const ErrNoLocation: Error
    const ErrServerClosed: Error
    const ErrSkipAltProtocol: Error
    const ErrUseLastResponse: Error
    const NoBody: noBody
    interface noBody extends io.ReadCloser, io.WriterTo { }

    function CanonicalHeaderKey(s: string): string
    function DetectContentType(data: Bytes): string
    function Error(w: ResponseWriter, error: string, code: Int | NumberLike): void
    function Handle(pattern: string, handler: Handler): void
    function HandleFunc(pattern: string, handler: (w: ResponseWriter, r: RequestPointer) => void): void
    function ListenAndServe(addr: string, handler: Handler): void
    function ListenAndServeTLS(addr: string, certFile: string, keyFile: string, handler: Handler): void
    function MaxBytesReader(w: ResponseWriter, r: io.ReadCloser, n: Int64 | NumberLike): io.ReadCloser
    function NotFound(w: ResponseWriter, r: RequestPointer): void
    /** return (major, minor int, ok bool) */
    function ParseHTTPVersion(vers: string): [Int, Int, boolean]
    function ParseTime(text: string): time.Time
    function ProxyFromEnvironment(req: RequestPointer): url.URLPointer
    function ProxyURL(fixedURL: url.URLPointer): (r: RequestPointer) => [url.URLPointer, Error]
    function Redirect(w: ResponseWriter, r: RequestPointer, url: string, code: Int | NumberLike): void
    function Serve(l: net.Listener, handler: Handler): void
    function ServeContent(w: ResponseWriter, req: RequestPointer, name: string, modtime: time.Time, content: io.ReadSeeker): void
    function ServeFile(w: ResponseWriter, r: RequestPointer, name: string): void
    function ServeTLS(l: net.Listener, handler: Handler, certFile: string, keyFile: string): void
    function SetCookie(w: ResponseWriter, cookie: CookiePointer): void
    function StatusText(code: Int | NumberLike): string

    function Client(): ClientPointer
    interface ClientPointer extends Native {
        readonly __ClientPointer: ClientPointer
        // Transport specifies the mechanism by which individual
        // HTTP requests are made.
        // If nil, DefaultTransport is used.
        Transport: RoundTripper

        // CheckRedirect specifies the policy for handling redirects.
        // If CheckRedirect is not nil, the client calls it before
        // following an HTTP redirect. The arguments req and via are
        // the upcoming request and the requests made already, oldest
        // first. If CheckRedirect returns an error, the Client's Get
        // method returns both the previous Response (with its Body
        // closed) and CheckRedirect's error (wrapped in a url.Error)
        // instead of issuing the Request req.
        // As a special case, if CheckRedirect returns ErrUseLastResponse,
        // then the most recent response is returned with its body
        // unclosed, along with a nil error.
        //
        // If CheckRedirect is nil, the Client uses its default policy,
        // which is to stop after 10 consecutive requests.
        CheckRedirect: (req: RequestPointer, via: Slice<RequestPointer>) => Error

        // Jar specifies the cookie jar.
        //
        // The Jar is used to insert relevant cookies into every
        // outbound Request and is updated with the cookie values
        // of every inbound Response. The Jar is consulted for every
        // redirect that the Client follows.
        //
        // If Jar is nil, cookies are only sent if they are explicitly
        // set on the Request.
        Jar: CookieJar

        // Timeout specifies a time limit for requests made by this
        // Client. The timeout includes connection time, any
        // redirects, and reading the response body. The timer remains
        // running after Get, Head, Post, or Do return and will
        // interrupt reading of the Response.Body.
        //
        // A Timeout of zero means no timeout.
        //
        // The Client cancels requests to the underlying Transport
        // as if the Request's Context ended.
        //
        // For compatibility, the Client will also use the deprecated
        // CancelRequest method on Transport if found. New
        // RoundTripper implementations should use the Request's Context
        // for cancellation instead of implementing CancelRequest.
        Timeout: time.Duration

        CloseIdleConnections()
        Do(req: RequestPointer): ResponsePointer
        Get(url: string): ResponsePointer
        Head(url: string): ResponsePointer
        Post(url: string, contentType: string, body: io.Reader): ResponsePointer
        PostForm(url: string, data: url.Values): ResponsePointer
    }

    interface ConnState extends Number {
        readonly __ConnState: ConnState
        String(): string
    }
    const StateNew: ConnState
    const StateActive: ConnState
    const StateIdle: ConnState
    const StateHijacked: ConnState
    const StateClosed: ConnState

    function Cookie(name: string, value: string): CookiePointer
    interface CookiePointer extends Native {
        readonly __CookiePointer: CookiePointer
        Name: string
        Value: string

        Path: string    // optional
        Domain: string    // optional
        Expires: time.Time // optional
        RawExpires: string    // for reading cookies only

        // MaxAge=0 means no 'Max-Age' attribute specified.
        // MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
        // MaxAge>0 means Max-Age attribute present and given in seconds
        MaxAge: Int
        Secure: boolean
        HttpOnly: boolean
        SameSite: SameSite
        Raw: string
        Unparsed: Array<string> // Raw text of unparsed attribute-value pairs

        String(): string
    }

    interface CookieJar extends Native {
        // SetCookies handles the receipt of the cookies in a reply for the
        // given URL.  It may or may not choose to save the cookies, depending
        // on the jar's policy and implementation.
        SetCookies(u: url.URLPointer, cookies: Slice<CookiePointer>): void

        // Cookies returns the cookies to send in a request for the given URL.
        // It is up to the implementation to honor the standard cookie use
        // restrictions such as in RFC 6265.
        Cookies(u: url.URLPointer): Slice<Cookie>
    }
    interface Dir extends Native {
        readonly __Dir: Dir

        Open(name: string): File
    }
    interface File extends io.Closer, io.Reader, io.Seeker {

        Readdir(count: Int | Number): Array<fs.FileInfo>
        Stat(): fs.FileInfo
    }

    function FS(fsys: fs.FS): FileSystem
    interface FileSystem extends Native {
        Open(name: string): File
    }
    interface Flusher extends Native {
        // Flush sends any buffered data to the client.
        Flush(): void
    }

    function FileServer(root: FileSystem): Handler
    function NotFoundHandler(): Handler
    function RedirectHandler(url: string, code: Int | NumberLike): Handler
    function StripPrefix(prefix: string, h: Handler): Handler
    function TimeoutHandler(h: Handler, dt: time.Duration, msg: string): Handler
    function Handler(f: (w: ResponseWriter, r: RequestPointer) => void | Promise<any>): Handler
    interface Handler extends Native {
        ServeHTTP(w: ResponseWriter, req: RequestPointer): void
    }
    interface HandlerFunc extends Handler {
        readonly __HandlerFunc: HandlerFunc
    }

    function Header(): Header
    interface Header extends Map<string, Array<string>> {
        Add(key: string, value: string): void
        Clone(): Header
        Del(key: string): void
        Get(key: string): string
        Set(key: string, value: string): void
        Values(key: string): Array<string>
        Write(w: io.Writer): void
        WriteSubset(w: io.Writer, exclude: Map<string, boolean>): void
    }
    interface Hijacker extends Native {
        /** return (net.Conn, *bufio.ReadWriter, error) */
        Hijack(): [net.Conn, bufio.ReadWriterPointer]
    }

    function PushOptions(method: string, header: Header): PushOptionsPointer
    interface PushOptionsPointer extends Native {
        readonly PushOptionsPointer: PushOptionsPointer
        // Method specifies the HTTP method for the promised request.
        // If set, it must be "GET" or "HEAD". Empty means "GET".
        Method: string

        // Header specifies additional promised request headers. This cannot
        // include HTTP/2 pseudo header fields like ":path" and ":scheme",
        // which will be added automatically.
        Header: Header
    }
    interface Pusher extends Native {
        // Push initiates an HTTP/2 server push. This constructs a synthetic
        // request using the given target and options, serializes that request
        // into a PUSH_PROMISE frame, then dispatches that request using the
        // server's request handler. If opts is nil, default options are used.
        //
        // The target must either be an absolute path (like "/path") or an absolute
        // URL that contains a valid host and the same scheme as the parent request.
        // If the target is a path, it will inherit the scheme and host of the
        // parent request.
        //
        // The HTTP/2 spec disallows recursive pushes and cross-authority pushes.
        // Push may or may not detect these invalid pushes; however, invalid
        // pushes will be detected and canceled by conforming clients.
        //
        // Handlers that wish to push URL X should call Push before sending any
        // data that may trigger a request for URL X. This avoids a race where the
        // client issues requests for X before receiving the PUSH_PROMISE for X.
        //
        // Push will run in a separate goroutine making the order of arrival
        // non-deterministic. Any required synchronization needs to be implemented
        // by the caller.
        //
        // Push returns ErrNotSupported if the client has disabled push or if push
        // is not supported on the underlying connection.
        Push(target: string, opts: PushOptionsPointer): void
    }

    function NewRequest(method: string, url: string, body: io.Reader): RequestPointer
    function NewRequestWithContext(ctx: context.Context, method: string, url: string, body: io.Reader): RequestPointer
    function ReadRequest(b: bufio.ReaderPointer): RequestPointer
    interface RequestPointer extends Native {
        readonly __RequestPointer: RequestPointer
        // Method specifies the HTTP method (GET, POST, PUT, etc.).
        // For client requests, an empty string means GET.
        //
        // Go's HTTP client does not support sending a request with
        // the CONNECT method. See the documentation on Transport for
        // details.
        Method: string

        // URL specifies either the URI being requested (for server
        // requests) or the URL to access (for client requests).
        //
        // For server requests, the URL is parsed from the URI
        // supplied on the Request-Line as stored in RequestURI.  For
        // most requests, fields other than Path and RawQuery will be
        // empty. (See RFC 7230, Section 5.3)
        //
        // For client requests, the URL's Host specifies the server to
        // connect to, while the Request's Host field optionally
        // specifies the Host header value to send in the HTTP
        // request.
        URL: url.URLPointer

        // The protocol version for incoming server requests.
        //
        // For client requests, these fields are ignored. The HTTP
        // client code always uses either HTTP/1.1 or HTTP/2.
        // See the docs on Transport for details.
        Proto: string // "HTTP/1.0"
        ProtoMajor: Int    // 1
        ProtoMinor: Int    // 0

        // Header contains the request header fields either received
        // by the server or to be sent by the client.
        //
        // If a server received a request with header lines,
        //
        //	Host: example.com
        //	accept-encoding: gzip, deflate
        //	Accept-Language: en-us
        //	fOO: Bar
        //	foo: two
        //
        // then
        //
        //	Header = map[string][]string{
        //		"Accept-Encoding": {"gzip, deflate"},
        //		"Accept-Language": {"en-us"},
        //		"Foo": {"Bar", "two"},
        //	}
        //
        // For incoming requests, the Host header is promoted to the
        // Request.Host field and removed from the Header map.
        //
        // HTTP defines that header names are case-insensitive. The
        // request parser implements this by using CanonicalHeaderKey,
        // making the first character and any characters following a
        // hyphen uppercase and the rest lowercase.
        //
        // For client requests, certain headers such as Content-Length
        // and Connection are automatically written when needed and
        // values in Header may be ignored. See the documentation
        // for the Request.Write method.
        Header: Header

        // Body is the request's body.
        //
        // For client requests, a nil body means the request has no
        // body, such as a GET request. The HTTP Client's Transport
        // is responsible for calling the Close method.
        //
        // For server requests, the Request Body is always non-nil
        // but will return EOF immediately when no body is present.
        // The Server will close the request body. The ServeHTTP
        // Handler does not need to.
        //
        // Body must allow Read to be called concurrently with Close.
        // In particular, calling Close should unblock a Read waiting
        // for input.
        Body: io.ReadCloser

        // GetBody defines an optional func to return a new copy of
        // Body. It is used for client requests when a redirect requires
        // reading the body more than once. Use of GetBody still
        // requires setting Body.
        //
        // For server requests, it is unused.
        GetBody: () => [io.ReadCloser, Error]

        // ContentLength records the length of the associated content.
        // The value -1 indicates that the length is unknown.
        // Values >= 0 indicate that the given number of bytes may
        // be read from Body.
        //
        // For client requests, a value of 0 with a non-nil Body is
        // also treated as unknown.
        ContentLength: Int64

        // TransferEncoding lists the transfer encodings from outermost to
        // innermost. An empty list denotes the "identity" encoding.
        // TransferEncoding can usually be ignored; chunked encoding is
        // automatically added and removed as necessary when sending and
        // receiving requests.
        TransferEncoding: Array<string>

        // Close indicates whether to close the connection after
        // replying to this request (for servers) or after sending this
        // request and reading its response (for clients).
        //
        // For server requests, the HTTP server handles this automatically
        // and this field is not needed by Handlers.
        //
        // For client requests, setting this field prevents re-use of
        // TCP connections between requests to the same hosts, as if
        // Transport.DisableKeepAlives were set.
        Close: boolean

        // For server requests, Host specifies the host on which the
        // URL is sought. For HTTP/1 (per RFC 7230, section 5.4), this
        // is either the value of the "Host" header or the host name
        // given in the URL itself. For HTTP/2, it is the value of the
        // ":authority" pseudo-header field.
        // It may be of the form "host:port". For international domain
        // names, Host may be in Punycode or Unicode form. Use
        // golang.org/x/net/idna to convert it to either format if
        // needed.
        // To prevent DNS rebinding attacks, server Handlers should
        // validate that the Host header has a value for which the
        // Handler considers itself authoritative. The included
        // ServeMux supports patterns registered to particular host
        // names and thus protects its registered Handlers.
        //
        // For client requests, Host optionally overrides the Host
        // header to send. If empty, the Request.Write method uses
        // the value of URL.Host. Host may contain an international
        // domain name.
        Host: string

        // Form contains the parsed form data, including both the URL
        // field's query parameters and the PATCH, POST, or PUT form data.
        // This field is only available after ParseForm is called.
        // The HTTP client ignores Form and uses Body instead.
        Form: url.Values

        // PostForm contains the parsed form data from PATCH, POST
        // or PUT body parameters.
        //
        // This field is only available after ParseForm is called.
        // The HTTP client ignores PostForm and uses Body instead.
        PostForm: url.Values

        // MultipartForm is the parsed multipart form, including file uploads.
        // This field is only available after ParseMultipartForm is called.
        // The HTTP client ignores MultipartForm and uses Body instead.
        MultipartForm: multipart.FormPointer

        // Trailer specifies additional headers that are sent after the request
        // body.
        //
        // For server requests, the Trailer map initially contains only the
        // trailer keys, with nil values. (The client declares which trailers it
        // will later send.)  While the handler is reading from Body, it must
        // not reference Trailer. After reading from Body returns EOF, Trailer
        // can be read again and will contain non-nil values, if they were sent
        // by the client.
        //
        // For client requests, Trailer must be initialized to a map containing
        // the trailer keys to later send. The values may be nil or their final
        // values. The ContentLength must be 0 or -1, to send a chunked request.
        // After the HTTP request is sent the map values can be updated while
        // the request body is read. Once the body returns EOF, the caller must
        // not mutate Trailer.
        //
        // Few HTTP clients, servers, or proxies support HTTP trailers.
        Trailer: Header

        // RemoteAddr allows HTTP servers and other software to record
        // the network address that sent the request, usually for
        // logging. This field is not filled in by ReadRequest and
        // has no defined format. The HTTP server in this package
        // sets RemoteAddr to an "IP:port" address before invoking a
        // handler.
        // This field is ignored by the HTTP client.
        RemoteAddr: string

        // RequestURI is the unmodified request-target of the
        // Request-Line (RFC 7230, Section 3.1.1) as sent by the client
        // to a server. Usually the URL field should be used instead.
        // It is an error to set this field in an HTTP client request.
        RequestURI: string

        // TLS allows HTTP servers and other software to record
        // information about the TLS connection on which the request
        // was received. This field is not filled in by ReadRequest.
        // The HTTP server in this package sets the field for
        // TLS-enabled connections before invoking a handler;
        // otherwise it leaves the field nil.
        // This field is ignored by the HTTP client.
        TLS: tls.ConnectionState

        // Cancel is an optional channel whose closure indicates that the client
        // request should be regarded as canceled. Not all implementations of
        // RoundTripper may support Cancel.
        //
        // For server requests, this field is not applicable.
        //
        // Deprecated: Set the Request's context with NewRequestWithContext
        // instead. If a Request's Cancel field and context are both
        // set, it is undefined whether Cancel is respected.
        Cancel: ReadChannel<null>

        // Response is the redirect response which caused this request
        // to be created. This field is only populated during client
        // redirects.
        Response: ResponsePointer
        // contains filtered or unexported fields

        AddCookie(c: CookiePointer): void
        /** return (username, password string, ok bool) */
        BasicAuth(): [string, string, boolean]
        Clone(ctx: context.Context): RequestPointer
        Context(): context.Context
        Cookie(name: string): CookiePointer
        Cookies(): Array<CookiePointer>
        /** return (multipart.File, * multipart.FileHeader, error) */
        FormFile(key: string): [multipart.File, multipart.FileHeaderPointer]
        FormValue(key: string): string
        MultipartReader(): multipart.ReaderPointer
        ParseForm(): void
        ParseMultipartForm(maxMemory: Int64 | NumberLike): void
        PostFormValue(key: string): string
        ProtoAtLeast(major: Int | NumberLike, minor: Int | NumberLike): boolean
        Referer(): string
        SetBasicAuth(username: string, password: string): void
        UserAgent(): string
        WithContext(ctx: context.Context): RequestPointer
        Write(w: io.Writer): void
        WriteProxy(w: io.Writer): void
    }

    function Get(url: string): ResponsePointer
    function Head(url: string): ResponsePointer
    function Post(url: string, contentType: string, body: io.Reader): ResponsePointer
    function PostForm(url: string, data: url.Values): ResponsePointer
    function ReadResponse(r: bufio.ReaderPointer, req: RequestPointer): ResponsePointer
    interface ResponsePointer extends Native {
        readonly __ResponsePointer: ResponsePointer

        Status: string // e.g. "200 OK"
        StatusCode: Int    // e.g. 200
        Proto: string // e.g. "HTTP/1.0"
        ProtoMajor: Int    // e.g. 1
        ProtoMinor: Int    // e.g. 0

        // Header maps header keys to values. If the response had multiple
        // headers with the same key, they may be concatenated, with comma
        // delimiters.  (RFC 7230, section 3.2.2 requires that multiple headers
        // be semantically equivalent to a comma-delimited sequence.) When
        // Header values are duplicated by other fields in this struct (e.g.,
        // ContentLength, TransferEncoding, Trailer), the field values are
        // authoritative.
        //
        // Keys in the map are canonicalized (see CanonicalHeaderKey).
        Header: Header

        // Body represents the response body.
        //
        // The response body is streamed on demand as the Body field
        // is read. If the network connection fails or the server
        // terminates the response, Body.Read calls return an error.
        //
        // The http Client and Transport guarantee that Body is always
        // non-nil, even on responses without a body or responses with
        // a zero-length body. It is the caller's responsibility to
        // close Body. The default HTTP client's Transport may not
        // reuse HTTP/1.x "keep-alive" TCP connections if the Body is
        // not read to completion and closed.
        //
        // The Body is automatically dechunked if the server replied
        // with a "chunked" Transfer-Encoding.
        //
        // As of Go 1.12, the Body will also implement io.Writer
        // on a successful "101 Switching Protocols" response,
        // as used by WebSockets and HTTP/2's "h2c" mode.
        Body: io.ReadCloser

        // ContentLength records the length of the associated content. The
        // value -1 indicates that the length is unknown. Unless Request.Method
        // is "HEAD", values >= 0 indicate that the given number of bytes may
        // be read from Body.
        ContentLength: Int64

        // Contains transfer encodings from outer-most to inner-most. Value is
        // nil, means that "identity" encoding is used.
        TransferEncoding: Array<string>

        // Close records whether the header directed that the connection be
        // closed after reading Body. The value is advice for clients: neither
        // ReadResponse nor Response.Write ever closes a connection.
        Close: boolean

        // Uncompressed reports whether the response was sent compressed but
        // was decompressed by the http package. When true, reading from
        // Body yields the uncompressed content instead of the compressed
        // content actually set from the server, ContentLength is set to -1,
        // and the "Content-Length" and "Content-Encoding" fields are deleted
        // from the responseHeader. To get the original response from
        // the server, set Transport.DisableCompression to true.
        Uncompressed: boolean

        // Trailer maps trailer keys to values in the same
        // format as Header.
        //
        // The Trailer initially contains only nil values, one for
        // each key specified in the server's "Trailer" header
        // value. Those values are not added to Header.
        //
        // Trailer must not be accessed concurrently with Read calls
        // on the Body.
        //
        // After Body.Read has returned io.EOF, Trailer will contain
        // any trailer values sent by the server.
        Trailer: Header

        // Request is the request that was sent to obtain this Response.
        // Request's Body is nil (having already been consumed).
        // This is only populated for Client requests.
        Request: RequestPointer

        // TLS contains information about the TLS connection on which the
        // response was received. It is nil for unencrypted responses.
        // The pointer is shared between responses and should not be
        // modified.
        TLS: tls.ConnectionState

        Cookies(): Array<Cookie>
        Location(): url.URLPointer
        ProtoAtLeast(major: Int | NumberLike, minor: Int | NumberLike): boolean
        Write(w: io.Writer): void
    }

    interface ResponseWriter extends Native {
        // Header returns the header map that will be sent by
        // WriteHeader. The Header map also is the mechanism with which
        // Handlers can set HTTP trailers.
        //
        // Changing the header map after a call to WriteHeader (or
        // Write) has no effect unless the modified headers are
        // trailers.
        //
        // There are two ways to set Trailers. The preferred way is to
        // predeclare in the headers which trailers you will later
        // send by setting the "Trailer" header to the names of the
        // trailer keys which will come later. In this case, those
        // keys of the Header map are treated as if they were
        // trailers. See the example. The second way, for trailer
        // keys not known to the Handler until after the first Write,
        // is to prefix the Header map keys with the TrailerPrefix
        // constant value. See TrailerPrefix.
        //
        // To suppress automatic response headers (such as "Date"), set
        // their value to nil.
        Header(): Header

        // Write writes the data to the connection as part of an HTTP reply.
        //
        // If WriteHeader has not yet been called, Write calls
        // WriteHeader(http.StatusOK) before writing the data. If the Header
        // does not contain a Content-Type line, Write adds a Content-Type set
        // to the result of passing the initial 512 bytes of written data to
        // DetectContentType. Additionally, if the total size of all written
        // data is under a few KB and there are no Flush calls, the
        // Content-Length header is added automatically.
        //
        // Depending on the HTTP protocol version and the client, calling
        // Write or WriteHeader may prevent future reads on the
        // Request.Body. For HTTP/1.x requests, handlers should read any
        // needed request body data before writing the response. Once the
        // headers have been flushed (due to either an explicit Flusher.Flush
        // call or writing enough data to trigger a flush), the request body
        // may be unavailable. For HTTP/2 requests, the Go HTTP server permits
        // handlers to continue to read the request body while concurrently
        // writing the response. However, such behavior may not be supported
        // by all HTTP/2 clients. Handlers should read before writing if
        // possible to maximize compatibility.
        Write(b: Bytes): Int

        // WriteHeader sends an HTTP response header with the provided
        // status code.
        //
        // If WriteHeader is not called explicitly, the first call to Write
        // will trigger an implicit WriteHeader(http.StatusOK).
        // Thus explicit calls to WriteHeader are mainly used to
        // send error codes.
        //
        // The provided code must be a valid HTTP 1xx-5xx status code.
        // Only one header may be written. Go does not currently
        // support sending user-defined 1xx informational headers,
        // with the exception of 100-continue response header that the
        // Server sends automatically when the Request.Body is read.
        WriteHeader(statusCode: Int | NumberLike): void
    }
    function NewFileTransport(fs: FileSystem): RoundTripper
    interface RoundTripper extends Native {
        // RoundTrip executes a single HTTP transaction, returning
        // a Response for the provided Request.
        //
        // RoundTrip should not attempt to interpret the response. In
        // particular, RoundTrip must return err == nil if it obtained
        // a response, regardless of the response's HTTP status code.
        // A non-nil err should be reserved for failure to obtain a
        // response. Similarly, RoundTrip should not attempt to
        // handle higher-level protocol details such as redirects,
        // authentication, or cookies.
        //
        // RoundTrip should not modify the request, except for
        // consuming and closing the Request's Body. RoundTrip may
        // read fields of the request in a separate goroutine. Callers
        // should not mutate or reuse the request until the Response's
        // Body has been closed.
        //
        // RoundTrip must always close the body, including on errors,
        // but depending on the implementation may do so in a separate
        // goroutine even after RoundTrip returns. This means that
        // callers wanting to reuse the body for subsequent requests
        // must arrange to wait for the Close call before doing so.
        //
        // The Request's URL and Header fields must be initialized.
        RoundTrip(req: RequestPointer): ResponsePointer
    }

    interface SameSite extends Number {
        readonly __SameSite: SameSite
    }
    const SameSiteDefaultMode: SameSite
    const SameSiteLaxMode: SameSite
    const SameSiteStrictMode: SameSite
    const SameSiteNoneMode: SameSite

    function NewServeMux(): ServeMuxPointer
    interface ServeMuxPointer extends Handler {
        readonly __ServeMuxPointer: ServeMuxPointer

        Handle(pattern: string, handler: Handler): void
        HandleFunc(pattern: string, handler: (w: ResponseWriter, req: RequestPointer) => void): void
        /** return (h Handler, pattern string) */
        Handler(r: RequestPointer): [Handler, string]
    }

    function Server(): ServerPointer
    interface ServerPointer extends Native {
        readonly __ServerPointer: ServerPointer
        // Addr optionally specifies the TCP address for the server to listen on,
        // in the form "host:port". If empty, ":http" (port 80) is used.
        // The service names are defined in RFC 6335 and assigned by IANA.
        // See net.Dial for details of the address format.
        Addr: string

        Handler: Handler // handler to invoke, http.DefaultServeMux if nil

        // TLSConfig optionally provides a TLS configuration for use
        // by ServeTLS and ListenAndServeTLS. Note that this value is
        // cloned by ServeTLS and ListenAndServeTLS, so it's not
        // possible to modify the configuration with methods like
        // tls.Config.SetSessionTicketKeys. To use
        // SetSessionTicketKeys, use Server.Serve with a TLS Listener
        // instead.
        TLSConfig: tls.ConfigPointer

        // ReadTimeout is the maximum duration for reading the entire
        // request, including the body. A zero or negative value means
        // there will be no timeout.
        //
        // Because ReadTimeout does not let Handlers make per-request
        // decisions on each request body's acceptable deadline or
        // upload rate, most users will prefer to use
        // ReadHeaderTimeout. It is valid to use them both.
        ReadTimeout: time.Duration

        // ReadHeaderTimeout is the amount of time allowed to read
        // request headers. The connection's read deadline is reset
        // after reading the headers and the Handler can decide what
        // is considered too slow for the body. If ReadHeaderTimeout
        // is zero, the value of ReadTimeout is used. If both are
        // zero, there is no timeout.
        ReadHeaderTimeout: time.Duration

        // WriteTimeout is the maximum duration before timing out
        // writes of the response. It is reset whenever a new
        // request's header is read. Like ReadTimeout, it does not
        // let Handlers make decisions on a per-request basis.
        // A zero or negative value means there will be no timeout.
        WriteTimeout: time.Duration

        // IdleTimeout is the maximum amount of time to wait for the
        // next request when keep-alives are enabled. If IdleTimeout
        // is zero, the value of ReadTimeout is used. If both are
        // zero, there is no timeout.
        IdleTimeout: time.Duration

        // MaxHeaderBytes controls the maximum number of bytes the
        // server will read parsing the request header's keys and
        // values, including the request line. It does not limit the
        // size of the request body.
        // If zero, DefaultMaxHeaderBytes is used.
        MaxHeaderBytes: Int

        // TLSNextProto optionally specifies a function to take over
        // ownership of the provided TLS connection when an ALPN
        // protocol upgrade has occurred. The map key is the protocol
        // name negotiated. The Handler argument should be used to
        // handle HTTP requests and will initialize the Request's TLS
        // and RemoteAddr if not already set. The connection is
        // automatically closed when the function returns.
        // If TLSNextProto is not nil, HTTP/2 support is not enabled
        // automatically.
        TLSNextProto: Map<string, (s: ServerPointer, conn: tls.ConnPointer, h: Handler) => void>

        // ConnState specifies an optional callback function that is
        // called when a client connection changes state. See the
        // ConnState type and associated constants for details.
        ConnState: (c: net.Conn, s: ConnState) => void

        // ErrorLog specifies an optional logger for errors accepting
        // connections, unexpected behavior from handlers, and
        // underlying FileSystem errors.
        // If nil, logging is done via the log package's standard logger.
        ErrorLog: log.Logger

        // BaseContext optionally specifies a function that returns
        // the base context for incoming requests on this server.
        // The provided Listener is the specific Listener that's
        // about to start accepting requests.
        // If BaseContext is nil, the default is context.Background().
        // If non-nil, it must return a non-nil context.
        BaseContext: (l: net.Listener) => context.Context

        // ConnContext optionally specifies a function that modifies
        // the context used for a new connection c. The provided ctx
        // is derived from the base context and has a ServerContextKey
        // value.
        ConnContext: (ctx: context.Context, c: net.Conn) => context.Context
        // contains filtered or unexported fields

        Close(): void
        ListenAndServe(): void
        ListenAndServeTLS(certFile: string, keyFile: string): void
        RegisterOnShutdown(f: () => void): void
        Serve(l: net.Listener): void
        ServeTLS(l: net.Listener, certFile: string, keyFile: string): void
        SetKeepAlivesEnabled(v: boolean): void
        Shutdown(ctx: context.Context): void
    }
    function Transport(): TransportPointer
    interface TransportPointer extends Native {
        readonly __TransportPointer: TransportPointer
        // Proxy specifies a function to return a proxy for a given
        // Request. If the function returns a non-nil error, the
        // request is aborted with the provided error.
        //
        // The proxy type is determined by the URL scheme. "http",
        // "https", and "socks5" are supported. If the scheme is empty,
        // "http" is assumed.
        //
        // If Proxy is nil or returns a nil *URL, no proxy is used.
        Proxy: (r: RequestPointer) => [url.URLPointer, Error]

        // DialContext specifies the dial function for creating unencrypted TCP connections.
        // If DialContext is nil (and the deprecated Dial below is also nil),
        // then the transport dials using package net.
        //
        // DialContext runs concurrently with calls to RoundTrip.
        // A RoundTrip call that initiates a dial may end up using
        // a connection dialed previously when the earlier connection
        // becomes idle before the later DialContext completes.
        DialContext: (ctx: context.Context, network: string, addr: string) => [net.Conn, Error]

        // Dial specifies the dial function for creating unencrypted TCP connections.
        //
        // Dial runs concurrently with calls to RoundTrip.
        // A RoundTrip call that initiates a dial may end up using
        // a connection dialed previously when the earlier connection
        // becomes idle before the later Dial completes.
        //
        // Deprecated: Use DialContext instead, which allows the transport
        // to cancel dials as soon as they are no longer needed.
        // If both are set, DialContext takes priority.
        Dial: (network: string, addr: string) => [net.Conn, Error]

        // DialTLSContext specifies an optional dial function for creating
        // TLS connections for non-proxied HTTPS requests.
        //
        // If DialTLSContext is nil (and the deprecated DialTLS below is also nil),
        // DialContext and TLSClientConfig are used.
        //
        // If DialTLSContext is set, the Dial and DialContext hooks are not used for HTTPS
        // requests and the TLSClientConfig and TLSHandshakeTimeout
        // are ignored. The returned net.Conn is assumed to already be
        // past the TLS handshake.
        DialTLSContext: (ctx: context.Context, network: string, addr: string) => [net.Conn, Error]

        // DialTLS specifies an optional dial function for creating
        // TLS connections for non-proxied HTTPS requests.
        //
        // Deprecated: Use DialTLSContext instead, which allows the transport
        // to cancel dials as soon as they are no longer needed.
        // If both are set, DialTLSContext takes priority.
        DialTLS: (network: string, addr: string) => [net.Conn, Error]

        // TLSClientConfig specifies the TLS configuration to use with
        // tls.Client.
        // If nil, the default configuration is used.
        // If non-nil, HTTP/2 support may not be enabled by default.
        TLSClientConfig: tls.ConfigPointer

        // TLSHandshakeTimeout specifies the maximum amount of time waiting to
        // wait for a TLS handshake. Zero means no timeout.
        TLSHandshakeTimeout: time.Duration

        // DisableKeepAlives, if true, disables HTTP keep-alives and
        // will only use the connection to the server for a single
        // HTTP request.
        //
        // This is unrelated to the similarly named TCP keep-alives.
        DisableKeepAlives: boolean

        // DisableCompression, if true, prevents the Transport from
        // requesting compression with an "Accept-Encoding: gzip"
        // request header when the Request contains no existing
        // Accept-Encoding value. If the Transport requests gzip on
        // its own and gets a gzipped response, it's transparently
        // decoded in the Response.Body. However, if the user
        // explicitly requested gzip it is not automatically
        // uncompressed.
        DisableCompression: boolean

        // MaxIdleConns controls the maximum number of idle (keep-alive)
        // connections across all hosts. Zero means no limit.
        MaxIdleConns: Int

        // MaxIdleConnsPerHost, if non-zero, controls the maximum idle
        // (keep-alive) connections to keep per-host. If zero,
        // DefaultMaxIdleConnsPerHost is used.
        MaxIdleConnsPerHost: Int

        // MaxConnsPerHost optionally limits the total number of
        // connections per host, including connections in the dialing,
        // active, and idle states. On limit violation, dials will block.
        //
        // Zero means no limit.
        MaxConnsPerHost; Int

        // IdleConnTimeout is the maximum amount of time an idle
        // (keep-alive) connection will remain idle before closing
        // itself.
        // Zero means no limit.
        IdleConnTimeout: time.Duration

        // ResponseHeaderTimeout, if non-zero, specifies the amount of
        // time to wait for a server's response headers after fully
        // writing the request (including its body, if any). This
        // time does not include the time to read the response body.
        ResponseHeaderTimeout: time.Duration

        // ExpectContinueTimeout, if non-zero, specifies the amount of
        // time to wait for a server's first response headers after fully
        // writing the request headers if the request has an
        // "Expect: 100-continue" header. Zero means no timeout and
        // causes the body to be sent immediately, without
        // waiting for the server to approve.
        // This time does not include the time to send the request header.
        ExpectContinueTimeout: time.Duration

        // TLSNextProto specifies how the Transport switches to an
        // alternate protocol (such as HTTP/2) after a TLS ALPN
        // protocol negotiation. If Transport dials an TLS connection
        // with a non-empty protocol name and TLSNextProto contains a
        // map entry for that key (such as "h2"), then the func is
        // called with the request's authority (such as "example.com"
        // or "example.com:1234") and the TLS connection. The function
        // must return a RoundTripper that then handles the request.
        // If TLSNextProto is not nil, HTTP/2 support is not enabled
        // automatically.
        TLSNextProto: Map<string, (authority: string, c: tls.ConnPointer) => RoundTripper>

        // ProxyConnectHeader optionally specifies headers to send to
        // proxies during CONNECT requests.
        // To set the header dynamically, see GetProxyConnectHeader.
        ProxyConnectHeader: Header

        // GetProxyConnectHeader optionally specifies a func to return
        // headers to send to proxyURL during a CONNECT request to the
        // ip:port target.
        // If it returns an error, the Transport's RoundTrip fails with
        // that error. It can return (nil, nil) to not add headers.
        // If GetProxyConnectHeader is non-nil, ProxyConnectHeader is
        // ignored.
        GetProxyConnectHeader: (ctx: context.Context, proxyURL: url.URLPointer, target: string) => [Header, Error]

        // MaxResponseHeaderBytes specifies a limit on how many
        // response bytes are allowed in the server's response
        // header.
        //
        // Zero means to use a default limit.
        MaxResponseHeaderBytes: Int64

        // WriteBufferSize specifies the size of the write buffer used
        // when writing to the transport.
        // If zero, a default (currently 4KB) is used.
        WriteBufferSize: Int

        // ReadBufferSize specifies the size of the read buffer used
        // when reading from the transport.
        // If zero, a default (currently 4KB) is used.
        ReadBufferSize: Int

        // ForceAttemptHTTP2 controls whether HTTP/2 is enabled when a non-zero
        // Dial, DialTLS, or DialContext func or TLSClientConfig is provided.
        // By default, use of any those fields conservatively disables HTTP/2.
        // To use a custom dialer or TLS config and still attempt HTTP/2
        // upgrades, set this to true.
        ForceAttemptHTTP2: boolean
        // contains filtered or unexported fields
    }
    function isClientPointer(v: any): v is ClientPointer
    function isConnState(v: any): v is ConnState
    function isCookiePointer(v: any): v is CookiePointer
    function isDir(v: any): v is Dir
    function isFile(v: any): v is File
    function isFileSystem(v: any): v is FileSystem
    function isFlusher(v: any): v is Flusher
    function isHandler(v: any): v is Handler
    function isHandlerFunc(v: any): v is HandlerFunc
    function isHeader(v: any): v is Header
    function isHijacker(v: any): v is Hijacker
    function isPushOptionsPointer(v: any): v is PushOptionsPointer
    function isPusher(v: any): v is Pusher
    function isRequestPointer(v: any): v is RequestPointer
    function isResponsePointer(v: any): v is ResponsePointer
    function isResponseWriter(v: any): v is ResponseWriter
    function isRoundTripper(v: any): v is RoundTripper
    function isSameSite(v: any): v is SameSite
    function isServeMuxPointer(v: any): v is ServeMuxPointer
    function isServerPointer(v: any): v is ServerPointer
}