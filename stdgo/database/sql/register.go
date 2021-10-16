package sql

import (
	"database/sql"

	"github.com/powerpuffpenguin/goja"
)

func (f *factory) register() {
	f.Set(`Open`, sql.Open)
	f.Set(`OpenDB`, sql.OpenDB)

	f.Accessor(`LevelDefault`, f.getLevelDefault, nil)
	f.Accessor(`LevelReadUncommitted`, f.getLevelReadUncommitted, nil)
	f.Accessor(`LevelReadCommitted`, f.getLevelReadCommitted, nil)
	f.Accessor(`LevelWriteCommitted`, f.getLevelWriteCommitted, nil)
	f.Accessor(`LevelRepeatableRead`, f.getLevelRepeatableRead, nil)
	f.Accessor(`LevelSnapshot`, f.getLevelSnapshot, nil)
	f.Accessor(`LevelSerializable`, f.getLevelSerializable, nil)
	f.Accessor(`LevelLinearizable`, f.getLevelLinearizable, nil)

	f.Set(`Named`, sql.Named)
	f.Set(`NullBool`, NullBool)
	f.Set(`NullFloat64`, NullFloat64)
	f.Set(`NullInt32`, NullInt32)
	f.Set(`NullInt64`, NullInt64)
	f.Set(`NullString`, NullString)
	f.Set(`NullTime`, NullTime)

	f.Set(`TxOptions`, TxOptions)

	f.Set(`isColumnTypePointer`, isColumnTypePointer)
	f.Set(`isConnPointer`, isConnPointer)
	f.Set(`isDBPointer`, isDBPointer)
	f.Set(`isDBStats`, isDBStats)
	f.Set(`isIsolationLevel`, isIsolationLevel)
	f.Set(`isNamedArg`, isNamedArg)
	f.Set(`isNullBoolPointer`, isNullBoolPointer)
	f.Set(`isNullFloat64Pointer`, isNullFloat64Pointer)
	f.Set(`isNullInt32Pointer`, isNullInt32Pointer)
	f.Set(`isNullInt64Pointer`, isNullInt64Pointer)
	f.Set(`isNullStringPointer`, isNullStringPointer)
	f.Set(`isNullTimePointer`, isNullTimePointer)
	f.Set(`isOut`, isOut)
	f.Set(`isRawBytes`, isRawBytes)
	f.Set(`isResult`, isResult)
	f.Set(`isRowPointer`, isRowPointer)
	f.Set(`isRowsPointer`, isRowsPointer)
	f.Set(`isScanner`, isScanner)
	f.Set(`isStmtPointer`, isStmtPointer)
	f.Set(`isTxPointer`, isTxPointer)
	f.Set(`isTxOptionsPointer`, isTxOptionsPointer)
}
func TxOptions(isolation sql.IsolationLevel, readOnly bool) *sql.TxOptions {
	return &sql.TxOptions{
		Isolation: isolation,
		ReadOnly:  readOnly,
	}
}
func NullBool() *sql.NullBool {
	return &sql.NullBool{}
}

func NullFloat64() *sql.NullFloat64 {
	return &sql.NullFloat64{}
}
func NullInt32() *sql.NullInt32 {
	return &sql.NullInt32{}
}
func NullInt64() *sql.NullInt64 {
	return &sql.NullInt64{}
}
func NullString() *sql.NullString {
	return &sql.NullString{}
}
func NullTime() *sql.NullTime {
	return &sql.NullTime{}
}
func (f *factory) getLevelDefault(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(sql.LevelDefault)
}
func (f *factory) getLevelReadUncommitted(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(sql.LevelReadUncommitted)
}
func (f *factory) getLevelReadCommitted(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(sql.LevelReadCommitted)
}
func (f *factory) getLevelWriteCommitted(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(sql.LevelWriteCommitted)
}
func (f *factory) getLevelRepeatableRead(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(sql.LevelRepeatableRead)
}
func (f *factory) getLevelSnapshot(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(sql.LevelSnapshot)
}
func (f *factory) getLevelSerializable(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(sql.LevelSerializable)
}
func (f *factory) getLevelLinearizable(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(sql.LevelLinearizable)
}
func isColumnTypePointer(i interface{}) bool {
	_, result := i.(*sql.ColumnType)
	return result
}
func isConnPointer(i interface{}) bool {
	_, result := i.(*sql.Conn)
	return result
}
func isDBPointer(i interface{}) bool {
	_, result := i.(*sql.DB)
	return result
}
func isDBStats(i interface{}) bool {
	_, result := i.(sql.DBStats)
	return result
}
func isIsolationLevel(i interface{}) bool {
	_, result := i.(sql.IsolationLevel)
	return result
}
func isNamedArg(i interface{}) bool {
	_, result := i.(sql.NamedArg)
	return result
}
func isNullBoolPointer(i interface{}) bool {
	_, result := i.(*sql.NullBool)
	return result
}
func isNullFloat64Pointer(i interface{}) bool {
	_, result := i.(*sql.NullFloat64)
	return result
}
func isNullInt32Pointer(i interface{}) bool {
	_, result := i.(*sql.NullInt32)
	return result
}
func isNullInt64Pointer(i interface{}) bool {
	_, result := i.(*sql.NullInt64)
	return result
}
func isNullStringPointer(i interface{}) bool {
	_, result := i.(*sql.NullString)
	return result
}
func isNullTimePointer(i interface{}) bool {
	_, result := i.(*sql.NullTime)
	return result
}
func isOut(i interface{}) bool {
	_, result := i.(sql.Out)
	return result
}
func isRawBytes(i interface{}) bool {
	_, result := i.(sql.RawBytes)
	return result
}
func isResult(i interface{}) bool {
	_, result := i.(sql.Result)
	return result
}
func isRowPointer(i interface{}) bool {
	_, result := i.(*sql.Row)
	return result
}
func isRowsPointer(i interface{}) bool {
	_, result := i.(*sql.Rows)
	return result
}
func isScanner(i interface{}) bool {
	_, result := i.(sql.Scanner)
	return result
}
func isStmtPointer(i interface{}) bool {
	_, result := i.(*sql.Stmt)
	return result
}
func isTxPointer(i interface{}) bool {
	_, result := i.(*sql.Tx)
	return result
}
func isTxOptionsPointer(i interface{}) bool {
	_, result := i.(*sql.TxOptions)
	return result
}
