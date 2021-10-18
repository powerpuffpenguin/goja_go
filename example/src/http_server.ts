import * as builtin from "stdgo/builtin";
import * as net from "stdgo/net";
import * as http from "stdgo/net/http";
import * as bytes from "stdgo/bytes";
import * as fmt from "stdgo/fmt";
declare const __filename: string
async function main() {
    const addr = "127.0.0.1:10000"
    const mux = http.NewServeMux()
    mux.Handle(`/`, http.Handler(home))
    mux.Handle(`/json`, http.Handler(json))
    mux.Handle(`/source`, http.Handler(source))

    const l = await builtin.async(net.Listen, `tcp`, addr)
    fmt.Println(`work at`, addr)
    await builtin.async(http.Serve, l, mux)
}
function home(w: http.ResponseWriter, r: http.RequestPointer): Promise<builtin.Int> {
    w.Header().Set(`Content-Type`, `text/plain; charset=utf-8`)
    return builtin.async(w.Write, bytes.Bytes(`cerberus is an idea`))
}
function json(w: http.ResponseWriter, r: http.RequestPointer): Promise<builtin.Int> {
    w.Header().Set(`Content-Type`, `application/json; charset=utf-8`)
    return builtin.async(w.Write, bytes.Bytes(`{"name":"cerberus","lv":1}`))
}
function source(w: http.ResponseWriter, r: http.RequestPointer): Promise<void> {
    return builtin.async(http.ServeFile, w, r, __filename)
}
main().catch((e) => {
    console.log("listen err->", e)
})