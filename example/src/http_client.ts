import * as builtin from "stdgo/builtin";
import * as http from "stdgo/net/http";
import * as io from "stdgo/io";
import * as fmt from "stdgo/fmt";
import * as bytes from "stdgo/bytes";
async function main() {
    const addr = "127.0.0.1:10000"
    const resp = await builtin.async(http.Get, `http://${addr}/json`)
    fmt.Println(`resp`, resp.StatusCode, resp.Status)
    const body = await builtin.async(io.ReadAll, resp.Body)
    fmt.Printf("body: %s\n", bytes.String(body))
}
main().catch((e) => {
    console.log("dial err->", e)
})