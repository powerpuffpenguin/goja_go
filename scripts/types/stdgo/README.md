# stdgo

[goja_go](https://github.com/powerpuffpenguin/goja_go) provides a js runtime environment implemented by golang and imports most of the golang standard library packages into js. This project is an api statement provided by goja_go

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
