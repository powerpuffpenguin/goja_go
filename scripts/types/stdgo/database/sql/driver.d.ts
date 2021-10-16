declare module "stdgo/database/sql/driver" {
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

    function IsScanValue(v: any): boolean
    function IsValue(v: any): boolean

    interface Conn extends Native {
        Prepare(query: string): Stmt
        Close(): void
        Begin(): Tx
    }
    interface ConnBeginTx extends Native {
        BeginTx(ctx: context.Context, opts: TxOptions): Tx
    }
    interface ConnPrepareContext extends Native {
        PrepareContext(ctx: context.Context, query: string): Stmt
    }
    interface Connector extends Native {
        Connect(ctx: context.Context): Conn
        Driver(): Driver
    }
    interface Driver extends Native {
        Open(name: string): Conn
    }
    interface DriverContext extends Native {
        OpenConnector(name: string): Connector
    }
    interface ExecerContext extends Native {
        ExecContext(ctx: context.Context, query: string, args: Array<NamedValue>): Result
    }
    interface IsolationLevel extends Number {
        readonly __IsolationLevel: IsolationLevel
    }
    interface NamedValuePointer extends Native {
        readonly __NamedValuePointer: NamedValuePointer
        Name: string
        Ordinal: Int
        Value: Value
    }
    interface NamedValue extends Native {
        readonly __NamedValue: NamedValue
        Name: string
        Ordinal: Int
        Value: Value
    }
    interface NamedValueChecker extends Native {
        CheckNamedValue(n: NamedValuePointer): void
    }
    interface NotNull extends Native {
        readonly __NotNull: NotNull
        Converter: ValueConverter
        ConvertValue(v: any): Value
    }
    interface Null extends Native {
        readonly __Null: Null
        Converter: ValueConverter
        ConvertValue(v: any): Value
    }
    interface Pinger extends Native {
        Ping(ctx: context.Context): void
    }
    interface QueryerContext extends Native {
        QueryContext(ctx: context.Context, query: string, args: Array<NamedValue>): Rows
    }
    interface Result extends Native {
        LastInsertId(): Int64
        RowsAffected(): Int64
    }
    interface Rows extends Native {
        Columns(): Array<string>
        Close(): void
        Next(dest: Array<Value>): void
    }
    interface RowsAffected extends Number {
        readonly __RowsAffected: RowsAffected

        LastInsertId(): Int64
        RowsAffected(): Int64
    }
    interface RowsColumnTypeDatabaseTypeName extends Rows {
        ColumnTypeDatabaseTypeName(index: Int | NumberLike): string
    }
    interface RowsColumnTypeLength extends Rows {
        /** return (length int64, ok bool) */
        ColumnTypeLength(index: Int | NumberLike): [Int64, boolean]
    }
    interface RowsColumnTypeNullable extends Rows {
        /** return (nullable, ok bool) */
        ColumnTypeNullable(index: Int | NumberLike): [boolean, boolean]
    }
    interface RowsColumnTypePrecisionScale extends Rows {
        /** return (precision, scale int64, ok bool) */
        ColumnTypePrecisionScale(index: Int | NumberLike): [Int64, Int64, boolean]
    }
    interface RowsColumnTypeScanType extends Rows {
        ColumnTypeScanType(index: Int | NumberLike): Type
    }
    interface RowsNextResultSet extends Rows {
        HasNextResultSet(): boolean
        NextResultSet(): void
    }
    interface SessionResetter extends Native {
        ResetSession(ctx: context.Context): void
    }
    interface Stmt extends Native {
        Close(): void
        NumInput(): Int
        Exec(args: Array<Value>): Result
        Query(args: Array<Value>): Rows
    }
    interface StmtExecContext extends Native {
        ExecContext(ctx: context.Context, args: Array<NamedValue>): Result
    }
    interface StmtQueryContext extends Native {
        QueryContext(ctx: context.Context, args: Array<NamedValue>): Rows
    }
    interface Tx extends Native {
        Commit(): void
        Rollback(): void
    }
    interface TxOptions extends Native {
        readonly __TxOptions: TxOptions
        Isolation: IsolationLevel
        ReadOnly: boolean
    }
    interface Validator extends Native {
        IsValid(): boolean
    }
    interface Value extends Native { }
    interface ValueConverter extends Native {
        ConvertValue(v: any): Value
    }
    interface Valuer extends Native {
        Value(): Value
    }

    function isConn(v: any): v is Conn
    function isConnBeginTx(v: any): v is ConnBeginTx
    function isConnPrepareContext(v: any): v is ConnPrepareContext
    function isConnector(v: any): v is Connector
    function isDriver(v: any): v is Driver
    function isDriverContext(v: any): v is DriverContext
    function isExecerContext(v: any): v is ExecerContext
    function isIsolationLevel(v: any): v is IsolationLevel
    function isNamedValue(v: any): v is NamedValue
    function isNamedValuePointer(v: any): v is NamedValuePointer
    function isNamedValueChecker(v: any): v is NamedValueChecker
    function isNotNull(v: any): v is NotNull
    function isNull(v: any): v is Null
    function isPinger(v: any): v is Pinger
    function isQueryerContext(v: any): v is QueryerContext
    function isResult(v: any): v is Result
    function isRows(v: any): v is Rows
    function isRowsAffected(v: any): v is RowsAffected
    function isRowsColumnTypeDatabaseTypeName(v: any): v is RowsColumnTypeDatabaseTypeName
    function isRowsColumnTypeLength(v: any): v is RowsColumnTypeLength
    function isRowsColumnTypeNullable(v: any): v is RowsColumnTypeNullable
    function isRowsColumnTypePrecisionScale(v: any): v is RowsColumnTypePrecisionScale
    function isRowsColumnTypeScanType(v: any): v is RowsColumnTypeScanType
    function isRowsNextResultSet(v: any): v is RowsNextResultSet
    function isSessionResetter(v: any): v is SessionResetter
    function isStmt(v: any): v is Stmt
    function isStmtExecContext(v: any): v is StmtExecContext
    function isStmtQueryContext(v: any): v is StmtQueryContext
    function isTx(v: any): v is Tx
    function isTxOptions(v: any): v is TxOptions
    function isValidator(v: any): v is Validator
    function isValue(v: any): v is Value
    function isValueConverter(v: any): v is ValueConverter
    function isValuer(v: any): v is Valuer
}