declare module "stdgo/crypto/elliptic" {
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
    import * as big from "stdgo/math/big";
    import * as io from "stdgo/io";

    /** return  (priv []byte, x, y *big.Int, err error) */
    function GenerateKey(curve: Curve, rand: io.Reader): [Bytes, big.IntPointer, big.IntPointer]
    function Marshal(curve: Curve, x: big.IntPointer, y: big.IntPointer): Bytes
    function MarshalCompressed(curve: Curve, x: big.IntPointer, y: big.IntPointer): Bytess
    /** return (x, y *big.Int) */
    function Unmarshal(curve: Curve, data: Bytes): [big.IntPointer, big.IntPointer]
    /** return (x, y *big.Int) */
    function UnmarshalCompressed(curve: Curve, data: Bytes): [big.IntPointer, big.IntPointer]

    function P224(): Curve
    function P256(): Curve
    function P384(): Curve
    function P521(): Curve
    interface Curve extends Native {
        // Params returns the parameters for the curve.
        Params(): CurveParamsPointer
        // IsOnCurve reports whether the given (x,y) lies on the curve.
        IsOnCurve(x: big.IntPointer, y: big.IntPointer): boolean
        // Add returns the sum of (x1,y1) and (x2,y2)
        Add(x1: big.IntPointer, y1: big.IntPointer, x2: big.IntPointer, y2: big.IntPointer): [big.IntPointer, big.IntPointer]
        // Double returns 2*(x,y)
        Double(x1: big.IntPointer, y1: big.IntPointer): [big.IntPointer, big.IntPointer]
        // ScalarMult returns k*(Bx,By) where k is a number in big-endian form.
        ScalarMult(x1: big.IntPointer, y1: big.IntPointer, k: Bytes): [big.IntPointer, big.IntPointer]
        // ScalarBaseMult returns k*G, where G is the base point of the group
        // and k is an integer in big-endian form.
        ScalarBaseMult(k: Bytes): [big.IntPointer, big.IntPointer]
    }

    interface CurveParamsPointer extends Native {
        readonly __CurveParamsPointer: CurveParamsPointer

        P: big.IntPointer // the order of the underlying field
        N: big.IntPointer // the order of the base point
        B: big.IntPointer // the constant of the curve equation

        Gx: big.IntPointer
        Gy: big.IntPointer // (x,y) of the base point

        BitSize: Int      // the size of the underlying field
        Name: string   // the canonical name of the curve

        Add(x1: big.IntPointer, y1: big.IntPointer, x2: big.IntPointer, y2: big.IntPointer): [big.IntPointer, big.IntPointer]
        Double(x1: big.IntPointer, y1: big.IntPointer): [big.IntPointer, big.IntPointer]
        IsOnCurve(x: big.IntPointer, y: big.IntPointer): boolean
        Params(): CurveParamsPointer
        ScalarBaseMult(k: Bytes): [big.IntPointer, big.IntPointer]
        ScalarMult(Bx: big.IntPointer, By: big.IntPointer, k: Bytes): [big.IntPointer, big.IntPointer]
    }
    function isCurve(v: any): v is Curve
    function isCurveParamsPointer(v: any): v is CurveParamsPointer
}