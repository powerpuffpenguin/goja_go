declare module "stdgo/database/sql" {
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
        Type,
    } from "stdgo/builtin";
    import * as context from "stdgo/context";
    import * as time from "stdgo/time";
    import * as driver from "stdgo/database/sql/driver";

    interface ColumnTypePointer extends Native {
        readonly __ColumnTypePointer: ColumnTypePointer
        DatabaseTypeName(): string
        /** return (precision, scale int64, ok bool) */
        DecimalSize(): [Int64, Int64, boolean]
        /** return (length int64, ok bool) */
        Length(): [Int64, boolean]
        Name(): string
        /** return (nullable, ok bool) */
        Nullable(): [boolean, boolean]
        ScanType(): Type
    }
    interface ConnPointer extends Native {
        readonly __ConnPointer: ConnPointer

        BeginTx(ctx: context.Context, opts: TxOptionsPointer): TxPointer
        Close(): void
        ExecContext(ctx: context.Context, query: string, ...args: Array<any>): Result
        PingContext(ctx: context.Context): void
        PrepareContext(ctx: context.Context, query: string): StmtPointer
        QueryContext(ctx: context.Context, query: string, ...args: Array<any>): RowsPointer
        QueryRowContext(ctx: context.Context, query: string, ...args: Array<any>): RowsPointer
        Raw(f: (driverConn: any) => Error): void
    }

    function Open(driverName: string, dataSourceName: string): DBPointer
    function OpenDB(c: driver.Connector): DBPointer
    interface DBPointer extends Native {
        readonly __DBPointer: DBPointer

        Begin(): TxPointer
        BeginTx(ctx: context.Context, opts: TxOptionsPointer): TxPointer
        Close(): void
        Conn(ctx: context.Context): ConnPointer
        Driver(): driver.Driver
        Exec(query: string, ...args: Array<any>): Result
        ExecContext(ctx: context.Context, query: string, ...args: Array<any>): Result
        Ping(): void
        PingContext(ctx: context.Context): void
        Prepare(query: string): StmtPointer
        PrepareContext(ctx: context.Context, query: string): StmtPointer
        Query(query: string, ...args: Array<any>): RowPointer
        QueryContext(ctx: context.Context, query: string, ...args: Array<any>): RowPointer
        QueryRow(query: string, ...args: Array<any>): RowPointer
        QueryRowContext(ctx: context.Context, query: string, ...args: Array<any>): RowPointer
        SetConnMaxIdleTime(d: time.Duration | NumberLike): void
        SetConnMaxLifetime(d: time.Duration | NumberLike): void
        SetMaxIdleConns(n: Int | NumberLike): void
        SetMaxOpenConns(n: Int | NumberLike): void
        Stats(): DBStats
    }
    interface DBStats extends Native {
        readonly __DBStats: DBStats
        MaxOpenConnections: Int // Maximum number of open connections to the database.

        // Pool Status
        OpenConnections: Int // The number of established connections both in use and idle.
        InUse: Int // The number of connections currently in use.
        Idle: Int // The number of idle connections.

        // Counters
        WaitCount: Int64         // The total number of connections waited for.
        WaitDuration: time.Duration // The total time blocked waiting for a new connection.
        MaxIdleClosed: Int64         // The total number of connections closed due to SetMaxIdleConns.
        MaxIdleTimeClosed: Int64         // The total number of connections closed due to SetConnMaxIdleTime.
        MaxLifetimeClosed: Int64         // The total number of connections closed due to SetConnMaxLifetime.
    }
    interface IsolationLevel extends Number {
        readonly __IsolationLevel: IsolationLevel

        String(): string
    }
    const LevelDefault: IsolationLevel
    const LevelReadUncommitted: IsolationLevel
    const LevelReadCommitted: IsolationLevel
    const LevelWriteCommitted: IsolationLevel
    const LevelRepeatableRead: IsolationLevel
    const LevelSnapshot: IsolationLevel
    const LevelSerializable: IsolationLevel
    const LevelLinearizable: IsolationLevel

    function Named(name: string, value: any): NamedArg
    interface NamedArg extends Native {
        Name: string
        Value: any
    }

    function NullBool(): NullBoolPointer
    interface NullBoolPointer extends Native {
        readonly __NullBoolPointer: NullBoolPointer
        Bool: boolean
        Valid: boolean // Valid is true if Bool is not NULL

        Scan(value: any): void
        Value(): driver.Value
    }
    function NullFloat64(): NullFloat64Pointer
    interface NullFloat64Pointer extends Native {
        readonly __NullFloat64Pointer: NullFloat64Pointer
        Float64: Float64
        Valid: boolean // Valid is true if Bool is not NULL

        Scan(value: any): void
        Value(): driver.Value
    }
    function NullInt32(): NullInt32Pointer
    interface NullInt32Pointer extends Native {
        readonly __NullInt32Pointer: NullInt32Pointer
        Int32: Int32
        Valid: boolean // Valid is true if Bool is not NULL

        Scan(value: any): void
        Value(): driver.Value
    }
    function NullInt64(): NullInt64Pointer
    interface NullInt64Pointer extends Native {
        readonly __NullInt64Pointer: NullInt64Pointer
        Int64: Int64
        Valid: boolean // Valid is true if Bool is not NULL

        Scan(value: any): void
        Value(): driver.Value
    }
    function NullString(): NullStringPointer
    interface NullStringPointer extends Native {
        readonly __NullStringPointer: NullStringPointer
        String: string
        Valid: boolean // Valid is true if Bool is not NULL

        Scan(value: any): void
        Value(): driver.Value
    }
    function NullTime(): NullTimePointer
    interface NullTimePointer extends Native {
        readonly __NullTimePointer: NullTimePointer
        Time: time.Time
        Valid: boolean // Valid is true if Bool is not NULL

        Scan(value: any): void
        Value(): driver.Value
    }
    interface Out extends Native {
        readonly __Out: Out
        Dest: any
        In: boolean
    }
    interface RawBytes extends Bytes {
        readonly __RawBytes: RawBytes
    }
    interface Result extends Native {
        LastInsertId(): Int64
        RowsAffected(): Int64
    }
    interface RowPointer extends Native {
        readonly __RowPointer: RowPointer

        Err(): Error
        Scan(...dest: Array<any>): void
    }
    interface RowsPointer extends Native {
        readonly __RowsPointer: RowsPointer

        Close(): void
        ColumnTypes(): Array<ColumnTypePointer>
        Columns(): Array<string>
        Err(): Error
        Next(): boolean
        NextResultSet(): boolean
        Scan(...dest: Array<any>): void
    }
    interface Scanner extends Native {
        Scan(src: any): void
    }
    interface StmtPointer extends Native {
        readonly __StmtPointer: StmtPointer

        Close(): void
        Exec(...args: Array<any>): Result
        ExecContext(ctx: context.Context, ...args: Array<any>): Result
        Query(...args: Array<any>): RowsPointer
        QueryContext(ctx: context.Context, ...args: Array<any>): RowsPointer
        QueryRow(...args: Array<any>): RowPointer
        QueryRowContext(ctx: context.Context, ...args: Array<any>): RowPointer
    }
    interface TxPointer extends Native {
        readonly __TxPointer: TxPointer

        Commit(): void
        Exec(query: string, ...args: Array<any>): Result
        ExecContext(ctx: context.Context, query: string, ...args: Array<any>): Result
        Prepare(query: string): StmtPointer
        PrepareContext(ctx: context.Context, query: string): StmtPointer
        Query(query: string, ...args: Array<any>): RowsPointer
        QueryContext(ctx: context.Context, query: string, ...args: Array<any>): RowsPointer
        QueryRow(query: string, ...args: Array<any>): RowPointer
        QueryRowContext(ctx: context.Context, query: string, ...args: Array<any>): RowPointer
        Rollback(): void
        Stmt(stmt: StmtPointer): StmtPointer
        StmtContext(ctx: context.Context, stmt: StmtPointer): StmtPointer
    }

    function TxOptions(isolation: IsolationLevel, readOnly: boolean): TxOptionsPointer
    interface TxOptionsPointer extends Native {
        readonly __TxOptionsPointer: TxOptionsPointer

        Isolation: IsolationLevel
        ReadOnly: boolean
    }

    function isColumnTypePointer(v: any): v is ColumnTypePointer
    function isConnPointer(v: any): v is ConnPointer
    function isDBPointer(v: any): v is DBPointer
    function isDBStats(v: any): v is DBStats
    function isIsolationLevel(v: any): v is IsolationLevel
    function isNamedArg(v: any): v is NamedArg
    function isNullBoolPointer(v: any): v is NullBoolPointer
    function isNullFloat64Pointer(v: any): v is NullFloat64Pointer
    function isNullInt32Pointer(v: any): v is NullInt32Pointer
    function isNullInt64Pointer(v: any): v is NullInt64Pointer
    function isNullStringPointer(v: any): v is NullStringPointer
    function isNullTimePointer(v: any): v is NullTimePointer
    function isOut(v: any): v is Out
    function isRawBytes(v: any): v is RawBytes
    function isResult(v: any): v is Result
    function isRowPointer(v: any): v is RowPointer
    function isRowsPointer(v: any): v is RowsPointer
    function isScanner(v: any): v is Scanner
    function isStmtPointer(v: any): v is StmtPointer
    function isTxPointer(v: any): v is TxPointer
    function isTxOptionsPointer(v: any): v is TxOptionsPointer
}