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
