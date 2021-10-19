# goja_go

some js extensions of github.com/powerpuffpenguin/goja

goja_go imports most of the golang standard library into goja for js call

You can compile main.go to use goja_go as an independent js engine, or you can refer to the writing in main.go to embed goja_go into your own program

# embed goja_go

Please refer to main.go for the complete code

1. Create goja.Runtime and set the Option provided by goja_go

   ```
   runtime := goja.New(
      goja.WithScheduler(loop.NewScheduler(32)),
      goja.WithFieldGetter(builtin.FieldGetter),
      goja.WithCallerFactory(builtin.NewCallerFactory()),
   )
   ```

2. Enable require

   ```
   // enable require
   registry := require.NewRegistry(
      require.WithGlobalFolders(`node_modules`),
      require.WithLoader(core.Loader),
   )
   registry.Enable(runtime)
   ```

3. Enable stdgo and register the golang standard library

   ```
   stdgo.Enable(runtime)
   stdgo.RegisterNativeModuleToRegistry(registry)
   ```

   The RegisterNativeModuleToRegistry function will inject all the implemented libraries into the js environment. You can also register the libraries under the stdgo folder on-demand instead of calling RegisterNativeModuleToRegistry

4. Call Run and Serve run scripts and event loops

   ```
   runtime.RunScriptAndServe(filename, string(source))
   ```

# stdgo

The following js environment supports the go standard library:

* archive
   * tar
   * zip
* bufio
* bytes
* compress
   * bzip2
   * flate
   * gzip
   * lzw
   * zlib
* ~~container~~
* context
* crypto
   * aes
   * cipher
   * des
   * dsa
   * ecdsa
   * ed25519
   * elliptic
   * hmac
   * md5
   * rand
   * rc4
   * rsa
   * sha1
   * sha256
   * sha512
   * subtle
   * tls
   * x509
      * pkix
* database
   * sql
      * driver
* ~~debug~~
* ~~embed~~
* encoding
   * ascii85
   * asn1
   * base32
   * base64
   * binary
   * csv
   * gob
   * hex
   * json
   * pem
   * xml
* errors
* ~~expvar~~
* ~~flag~~
* fmt
* ~~go~~
* hash
   * adler32
   * crc32
   * crc64
   * fnv
   * maphash
* html
   * template
* image
   * color
      * palette
   * draw
   * gif
   * jpeg
   * png
* index
   * suffixarray
* io
   * fs
   * ioutil
* ~~log~~
* math
   * big
   * bits
   * cmplx
   * rand
* mime
   * multipart
   * quotedprintable
* net
   * http
   * mail
   * smtp
   * textproto
   * url
* os
   * exec
   * signal
* path
   * filepath
* ~~plugin~~
* ~~reflect~~
* regexp
* ~~runtime~~
* sort
* strconv
* strings
* ~~sync~~
* ~~syscall~~
* ~~testing~~
* text
   * template
      * parse
* time
* ~~unicode~~
* ~~unsafe~~

1. The typescript declaration of the js function is located in **scripts/types/stdgo**
2. The module name in js is the prefix of stdgo + ${go std module name}

# Example 

Example of calling the go standard library in js

more [example](https://github.com/powerpuffpenguin/goja_go/tree/main/example/src)

```
// This is an example of an echo server
import * as builtin from "stdgo/builtin";
import * as net from "stdgo/net";
import * as io from "stdgo/io";
import * as time from "stdgo/time";
import * as fmt from "stdgo/fmt";

async function main() {
    const addr = "127.0.0.1:10000"

    const l = await builtin.async(net.Listen, 'tcp', addr)
    console.log("listen on->", addr)
    let tempDelay = 0
    while (true) {
        try {
            const c = await builtin.async(l.Accept)

            const remote = c.RemoteAddr().String()
            console.log("one in:", remote)
            builtin.async(io.Copy, c, c).catch((e) => {
                console.log("io.Copy err->", e)
            }).finally(() => {
                console.log("one out:", remote)
                c.Close()
            })
        } catch (e) {
            if (e instanceof GoError && net.isError(e.value)) {
                const ne = e.value
                if (ne.Temporary()) {
                    if (tempDelay == 0) {
                        tempDelay = 5
                    } else {
                        tempDelay *= 2
                        if (tempDelay > 1000) {
                            tempDelay = 1000
                        }
                    }

                    const duration = time.Duration(builtin.Int64(time.Millisecond).Mul(tempDelay))
                    fmt.Printf("Accept error: %v; retrying in %v\n", ne, tempDelay)
                    await builtin.async(time.Sleep, duration)
                    continue
                }
            }
            console.log("accept err->", e)
            return
        }
    }
}
main().catch((e) => {
    console.log("listen err->", e)
})
```

```
// This example connects to the echo server and disconnects after sending and receiving 5 pieces of data
import * as builtin from "stdgo/builtin";
import * as net from "stdgo/net";
import * as io from "stdgo/io";
import * as fmt from "stdgo/fmt";
import * as bytes from "stdgo/bytes";
import * as time from "stdgo/time";

async function main(c: net.Conn) {
    const data = builtin.Uint8Slice(1024)
    for (let i = 0; i < 5; i++) {
        if (i != 0) {
            await builtin.async(time.Sleep, time.Second)
        }

        const msg = fmt.Sprintf("message %d", i)
        const b = bytes.Bytes(msg)
        await builtin.async(c.Write, b)
        console.log("send->", msg)

        const read = data.Slice2(0, msg.length)
        await builtin.async(io.ReadAtLeast, c, read, read.Len())
        console.log("recv->", bytes.String(read))
    }
}
const addr = "127.0.0.1:10000"
builtin.async(net.Dial, 'tcp', addr).then((c) => {
    main(c).catch((e) => {
        console.log("msg err->", e)
    }).finally(() => {
        c.Close()
    })
}).catch((e) => {
    console.log("dial err->", e)
})
```