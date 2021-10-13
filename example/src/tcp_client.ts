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