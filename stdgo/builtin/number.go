package builtin

import (
	"errors"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/powerpuffpenguin/goja"
)

type Int int

func (v Int) String() string {
	return fmt.Sprint(int(v))
}
func (v Int) ToNumber() goja.Value {
	return goja.Int(int64(v))
}
func (v Int) ToInt() int {
	return int(v)
}
func (v Int) ToInt64() int64 {
	return int64(v)
}
func (v Int) ToInt32() int32 {
	return int32(v)
}
func (v Int) ToInt16() int16 {
	return int16(v)
}
func (v Int) ToInt8() int8 {
	return int8(v)
}
func (v Int) ToUint() uint {
	return uint(v)
}
func (v Int) ToUint64() uint64 {
	return uint64(v)
}
func (v Int) ToUint32() uint32 {
	return uint32(v)
}
func (v Int) ToUint16() uint16 {
	return uint16(v)
}
func (v Int) ToUint8() uint8 {
	return uint8(v)
}
func (v Int) ToFloat64() float64 {
	return float64(v)
}
func (v Int) ToFloat32() float32 {
	return float32(v)
}
func (v Int) ABS() int {
	result := int(v)
	if result < 0 {
		return -result
	}
	return result
}
func (v Int) Negate() int {
	return int(-v)
}
func (v Int) Add(vals ...int) int {
	result := int(v)
	for _, val := range vals {
		result += val
	}
	return result
}
func (v Int) Sub(vals ...int) int {
	result := int(v)
	for _, val := range vals {
		result -= val
	}
	return result
}
func (v Int) Mul(vals ...int) int {
	result := int(v)
	for _, val := range vals {
		result *= val
	}
	return result
}
func (v Int) Div(vals ...int) (int, error) {
	result := int(v)
	if result == 0 {
		return 0, nil
	}
	for _, val := range vals {
		if val == 0 {
			return 0, errors.New("divide by zero")
		}
		result /= val
	}
	return result, nil
}
func (v Int) Mod(vals ...int) (int, error) {
	result := int(v)
	if result == 0 {
		return 0, nil
	}
	for _, val := range vals {
		if val == 0 {
			return 0, errors.New("divide by zero")
		}
		result %= val
		if result == 0 {
			return 0, nil
		}
	}
	return result, nil
}
func (v Int) And(vals ...int) int {
	result := int(v)
	for _, val := range vals {
		result &= val
	}
	return result
}
func (v Int) AndNot(vals ...int) int {
	result := int(v)
	for _, val := range vals {
		result &= (^val)
	}
	return result
}
func (v Int) Not() int {
	return ^int(v)
}
func (v Int) Or(vals ...int) int {
	result := int(v)
	for _, val := range vals {
		result |= val
	}
	return result
}
func (v Int) Xor(vals ...int) int {
	result := int(v)
	for _, val := range vals {
		result ^= val
	}
	return result
}
func (v Int) ShiftLeft(vals ...int) int {
	result := int(v)
	for _, val := range vals {
		result <<= val
	}
	return result
}
func (v Int) ShiftRight(vals ...int) int {
	result := int(v)
	for _, val := range vals {
		result >>= val
	}
	return result
}
func (v Int) Compare(val int) goja.Value {
	current := int(v)
	if current == val {
		return goja.Int(0)
	} else if current < val {
		return goja.Int(-1)
	}
	return goja.Int(1)
}
func (v Int) Max(vals ...int) int {
	result := int(v)
	for _, val := range vals {
		if val > result {
			result = val
		}
	}
	return result
}
func (v Int) Min(vals ...int) int {
	result := int(v)
	for _, val := range vals {
		if val < result {
			result = val
		}
	}
	return result
}
func (f *factory) isInt(call goja.FunctionCall) goja.Value {
	_, result := call.Argument(0).Export().(Int)
	return goja.Boolean(result)
}
func (f *factory) Int(call goja.FunctionCall) goja.Value {
	var (
		r      = f.runtime
		result int
		args   = call.Arguments
		count  = len(call.Arguments)
	)
	if count > 0 {
		if s, ok := args[0].Export().(string); ok {
			var base = 10
			if count > 1 {
				e := r.ExportTo(args[1], &base)
				if e != nil {
					panic(r.NewGoError(e))
				}
			}
			val, e := strconv.ParseInt(s, base, 64)
			if e != nil {
				panic(r.NewGoError(e))
			}
			result = int(val)
		} else {
			e := r.ExportTo(args[0], &result)
			if e != nil {
				panic(r.NewGoError(e))
			}
		}
	}
	return r.ToValue(Int(result))
}

type Int64 int64

func (v Int64) String() string {
	return fmt.Sprint(int64(v))
}
func (v Int64) ToNumber() goja.Value {
	return goja.Int(int64(v))
}
func (v Int64) ToInt() int {
	return int(v)
}
func (v Int64) ToInt64() int64 {
	return int64(v)
}
func (v Int64) ToInt32() int32 {
	return int32(v)
}
func (v Int64) ToInt16() int16 {
	return int16(v)
}
func (v Int64) ToInt8() int8 {
	return int8(v)
}
func (v Int64) ToUint() uint {
	return uint(v)
}
func (v Int64) ToUint64() uint64 {
	return uint64(v)
}
func (v Int64) ToUint32() uint32 {
	return uint32(v)
}
func (v Int64) ToUint16() uint16 {
	return uint16(v)
}
func (v Int64) ToUint8() uint8 {
	return uint8(v)
}
func (v Int64) ToFloat64() float64 {
	return float64(v)
}
func (v Int64) ToFloat32() float32 {
	return float32(v)
}
func (v Int64) ABS() int64 {
	result := int64(v)
	if result < 0 {
		return -result
	}
	return result
}
func (v Int64) Negate() int64 {
	return int64(-v)
}
func (v Int64) Add(vals ...int64) int64 {
	result := int64(v)
	for _, val := range vals {
		result += val
	}
	return result
}
func (v Int64) Sub(vals ...int64) int64 {
	result := int64(v)
	for _, val := range vals {
		result -= val
	}
	return result
}
func (v Int64) Mul(vals ...int64) int64 {
	result := int64(v)
	for _, val := range vals {
		result *= val
	}
	return result
}
func (v Int64) Div(vals ...int64) (int64, error) {
	result := int64(v)
	if result == 0 {
		return 0, nil
	}
	for _, val := range vals {
		if val == 0 {
			return 0, errors.New("divide by zero")
		}
		result /= val
	}
	return result, nil
}
func (v Int64) Mod(vals ...int64) (int64, error) {
	result := int64(v)
	if result == 0 {
		return 0, nil
	}
	for _, val := range vals {
		if val == 0 {
			return 0, errors.New("divide by zero")
		}
		result %= val
		if result == 0 {
			return 0, nil
		}
	}
	return result, nil
}
func (v Int64) And(vals ...int64) int64 {
	result := int64(v)
	for _, val := range vals {
		result &= val
	}
	return result
}
func (v Int64) AndNot(vals ...int64) int64 {
	result := int64(v)
	for _, val := range vals {
		result &= (^val)
	}
	return result
}
func (v Int64) Not() int64 {
	return ^int64(v)
}
func (v Int64) Or(vals ...int64) int64 {
	result := int64(v)
	for _, val := range vals {
		result |= val
	}
	return result
}
func (v Int64) Xor(vals ...int64) int64 {
	result := int64(v)
	for _, val := range vals {
		result ^= val
	}
	return result
}
func (v Int64) ShiftLeft(vals ...int) int64 {
	result := int64(v)
	for _, val := range vals {
		result <<= val
	}
	return result
}
func (v Int64) ShiftRight(vals ...int) int64 {
	result := int64(v)
	for _, val := range vals {
		result >>= val
	}
	return result
}
func (v Int64) Compare(val int64) goja.Value {
	current := int64(v)
	if current == val {
		return goja.Int(0)
	} else if current < val {
		return goja.Int(-1)
	}
	return goja.Int(1)
}
func (v Int64) Max(vals ...int64) int64 {
	result := int64(v)
	for _, val := range vals {
		if val > result {
			result = val
		}
	}
	return result
}
func (v Int64) Min(vals ...int64) int64 {
	result := int64(v)
	for _, val := range vals {
		if val < result {
			result = val
		}
	}
	return result
}
func (f *factory) isInt64(call goja.FunctionCall) goja.Value {
	_, result := call.Argument(0).Export().(Int64)
	return goja.Boolean(result)
}
func (f *factory) Int64(call goja.FunctionCall) goja.Value {
	var (
		r      = f.runtime
		result int64
		args   = call.Arguments
		count  = len(call.Arguments)
	)
	if count > 0 {
		if s, ok := args[0].Export().(string); ok {
			var base = 10
			if count > 1 {
				e := r.ExportTo(args[1], &base)
				if e != nil {
					panic(r.NewGoError(e))
				}
			}
			val, e := strconv.ParseInt(s, base, 64)
			if e != nil {
				panic(r.NewGoError(e))
			}
			result = val
		} else {
			e := r.ExportTo(args[0], &result)
			if e != nil {
				panic(r.NewGoError(e))
			}
		}
	}
	return r.ToValue(Int64(result))
}

type Int32 int32

func (v Int32) String() string {
	return fmt.Sprint(int32(v))
}
func (v Int32) ToNumber() goja.Value {
	return goja.Int(int64(v))
}
func (v Int32) ToInt() int {
	return int(v)
}
func (v Int32) ToInt64() int64 {
	return int64(v)
}
func (v Int32) ToInt32() int32 {
	return int32(v)
}
func (v Int32) ToInt16() int16 {
	return int16(v)
}
func (v Int32) ToInt8() int8 {
	return int8(v)
}
func (v Int32) ToUint() uint {
	return uint(v)
}
func (v Int32) ToUint64() uint64 {
	return uint64(v)
}
func (v Int32) ToUint32() uint32 {
	return uint32(v)
}
func (v Int32) ToUint16() uint16 {
	return uint16(v)
}
func (v Int32) ToUint8() uint8 {
	return uint8(v)
}
func (v Int32) ToFloat64() float64 {
	return float64(v)
}
func (v Int32) ToFloat32() float32 {
	return float32(v)
}
func (v Int32) ABS() int32 {
	result := int32(v)
	if result < 0 {
		return -result
	}
	return result
}
func (v Int32) Negate() int32 {
	return int32(-v)
}
func (v Int32) Add(vals ...int32) int32 {
	result := int32(v)
	for _, val := range vals {
		result += val
	}
	return result
}
func (v Int32) Sub(vals ...int32) int32 {
	result := int32(v)
	for _, val := range vals {
		result -= val
	}
	return result
}
func (v Int32) Mul(vals ...int32) int32 {
	result := int32(v)
	for _, val := range vals {
		result *= val
	}
	return result
}
func (v Int32) Div(vals ...int32) (int32, error) {
	result := int32(v)
	if result == 0 {
		return 0, nil
	}
	for _, val := range vals {
		if val == 0 {
			return 0, errors.New("divide by zero")
		}
		result /= val
	}
	return result, nil
}
func (v Int32) Mod(vals ...int32) (int32, error) {
	result := int32(v)
	if result == 0 {
		return 0, nil
	}
	for _, val := range vals {
		if val == 0 {
			return 0, errors.New("divide by zero")
		}
		result %= val
		if result == 0 {
			return 0, nil
		}
	}
	return result, nil
}
func (v Int32) And(vals ...int32) int32 {
	result := int32(v)
	for _, val := range vals {
		result &= val
	}
	return result
}
func (v Int32) AndNot(vals ...int32) int32 {
	result := int32(v)
	for _, val := range vals {
		result &= (^val)
	}
	return result
}
func (v Int32) Not() int32 {
	return ^int32(v)
}
func (v Int32) Or(vals ...int32) int32 {
	result := int32(v)
	for _, val := range vals {
		result |= val
	}
	return result
}
func (v Int32) Xor(vals ...int32) int32 {
	result := int32(v)
	for _, val := range vals {
		result ^= val
	}
	return result
}
func (v Int32) ShiftLeft(vals ...int) int32 {
	result := int32(v)
	for _, val := range vals {
		result <<= val
	}
	return result
}
func (v Int32) ShiftRight(vals ...int) int32 {
	result := int32(v)
	for _, val := range vals {
		result >>= val
	}
	return result
}
func (v Int32) Compare(val int32) goja.Value {
	current := int32(v)
	if current == val {
		return goja.Int(0)
	} else if current < val {
		return goja.Int(-1)
	}
	return goja.Int(1)
}
func (v Int32) Max(vals ...int32) int32 {
	result := int32(v)
	for _, val := range vals {
		if val > result {
			result = val
		}
	}
	return result
}
func (v Int32) Min(vals ...int32) int32 {
	result := int32(v)
	for _, val := range vals {
		if val < result {
			result = val
		}
	}
	return result
}
func (f *factory) isInt32(call goja.FunctionCall) goja.Value {
	_, result := call.Argument(0).Export().(Int32)
	return goja.Boolean(result)
}
func (f *factory) Int32(call goja.FunctionCall) goja.Value {
	var (
		r      = f.runtime
		result int32
		args   = call.Arguments
		count  = len(call.Arguments)
	)
	if count > 0 {
		if s, ok := args[0].Export().(string); ok {
			var base = 10
			if count > 1 {
				e := r.ExportTo(args[1], &base)
				if e != nil {
					panic(r.NewGoError(e))
				}
			}
			val, e := strconv.ParseInt(s, base, 32)
			if e != nil {
				panic(r.NewGoError(e))
			}
			result = int32(val)
		} else {
			e := r.ExportTo(args[0], &result)
			if e != nil {
				panic(r.NewGoError(e))
			}
		}
	}
	return r.ToValue(Int32(result))
}

