package driver

import "database/sql/driver"

func (f *factory) register() {
	f.Set(`IsScanValue`, driver.IsScanValue)
	f.Set(`IsValue`, driver.IsValue)

	f.Set(`isConn`, isConn)
	f.Set(`isConnBeginTx`, isConnBeginTx)
	f.Set(`isConnPrepareContext`, isConnPrepareContext)
	f.Set(`isConnector`, isConnector)
	f.Set(`isDriver`, isDriver)
	f.Set(`isDriverContext`, isDriverContext)
	f.Set(`isExecerContext`, isExecerContext)
	f.Set(`isIsolationLevel`, isIsolationLevel)
	f.Set(`isNamedValue`, isNamedValue)
	f.Set(`isNamedValuePointer`, isNamedValuePointer)
	f.Set(`isNamedValueChecker`, isNamedValueChecker)
	f.Set(`isNotNull`, isNotNull)
	f.Set(`isNull`, isNull)
	f.Set(`isPinger`, isPinger)
	f.Set(`isQueryerContext`, isQueryerContext)
	f.Set(`isResult`, isResult)
	f.Set(`isRows`, isRows)
	f.Set(`isRowsAffected`, isRowsAffected)
	f.Set(`isRowsColumnTypeDatabaseTypeName`, isRowsColumnTypeDatabaseTypeName)
	f.Set(`isRowsColumnTypeLength`, isRowsColumnTypeLength)
	f.Set(`isRowsColumnTypeNullable`, isRowsColumnTypeNullable)
	f.Set(`isRowsColumnTypePrecisionScale`, isRowsColumnTypePrecisionScale)
	f.Set(`isRowsColumnTypeScanType`, isRowsColumnTypeScanType)
	f.Set(`isRowsNextResultSet`, isRowsNextResultSet)
	f.Set(`isSessionResetter`, isSessionResetter)
	f.Set(`isStmt`, isStmt)
	f.Set(`isStmtExecContext`, isStmtExecContext)
	f.Set(`isStmtQueryContext`, isStmtQueryContext)
	f.Set(`isTx`, isTx)
	f.Set(`isTxOptions`, isTxOptions)
	f.Set(`isValidator`, isValidator)
	f.Set(`isValue`, isValue)
	f.Set(`isValueConverter`, isValueConverter)
	f.Set(`isValuer`, isValuer)
}
func isConn(i interface{}) bool {
	_, result := i.(driver.Conn)
	return result
}
func isConnBeginTx(i interface{}) bool {
	_, result := i.(driver.ConnBeginTx)
	return result
}
func isConnPrepareContext(i interface{}) bool {
	_, result := i.(driver.ConnPrepareContext)
	return result
}
func isConnector(i interface{}) bool {
	_, result := i.(driver.Connector)
	return result
}
func isDriver(i interface{}) bool {
	_, result := i.(driver.Driver)
	return result
}
func isDriverContext(i interface{}) bool {
	_, result := i.(driver.DriverContext)
	return result
}
func isExecerContext(i interface{}) bool {
	_, result := i.(driver.ExecerContext)
	return result
}
func isIsolationLevel(i interface{}) bool {
	_, result := i.(driver.IsolationLevel)
	return result
}
func isNamedValue(i interface{}) bool {
	_, result := i.(driver.NamedValue)
	return result
}
func isNamedValuePointer(i interface{}) bool {
	_, result := i.(*driver.NamedValue)
	return result
}
func isNamedValueChecker(i interface{}) bool {
	_, result := i.(driver.NamedValueChecker)
	return result
}
func isNotNull(i interface{}) bool {
	_, result := i.(driver.NotNull)
	return result
}
func isNull(i interface{}) bool {
	_, result := i.(driver.Null)
	return result
}
func isPinger(i interface{}) bool {
	_, result := i.(driver.Pinger)
	return result
}
func isQueryerContext(i interface{}) bool {
	_, result := i.(driver.QueryerContext)
	return result
}
func isResult(i interface{}) bool {
	_, result := i.(driver.Result)
	return result
}
func isRows(i interface{}) bool {
	_, result := i.(driver.Rows)
	return result
}
func isRowsAffected(i interface{}) bool {
	_, result := i.(driver.RowsAffected)
	return result
}
func isRowsColumnTypeDatabaseTypeName(i interface{}) bool {
	_, result := i.(driver.RowsColumnTypeDatabaseTypeName)
	return result
}
func isRowsColumnTypeLength(i interface{}) bool {
	_, result := i.(driver.RowsColumnTypeLength)
	return result
}
func isRowsColumnTypeNullable(i interface{}) bool {
	_, result := i.(driver.RowsColumnTypeNullable)
	return result
}
func isRowsColumnTypePrecisionScale(i interface{}) bool {
	_, result := i.(driver.RowsColumnTypePrecisionScale)
	return result
}
func isRowsColumnTypeScanType(i interface{}) bool {
	_, result := i.(driver.RowsColumnTypeScanType)
	return result
}
func isRowsNextResultSet(i interface{}) bool {
	_, result := i.(driver.RowsNextResultSet)
	return result
}
func isSessionResetter(i interface{}) bool {
	_, result := i.(driver.SessionResetter)
	return result
}
func isStmt(i interface{}) bool {
	_, result := i.(driver.Stmt)
	return result
}
func isStmtExecContext(i interface{}) bool {
	_, result := i.(driver.StmtExecContext)
	return result
}
func isStmtQueryContext(i interface{}) bool {
	_, result := i.(driver.StmtQueryContext)
	return result
}
func isTx(i interface{}) bool {
	_, result := i.(driver.Tx)
	return result
}
func isTxOptions(i interface{}) bool {
	_, result := i.(driver.TxOptions)
	return result
}
func isValidator(i interface{}) bool {
	_, result := i.(driver.Validator)
	return result
}
func isValue(i interface{}) bool {
	_, result := i.(driver.Value)
	return result
}
func isValueConverter(i interface{}) bool {
	_, result := i.(driver.ValueConverter)
	return result
}
func isValuer(i interface{}) bool {
	_, result := i.(driver.Valuer)
	return result
}
