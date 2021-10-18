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
    import * as bufio from "stdgo/bufio";
    import * as io from "stdgo/io";
    import * as fs from "stdgo/io/fs";
    import * as url from "stdgo/net/url";
    import * as tls from "stdgo/crypto/tls";
    import * as multipart from "stdgo/mime/multipart"

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

    interface ServerPointer extends Native {
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
        TLSNextProto: Map<string, (s: ServerPointer, conn: tls.ConnPoint, h: Handler) => void>

            // ConnState specifies an optional callback function that is
            // called when a client connection changes state. See the
            // ConnState type and associated constants for details.
            ConnState func(net.Conn, ConnState)

    // ErrorLog specifies an optional logger for errors accepting
    // connections, unexpected behavior from handlers, and
    // underlying FileSystem errors.
    // If nil, logging is done via the log package's standard logger.
    ErrorLog * log.Logger

        // BaseContext optionally specifies a function that returns
        // the base context for incoming requests on this server.
        // The provided Listener is the specific Listener that's
        // about to start accepting requests.
        // If BaseContext is nil, the default is context.Background().
        // If non-nil, it must return a non-nil context.
        BaseContext func(net.Listener) context.Context

        // ConnContext optionally specifies a function that modifies
        // the context used for a new connection c. The provided ctx
        // is derived from the base context and has a ServerContextKey
        // value.
        ConnContext func(ctx context.Context, c net.Conn) context.Context
    // contains filtered or unexported fields
}
    // ClientPointer
    // ConnState
    // CookiePointer
    // Dir
    // File
    // FileSystem
    // Flusher
    // Handler
    // HandlerFunc
    // Header
    // Hijacker
    // PushOptionsPointer
    // Pusher
    // RequestPointer
    // ResponsePointer
    // ResponseWriter
    // RoundTripper
    // SameSite
    // ServeMuxPointer
    // ServerPointer
}