type Int16 int16

func (v Int16) String() string {
	return fmt.Sprint(int16(v))
}
func (v Int16) ToNumber() goja.Value {
	return goja.Int(int64(v))
}
func (v Int16) ToInt() int {
	return int(v)
}
func (v Int16) ToInt64() int64 {
	return int64(v)
}
func (v Int16) ToInt32() int32 {
	return int32(v)
}
func (v Int16) ToInt16() int16 {
	return int16(v)
}
func (v Int16) ToInt8() int8 {
	return int8(v)
}
func (v Int16) ToUint() uint {
	return uint(v)
}
func (v Int16) ToUint64() uint64 {
	return uint64(v)
}
func (v Int16) ToUint32() uint32 {
	return uint32(v)
}
func (v Int16) ToUint16() uint16 {
	return uint16(v)
}
func (v Int16) ToUint8() uint8 {
	return uint8(v)
}
func (v Int16) ToFloat64() float64 {
	return float64(v)
}
func (v Int16) ToFloat32() float32 {
	return float32(v)
}
func (v Int16) ABS() int16 {
	result := int16(v)
	if result < 0 {
		return -result
	}
	return result
}
func (v Int16) Negate() int16 {
	return int16(-v)
}
func (v Int16) Add(vals ...int16) int16 {
	result := int16(v)
	for _, val := range vals {
		result += val
	}
	return result
}
func (v Int16) Sub(vals ...int16) int16 {
	result := int16(v)
	for _, val := range vals {
		result -= val
	}
	return result
}
func (v Int16) Mul(vals ...int16) int16 {
	result := int16(v)
	for _, val := range vals {
		result *= val
	}
	return result
}
func (v Int16) Div(vals ...int16) (int16, error) {
	result := int16(v)
	if result == 0 {
		return 0, nil
	}
	for _, val := range vals {
		if val == 0 {
			return 0, errors.New("divide by zero")
		}
		result /= val
	}
	return result, nil
}
func (v Int16) Mod(vals ...int16) (int16, error) {
	result := int16(v)
	if result == 0 {
		return 0, nil
	}
	for _, val := range vals {
		if val == 0 {
			return 0, errors.New("divide by zero")
		}
		result %= val
		if result == 0 {
			return 0, nil
		}
	}
	return result, nil
}
func (v Int16) And(vals ...int16) int16 {
	result := int16(v)
	for _, val := range vals {
		result &= val
	}
	return result
}
func (v Int16) AndNot(vals ...int16) int16 {
	result := int16(v)
	for _, val := range vals {
		result &= (^val)
	}
	return result
}
func (v Int16) Not() int16 {
	return ^int16(v)
}
func (v Int16) Or(vals ...int16) int16 {
	result := int16(v)
	for _, val := range vals {
		result |= val
	}
	return result
}
func (v Int16) Xor(vals ...int16) int16 {
	result := int16(v)
	for _, val := range vals {
		result ^= val
	}
	return result
}
func (v Int16) ShiftLeft(vals ...int) int16 {
	result := int16(v)
	for _, val := range vals {
		result <<= val
	}
	return result
}
func (v Int16) ShiftRight(vals ...int) int16 {
	result := int16(v)
	for _, val := range vals {
		result >>= val
	}
	return result
}
func (v Int16) Compare(val int16) goja.Value {
	current := int16(v)
	if current == val {
		return goja.Int(0)
	} else if current < val {
		return goja.Int(-1)
	}
	return goja.Int(1)
}
func (v Int16) Max(vals ...int16) int16 {
	result := int16(v)
	for _, val := range vals {
		if val > result {
			result = val
		}
	}
	return result
}
func (v Int16) Min(vals ...int16) int16 {
	result := int16(v)
	for _, val := range vals {
		if val < result {
			result = val
		}
	}
	return result
}
func (f *factory) isInt16(call goja.FunctionCall) goja.Value {
	_, result := call.Argument(0).Export().(Int16)
	return goja.Boolean(result)
}
func (f *factory) Int16(call goja.FunctionCall) goja.Value {
	var (
		r      = f.runtime
		result int16
		args   = call.Arguments
		count  = len(call.Arguments)
	)
	if count > 0 {
		if s, ok := args[0].Export().(string); ok {
			var base = 10
			if count > 1 {
				e := r.ExportTo(args[1], &base)
				if e != nil {
					panic(r.NewGoError(e))
				}
			}
			val, e := strconv.ParseInt(s, base, 16)
			if e != nil {
				panic(r.NewGoError(e))
			}
			result = int16(val)
		} else {
			e := r.ExportTo(args[0], &result)
			if e != nil {
				panic(r.NewGoError(e))
			}
		}
	}
	return r.ToValue(Int16(result))
}

type Int8 int8

func (v Int8) String() string {
	return fmt.Sprint(int8(v))
}
func (v Int8) ToNumber() goja.Value {
	return goja.Int(int64(v))
}
func (v Int8) ToInt() int {
	return int(v)
}
func (v Int8) ToInt64() int64 {
	return int64(v)
}
func (v Int8) ToInt32() int32 {
	return int32(v)
}
func (v Int8) ToInt16() int16 {
	return int16(v)
}
func (v Int8) ToInt8() int8 {
	return int8(v)
}
func (v Int8) ToUint() uint {
	return uint(v)
}
func (v Int8) ToUint64() uint64 {
	return uint64(v)
}
func (v Int8) ToUint32() uint32 {
	return uint32(v)
}
func (v Int8) ToUint16() uint16 {
	return uint16(v)
}
func (v Int8) ToUint8() uint8 {
	return uint8(v)
}
func (v Int8) ToFloat64() float64 {
	return float64(v)
}
func (v Int8) ToFloat32() float32 {
	return float32(v)
}
func (v Int8) ABS() int8 {
	result := int8(v)
	if result < 0 {
		return -result
	}
	return result
}
func (v Int8) Negate() int8 {
	return int8(-v)
}
func (v Int8) Add(vals ...int8) int8 {
	result := int8(v)
	for _, val := range vals {
		result += val
	}
	return result
}
func (v Int8) Sub(vals ...int8) int8 {
	result := int8(v)
	for _, val := range vals {
		result -= val
	}
	return result
}
func (v Int8) Mul(vals ...int8) int8 {
	result := int8(v)
	for _, val := range vals {
		result *= val
	}
	return result
}
func (v Int8) Div(vals ...int8) (int8, error) {
	result := int8(v)
	if result == 0 {
		return 0, nil
	}
	for _, val := range vals {
		if val == 0 {
			return 0, errors.New("divide by zero")
		}
		result /= val
	}
	return result, nil
}
func (v Int8) Mod(vals ...int8) (int8, error) {
	result := int8(v)
	if result == 0 {
		return 0, nil
	}
	for _, val := range vals {
		if val == 0 {
			return 0, errors.New("divide by zero")
		}
		result %= val
		if result == 0 {
			return 0, nil
		}
	}
	return result, nil
}
func (v Int8) And(vals ...int8) int8 {
	result := int8(v)
	for _, val := range vals {
		result &= val
	}
	return result
}
func (v Int8) AndNot(vals ...int8) int8 {
	result := int8(v)
	for _, val := range vals {
		result &= (^val)
	}
	return result
}
func (v Int8) Not() int8 {
	return ^int8(v)
}
func (v Int8) Or(vals ...int8) int8 {
	result := int8(v)
	for _, val := range vals {
		result |= val
	}
	return result
}
func (v Int8) Xor(vals ...int8) int8 {
	result := int8(v)
	for _, val := range vals {
		result ^= val
	}
	return result
}
func (v Int8) ShiftLeft(vals ...int) int8 {
	result := int8(v)
	for _, val := range vals {
		result <<= val
	}
	return result
}
func (v Int8) ShiftRight(vals ...int) int8 {
	result := int8(v)
	for _, val := range vals {
		result >>= val
	}
	return result
}
func (v Int8) Compare(val int8) goja.Value {
	current := int8(v)
	if current == val {
		return goja.Int(0)
	} else if current < val {
		return goja.Int(-1)
	}
	return goja.Int(1)
}
func (v Int8) Max(vals ...int8) int8 {
	result := int8(v)
	for _, val := range vals {
		if val > result {
			result = val
		}
	}
	return result
}
func (v Int8) Min(vals ...int8) int8 {
	result := int8(v)
	for _, val := range vals {
		if val < result {
			result = val
		}
	}
	return result
}
func (f *factory) isInt8(call goja.FunctionCall) goja.Value {
	_, result := call.Argument(0).Export().(Int8)
	return goja.Boolean(result)
}
func (f *factory) Int8(call goja.FunctionCall) goja.Value {
	var (
		r      = f.runtime
		result int8
		args   = call.Arguments
		count  = len(call.Arguments)
	)
	if count > 0 {
		if s, ok := args[0].Export().(string); ok {
			var base = 10
			if count > 1 {
				e := r.ExportTo(args[1], &base)
				if e != nil {
					panic(r.NewGoError(e))
				}
			}
			val, e := strconv.ParseInt(s, base, 8)
			if e != nil {
				panic(r.NewGoError(e))
			}
			result = int8(val)
		} else {
			e := r.ExportTo(args[0], &result)
			if e != nil {
				panic(r.NewGoError(e))
			}
		}
	}
	return r.ToValue(Int8(result))
}

type Uint uint

