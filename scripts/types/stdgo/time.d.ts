declare module "stdgo/time" {
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

    const Layout = "01/02 03:04:05PM '06 -0700" // The reference time, in numerical order.
    const ANSIC = "Mon Jan _2 15:04:05 2006"
    const UnixDate = "Mon Jan _2 15:04:05 MST 2006"
    const RubyDate = "Mon Jan 02 15:04:05 -0700 2006"
    const RFC822 = "02 Jan 06 15:04 MST"
    const RFC822Z = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
    const RFC850 = "Monday, 02-Jan-06 15:04:05 MST"
    const RFC1123 = "Mon, 02 Jan 2006 15:04:05 MST"
    const RFC1123Z = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
    const RFC3339 = "2006-01-02T15:04:05Z07:00"
    const RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
    const Kitchen = "3:04PM"
    // Handy time stamps.
    const Stamp = "Jan _2 15:04:05"
    const StampMilli = "Jan _2 15:04:05.000"
    const StampMicro = "Jan _2 15:04:05.000000"
    const StampNano = "Jan _2 15:04:05.000000000"

    const Nanosecond: Duration
    const Microsecond: Duration
    const Millisecond: Duration
    const Second: Duration
    const Minute: Duration
    const Hour: Duration

    function After(d: Duration | NumberLike): ReadChannel<Time>
    function Sleep(d: Duration | NumberLike): void
    function Tick(d: Duration | NumberLike): ReadChannel<Time> | null

    function Duration(val: Int64 | NumberLike): Duration
    function AddDuration(val: Duration | NumberLike, ...vals: Array<Duration | NumberLike>): Duration
    function SubDuration(val: Duration | NumberLike, ...vals: Array<Duration | NumberLike>): Duration
    function ParseDuration(s: string): Duration
    function Since(t: Time): Duration
    function Until(t: Time): Duration
    interface Duration extends Number {
        readonly __Duration: Duration

        String(): string
        Hours(): Float64
        Microseconds(): Int64
        Milliseconds(): Int64
        Minutes(): Float64
        Nanoseconds(): Int64
        Round(m: Duration | NumberLike): Duration
        Seconds(): Float64
        Truncate(m: Duration | NumberLike): Duration
    }

    function FixedZone(name: string, offset: Int | NumberLike): LocationPointer
    function LoadLocation(name: string): LocationPointer
    function LoadLocationFromTZData(name: string, data: Bytes): LocationPointer
    interface LocationPointer extends Native {
        readonly __LocationPointer: LocationPointer
        String(): string
    }

    const January: Month
    const February: Month
    const March: Month
    const April: Month
    const May: Month
    const June: Month
    const July: Month
    const August: Month
    const September: Month
    const October: Month
    const November: Month
    const December: Month
    function Month(v: Int | NumberLike): Month
    interface Month extends Number {
        readonly __Month: Month

        String(): string
    }

    interface ParseErrorPointer extends Error {
        readonly __ParseErrorPointer: ParseErrorPointer

        Layout: string
        Value: string
        LayoutElem: string
        ValueElem: string
        Message: string
    }

    function NewTicker(d: Duration | NumberLike): TickerPointer
    interface TickerPointer extends Native {
        readonly __TickerPointer: TickerPointer
        C: ReadChannel<Time>

        Reset(d: Duration | NumberLike)
        Stop()
    }

    function Date(year: Int | NumberLike, month: Month | NumberLike, day: Int | NumberLike,
        hour: Int | NumberLike, min: Int | NumberLike, sec: Int | NumberLike,
        nsec: Int | NumberLike, loc: LocationPointer): Time
    function Now(): Time
    function Parse(layout: string, value: string): Time
    function ParseInLocation(layout: string, value: string, loc: LocationPointer): Time
    function Unix(sec: Int64 | NumberLike, nsec: Int64 | NumberLike): Time
    function UnixMicro(usec: Int64 | NumberLike): Time
    function UnixMilli(msec: Int64 | NumberLike): Time

    interface Time extends Native {
        private readonly __Time: Time
        private constructor()

        Add(d: Duration | NumberLike): Time
        AddDate(years: Int | NumberLike, months: Int | NumberLike, days: Int | NumberLike): Time
        After(u: Time): boolean
        /** added in go1.5 */
        AppendFormat(b: Bytes, layout: string): Bytes
        Before(u: Time): boolean
        /** return (hour, min, sec int)*/
        Clock(): [Int, Int, Int]
        /** return (year int, month Month, day int)*/
        Date(): [Int, Month, Int]
        Day(): Int
        Equal(u: Time): boolean
        Format(layout: string): string
        /** added in go1.17 */
        GoString(): string
        GobEncode(): Bytes
        Hour(): Int
        /** return (year, week int)*/
        ISOWeek(): [Int, Int]
        In(loc: LocationPointer): Time
        /** added in go1.17 */
        IsDST(): boolean
        IsZero(): boolean
        Local(): Time
        Location(): LocationPointer
        /** added in go1.2 */
        MarshalBinary(): Bytes
        MarshalJSON(): Bytes
        /** added in go1.2 */
        MarshalText(): Bytes
        Minute(): Int
        Month(): Month
        Nanosecond(): Int
        /** added in go1.1 */
        Round(d: Duration | NumberLike): Time
        Second(): Int
        String(): string
        Sub(u: Time): Duration
        /** added in go1.1 */
        Truncate(d: Duration): Time
        UTC(): Time
        Unix(): Int64
        /** added in go1.17 */
        UnixMicro(): Int64
        /** added in go1.17 */
        UnixMilli(): Int64
        UnixNano(): Int64

        Weekday(): Weekday
        Year(): Int
        YearDay(): Int
        /** return (name string, offset int) */
        Zone(): [string, Int]
    }

    function AfterFunc(d: Duration | NumberLike, f: () => void): TimerPointer
    function NewTimer(d: Duration | NumberLike): TimerPointer
    interface TimerPointer extends Native {
        readonly __TimerPointer: TimerPointer
        Reset(d: Duration): boolean
        Stop(): boolean
    }

    function Weekday(v: Int | NumberLike): Weekday
    interface Weekday extends Number {
        readonly __Weekday: Weekday
        String(): string
    }
    const Sunday: Weekday
    const Monday: Weekday
    const Tuesday: Weekday
    const Wednesday: Weekday
    const Thursday: Weekday
    const Friday: Weekday
    const Saturday: Weekday

    function isDuration(v: any): v is Duration
    function isLocationPointer(v: any): v is LocationPointer
    function isMonth(v: any): v is Month
    function isParseErrorPointer(v: any): v is ParseErrorPointer
    function isTickerTickerPointer(v: any): v is TickerPointer
    function isTime(v: any): v is Time
    function isTimerPointer(v: any, isptr?: boolean): v is TimerPointer
    function isWeekday(v: any, isptr?: boolean): v is Weekday
}