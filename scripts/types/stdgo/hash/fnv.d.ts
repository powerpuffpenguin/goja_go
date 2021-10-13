declare module "stdgo/hash/fnv" {
    import * as hash from "stdgo/hash";

    function New128(): hash.Hash
    function New128a(): hash.Hash
    function New32(): hash.Hash32
    function New32a(): hash.Hash32
    function New64(): hash.Hash64
    function New64a(): hash.Hash64
}