func (v Uint) String() string {
	return fmt.Sprint(uint(v))
}
func (v Uint) ToNumber() goja.Value {
	return goja.Int(int64(v))
}
func (v Uint) ToInt() int {
	return int(v)
}
func (v Uint) ToInt64() int64 {
	return int64(v)
}
func (v Uint) ToInt32() int32 {
	return int32(v)
}
func (v Uint) ToInt16() int16 {
	return int16(v)
}
func (v Uint) ToInt8() int8 {
	return int8(v)
}
func (v Uint) ToUint() uint {
	return uint(v)
}
func (v Uint) ToUint64() uint64 {
	return uint64(v)
}
func (v Uint) ToUint32() uint32 {
	return uint32(v)
}
func (v Uint) ToUint16() uint16 {
	return uint16(v)
}
func (v Uint) ToUint8() uint8 {
	return uint8(v)
}
func (v Uint) ToFloat64() float64 {
	return float64(v)
}
func (v Uint) ToFloat32() float32 {
	return float32(v)
}
func (v Uint) Add(vals ...uint) uint {
	result := uint(v)
	for _, val := range vals {
		result += val
	}
	return result
}
func (v Uint) Sub(vals ...uint) uint {
	result := uint(v)
	for _, val := range vals {
		result -= val
	}
	return result
}
func (v Uint) Mul(vals ...uint) uint {
	result := uint(v)
	for _, val := range vals {
		result *= val
	}
	return result
}
func (v Uint) Div(vals ...uint) (uint, error) {
	result := uint(v)
	if result == 0 {
		return 0, nil
	}
	for _, val := range vals {
		if val == 0 {
			return 0, errors.New("divide by zero")
		}
		result /= val
	}
	return result, nil
}
func (v Uint) Mod(vals ...uint) (uint, error) {
	result := uint(v)
	if result == 0 {
		return 0, nil
	}
	for _, val := range vals {
		if val == 0 {
			return 0, errors.New("divide by zero")
		}
		result %= val
		if result == 0 {
			return 0, nil
		}
	}
	return result, nil
}
func (v Uint) And(vals ...uint) uint {
	result := uint(v)
	for _, val := range vals {
		result &= val
	}
	return result
}
func (v Uint) AndNot(vals ...uint) uint {
	result := uint(v)
	for _, val := range vals {
		result &= (^val)
	}
	return result
}
func (v Uint) Not() uint {
	return ^uint(v)
}
func (v Uint) Or(vals ...uint) uint {
	result := uint(v)
	for _, val := range vals {
		result |= val
	}
	return result
}
func (v Uint) Xor(vals ...uint) uint {
	result := uint(v)
	for _, val := range vals {
		result ^= val
	}
	return result
}
func (v Uint) ShiftLeft(vals ...int) uint {
	result := uint(v)
	for _, val := range vals {
		result <<= val
	}
	return result
}
func (v Uint) ShiftRight(vals ...int) uint {
	result := uint(v)
	for _, val := range vals {
		result >>= val
	}
	return result
}
func (v Uint) Compare(val uint) goja.Value {
	current := uint(v)
	if current == val {
		return goja.Int(0)
	} else if current < val {
		return goja.Int(-1)
	}
	return goja.Int(1)
}
func (v Uint) Max(vals ...uint) uint {
	result := uint(v)
	for _, val := range vals {
		if val > result {
			result = val
		}
	}
	return result
}
func (v Uint) Min(vals ...uint) uint {
	result := uint(v)
	for _, val := range vals {
		if val < result {
			result = val
		}
	}
	return result
}
func (f *factory) isUint(call goja.FunctionCall) goja.Value {
	_, result := call.Argument(0).Export().(Uint)
	return goja.Boolean(result)
}
func (f *factory) Uint(call goja.FunctionCall) goja.Value {
	var (
		r      = f.runtime
		result uint
		args   = call.Arguments
		count  = len(call.Arguments)
	)
	if count > 0 {
		if s, ok := args[0].Export().(string); ok {
			var base = 10
			if count > 1 {
				e := r.ExportTo(args[1], &base)
				if e != nil {
					panic(r.NewGoError(e))
				}
			}
			val, e := strconv.ParseUint(s, base, 64)
			if e != nil {
				panic(r.NewGoError(e))
			}
			result = uint(val)
		} else {
			e := r.ExportTo(args[0], &result)
			if e != nil {
				panic(r.NewGoError(e))
			}
		}
	}
	return r.ToValue(Uint(result))
}

type Uint64 uint64

func (v Uint64) String() string {
	return fmt.Sprint(uint64(v))
}
func (v Uint64) ToNumber() goja.Value {
	return goja.Int(int64(v))
}
func (v Uint64) ToInt() int {
	return int(v)
}
func (v Uint64) ToInt64() int64 {
	return int64(v)
}
func (v Uint64) ToInt32() int32 {
	return int32(v)
}
func (v Uint64) ToInt16() int16 {
	return int16(v)
}
func (v Uint64) ToInt8() int8 {
	return int8(v)
}
func (v Uint64) ToUint() uint {
	return uint(v)
}
func (v Uint64) ToUint64() uint64 {
	return uint64(v)
}
func (v Uint64) ToUint32() uint32 {
	return uint32(v)
}
func (v Uint64) ToUint16() uint16 {
	return uint16(v)
}
func (v Uint64) ToUint8() uint8 {
	return uint8(v)
}
func (v Uint64) ToFloat64() float64 {
	return float64(v)
}
func (v Uint64) ToFloat32() float32 {
	return float32(v)
}
func (v Uint64) Add(vals ...uint64) uint64 {
	result := uint64(v)
	for _, val := range vals {
		result += val
	}
	return result
}
func (v Uint64) Sub(vals ...uint64) uint64 {
	result := uint64(v)
	for _, val := range vals {
		result -= val
	}
	return result
}
func (v Uint64) Mul(vals ...uint64) uint64 {
	result := uint64(v)
	for _, val := range vals {
		result *= val
	}
	return result
}
func (v Uint64) Div(vals ...uint64) (uint64, error) {
	result := uint64(v)
	if result == 0 {
		return 0, nil
	}
	for _, val := range vals {
		if val == 0 {
			return 0, errors.New("divide by zero")
		}
		result /= val
	}
	return result, nil
}
func (v Uint64) Mod(vals ...uint64) (uint64, error) {
	result := uint64(v)
	if result == 0 {
		return 0, nil
	}
	for _, val := range vals {
		if val == 0 {
			return 0, errors.New("divide by zero")
		}
		result %= val
		if result == 0 {
			return 0, nil
		}
	}
	return result, nil
}
func (v Uint64) And(vals ...uint64) uint64 {
	result := uint64(v)
	for _, val := range vals {
		result &= val
	}
	return result
}
func (v Uint64) AndNot(vals ...uint64) uint64 {
	result := uint64(v)
	for _, val := range vals {
		result &= (^val)
	}
	return result
}
func (v Uint64) Not() uint64 {
	return ^uint64(v)
}
func (v Uint64) Or(vals ...uint64) uint64 {
	result := uint64(v)
	for _, val := range vals {
		result |= val
	}
	return result
}
func (v Uint64) Xor(vals ...uint64) uint64 {
	result := uint64(v)
	for _, val := range vals {
		result ^= val
	}
	return result
}
func (v Uint64) ShiftLeft(vals ...int) uint64 {
	result := uint64(v)
	for _, val := range vals {
		result <<= val
	}
	return result
}
func (v Uint64) ShiftRight(vals ...int) uint64 {
	result := uint64(v)
	for _, val := range vals {
		result >>= val
	}
	return result
}
func (v Uint64) Compare(val uint64) goja.Value {
	current := uint64(v)
	if current == val {
		return goja.Int(0)
	} else if current < val {
		return goja.Int(-1)
	}
	return goja.Int(1)
}
func (v Uint64) Max(vals ...uint64) uint64 {
	result := uint64(v)
	for _, val := range vals {
		if val > result {
			result = val
		}
	}
	return result
}
func (v Uint64) Min(vals ...uint64) uint64 {
	result := uint64(v)
	for _, val := range vals {
		if val < result {
			result = val
		}
	}
	return result
}
func (f *factory) isUint64(call goja.FunctionCall) goja.Value {
	_, result := call.Argument(0).Export().(Uint64)
	return goja.Boolean(result)
}
func (f *factory) Uint64(call goja.FunctionCall) goja.Value {
	var (
		r      = f.runtime
		result uint64
		args   = call.Arguments
		count  = len(call.Arguments)
	)
	if count > 0 {
		if s, ok := args[0].Export().(string); ok {
			var base = 10
			if count > 1 {
				e := r.ExportTo(args[1], &base)
				if e != nil {
					panic(r.NewGoError(e))
				}
			}
			val, e := strconv.ParseUint(s, base, 64)
			if e != nil {
				panic(r.NewGoError(e))
			}
			result = val
		} else {
			e := r.ExportTo(args[0], &result)
			if e != nil {
				panic(r.NewGoError(e))
			}
		}
	}
	return r.ToValue(Uint64(result))
}

type Uint32 uint32

func (v Uint32) String() string {
	return fmt.Sprint(uint32(v))
}
func (v Uint32) ToNumber() goja.Value {
	return goja.Int(int64(v))
}
func (v Uint32) ToInt() int {
	return int(v)
}
func (v Uint32) ToInt64() int64 {
	return int64(v)
}
func (v Uint32) ToInt32() int32 {
	return int32(v)
}
func (v Uint32) ToInt16() int16 {
	return int16(v)
}
func (v Uint32) ToInt8() int8 {
	return int8(v)
}
func (v Uint32) ToUint() uint {
	return uint(v)
}
func (v Uint32) ToUint64() uint64 {
	return uint64(v)
}
func (v Uint32) ToUint32() uint32 {
	return uint32(v)
}
func (v Uint32) ToUint16() uint16 {
	return uint16(v)
}
func (v Uint32) ToUint8() uint8 {
	return uint8(v)
}
func (v Uint32) ToFloat64() float64 {
	return float64(v)
}
func (v Uint32) ToFloat32() float32 {
	return float32(v)
}
func (v Uint32) Add(vals ...uint32) uint32 {
	result := uint32(v)
	for _, val := range vals {
		result += val
	}
	return result
}
func (v Uint32) Sub(vals ...uint32) uint32 {
	result := uint32(v)
	for _, val := range vals {
		result -= val
	}
	return result
}
func (v Uint32) Mul(vals ...uint32) uint32 {
	result := uint32(v)
	for _, val := range vals {
		result *= val
	}
	return result
}
func (v Uint32) Div(vals ...uint32) (uint32, error) {
	result := uint32(v)
	if result == 0 {
		return 0, nil
	}
	for _, val := range vals {
		if val == 0 {
			return 0, errors.New("divide by zero")
		}
		result /= val
	}
	return result, nil
}
func (v Uint32) Mod(vals ...uint32) (uint32, error) {
	result := uint32(v)
	if result == 0 {
		return 0, nil
	}
	for _, val := range vals {
		if val == 0 {
			return 0, errors.New("divide by zero")
		}
		result %= val
		if result == 0 {
			return 0, nil
		}
	}
	return result, nil
}
func (v Uint32) And(vals ...uint32) uint32 {
	result := uint32(v)
	for _, val := range vals {
		result &= val
	}
	return result
}
func (v Uint32) AndNot(vals ...uint32) uint32 {
	result := uint32(v)
	for _, val := range vals {
		result &= (^val)
	}
	return result
}
func (v Uint32) Not() uint32 {
	return ^uint32(v)
}
func (v Uint32) Or(vals ...uint32) uint32 {
	result := uint32(v)
	for _, val := range vals {
		result |= val
	}
	return result
}
func (v Uint32) Xor(vals ...uint32) uint32 {
	result := uint32(v)
	for _, val := range vals {
		result ^= val
	}
	return result
}
func (v Uint32) ShiftLeft(vals ...int) uint32 {
	result := uint32(v)
	for _, val := range vals {
		result <<= val
	}
	return result
}
func (v Uint32) ShiftRight(vals ...int) uint32 {
	result := uint32(v)
	for _, val := range vals {
		result >>= val
	}
	return result
}
func (v Uint32) Compare(val uint32) goja.Value {
	current := uint32(v)
	if current == val {
		return goja.Int(0)
	} else if current < val {
		return goja.Int(-1)
	}
	return goja.Int(1)
}
func (v Uint32) Max(vals ...uint32) uint32 {
	result := uint32(v)
	for _, val := range vals {
		if val > result {
			result = val
		}
	}
	return result
}
func (v Uint32) Min(vals ...uint32) uint32 {
	result := uint32(v)
	for _, val := range vals {
		if val < result {
			result = val
		}
	}
	return result
}
func (f *factory) isUint32(call goja.FunctionCall) goja.Value {
	_, result := call.Argument(0).Export().(Uint32)
	return goja.Boolean(result)
}
func (f *factory) Uint32(call goja.FunctionCall) goja.Value {
	var (
		r      = f.runtime
		result uint32
		args   = call.Arguments
		count  = len(call.Arguments)
	)
	if count > 0 {
		if s, ok := args[0].Export().(string); ok {
			var base = 10
			if count > 1 {
				e := r.ExportTo(args[1], &base)
				if e != nil {
					panic(r.NewGoError(e))
				}
			}
			val, e := strconv.ParseUint(s, base, 32)
			if e != nil {
				panic(r.NewGoError(e))
			}
			result = uint32(val)
		} else {
			e := r.ExportTo(args[0], &result)
			if e != nil {
				panic(r.NewGoError(e))
			}
		}
	}
	return r.ToValue(Uint32(result))
}

