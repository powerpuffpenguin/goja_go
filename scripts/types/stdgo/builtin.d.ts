declare module "stdgo/builtin" {
    interface Native {
        readonly __Native: Native
    }
    interface Type {
        readonly __Type: Type
    }
    interface Uintptr extends Native {
        readonly __Uintptr: Uintptr
    }
    interface Number extends Native {
        readonly __Number: Number
    }
    type NumberLike = number | string | null | undefined | Number
    interface Slice<T> extends Native {
        readonly __Slice: Slice<T>
    }
    interface Map<K, V> extends Native {
        readonly __Map: Map<K, V>
    }

    interface ReadChannel<T> extends Native {
        readonly __ReadChannel: ReadChannel<T>
    }
    interface WriteChannel<T> extends Native {
        readonly __WriteChannel: WriteChannel<T>
    }
    interface Channel<T> extends ReadChannel<T>, WriteChannel<T> { }
    interface SelectCase extends Native {
        readonly __SelectCase: SelectCase
    }

    function print(...args: Array<any>): void
    function printType(...args: Array<any>): void
    function async<T>(f: T, ...args: Parameters<T>): Promise<ReturnType<T>>
    function error<T>(f: T, ...args: Parameters<T>): Error


    function append<T>(slice: Slice<T>, ...elems: Array<T>): Slice<T>
    function cap<T>(slice: Slice<T>): Int
    function len<T>(slice: Slice<T>): Int
    function copy<T>(dst: Slice<T>, src: Slice<T>): Int
    function getIndex<T>(slice: Slice<T>, index: NumberLike): T
    function setIndex<T>(slice: Slice<T>, index: NumberLike, x: T): void
    function slice<T>(slice: Slice<T>, i?: NumberLike, j?: NumberLike, k?: NumberLike): Slice<T>
    function getKey<K, V>(m: Map<K, V>, key: K): [V, boolean]
    function setKey<K, V>(m: Map<K, V>, key: K, elem: V): void
    function deleteKey<K, V>(m: Map<K, V>, key: K): void
    function typeOf(native: Native): Type
    function newType(t: Type): Native
    function makeSlice(t: Type, len?: Int | NumberLike, cap?: Int | NumberLike): Slice
    function makeMap(t: Type): Map
    function makeChan(t: Type, buffer?: Int | NumberLike): Channel
    function recv<T>(ch: ReadChannel<T>): [T, boolean]
    function tryRecv<T>(ch: ReadChannel<T>): [T, boolean]
    function send<T>(ch: WriteChannel<T>, x: T): void
    function trySend<T>(ch: WriteChannel<T>, x: T): boolean
    function selectDefault(): SelectCase
    function selectRecv<T>(ch: ReadChannel<T>): SelectCase
    function selectSend<T>(ch: WriteChannel<T>, send: T): SelectCase
    /** return (chosen int, recv interface{}, recvOK bool)  */
    function select(...cases: Array<SelectCase>): [Int, any, boolean]
    function close(ch: WriteChannel): void

    interface Error extends Native {
        readonly __Error: Error
        Error(): string
    }


    const MaxInt64: Int64
    const MaxInt32: Int32
    const MaxInt16: Int16
    const MaxInt8: Int8
    const MaxUint64: Uint64
    const MaxUint32: Uint32
    const MaxUint16: Uint16
    const MaxUint8: Uint8
    const MaxFloat64: Float64
    const MaxFloat32: Float32
    const MinInt64: Int64
    const MinInt32: Int32
    const MinInt16: Int16
    const MinInt8: Int8
    const IntSize: number

    interface Int extends Number {
        readonly __Int: Int

        String(): string
        ToNumber(): number
        ToInt(): Int
        ToInt64(): Int64
        ToInt32(): Int32
        ToInt16(): Int16
        ToInt8(): Int8
        ToUint(): Uint
        ToUint64(): Uint64
        ToUint32(): Uint32
        ToUint16(): Uint16
        ToUint8(): Uint8
        ToFloat64(): Float64
        ToFloat32(): Float32
        ABS(): Int
        Negate(): Int
        Add(...value: Array<NumberLike>): Int
        Sub(...value: Array<NumberLike>): Int
        Mul(...value: Array<NumberLike>): Int
        Div(...value: Array<NumberLike>): Int
        Mod(...value: Array<NumberLike>): Int
        And(...value: Array<NumberLike>): Int
        AndNot(...value: Array<NumberLike>): Int
        Not(): Int
        Or(...value: Array<NumberLike>): Int
        Xor(...value: Array<NumberLike>): Int
        ShiftLeft(...value: Array<NumberLike>): Int
        ShiftRight(...value: Array<NumberLike>): Int
        Compare(val: NumberLike): number
        Max(...value: Array<NumberLike>): Int
        Min(...value: Array<NumberLike>): Int
    }
    function Int(val: NumberLike): Int
    function Int(val: string, base: number | string): Int
    function isInt(v: any): v is Int

    interface Int64 extends Number {
        readonly __Int64: Int64

        String(): string
        ToNumber(): number
        ToInt(): Int
        ToInt64(): Int64
        ToInt32(): Int32
        ToInt16(): Int16
        ToInt8(): Int8
        ToUint(): Uint
        ToUint64(): Uint64
        ToUint32(): Uint32
        ToUint16(): Uint16
        ToUint8(): Uint8
        ToFloat64(): Float64
        ToFloat32(): Float32
        ABS(): Int64
        Negate(): Int64
        Add(...value: Array<NumberLike>): Int64
        Sub(...value: Array<NumberLike>): Int64
        Mul(...value: Array<NumberLike>): Int64
        Div(...value: Array<NumberLike>): Int64
        Mod(...value: Array<NumberLike>): Int64
        And(...value: Array<NumberLike>): Int64
        AndNot(...value: Array<NumberLike>): Int64
        Not(): Int64
        Or(...value: Array<NumberLike>): Int64
        Xor(...value: Array<NumberLike>): Int64
        ShiftLeft(...value: Array<NumberLike>): Int64
        ShiftRight(...value: Array<NumberLike>): Int64
        Compare(val: NumberLike): number
        Max(...value: Array<NumberLike>): Int64
        Min(...value: Array<NumberLike>): Int64
    }
    function Int64(val: NumberLike): Int64
    function Int64(val: string, base: number | string): Int64
    function isInt64(v: any): v is Int64

    interface Int32 extends Number {
        readonly __Int32: Int32

        String(): string
        ToNumber(): number
        ToInt(): Int
        ToInt64(): Int64
        ToInt32(): Int32
        ToInt16(): Int16
        ToInt8(): Int8
        ToUint(): Uint
        ToUint64(): Uint64
        ToUint32(): Uint32
        ToUint16(): Uint16
        ToUint8(): Uint8
        ToFloat64(): Float64
        ToFloat32(): Float32
        ABS(): Int32
        Negate(): Int32
        Add(...value: Array<NumberLike>): Int32
        Sub(...value: Array<NumberLike>): Int32
        Mul(...value: Array<NumberLike>): Int32
        Div(...value: Array<NumberLike>): Int32
        Mod(...value: Array<NumberLike>): Int32
        And(...value: Array<NumberLike>): Int32
        AndNot(...value: Array<NumberLike>): Int32
        Not(): Int32
        Or(...value: Array<NumberLike>): Int32
        Xor(...value: Array<NumberLike>): Int32
        ShiftLeft(...value: Array<NumberLike>): Int32
        ShiftRight(...value: Array<NumberLike>): Int32
        Compare(val: NumberLike): number
        Max(...value: Array<NumberLike>): Int32
        Min(...value: Array<NumberLike>): Int32
    }
    function Int32(val: NumberLike): Int32
    function Int32(val: string, base: number | string): Int32
    function isInt32(v: any): v is Int32

    interface Int16 extends Number {
        readonly __Int16: Int16

        String(): string
        ToNumber(): number
        ToInt(): Int
        ToInt64(): Int64
        ToInt32(): Int32
        ToInt16(): Int16
        ToInt8(): Int8
        ToUint(): Uint
        ToUint64(): Uint64
        ToUint32(): Uint32
        ToUint16(): Uint16
        ToUint8(): Uint8
        ToFloat64(): Float64
        ToFloat32(): Float32
        ABS(): Int16
        Negate(): Int16
        Add(...value: Array<NumberLike>): Int16
        Sub(...value: Array<NumberLike>): Int16
        Mul(...value: Array<NumberLike>): Int16
        Div(...value: Array<NumberLike>): Int16
        Mod(...value: Array<NumberLike>): Int16
        And(...value: Array<NumberLike>): Int16
        AndNot(...value: Array<NumberLike>): Int16
        Not(): Int16
        Or(...value: Array<NumberLike>): Int16
        Xor(...value: Array<NumberLike>): Int16
        ShiftLeft(...value: Array<NumberLike>): Int16
        ShiftRight(...value: Array<NumberLike>): Int16
        Compare(val: NumberLike): number
        Max(...value: Array<NumberLike>): Int16
        Min(...value: Array<NumberLike>): Int16
    }
    function Int16(val: NumberLike): Int16
    function Int16(val: string, base: number | string): Int16
    function isInt16(v: any): v is Int16

    interface Int8 extends Number {
        readonly __Int8: Int8

        String(): string
        ToNumber(): number
        ToInt(): Int
        ToInt64(): Int64
        ToInt32(): Int32
        ToInt16(): Int16
        ToInt8(): Int8
        ToUint(): Uint
        ToUint64(): Uint64
        ToUint32(): Uint32
        ToUint16(): Uint16
        ToUint8(): Uint8
        ToFloat64(): Float64
        ToFloat32(): Float32
        ABS(): Int8
        Negate(): Int8
        Add(...value: Array<NumberLike>): Int8
        Sub(...value: Array<NumberLike>): Int8
        Mul(...value: Array<NumberLike>): Int8
        Div(...value: Array<NumberLike>): Int8
        Mod(...value: Array<NumberLike>): Int8
        And(...value: Array<NumberLike>): Int8
        AndNot(...value: Array<NumberLike>): Int8
        Not(): Int8
        Or(...value: Array<NumberLike>): Int8
        Xor(...value: Array<NumberLike>): Int8
        ShiftLeft(...value: Array<NumberLike>): Int8
        ShiftRight(...value: Array<NumberLike>): Int8
        Compare(val: NumberLike): number
        Max(...value: Array<NumberLike>): Int8
        Min(...value: Array<NumberLike>): Int8
    }
    function Int8(val: NumberLike): Int8
    function Int8(val: string, base: number | string): Int8
    function isInt8(v: any): v is Int8

    interface Uint extends Number {
        readonly __Uint: Uint

        String(): string
        ToNumber(): number
        ToInt(): Int
        ToInt64(): Int64
        ToInt32(): Int32
        ToInt16(): Int16
        ToInt8(): Int8
        ToUint(): Uint
        ToUint64(): Uint64
        ToUint32(): Uint32
        ToUint16(): Uint16
        ToUint8(): Uint8
        ToFloat64(): Float64
        ToFloat32(): Float32
        Add(...value: Array<NumberLike>): Uint
        Sub(...value: Array<NumberLike>): Uint
        Mul(...value: Array<NumberLike>): Uint
        Div(...value: Array<NumberLike>): Uint
        Mod(...value: Array<NumberLike>): Uint
        And(...value: Array<NumberLike>): Uint
        AndNot(...value: Array<NumberLike>): Uint
        Not(): Uint
        Or(...value: Array<NumberLike>): Uint
        Xor(...value: Array<NumberLike>): Uint
        ShiftLeft(...value: Array<NumberLike>): Uint
        ShiftRight(...value: Array<NumberLike>): Uint
        Compare(val: NumberLike): number
        Max(...value: Array<NumberLike>): Uint
        Min(...value: Array<NumberLike>): Uint
    }
    function Uint(val: NumberLike): Uint
    function Uint(val: string, base: number | string): Uint
    function isUint(v: any): v is Uint

    interface Uint64 extends Number {
        readonly __Uint64: Uint64

        String(): string
        ToNumber(): number
        ToInt(): Int
        ToInt64(): Int64
        ToInt32(): Int32
        ToInt16(): Int16
        ToInt8(): Int8
        ToUint(): Uint
        ToUint64(): Uint64
        ToUint32(): Uint32
        ToUint16(): Uint16
        ToUint8(): Uint8
        ToFloat64(): Float64
        ToFloat32(): Float32
        Add(...value: Array<NumberLike>): Uint64
        Sub(...value: Array<NumberLike>): Uint64
        Mul(...value: Array<NumberLike>): Uint64
        Div(...value: Array<NumberLike>): Uint64
        Mod(...value: Array<NumberLike>): Uint64
        And(...value: Array<NumberLike>): Uint64
        AndNot(...value: Array<NumberLike>): Uint64
        Not(): Uint64
        Or(...value: Array<NumberLike>): Uint64
        Xor(...value: Array<NumberLike>): Uint64
        ShiftLeft(...value: Array<NumberLike>): Uint64
        ShiftRight(...value: Array<NumberLike>): Uint64
        Compare(val: NumberLike): number
        Max(...value: Array<NumberLike>): Uint64
        Min(...value: Array<NumberLike>): Uint64
    }
    function Uint64(val: NumberLike): Uint64
    function Uint64(val: string, base: number | string): Uint64
    function isUint64(v: any): v is Uint64

    interface Uint32 extends Number {
        readonly __Uint32: Uint32

        String(): string
        ToNumber(): number
        ToInt(): Int
        ToInt64(): Int64
        ToInt32(): Int32
        ToInt16(): Int16
        ToInt8(): Int8
        ToUint(): Uint
        ToUint64(): Uint64
        ToUint32(): Uint32
        ToUint16(): Uint16
        ToUint8(): Uint8
        ToFloat64(): Float64
        ToFloat32(): Float32
        Add(...value: Array<NumberLike>): Uint32
        Sub(...value: Array<NumberLike>): Uint32
        Mul(...value: Array<NumberLike>): Uint32
        Div(...value: Array<NumberLike>): Uint32
        Mod(...value: Array<NumberLike>): Uint32
        And(...value: Array<NumberLike>): Uint32
        AndNot(...value: Array<NumberLike>): Uint32
        Not(): Uint32
        Or(...value: Array<NumberLike>): Uint32
        Xor(...value: Array<NumberLike>): Uint32
        ShiftLeft(...value: Array<NumberLike>): Uint32
        ShiftRight(...value: Array<NumberLike>): Uint32
        Compare(val: NumberLike): number
        Max(...value: Array<NumberLike>): Uint32
        Min(...value: Array<NumberLike>): Uint32
    }
    function Uint32(val: NumberLike): Uint32
    function Uint32(val: string, base: number | string): Uint32
    function isUint32(v: any): v is Uint32

    interface Uint16 extends Number {
        readonly __Uint16: Uint16

        String(): string
        ToNumber(): number
        ToInt(): Int
        ToInt64(): Int64
        ToInt32(): Int32
        ToInt16(): Int16
        ToInt8(): Int8
        ToUint(): Uint
        ToUint64(): Uint64
        ToUint32(): Uint32
        ToUint16(): Uint16
        ToUint8(): Uint8
        ToFloat64(): Float64
        ToFloat32(): Float32
        Add(...value: Array<NumberLike>): Uint16
        Sub(...value: Array<NumberLike>): Uint16
        Mul(...value: Array<NumberLike>): Uint16
        Div(...value: Array<NumberLike>): Uint16
        Mod(...value: Array<NumberLike>): Uint16
        And(...value: Array<NumberLike>): Uint16
        AndNot(...value: Array<NumberLike>): Uint16
        Not(): Uint16
        Or(...value: Array<NumberLike>): Uint16
        Xor(...value: Array<NumberLike>): Uint16
        ShiftLeft(...value: Array<NumberLike>): Uint16
        ShiftRight(...value: Array<NumberLike>): Uint16
        Compare(val: NumberLike): number
        Max(...value: Array<NumberLike>): Uint16
        Min(...value: Array<NumberLike>): Uint16
    }
    function Uint16(val: NumberLike): Uint16
    function Uint16(val: string, base: number | string): Uint16
    function isUint16(v: any): v is Uint16

    interface Uint8 extends Number {
        readonly __Uint8: Uint8

        String(): string
        ToNumber(): number
        ToInt(): Int
        ToInt64(): Int64
        ToInt32(): Int32
        ToInt16(): Int16
        ToInt8(): Int8
        ToUint(): Uint
        ToUint64(): Uint64
        ToUint32(): Uint32
        ToUint16(): Uint16
        ToUint8(): Uint8
        ToFloat64(): Float64
        ToFloat32(): Float32
        Add(...value: Array<NumberLike>): Uint8
        Sub(...value: Array<NumberLike>): Uint8
        Mul(...value: Array<NumberLike>): Uint8
        Div(...value: Array<NumberLike>): Uint8
        Mod(...value: Array<NumberLike>): Uint8
        And(...value: Array<NumberLike>): Uint8
        AndNot(...value: Array<NumberLike>): Uint8
        Not(): Uint8
        Or(...value: Array<NumberLike>): Uint8
        Xor(...value: Array<NumberLike>): Uint8
        ShiftLeft(...value: Array<NumberLike>): Uint8
        ShiftRight(...value: Array<NumberLike>): Uint8
        Compare(val: NumberLike): number
        Max(...value: Array<NumberLike>): Uint8
        Min(...value: Array<NumberLike>): Uint8
    }
    function Uint8(val: NumberLike): Uint8
    function Uint8(val: string, base: number | string): Uint8
    function isUint8(v: any): v is Uint8

    interface Float64 extends Number {
        readonly __Float64: Float64

        String(): string
        ToNumber(): number
        ToInt(): Int
        ToInt64(): Int64
        ToInt32(): Int32
        ToInt16(): Int16
        ToInt8(): Int8
        ToUint(): Uint
        ToUint64(): Uint64
        ToUint32(): Uint32
        ToUint16(): Uint16
        ToUint8(): Uint8
        ToFloat64(): Float64
        ToFloat32(): Float32
        ABS(): Float64
        Negate(): Float64
        Add(...value: Array<NumberLike>): Float64
        Sub(...value: Array<NumberLike>): Float64
        Mul(...value: Array<NumberLike>): Float64
        Div(...value: Array<NumberLike>): Float64
        Sqrt(): Float64
        Compare(val: NumberLike): number
        Max(...value: Array<NumberLike>): Float64
        Min(...value: Array<NumberLike>): Float64
    }
    function Float64(val: NumberLike): Float64
    function Float64(val: string): Float64
    function isFloat64(v: any): v is Float64

    interface Float32 extends Number {
        readonly __Float32: Float32

        String(): string
        ToNumber(): number
        ToInt(): Int
        ToInt64(): Int64
        ToInt32(): Int32
        ToInt16(): Int16
        ToInt8(): Int8
        ToUint(): Uint
        ToUint64(): Uint64
        ToUint32(): Uint32
        ToUint16(): Uint16
        ToUint8(): Uint8
        ToFloat64(): Float64
        ToFloat32(): Float32
        ABS(): Float32
        Negate(): Float32
        Add(...value: Array<NumberLike>): Float32
        Sub(...value: Array<NumberLike>): Float32
        Mul(...value: Array<NumberLike>): Float32
        Div(...value: Array<NumberLike>): Float32
        Sqrt(): Float32
        Compare(val: NumberLike): number
        Max(...value: Array<NumberLike>): Float32
        Min(...value: Array<NumberLike>): Float32
    }
    function Float32(val: NumberLike): Float32
    function Float32(val: string): Float32
    function isFloat32(v: any): v is Float32

    interface IntSlice extends Slice<Int> {
        readonly __IntSlice: IntSlice
        String(): string
        Len(): Int
        Swap(i: NumberLike, j: NumberLike)
        Less(i: NumberLike, j: NumberLike): boolean
        Cap(): Int
        Copy(src: IntSlice): Int
        Slice(start: NumberLike): IntSlice
        Slice2(start: NumberLike, end: NumberLike): IntSlice
        Append(...data: Array<NumberLike>): IntSlice
        Get(index: NumberLike): Int
        Set(index: NumberLike, val: NumberLike)
        Join(sep: string): string
        Asc()
        Desc()
    }
    function IntSlice(): IntSlice
    function IntSlice(len: NumberLike): IntSlice
    function IntSlice(len: NumberLike, cap: NumberLike): IntSlice
    function isIntSlice(v: any): v is IntSlice
    interface Int64Slice extends Slice<Int64> {
        readonly __Int64Slice: Int64Slice
        String(): string
        Len(): Int
        Swap(i: NumberLike, j: NumberLike)
        Less(i: NumberLike, j: NumberLike): boolean
        Cap(): Int
        Copy(src: Int64Slice): Int
        Slice(start: NumberLike): Int64Slice
        Slice2(start: NumberLike, end: NumberLike): Int64Slice
        Append(...data: Array<NumberLike>): Int64Slice
        Get(index: NumberLike): Int64
        Set(index: NumberLike, val: NumberLike)
        Join(sep: string): string
        Asc()
        Desc()
    }
    function Int64Slice(): Int64Slice
    function Int64Slice(len: NumberLike): Int64Slice
    function Int64Slice(len: NumberLike, cap: NumberLike): Int64Slice
    function isInt64Slice(v: any): v is Int64Slice
    interface Int32Slice extends Slice<Int32> {
        readonly __Int32Slice: Int32Slice
        String(): string
        Len(): Int
        Swap(i: NumberLike, j: NumberLike)
        Less(i: NumberLike, j: NumberLike): boolean
        Cap(): Int
        Copy(src: Int32Slice): Int
        Slice(start: NumberLike): Int32Slice
        Slice2(start: NumberLike, end: NumberLike): Int32Slice
        Append(...data: Array<NumberLike>): Int32Slice
        Get(index: NumberLike): Int32
        Set(index: NumberLike, val: NumberLike)
        Join(sep: string): string
        Asc()
        Desc()
    }
    function Int32Slice(): Int32Slice
    function Int32Slice(len: NumberLike): Int32Slice
    function Int32Slice(len: NumberLike, cap: NumberLike): Int32Slice
    function isInt32Slice(v: any): v is Int32Slice
    interface Int16Slice extends Slice<Int16> {
        readonly __Int16Slice: Int16Slice
        String(): string
        Len(): Int
        Swap(i: NumberLike, j: NumberLike)
        Less(i: NumberLike, j: NumberLike): boolean
        Cap(): Int
        Copy(src: Int16Slice): Int
        Slice(start: NumberLike): Int16Slice
        Slice2(start: NumberLike, end: NumberLike): Int16Slice
        Append(...data: Array<NumberLike>): Int16Slice
        Get(index: NumberLike): Int16
        Set(index: NumberLike, val: NumberLike)
        Join(sep: string): string
        Asc()
        Desc()
    }
    function Int16Slice(): Int16Slice
    function Int16Slice(len: NumberLike): Int16Slice
    function Int16Slice(len: NumberLike, cap: NumberLike): Int16Slice
    function isInt16Slice(v: any): v is Int16Slice
    interface Int8Slice extends Slice<Int8> {
        readonly __Int8Slice: Int8Slice
        String(): string
        Len(): Int
        Swap(i: NumberLike, j: NumberLike)
        Less(i: NumberLike, j: NumberLike): boolean
        Cap(): Int
        Copy(src: Int8Slice): Int
        Slice(start: NumberLike): Int8Slice
        Slice2(start: NumberLike, end: NumberLike): Int8Slice
        Append(...data: Array<NumberLike>): Int8Slice
        Get(index: NumberLike): Int8
        Set(index: NumberLike, val: NumberLike)
        Join(sep: string): string
        Asc()
        Desc()
    }
    function Int8Slice(): Int8Slice
    function Int8Slice(len: NumberLike): Int8Slice
    function Int8Slice(len: NumberLike, cap: NumberLike): Int8Slice
    function isInt8Slice(v: any): v is Int8Slice
    interface UintSlice extends Slice<Uint> {
        readonly __UintSlice: UintSlice
        String(): string
        Len(): Int
        Swap(i: NumberLike, j: NumberLike)
        Less(i: NumberLike, j: NumberLike): boolean
        Cap(): Int
        Copy(src: UintSlice): Int
        Slice(start: NumberLike): UintSlice
        Slice2(start: NumberLike, end: NumberLike): UintSlice
        Append(...data: Array<NumberLike>): UintSlice
        Get(index: NumberLike): Uint
        Set(index: NumberLike, val: NumberLike)
        Join(sep: string): string
        Asc()
        Desc()
    }
    function UintSlice(): UintSlice
    function UintSlice(len: NumberLike): UintSlice
    function UintSlice(len: NumberLike, cap: NumberLike): UintSlice
    function isUintSlice(v: any): v is UintSlice
    interface Uint64Slice extends Slice<Uint64> {
        readonly __Uint64Slice: Uint64Slice
        String(): string
        Len(): Int
        Swap(i: NumberLike, j: NumberLike)
        Less(i: NumberLike, j: NumberLike): boolean
        Cap(): Int
        Copy(src: Uint64Slice): Int
        Slice(start: NumberLike): Uint64Slice
        Slice2(start: NumberLike, end: NumberLike): Uint64Slice
        Append(...data: Array<NumberLike>): Uint64Slice
        Get(index: NumberLike): Uint64
        Set(index: NumberLike, val: NumberLike)
        Join(sep: string): string
        Asc()
        Desc()
    }
    function Uint64Slice(): Uint64Slice
    function Uint64Slice(len: NumberLike): Uint64Slice
    function Uint64Slice(len: NumberLike, cap: NumberLike): Uint64Slice
    function isUint64Slice(v: any): v is Uint64Slice
    interface Uint32Slice extends Slice<Uint32> {
        readonly __Uint32Slice: Uint32Slice
        String(): string
        Len(): Int
        Swap(i: NumberLike, j: NumberLike)
        Less(i: NumberLike, j: NumberLike): boolean
        Cap(): Int
        Copy(src: Uint32Slice): Int
        Slice(start: NumberLike): Uint32Slice
        Slice2(start: NumberLike, end: NumberLike): Uint32Slice
        Append(...data: Array<NumberLike>): Uint32Slice
        Get(index: NumberLike): Uint32
        Set(index: NumberLike, val: NumberLike)
        Join(sep: string): string
        Asc()
        Desc()
    }
    function Uint32Slice(): Uint32Slice
    function Uint32Slice(len: NumberLike): Uint32Slice
    function Uint32Slice(len: NumberLike, cap: NumberLike): Uint32Slice
    function isUint32Slice(v: any): v is Uint32Slice
    interface Uint16Slice extends Slice<Uint16> {
        readonly __Uint16Slice: Uint16Slice
        String(): string
        Len(): Int
        Swap(i: NumberLike, j: NumberLike)
        Less(i: NumberLike, j: NumberLike): boolean
        Cap(): Int
        Copy(src: Uint16Slice): Int
        Slice(start: NumberLike): Uint16Slice
        Slice2(start: NumberLike, end: NumberLike): Uint16Slice
        Append(...data: Array<NumberLike>): Uint16Slice
        Get(index: NumberLike): Uint16
        Set(index: NumberLike, val: NumberLike)
        Join(sep: string): string
        Asc()
        Desc()
    }
    function Uint16Slice(): Uint16Slice
    function Uint16Slice(len: NumberLike): Uint16Slice
    function Uint16Slice(len: NumberLike, cap: NumberLike): Uint16Slice
    function isUint16Slice(v: any): v is Uint16Slice
    interface Uint8Slice extends Slice<Uint8> {
        readonly __Uint8Slice: Uint8Slice
        String(): string
        Len(): Int
        Swap(i: NumberLike, j: NumberLike)
        Less(i: NumberLike, j: NumberLike): boolean
        Cap(): Int
        Copy(src: Uint8Slice): Int
        Slice(start: NumberLike): Uint8Slice
        Slice2(start: NumberLike, end: NumberLike): Uint8Slice
        Append(...data: Array<NumberLike>): Uint8Slice
        Get(index: NumberLike): Uint8
        Set(index: NumberLike, val: NumberLike)
        Join(sep: string): string
        Asc()
        Desc()
    }
    function Uint8Slice(): Uint8Slice
    function Uint8Slice(len: NumberLike): Uint8Slice
    function Uint8Slice(len: NumberLike, cap: NumberLike): Uint8Slice
    function isUint8Slice(v: any): v is Uint8Slice
    interface Float64Slice extends Slice<Float64> {
        readonly __Float64Slice: Float64Slice
        String(): string
        Len(): Int
        Swap(i: NumberLike, j: NumberLike)
        Less(i: NumberLike, j: NumberLike): boolean
        Cap(): Int
        Copy(src: Float64Slice): Int
        Slice(start: NumberLike): Float64Slice
        Slice2(start: NumberLike, end: NumberLike): Float64Slice
        Append(...data: Array<NumberLike>): Float64Slice
        Get(index: NumberLike): Float64
        Set(index: NumberLike, val: NumberLike)
        Join(sep: string): string
        Asc()
        Desc()
    }
    function Float64Slice(): Float64Slice
    function Float64Slice(len: NumberLike): Float64Slice
    function Float64Slice(len: NumberLike, cap: NumberLike): Float64Slice
    function isFloat64Slice(v: any): v is Float64Slice
    interface Float32Slice extends Slice<Float32> {
        readonly __Float32Slice: Float32Slice
        String(): string
        Len(): Int
        Swap(i: NumberLike, j: NumberLike)
        Less(i: NumberLike, j: NumberLike): boolean
        Cap(): Int
        Copy(src: Float32Slice): Int
        Slice(start: NumberLike): Float32Slice
        Slice2(start: NumberLike, end: NumberLike): Float32Slice
        Append(...data: Array<NumberLike>): Float32Slice
        Get(index: NumberLike): Float32
        Set(index: NumberLike, val: NumberLike)
        Join(sep: string): string
        Asc()
        Desc()
    }
    function Float32Slice(): Float32Slice
    function Float32Slice(len: NumberLike): Float32Slice
    function Float32Slice(len: NumberLike, cap: NumberLike): Float32Slice
    function isFloat32Slice(v: any): v is Float32Slice

    type Rune = Int32
    type Runes = Int32Slice
    type Byte = Uint8
    type Bytes = Uint8Slice
}