type Uint16 uint16

func (v Uint16) String() string {
	return fmt.Sprint(uint16(v))
}
func (v Uint16) ToNumber() goja.Value {
	return goja.Int(int64(v))
}
func (v Uint16) ToInt() int {
	return int(v)
}
func (v Uint16) ToInt64() int64 {
	return int64(v)
}
func (v Uint16) ToInt32() int32 {
	return int32(v)
}
func (v Uint16) ToInt16() int16 {
	return int16(v)
}
func (v Uint16) ToInt8() int8 {
	return int8(v)
}
func (v Uint16) ToUint() uint {
	return uint(v)
}
func (v Uint16) ToUint64() uint64 {
	return uint64(v)
}
func (v Uint16) ToUint32() uint32 {
	return uint32(v)
}
func (v Uint16) ToUint16() uint16 {
	return uint16(v)
}
func (v Uint16) ToUint8() uint8 {
	return uint8(v)
}
func (v Uint16) ToFloat64() float64 {
	return float64(v)
}
func (v Uint16) ToFloat32() float32 {
	return float32(v)
}
func (v Uint16) Add(vals ...uint16) uint16 {
	result := uint16(v)
	for _, val := range vals {
		result += val
	}
	return result
}
func (v Uint16) Sub(vals ...uint16) uint16 {
	result := uint16(v)
	for _, val := range vals {
		result -= val
	}
	return result
}
func (v Uint16) Mul(vals ...uint16) uint16 {
	result := uint16(v)
	for _, val := range vals {
		result *= val
	}
	return result
}
func (v Uint16) Div(vals ...uint16) (uint16, error) {
	result := uint16(v)
	if result == 0 {
		return 0, nil
	}
	for _, val := range vals {
		if val == 0 {
			return 0, errors.New("divide by zero")
		}
		result /= val
	}
	return result, nil
}
func (v Uint16) Mod(vals ...uint16) (uint16, error) {
	result := uint16(v)
	if result == 0 {
		return 0, nil
	}
	for _, val := range vals {
		if val == 0 {
			return 0, errors.New("divide by zero")
		}
		result %= val
		if result == 0 {
			return 0, nil
		}
	}
	return result, nil
}
func (v Uint16) And(vals ...uint16) uint16 {
	result := uint16(v)
	for _, val := range vals {
		result &= val
	}
	return result
}
func (v Uint16) AndNot(vals ...uint16) uint16 {
	result := uint16(v)
	for _, val := range vals {
		result &= (^val)
	}
	return result
}
func (v Uint16) Not() uint16 {
	return ^uint16(v)
}
func (v Uint16) Or(vals ...uint16) uint16 {
	result := uint16(v)
	for _, val := range vals {
		result |= val
	}
	return result
}
func (v Uint16) Xor(vals ...uint16) uint16 {
	result := uint16(v)
	for _, val := range vals {
		result ^= val
	}
	return result
}
func (v Uint16) ShiftLeft(vals ...int) uint16 {
	result := uint16(v)
	for _, val := range vals {
		result <<= val
	}
	return result
}
func (v Uint16) ShiftRight(vals ...int) uint16 {
	result := uint16(v)
	for _, val := range vals {
		result >>= val
	}
	return result
}
func (v Uint16) Compare(val uint16) goja.Value {
	current := uint16(v)
	if current == val {
		return goja.Int(0)
	} else if current < val {
		return goja.Int(-1)
	}
	return goja.Int(1)
}
func (v Uint16) Max(vals ...uint16) uint16 {
	result := uint16(v)
	for _, val := range vals {
		if val > result {
			result = val
		}
	}
	return result
}
func (v Uint16) Min(vals ...uint16) uint16 {
	result := uint16(v)
	for _, val := range vals {
		if val < result {
			result = val
		}
	}
	return result
}
func (f *factory) isUint16(call goja.FunctionCall) goja.Value {
	_, result := call.Argument(0).Export().(Uint16)
	return goja.Boolean(result)
}
func (f *factory) Uint16(call goja.FunctionCall) goja.Value {
	var (
		r      = f.runtime
		result uint16
		args   = call.Arguments
		count  = len(call.Arguments)
	)
	if count > 0 {
		if s, ok := args[0].Export().(string); ok {
			var base = 10
			if count > 1 {
				e := r.ExportTo(args[1], &base)
				if e != nil {
					panic(r.NewGoError(e))
				}
			}
			val, e := strconv.ParseUint(s, base, 16)
			if e != nil {
				panic(r.NewGoError(e))
			}
			result = uint16(val)
		} else {
			e := r.ExportTo(args[0], &result)
			if e != nil {
				panic(r.NewGoError(e))
			}
		}
	}
	return r.ToValue(Uint16(result))
}

type Uint8 uint8

func (v Uint8) String() string {
	return fmt.Sprint(uint8(v))
}
func (v Uint8) ToNumber() goja.Value {
	return goja.Int(int64(v))
}
func (v Uint8) ToInt() int {
	return int(v)
}
func (v Uint8) ToInt64() int64 {
	return int64(v)
}
func (v Uint8) ToInt32() int32 {
	return int32(v)
}
func (v Uint8) ToInt16() int16 {
	return int16(v)
}
func (v Uint8) ToInt8() int8 {
	return int8(v)
}
func (v Uint8) ToUint() uint {
	return uint(v)
}
func (v Uint8) ToUint64() uint64 {
	return uint64(v)
}
func (v Uint8) ToUint32() uint32 {
	return uint32(v)
}
func (v Uint8) ToUint16() uint16 {
	return uint16(v)
}
func (v Uint8) ToUint8() uint8 {
	return uint8(v)
}
func (v Uint8) ToFloat64() float64 {
	return float64(v)
}
func (v Uint8) ToFloat32() float32 {
	return float32(v)
}
func (v Uint8) Add(vals ...uint8) uint8 {
	result := uint8(v)
	for _, val := range vals {
		result += val
	}
	return result
}
func (v Uint8) Sub(vals ...uint8) uint8 {
	result := uint8(v)
	for _, val := range vals {
		result -= val
	}
	return result
}
func (v Uint8) Mul(vals ...uint8) uint8 {
	result := uint8(v)
	for _, val := range vals {
		result *= val
	}
	return result
}
func (v Uint8) Div(vals ...uint8) (uint8, error) {
	result := uint8(v)
	if result == 0 {
		return 0, nil
	}
	for _, val := range vals {
		if val == 0 {
			return 0, errors.New("divide by zero")
		}
		result /= val
	}
	return result, nil
}
func (v Uint8) Mod(vals ...uint8) (uint8, error) {
	result := uint8(v)
	if result == 0 {
		return 0, nil
	}
	for _, val := range vals {
		if val == 0 {
			return 0, errors.New("divide by zero")
		}
		result %= val
		if result == 0 {
			return 0, nil
		}
	}
	return result, nil
}
func (v Uint8) And(vals ...uint8) uint8 {
	result := uint8(v)
	for _, val := range vals {
		result &= val
	}
	return result
}
func (v Uint8) AndNot(vals ...uint8) uint8 {
	result := uint8(v)
	for _, val := range vals {
		result &= (^val)
	}
	return result
}
func (v Uint8) Not() uint8 {
	return ^uint8(v)
}
func (v Uint8) Or(vals ...uint8) uint8 {
	result := uint8(v)
	for _, val := range vals {
		result |= val
	}
	return result
}
func (v Uint8) Xor(vals ...uint8) uint8 {
	result := uint8(v)
	for _, val := range vals {
		result ^= val
	}
	return result
}
func (v Uint8) ShiftLeft(vals ...int) uint8 {
	result := uint8(v)
	for _, val := range vals {
		result <<= val
	}
	return result
}
func (v Uint8) ShiftRight(vals ...int) uint8 {
	result := uint8(v)
	for _, val := range vals {
		result >>= val
	}
	return result
}
func (v Uint8) Compare(val uint8) goja.Value {
	current := uint8(v)
	if current == val {
		return goja.Int(0)
	} else if current < val {
		return goja.Int(-1)
	}
	return goja.Int(1)
}
func (v Uint8) Max(vals ...uint8) uint8 {
	result := uint8(v)
	for _, val := range vals {
		if val > result {
			result = val
		}
	}
	return result
}
func (v Uint8) Min(vals ...uint8) uint8 {
	result := uint8(v)
	for _, val := range vals {
		if val < result {
			result = val
		}
	}
	return result
}
func (f *factory) isUint8(call goja.FunctionCall) goja.Value {
	_, result := call.Argument(0).Export().(Uint8)
	return goja.Boolean(result)
}
func (f *factory) Uint8(call goja.FunctionCall) goja.Value {
	var (
		r      = f.runtime
		result uint8
		args   = call.Arguments
		count  = len(call.Arguments)
	)
	if count > 0 {
		if s, ok := args[0].Export().(string); ok {
			var base = 10
			if count > 1 {
				e := r.ExportTo(args[1], &base)
				if e != nil {
					panic(r.NewGoError(e))
				}
			}
			val, e := strconv.ParseUint(s, base, 8)
			if e != nil {
				panic(r.NewGoError(e))
			}
			result = uint8(val)
		} else {
			e := r.ExportTo(args[0], &result)
			if e != nil {
				panic(r.NewGoError(e))
			}
		}
	}
	return r.ToValue(Uint8(result))
}

type Float64 float64

func (v Float64) String() string {
	return fmt.Sprint(float64(v))
}
func (v Float64) ToNumber() goja.Value {
	return goja.Float(float64(v))
}
func (v Float64) ToInt() int {
	return int(v)
}
func (v Float64) ToInt64() int64 {
	return int64(v)
}
func (v Float64) ToInt32() int32 {
	return int32(v)
}
func (v Float64) ToInt16() int16 {
	return int16(v)
}
func (v Float64) ToInt8() int8 {
	return int8(v)
}
func (v Float64) ToUint() uint {
	return uint(v)
}
func (v Float64) ToUint64() uint64 {
	return uint64(v)
}
func (v Float64) ToUint32() uint32 {
	return uint32(v)
}
func (v Float64) ToUint16() uint16 {
	return uint16(v)
}
func (v Float64) ToUint8() uint8 {
	return uint8(v)
}
func (v Float64) ToFloat64() float64 {
	return float64(v)
}
func (v Float64) ToFloat32() float32 {
	return float32(v)
}
func (v Float64) ABS() float64 {
	result := float64(v)
	if result < 0 {
		return -result
	}
	return result
}
func (v Float64) Negate() float64 {
	return float64(-v)
}
func (v Float64) Add(vals ...float64) float64 {
	result := float64(v)
	for _, val := range vals {
		result += val
	}
	return result
}
func (v Float64) Sub(vals ...float64) float64 {
	result := float64(v)
	for _, val := range vals {
		result -= val
	}
	return result
}
func (v Float64) Mul(vals ...float64) float64 {
	result := float64(v)
	for _, val := range vals {
		result *= val
	}
	return result
}
func (v Float64) Div(vals ...float64) (float64, error) {
	result := float64(v)
	if result == 0 {
		return 0, nil
	}
	for _, val := range vals {
		if val == 0 {
			return 0, errors.New("divide by zero")
		}
		result /= val
	}
	return result, nil
}
func (v Float64) Sqrt() float64 {
	result := math.Sqrt(float64(v))
	return result
}
func (v Float64) Compare(val float64) goja.Value {
	current := float64(v)
	if current == val {
		return goja.Int(0)
	} else if current < val {
		return goja.Int(-1)
	}
	return goja.Int(1)
}
func (v Float64) Max(vals ...float64) float64 {
	result := float64(v)
	for _, val := range vals {
		if val > result {
			result = val
		}
	}
	return result
}
func (v Float64) Min(vals ...float64) float64 {
	result := float64(v)
	for _, val := range vals {
		if val < result {
			result = val
		}
	}
	return result
}
func (f *factory) isFloat64(call goja.FunctionCall) goja.Value {
	_, result := call.Argument(0).Export().(Float64)
	return goja.Boolean(result)
}
func (f *factory) Float64(call goja.FunctionCall) goja.Value {
	var (
		r      = f.runtime
		result float64
		args   = call.Arguments
		count  = len(call.Arguments)
	)
	if count > 0 {
		if s, ok := args[0].Export().(string); ok {
			val, e := strconv.ParseFloat(s, 64)
			if e != nil {
				panic(r.NewGoError(e))
			}
			result = val
		} else {
			e := r.ExportTo(args[0], &result)
			if e != nil {
				panic(r.NewGoError(e))
			}
		}
	}
	return r.ToValue(Float64(result))
}

type Float32 float32

func (v Float32) String() string {
	return fmt.Sprint(float32(v))
}
func (v Float32) ToNumber() goja.Value {
	return goja.Float(float64(v))
}
func (v Float32) ToInt() int {
	return int(v)
}
func (v Float32) ToInt64() int64 {
	return int64(v)
}
func (v Float32) ToInt32() int32 {
	return int32(v)
}
func (v Float32) ToInt16() int16 {
	return int16(v)
}
func (v Float32) ToInt8() int8 {
	return int8(v)
}
func (v Float32) ToUint() uint {
	return uint(v)
}
func (v Float32) ToUint64() uint64 {
	return uint64(v)
}
func (v Float32) ToUint32() uint32 {
	return uint32(v)
}
func (v Float32) ToUint16() uint16 {
	return uint16(v)
}
func (v Float32) ToUint8() uint8 {
	return uint8(v)
}
func (v Float32) ToFloat64() float64 {
	return float64(v)
}
func (v Float32) ToFloat32() float32 {
	return float32(v)
}
func (v Float32) ABS() float32 {
	result := float32(v)
	if result < 0 {
		return -result
	}
	return result
}
func (v Float32) Negate() float32 {
	return float32(-v)
}
func (v Float32) Add(vals ...float32) float32 {
	result := float32(v)
	for _, val := range vals {
		result += val
	}
	return result
}
func (v Float32) Sub(vals ...float32) float32 {
	result := float32(v)
	for _, val := range vals {
		result -= val
	}
	return result
}
func (v Float32) Mul(vals ...float32) float32 {
	result := float32(v)
	for _, val := range vals {
		result *= val
	}
	return result
}
func (v Float32) Div(vals ...float32) (float32, error) {
	result := float32(v)
	if result == 0 {
		return 0, nil
	}
	for _, val := range vals {
		if val == 0 {
			return 0, errors.New("divide by zero")
		}
		result /= val
	}
	return result, nil
}
func (v Float32) Sqrt() float32 {
	result := math.Sqrt(float64(v))
	return float32(result)
}
func (v Float32) Compare(val float32) goja.Value {
	current := float32(v)
	if current == val {
		return goja.Int(0)
	} else if current < val {
		return goja.Int(-1)
	}
	return goja.Int(1)
}
func (v Float32) Max(vals ...float32) float32 {
	result := float32(v)
	for _, val := range vals {
		if val > result {
			result = val
		}
	}
	return result
}
func (v Float32) Min(vals ...float32) float32 {
	result := float32(v)
	for _, val := range vals {
		if val < result {
			result = val
		}
	}
	return result
}
func (f *factory) isFloat32(call goja.FunctionCall) goja.Value {
	_, result := call.Argument(0).Export().(Float32)
	return goja.Boolean(result)
}
func (f *factory) Float32(call goja.FunctionCall) goja.Value {
	var (
		r      = f.runtime
		result float32
		args   = call.Arguments
		count  = len(call.Arguments)
	)
	if count > 0 {
		if s, ok := args[0].Export().(string); ok {
			val, e := strconv.ParseFloat(s, 32)
			if e != nil {
				panic(r.NewGoError(e))
			}
			result = float32(val)
		} else {
			e := r.ExportTo(args[0], &result)
			if e != nil {
				panic(r.NewGoError(e))
			}
		}
	}
	return r.ToValue(Float32(result))
}

type IntSlice []int

func NewIntSlice(val []int) IntSlice {
	return IntSlice(val)
}
func (v IntSlice) String() string {
	return fmt.Sprint([]int(v))
}
func (v IntSlice) Len() int {
	return len(v)
}
func (v IntSlice) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}
func (v IntSlice) Less(i, j int) bool {
	return v[i] < v[j]
}
func (v IntSlice) Cap() int {
	return cap(v)
}
func (v IntSlice) Copy(src []int) int {
	return copy(v, src)
}
func (v IntSlice) Slice(start int) []int {
	count := len(v)
	if count == 0 {
		return v[:]
	}
	if start < 0 {
		start = 0
	}
	return v[start:]
}
func (v IntSlice) Slice2(start, end int) []int {
	count := len(v)
	if count == 0 {
		return v[:]
	}
	if start < 0 {
		start = 0
	} else if start >= count {
		return v[count:]
	} else if end < start {
		return v[start:start]
	}
	return v[start:end]
}
func (v IntSlice) Append(data ...int) []int {
	return append(v, data...)
}
func (v IntSlice) Get(index int) (int, error) {
	if index < 0 || index >= len(v) {
		return 0, fmt.Errorf("slice bounds out of range get(%d)", index)
	}
	return int(v[index]), nil
}
func (v IntSlice) Set(index int, val int) error {
	if index < 0 || index >= len(v) {
		return fmt.Errorf("slice bounds out of range get(%d)", index)
	}
	v[index] = val
	return nil
}
func (v IntSlice) Join(sep string) string {
	var (
		result string
		count  = len(v)
	)
	if count > 0 {
		strs := make([]string, count)
		for i, val := range v {
			strs[i] = fmt.Sprint(val)
		}
		result = strings.Join(strs, sep)
	}
	return result
}
func (v IntSlice) Asc() {
	sort.Sort(v)
}
func (v IntSlice) Desc() {
	sort.Sort(sortIntSlice(v))
}

type sortIntSlice []int

func (a sortIntSlice) Len() int           { return len(a) }
func (a sortIntSlice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortIntSlice) Less(i, j int) bool { return a[i] > a[j] }

func (f *factory) isIntSlice(call goja.FunctionCall) goja.Value {
	_, result := call.Argument(0).Export().(IntSlice)
	return goja.Boolean(result)
}
func (f *factory) IntSlice(call goja.FunctionCall) goja.Value {
	var (
		r      = f.runtime
		result []int
		args   = call.Arguments
		count  = len(call.Arguments)
	)
	if count > 0 {
		var l int
		e := r.ExportTo(args[0], &l)
		if e != nil {
			panic(r.NewGoError(e))
		} else if l < 0 {
			l = 0
		}
		if count > 1 {
			var c int
			e := r.ExportTo(args[1], &c)
			if e != nil {
				panic(r.NewGoError(e))
			} else if c < l {
				c = l
			}
			result = make([]int, l, c)
		} else {
			result = make([]int, l)
		}
	}
	return r.ToValue(IntSlice(result))
}

type Int64Slice []int64

func NewInt64Slice(val []int64) Int64Slice {
	return Int64Slice(val)
}
func (v Int64Slice) String() string {
	return fmt.Sprint([]int64(v))
}
func (v Int64Slice) Len() int {
	return len(v)
}
func (v Int64Slice) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}
func (v Int64Slice) Less(i, j int) bool {
	return v[i] < v[j]
}
func (v Int64Slice) Cap() int {
	return cap(v)
}
func (v Int64Slice) Copy(src []int64) int {
	return copy(v, src)
}
func (v Int64Slice) Slice(start int) []int64 {
	count := len(v)
	if count == 0 {
		return v[:]
	}
	if start < 0 {
		start = 0
	}
	return v[start:]
}
func (v Int64Slice) Slice2(start, end int) []int64 {
	count := len(v)
	if count == 0 {
		return v[:]
	}
	if start < 0 {
		start = 0
	} else if start >= count {
		return v[count:]
	} else if end < start {
		return v[start:start]
	}
	return v[start:end]
}
func (v Int64Slice) Append(data ...int64) []int64 {
	return append(v, data...)
}
func (v Int64Slice) Get(index int) (int64, error) {
	if index < 0 || index >= len(v) {
		return 0, fmt.Errorf("slice bounds out of range get(%d)", index)
	}
	return int64(v[index]), nil
}
func (v Int64Slice) Set(index int, val int64) error {
	if index < 0 || index >= len(v) {
		return fmt.Errorf("slice bounds out of range get(%d)", index)
	}
	v[index] = val
	return nil
}
func (v Int64Slice) Join(sep string) string {
	var (
		result string
		count  = len(v)
	)
	if count > 0 {
		strs := make([]string, count)
		for i, val := range v {
			strs[i] = fmt.Sprint(val)
		}
		result = strings.Join(strs, sep)
	}
	return result
}
func (v Int64Slice) Asc() {
	sort.Sort(v)
}
func (v Int64Slice) Desc() {
	sort.Sort(sortInt64Slice(v))
}

type sortInt64Slice []int64

func (a sortInt64Slice) Len() int           { return len(a) }
func (a sortInt64Slice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortInt64Slice) Less(i, j int) bool { return a[i] > a[j] }

func (f *factory) isInt64Slice(call goja.FunctionCall) goja.Value {
	_, result := call.Argument(0).Export().(Int64Slice)
	return goja.Boolean(result)
}
func (f *factory) Int64Slice(call goja.FunctionCall) goja.Value {
	var (
		r      = f.runtime
		result []int64
		args   = call.Arguments
		count  = len(call.Arguments)
	)
	if count > 0 {
		var l int
		e := r.ExportTo(args[0], &l)
		if e != nil {
			panic(r.NewGoError(e))
		} else if l < 0 {
			l = 0
		}
		if count > 1 {
			var c int
			e := r.ExportTo(args[1], &c)
			if e != nil {
				panic(r.NewGoError(e))
			} else if c < l {
				c = l
			}
			result = make([]int64, l, c)
		} else {
			result = make([]int64, l)
		}
	}
	return r.ToValue(Int64Slice(result))
}

type Int32Slice []int32

func NewInt32Slice(val []int32) Int32Slice {
	return Int32Slice(val)
}
func (v Int32Slice) String() string {
	return fmt.Sprint([]int32(v))
}
func (v Int32Slice) Len() int {
	return len(v)
}
func (v Int32Slice) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}
func (v Int32Slice) Less(i, j int) bool {
	return v[i] < v[j]
}
func (v Int32Slice) Cap() int {
	return cap(v)
}
func (v Int32Slice) Copy(src []int32) int {
	return copy(v, src)
}
func (v Int32Slice) Slice(start int) []int32 {
	count := len(v)
	if count == 0 {
		return v[:]
	}
	if start < 0 {
		start = 0
	}
	return v[start:]
}
func (v Int32Slice) Slice2(start, end int) []int32 {
	count := len(v)
	if count == 0 {
		return v[:]
	}
	if start < 0 {
		start = 0
	} else if start >= count {
		return v[count:]
	} else if end < start {
		return v[start:start]
	}
	return v[start:end]
}
func (v Int32Slice) Append(data ...int32) []int32 {
	return append(v, data...)
}
func (v Int32Slice) Get(index int) (int32, error) {
	if index < 0 || index >= len(v) {
		return 0, fmt.Errorf("slice bounds out of range get(%d)", index)
	}
	return int32(v[index]), nil
}
func (v Int32Slice) Set(index int, val int32) error {
	if index < 0 || index >= len(v) {
		return fmt.Errorf("slice bounds out of range get(%d)", index)
	}
	v[index] = val
	return nil
}
func (v Int32Slice) Join(sep string) string {
	var (
		result string
		count  = len(v)
	)
	if count > 0 {
		strs := make([]string, count)
		for i, val := range v {
			strs[i] = fmt.Sprint(val)
		}
		result = strings.Join(strs, sep)
	}
	return result
}
func (v Int32Slice) Asc() {
	sort.Sort(v)
}
func (v Int32Slice) Desc() {
	sort.Sort(sortInt32Slice(v))
}

type sortInt32Slice []int32

func (a sortInt32Slice) Len() int           { return len(a) }
func (a sortInt32Slice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortInt32Slice) Less(i, j int) bool { return a[i] > a[j] }

func (f *factory) isInt32Slice(call goja.FunctionCall) goja.Value {
	_, result := call.Argument(0).Export().(Int32Slice)
	return goja.Boolean(result)
}
func (f *factory) Int32Slice(call goja.FunctionCall) goja.Value {
	var (
		r      = f.runtime
		result []int32
		args   = call.Arguments
		count  = len(call.Arguments)
	)
	if count > 0 {
		var l int
		e := r.ExportTo(args[0], &l)
		if e != nil {
			panic(r.NewGoError(e))
		} else if l < 0 {
			l = 0
		}
		if count > 1 {
			var c int
			e := r.ExportTo(args[1], &c)
			if e != nil {
				panic(r.NewGoError(e))
			} else if c < l {
				c = l
			}
			result = make([]int32, l, c)
		} else {
			result = make([]int32, l)
		}
	}
	return r.ToValue(Int32Slice(result))
}

type Int16Slice []int16

func NewInt16Slice(val []int16) Int16Slice {
	return Int16Slice(val)
}
func (v Int16Slice) String() string {
	return fmt.Sprint([]int16(v))
}
func (v Int16Slice) Len() int {
	return len(v)
}
func (v Int16Slice) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}
func (v Int16Slice) Less(i, j int) bool {
	return v[i] < v[j]
}
func (v Int16Slice) Cap() int {
	return cap(v)
}
func (v Int16Slice) Copy(src []int16) int {
	return copy(v, src)
}
func (v Int16Slice) Slice(start int) []int16 {
	count := len(v)
	if count == 0 {
		return v[:]
	}
	if start < 0 {
		start = 0
	}
	return v[start:]
}
func (v Int16Slice) Slice2(start, end int) []int16 {
	count := len(v)
	if count == 0 {
		return v[:]
	}
	if start < 0 {
		start = 0
	} else if start >= count {
		return v[count:]
	} else if end < start {
		return v[start:start]
	}
	return v[start:end]
}
func (v Int16Slice) Append(data ...int16) []int16 {
	return append(v, data...)
}
func (v Int16Slice) Get(index int) (int16, error) {
	if index < 0 || index >= len(v) {
		return 0, fmt.Errorf("slice bounds out of range get(%d)", index)
	}
	return int16(v[index]), nil
}
func (v Int16Slice) Set(index int, val int16) error {
	if index < 0 || index >= len(v) {
		return fmt.Errorf("slice bounds out of range get(%d)", index)
	}
	v[index] = val
	return nil
}
func (v Int16Slice) Join(sep string) string {
	var (
		result string
		count  = len(v)
	)
	if count > 0 {
		strs := make([]string, count)
		for i, val := range v {
			strs[i] = fmt.Sprint(val)
		}
		result = strings.Join(strs, sep)
	}
	return result
}
func (v Int16Slice) Asc() {
	sort.Sort(v)
}
func (v Int16Slice) Desc() {
	sort.Sort(sortInt16Slice(v))
}

type sortInt16Slice []int16

func (a sortInt16Slice) Len() int           { return len(a) }
func (a sortInt16Slice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortInt16Slice) Less(i, j int) bool { return a[i] > a[j] }

func (f *factory) isInt16Slice(call goja.FunctionCall) goja.Value {
	_, result := call.Argument(0).Export().(Int16Slice)
	return goja.Boolean(result)
}
func (f *factory) Int16Slice(call goja.FunctionCall) goja.Value {
	var (
		r      = f.runtime
		result []int16
		args   = call.Arguments
		count  = len(call.Arguments)
	)
	if count > 0 {
		var l int
		e := r.ExportTo(args[0], &l)
		if e != nil {
			panic(r.NewGoError(e))
		} else if l < 0 {
			l = 0
		}
		if count > 1 {
			var c int
			e := r.ExportTo(args[1], &c)
			if e != nil {
				panic(r.NewGoError(e))
			} else if c < l {
				c = l
			}
			result = make([]int16, l, c)
		} else {
			result = make([]int16, l)
		}
	}
	return r.ToValue(Int16Slice(result))
}

type Int8Slice []int8

func NewInt8Slice(val []int8) Int8Slice {
	return Int8Slice(val)
}
func (v Int8Slice) String() string {
	return fmt.Sprint([]int8(v))
}
func (v Int8Slice) Len() int {
	return len(v)
}
func (v Int8Slice) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}
func (v Int8Slice) Less(i, j int) bool {
	return v[i] < v[j]
}
func (v Int8Slice) Cap() int {
	return cap(v)
}
func (v Int8Slice) Copy(src []int8) int {
	return copy(v, src)
}
func (v Int8Slice) Slice(start int) []int8 {
	count := len(v)
	if count == 0 {
		return v[:]
	}
	if start < 0 {
		start = 0
	}
	return v[start:]
}
func (v Int8Slice) Slice2(start, end int) []int8 {
	count := len(v)
	if count == 0 {
		return v[:]
	}
	if start < 0 {
		start = 0
	} else if start >= count {
		return v[count:]
	} else if end < start {
		return v[start:start]
	}
	return v[start:end]
}
func (v Int8Slice) Append(data ...int8) []int8 {
	return append(v, data...)
}
func (v Int8Slice) Get(index int) (int8, error) {
	if index < 0 || index >= len(v) {
		return 0, fmt.Errorf("slice bounds out of range get(%d)", index)
	}
	return int8(v[index]), nil
}
func (v Int8Slice) Set(index int, val int8) error {
	if index < 0 || index >= len(v) {
		return fmt.Errorf("slice bounds out of range get(%d)", index)
	}
	v[index] = val
	return nil
}
func (v Int8Slice) Join(sep string) string {
	var (
		result string
		count  = len(v)
	)
	if count > 0 {
		strs := make([]string, count)
		for i, val := range v {
			strs[i] = fmt.Sprint(val)
		}
		result = strings.Join(strs, sep)
	}
	return result
}
func (v Int8Slice) Asc() {
	sort.Sort(v)
}
func (v Int8Slice) Desc() {
	sort.Sort(sortInt8Slice(v))
}

type sortInt8Slice []int8

func (a sortInt8Slice) Len() int           { return len(a) }
func (a sortInt8Slice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortInt8Slice) Less(i, j int) bool { return a[i] > a[j] }

func (f *factory) isInt8Slice(call goja.FunctionCall) goja.Value {
	_, result := call.Argument(0).Export().(Int8Slice)
	return goja.Boolean(result)
}
func (f *factory) Int8Slice(call goja.FunctionCall) goja.Value {
	var (
		r      = f.runtime
		result []int8
		args   = call.Arguments
		count  = len(call.Arguments)
	)
	if count > 0 {
		var l int
		e := r.ExportTo(args[0], &l)
		if e != nil {
			panic(r.NewGoError(e))
		} else if l < 0 {
			l = 0
		}
		if count > 1 {
			var c int
			e := r.ExportTo(args[1], &c)
			if e != nil {
				panic(r.NewGoError(e))
			} else if c < l {
				c = l
			}
			result = make([]int8, l, c)
		} else {
			result = make([]int8, l)
		}
	}
	return r.ToValue(Int8Slice(result))
}

type UintSlice []uint

func NewUintSlice(val []uint) UintSlice {
	return UintSlice(val)
}
func (v UintSlice) String() string {
	return fmt.Sprint([]uint(v))
}
func (v UintSlice) Len() int {
	return len(v)
}
func (v UintSlice) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}
func (v UintSlice) Less(i, j int) bool {
	return v[i] < v[j]
}
func (v UintSlice) Cap() int {
	return cap(v)
}
func (v UintSlice) Copy(src []uint) int {
	return copy(v, src)
}
func (v UintSlice) Slice(start int) []uint {
	count := len(v)
	if count == 0 {
		return v[:]
	}
	if start < 0 {
		start = 0
	}
	return v[start:]
}
func (v UintSlice) Slice2(start, end int) []uint {
	count := len(v)
	if count == 0 {
		return v[:]
	}
	if start < 0 {
		start = 0
	} else if start >= count {
		return v[count:]
	} else if end < start {
		return v[start:start]
	}
	return v[start:end]
}
func (v UintSlice) Append(data ...uint) []uint {
	return append(v, data...)
}
func (v UintSlice) Get(index int) (uint, error) {
	if index < 0 || index >= len(v) {
		return 0, fmt.Errorf("slice bounds out of range get(%d)", index)
	}
	return uint(v[index]), nil
}
func (v UintSlice) Set(index int, val uint) error {
	if index < 0 || index >= len(v) {
		return fmt.Errorf("slice bounds out of range get(%d)", index)
	}
	v[index] = val
	return nil
}
func (v UintSlice) Join(sep string) string {
	var (
		result string
		count  = len(v)
	)
	if count > 0 {
		strs := make([]string, count)
		for i, val := range v {
			strs[i] = fmt.Sprint(val)
		}
		result = strings.Join(strs, sep)
	}
	return result
}
func (v UintSlice) Asc() {
	sort.Sort(v)
}
func (v UintSlice) Desc() {
	sort.Sort(sortUintSlice(v))
}

type sortUintSlice []uint

func (a sortUintSlice) Len() int           { return len(a) }
func (a sortUintSlice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortUintSlice) Less(i, j int) bool { return a[i] > a[j] }

func (f *factory) isUintSlice(call goja.FunctionCall) goja.Value {
	_, result := call.Argument(0).Export().(UintSlice)
	return goja.Boolean(result)
}
func (f *factory) UintSlice(call goja.FunctionCall) goja.Value {
	var (
		r      = f.runtime
		result []uint
		args   = call.Arguments
		count  = len(call.Arguments)
	)
	if count > 0 {
		var l int
		e := r.ExportTo(args[0], &l)
		if e != nil {
			panic(r.NewGoError(e))
		} else if l < 0 {
			l = 0
		}
		if count > 1 {
			var c int
			e := r.ExportTo(args[1], &c)
			if e != nil {
				panic(r.NewGoError(e))
			} else if c < l {
				c = l
			}
			result = make([]uint, l, c)
		} else {
			result = make([]uint, l)
		}
	}
	return r.ToValue(UintSlice(result))
}

type Uint64Slice []uint64

func NewUint64Slice(val []uint64) Uint64Slice {
	return Uint64Slice(val)
}
func (v Uint64Slice) String() string {
	return fmt.Sprint([]uint64(v))
}
func (v Uint64Slice) Len() int {
	return len(v)
}
func (v Uint64Slice) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}
func (v Uint64Slice) Less(i, j int) bool {
	return v[i] < v[j]
}
func (v Uint64Slice) Cap() int {
	return cap(v)
}
func (v Uint64Slice) Copy(src []uint64) int {
	return copy(v, src)
}
func (v Uint64Slice) Slice(start int) []uint64 {
	count := len(v)
	if count == 0 {
		return v[:]
	}
	if start < 0 {
		start = 0
	}
	return v[start:]
}
func (v Uint64Slice) Slice2(start, end int) []uint64 {
	count := len(v)
	if count == 0 {
		return v[:]
	}
	if start < 0 {
		start = 0
	} else if start >= count {
		return v[count:]
	} else if end < start {
		return v[start:start]
	}
	return v[start:end]
}
func (v Uint64Slice) Append(data ...uint64) []uint64 {
	return append(v, data...)
}
func (v Uint64Slice) Get(index int) (uint64, error) {
	if index < 0 || index >= len(v) {
		return 0, fmt.Errorf("slice bounds out of range get(%d)", index)
	}
	return uint64(v[index]), nil
}
func (v Uint64Slice) Set(index int, val uint64) error {
	if index < 0 || index >= len(v) {
		return fmt.Errorf("slice bounds out of range get(%d)", index)
	}
	v[index] = val
	return nil
}
func (v Uint64Slice) Join(sep string) string {
	var (
		result string
		count  = len(v)
	)
	if count > 0 {
		strs := make([]string, count)
		for i, val := range v {
			strs[i] = fmt.Sprint(val)
		}
		result = strings.Join(strs, sep)
	}
	return result
}
func (v Uint64Slice) Asc() {
	sort.Sort(v)
}
func (v Uint64Slice) Desc() {
	sort.Sort(sortUint64Slice(v))
}

type sortUint64Slice []uint64

func (a sortUint64Slice) Len() int           { return len(a) }
func (a sortUint64Slice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortUint64Slice) Less(i, j int) bool { return a[i] > a[j] }

func (f *factory) isUint64Slice(call goja.FunctionCall) goja.Value {
	_, result := call.Argument(0).Export().(Uint64Slice)
	return goja.Boolean(result)
}
func (f *factory) Uint64Slice(call goja.FunctionCall) goja.Value {
	var (
		r      = f.runtime
		result []uint64
		args   = call.Arguments
		count  = len(call.Arguments)
	)
	if count > 0 {
		var l int
		e := r.ExportTo(args[0], &l)
		if e != nil {
			panic(r.NewGoError(e))
		} else if l < 0 {
			l = 0
		}
		if count > 1 {
			var c int
			e := r.ExportTo(args[1], &c)
			if e != nil {
				panic(r.NewGoError(e))
			} else if c < l {
				c = l
			}
			result = make([]uint64, l, c)
		} else {
			result = make([]uint64, l)
		}
	}
	return r.ToValue(Uint64Slice(result))
}

type Uint32Slice []uint32

func NewUint32Slice(val []uint32) Uint32Slice {
	return Uint32Slice(val)
}
func (v Uint32Slice) String() string {
	return fmt.Sprint([]uint32(v))
}
func (v Uint32Slice) Len() int {
	return len(v)
}
func (v Uint32Slice) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}
func (v Uint32Slice) Less(i, j int) bool {
	return v[i] < v[j]
}
func (v Uint32Slice) Cap() int {
	return cap(v)
}
func (v Uint32Slice) Copy(src []uint32) int {
	return copy(v, src)
}
func (v Uint32Slice) Slice(start int) []uint32 {
	count := len(v)
	if count == 0 {
		return v[:]
	}
	if start < 0 {
		start = 0
	}
	return v[start:]
}
func (v Uint32Slice) Slice2(start, end int) []uint32 {
	count := len(v)
	if count == 0 {
		return v[:]
	}
	if start < 0 {
		start = 0
	} else if start >= count {
		return v[count:]
	} else if end < start {
		return v[start:start]
	}
	return v[start:end]
}
func (v Uint32Slice) Append(data ...uint32) []uint32 {
	return append(v, data...)
}
func (v Uint32Slice) Get(index int) (uint32, error) {
	if index < 0 || index >= len(v) {
		return 0, fmt.Errorf("slice bounds out of range get(%d)", index)
	}
	return uint32(v[index]), nil
}
func (v Uint32Slice) Set(index int, val uint32) error {
	if index < 0 || index >= len(v) {
		return fmt.Errorf("slice bounds out of range get(%d)", index)
	}
	v[index] = val
	return nil
}
func (v Uint32Slice) Join(sep string) string {
	var (
		result string
		count  = len(v)
	)
	if count > 0 {
		strs := make([]string, count)
		for i, val := range v {
			strs[i] = fmt.Sprint(val)
		}
		result = strings.Join(strs, sep)
	}
	return result
}
func (v Uint32Slice) Asc() {
	sort.Sort(v)
}
func (v Uint32Slice) Desc() {
	sort.Sort(sortUint32Slice(v))
}

type sortUint32Slice []uint32

func (a sortUint32Slice) Len() int           { return len(a) }
func (a sortUint32Slice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortUint32Slice) Less(i, j int) bool { return a[i] > a[j] }

func (f *factory) isUint32Slice(call goja.FunctionCall) goja.Value {
	_, result := call.Argument(0).Export().(Uint32Slice)
	return goja.Boolean(result)
}
func (f *factory) Uint32Slice(call goja.FunctionCall) goja.Value {
	var (
		r      = f.runtime
		result []uint32
		args   = call.Arguments
		count  = len(call.Arguments)
	)
	if count > 0 {
		var l int
		e := r.ExportTo(args[0], &l)
		if e != nil {
			panic(r.NewGoError(e))
		} else if l < 0 {
			l = 0
		}
		if count > 1 {
			var c int
			e := r.ExportTo(args[1], &c)
			if e != nil {
				panic(r.NewGoError(e))
			} else if c < l {
				c = l
			}
			result = make([]uint32, l, c)
		} else {
			result = make([]uint32, l)
		}
	}
	return r.ToValue(Uint32Slice(result))
}

type Uint16Slice []uint16

func NewUint16Slice(val []uint16) Uint16Slice {
	return Uint16Slice(val)
}
func (v Uint16Slice) String() string {
	return fmt.Sprint([]uint16(v))
}
func (v Uint16Slice) Len() int {
	return len(v)
}
func (v Uint16Slice) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}
func (v Uint16Slice) Less(i, j int) bool {
	return v[i] < v[j]
}
func (v Uint16Slice) Cap() int {
	return cap(v)
}
func (v Uint16Slice) Copy(src []uint16) int {
	return copy(v, src)
}
func (v Uint16Slice) Slice(start int) []uint16 {
	count := len(v)
	if count == 0 {
		return v[:]
	}
	if start < 0 {
		start = 0
	}
	return v[start:]
}
func (v Uint16Slice) Slice2(start, end int) []uint16 {
	count := len(v)
	if count == 0 {
		return v[:]
	}
	if start < 0 {
		start = 0
	} else if start >= count {
		return v[count:]
	} else if end < start {
		return v[start:start]
	}
	return v[start:end]
}
func (v Uint16Slice) Append(data ...uint16) []uint16 {
	return append(v, data...)
}
func (v Uint16Slice) Get(index int) (uint16, error) {
	if index < 0 || index >= len(v) {
		return 0, fmt.Errorf("slice bounds out of range get(%d)", index)
	}
	return uint16(v[index]), nil
}
func (v Uint16Slice) Set(index int, val uint16) error {
	if index < 0 || index >= len(v) {
		return fmt.Errorf("slice bounds out of range get(%d)", index)
	}
	v[index] = val
	return nil
}
func (v Uint16Slice) Join(sep string) string {
	var (
		result string
		count  = len(v)
	)
	if count > 0 {
		strs := make([]string, count)
		for i, val := range v {
			strs[i] = fmt.Sprint(val)
		}
		result = strings.Join(strs, sep)
	}
	return result
}
func (v Uint16Slice) Asc() {
	sort.Sort(v)
}
func (v Uint16Slice) Desc() {
	sort.Sort(sortUint16Slice(v))
}

type sortUint16Slice []uint16

func (a sortUint16Slice) Len() int           { return len(a) }
func (a sortUint16Slice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortUint16Slice) Less(i, j int) bool { return a[i] > a[j] }

func (f *factory) isUint16Slice(call goja.FunctionCall) goja.Value {
	_, result := call.Argument(0).Export().(Uint16Slice)
	return goja.Boolean(result)
}
func (f *factory) Uint16Slice(call goja.FunctionCall) goja.Value {
	var (
		r      = f.runtime
		result []uint16
		args   = call.Arguments
		count  = len(call.Arguments)
	)
	if count > 0 {
		var l int
		e := r.ExportTo(args[0], &l)
		if e != nil {
			panic(r.NewGoError(e))
		} else if l < 0 {
			l = 0
		}
		if count > 1 {
			var c int
			e := r.ExportTo(args[1], &c)
			if e != nil {
				panic(r.NewGoError(e))
			} else if c < l {
				c = l
			}
			result = make([]uint16, l, c)
		} else {
			result = make([]uint16, l)
		}
	}
	return r.ToValue(Uint16Slice(result))
}

type Uint8Slice []uint8

func NewUint8Slice(val []uint8) Uint8Slice {
	return Uint8Slice(val)
}
func (v Uint8Slice) String() string {
	return fmt.Sprint([]uint8(v))
}
func (v Uint8Slice) Len() int {
	return len(v)
}
func (v Uint8Slice) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}
func (v Uint8Slice) Less(i, j int) bool {
	return v[i] < v[j]
}
func (v Uint8Slice) Cap() int {
	return cap(v)
}
func (v Uint8Slice) Copy(src []uint8) int {
	return copy(v, src)
}
func (v Uint8Slice) Slice(start int) []uint8 {
	count := len(v)
	if count == 0 {
		return v[:]
	}
	if start < 0 {
		start = 0
	}
	return v[start:]
}
func (v Uint8Slice) Slice2(start, end int) []uint8 {
	count := len(v)
	if count == 0 {
		return v[:]
	}
	if start < 0 {
		start = 0
	} else if start >= count {
		return v[count:]
	} else if end < start {
		return v[start:start]
	}
	return v[start:end]
}
func (v Uint8Slice) Append(data ...uint8) []uint8 {
	return append(v, data...)
}
func (v Uint8Slice) Get(index int) (uint8, error) {
	if index < 0 || index >= len(v) {
		return 0, fmt.Errorf("slice bounds out of range get(%d)", index)
	}
	return uint8(v[index]), nil
}
func (v Uint8Slice) Set(index int, val uint8) error {
	if index < 0 || index >= len(v) {
		return fmt.Errorf("slice bounds out of range get(%d)", index)
	}
	v[index] = val
	return nil
}
func (v Uint8Slice) Join(sep string) string {
	var (
		result string
		count  = len(v)
	)
	if count > 0 {
		strs := make([]string, count)
		for i, val := range v {
			strs[i] = fmt.Sprint(val)
		}
		result = strings.Join(strs, sep)
	}
	return result
}
func (v Uint8Slice) Asc() {
	sort.Sort(v)
}
func (v Uint8Slice) Desc() {
	sort.Sort(sortUint8Slice(v))
}

type sortUint8Slice []uint8

func (a sortUint8Slice) Len() int           { return len(a) }
func (a sortUint8Slice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortUint8Slice) Less(i, j int) bool { return a[i] > a[j] }

func (f *factory) isUint8Slice(call goja.FunctionCall) goja.Value {
	_, result := call.Argument(0).Export().(Uint8Slice)
	return goja.Boolean(result)
}
func (f *factory) Uint8Slice(call goja.FunctionCall) goja.Value {
	var (
		r      = f.runtime
		result []uint8
		args   = call.Arguments
		count  = len(call.Arguments)
	)
	if count > 0 {
		var l int
		e := r.ExportTo(args[0], &l)
		if e != nil {
			panic(r.NewGoError(e))
		} else if l < 0 {
			l = 0
		}
		if count > 1 {
			var c int
			e := r.ExportTo(args[1], &c)
			if e != nil {
				panic(r.NewGoError(e))
			} else if c < l {
				c = l
			}
			result = make([]uint8, l, c)
		} else {
			result = make([]uint8, l)
		}
	}
	return r.ToValue(Uint8Slice(result))
}

type Float64Slice []float64

func NewFloat64Slice(val []float64) Float64Slice {
	return Float64Slice(val)
}
func (v Float64Slice) String() string {
	return fmt.Sprint([]float64(v))
}
func (v Float64Slice) Len() int {
	return len(v)
}
func (v Float64Slice) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}
func (v Float64Slice) Less(i, j int) bool {
	return v[i] < v[j]
}
func (v Float64Slice) Cap() int {
	return cap(v)
}
func (v Float64Slice) Copy(src []float64) int {
	return copy(v, src)
}
func (v Float64Slice) Slice(start int) []float64 {
	count := len(v)
	if count == 0 {
		return v[:]
	}
	if start < 0 {
		start = 0
	}
	return v[start:]
}
func (v Float64Slice) Slice2(start, end int) []float64 {
	count := len(v)
	if count == 0 {
		return v[:]
	}
	if start < 0 {
		start = 0
	} else if start >= count {
		return v[count:]
	} else if end < start {
		return v[start:start]
	}
	return v[start:end]
}
func (v Float64Slice) Append(data ...float64) []float64 {
	return append(v, data...)
}
func (v Float64Slice) Get(index int) (float64, error) {
	if index < 0 || index >= len(v) {
		return 0, fmt.Errorf("slice bounds out of range get(%d)", index)
	}
	return float64(v[index]), nil
}
func (v Float64Slice) Set(index int, val float64) error {
	if index < 0 || index >= len(v) {
		return fmt.Errorf("slice bounds out of range get(%d)", index)
	}
	v[index] = val
	return nil
}
func (v Float64Slice) Join(sep string) string {
	var (
		result string
		count  = len(v)
	)
	if count > 0 {
		strs := make([]string, count)
		for i, val := range v {
			strs[i] = fmt.Sprint(val)
		}
		result = strings.Join(strs, sep)
	}
	return result
}
func (v Float64Slice) Asc() {
	sort.Sort(v)
}
func (v Float64Slice) Desc() {
	sort.Sort(sortFloat64Slice(v))
}

type sortFloat64Slice []float64

func (a sortFloat64Slice) Len() int           { return len(a) }
func (a sortFloat64Slice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortFloat64Slice) Less(i, j int) bool { return a[i] > a[j] }

func (f *factory) isFloat64Slice(call goja.FunctionCall) goja.Value {
	_, result := call.Argument(0).Export().(Float64Slice)
	return goja.Boolean(result)
}
func (f *factory) Float64Slice(call goja.FunctionCall) goja.Value {
	var (
		r      = f.runtime
		result []float64
		args   = call.Arguments
		count  = len(call.Arguments)
	)
	if count > 0 {
		var l int
		e := r.ExportTo(args[0], &l)
		if e != nil {
			panic(r.NewGoError(e))
		} else if l < 0 {
			l = 0
		}
		if count > 1 {
			var c int
			e := r.ExportTo(args[1], &c)
			if e != nil {
				panic(r.NewGoError(e))
			} else if c < l {
				c = l
			}
			result = make([]float64, l, c)
		} else {
			result = make([]float64, l)
		}
	}
	return r.ToValue(Float64Slice(result))
}

type Float32Slice []float32

func NewFloat32Slice(val []float32) Float32Slice {
	return Float32Slice(val)
}
func (v Float32Slice) String() string {
	return fmt.Sprint([]float32(v))
}
func (v Float32Slice) Len() int {
	return len(v)
}
func (v Float32Slice) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}
func (v Float32Slice) Less(i, j int) bool {
	return v[i] < v[j]
}
func (v Float32Slice) Cap() int {
	return cap(v)
}
func (v Float32Slice) Copy(src []float32) int {
	return copy(v, src)
}
func (v Float32Slice) Slice(start int) []float32 {
	count := len(v)
	if count == 0 {
		return v[:]
	}
	if start < 0 {
		start = 0
	}
	return v[start:]
}
func (v Float32Slice) Slice2(start, end int) []float32 {
	count := len(v)
	if count == 0 {
		return v[:]
	}
	if start < 0 {
		start = 0
	} else if start >= count {
		return v[count:]
	} else if end < start {
		return v[start:start]
	}
	return v[start:end]
}
func (v Float32Slice) Append(data ...float32) []float32 {
	return append(v, data...)
}
func (v Float32Slice) Get(index int) (float32, error) {
	if index < 0 || index >= len(v) {
		return 0, fmt.Errorf("slice bounds out of range get(%d)", index)
	}
	return float32(v[index]), nil
}
func (v Float32Slice) Set(index int, val float32) error {
	if index < 0 || index >= len(v) {
		return fmt.Errorf("slice bounds out of range get(%d)", index)
	}
	v[index] = val
	return nil
}
func (v Float32Slice) Join(sep string) string {
	var (
		result string
		count  = len(v)
	)
	if count > 0 {
		strs := make([]string, count)
		for i, val := range v {
			strs[i] = fmt.Sprint(val)
		}
		result = strings.Join(strs, sep)
	}
	return result
}
func (v Float32Slice) Asc() {
	sort.Sort(v)
}
func (v Float32Slice) Desc() {
	sort.Sort(sortFloat32Slice(v))
}

type sortFloat32Slice []float32

func (a sortFloat32Slice) Len() int           { return len(a) }
func (a sortFloat32Slice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortFloat32Slice) Less(i, j int) bool { return a[i] > a[j] }

func (f *factory) isFloat32Slice(call goja.FunctionCall) goja.Value {
	_, result := call.Argument(0).Export().(Float32Slice)
	return goja.Boolean(result)
}
func (f *factory) Float32Slice(call goja.FunctionCall) goja.Value {
	var (
		r      = f.runtime
		result []float32
		args   = call.Arguments
		count  = len(call.Arguments)
	)
	if count > 0 {
		var l int
		e := r.ExportTo(args[0], &l)
		if e != nil {
			panic(r.NewGoError(e))
		} else if l < 0 {
			l = 0
		}
		if count > 1 {
			var c int
			e := r.ExportTo(args[1], &c)
			if e != nil {
				panic(r.NewGoError(e))
			} else if c < l {
				c = l
			}
			result = make([]float32, l, c)
		} else {
			result = make([]float32, l)
		}
	}
	return r.ToValue(Float32Slice(result))
}

func (f *factory) getMaxInt64(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(Int64(math.MaxInt64))
}
func (f *factory) getMaxInt32(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(Int32(math.MaxInt32))
}
func (f *factory) getMaxInt16(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(Int16(math.MaxInt16))
}
func (f *factory) getMaxInt8(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(Int8(math.MaxInt8))
}
func (f *factory) getMaxUint64(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(Uint64(math.MaxUint64))
}
func (f *factory) getMaxUint32(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(Uint32(math.MaxUint32))
}
func (f *factory) getMaxUint16(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(Uint16(math.MaxUint16))
}
func (f *factory) getMaxUint8(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(Uint8(math.MaxUint8))
}
func (f *factory) getMaxFloat64(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(Float64(math.MaxFloat64))
}
func (f *factory) getMaxFloat32(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(Float32(math.MaxFloat32))
}

func (f *factory) getMinInt64(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(Int64(math.MinInt64))
}
func (f *factory) getMinInt32(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(Int32(math.MinInt32))
}
func (f *factory) getMinInt16(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(Int16(math.MinInt16))
}
func (f *factory) getMinInt8(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(Int8(math.MinInt8))
}

func (f *factory) getIntSize(call goja.FunctionCall) goja.Value {
	return goja.Int(strconv.IntSize)
}
func (f *factory) registerNumber() {
	f.Set(`Int`, f.Int)
	f.Set(`Int64`, f.Int64)
	f.Set(`Int32`, f.Int32)
	f.Set(`Int16`, f.Int16)
	f.Set(`Int8`, f.Int8)
	f.Set(`Uint`, f.Uint)
	f.Set(`Uint64`, f.Uint64)
	f.Set(`Uint32`, f.Uint32)
	f.Set(`Uint16`, f.Uint16)
	f.Set(`Uint8`, f.Uint8)
	f.Set(`Float64`, f.Float64)
	f.Set(`Float32`, f.Float32)
	f.Accessor(`MaxInt64`, f.getMaxInt64, nil)
	f.Accessor(`MaxInt32`, f.getMaxInt32, nil)
	f.Accessor(`MaxInt16`, f.getMaxInt16, nil)
	f.Accessor(`MaxInt8`, f.getMaxInt8, nil)
	f.Accessor(`MaxUint64`, f.getMaxUint64, nil)
	f.Accessor(`MaxUint32`, f.getMaxUint32, nil)
	f.Accessor(`MaxUint16`, f.getMaxUint16, nil)
	f.Accessor(`MaxUint8`, f.getMaxUint8, nil)
	f.Accessor(`MaxFloat64`, f.getMaxFloat64, nil)
	f.Accessor(`MaxFloat32`, f.getMaxFloat32, nil)
	f.Accessor(`MinInt64`, f.getMinInt64, nil)
	f.Accessor(`MinInt32`, f.getMinInt32, nil)
	f.Accessor(`MinInt16`, f.getMinInt16, nil)
	f.Accessor(`MinInt8`, f.getMinInt8, nil)
	f.Accessor(`IntSize`, f.getIntSize, nil)
	f.Set(`IntSlice`, f.IntSlice)
	f.Set(`Int64Slice`, f.Int64Slice)
	f.Set(`Int32Slice`, f.Int32Slice)
	f.Set(`Int16Slice`, f.Int16Slice)
	f.Set(`Int8Slice`, f.Int8Slice)
	f.Set(`UintSlice`, f.UintSlice)
	f.Set(`Uint64Slice`, f.Uint64Slice)
	f.Set(`Uint32Slice`, f.Uint32Slice)
	f.Set(`Uint16Slice`, f.Uint16Slice)
	f.Set(`Uint8Slice`, f.Uint8Slice)
	f.Set(`Float64Slice`, f.Float64Slice)
	f.Set(`Float32Slice`, f.Float32Slice)

	f.Set(`isInt`, f.isInt)
	f.Set(`isIntSlice`, f.isIntSlice)
	f.Set(`isInt64`, f.isInt64)
	f.Set(`isInt64Slice`, f.isInt64Slice)
	f.Set(`isInt32`, f.isInt32)
	f.Set(`isInt32Slice`, f.isInt32Slice)
	f.Set(`isInt16`, f.isInt16)
	f.Set(`isInt16Slice`, f.isInt16Slice)
	f.Set(`isInt8`, f.isInt8)
	f.Set(`isInt8Slice`, f.isInt8Slice)
	f.Set(`isUint`, f.isUint)
	f.Set(`isUintSlice`, f.isUintSlice)
	f.Set(`isUint64`, f.isUint64)
	f.Set(`isUint64Slice`, f.isUint64Slice)
	f.Set(`isUint32`, f.isUint32)
	f.Set(`isUint32Slice`, f.isUint32Slice)
	f.Set(`isUint16`, f.isUint16)
	f.Set(`isUint16Slice`, f.isUint16Slice)
	f.Set(`isUint8`, f.isUint8)
	f.Set(`isUint8Slice`, f.isUint8Slice)
	f.Set(`isFloat64`, f.isFloat64)
	f.Set(`isFloat64Slice`, f.isFloat64Slice)
	f.Set(`isFloat32`, f.isFloat32)
	f.Set(`isFloat32Slice`, f.isFloat32Slice)
}
