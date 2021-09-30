package builtin

import (
	"reflect"
	"unsafe"
)

type Float32 float32
type Float64 float64
type Int int
type Int8 int8
type Int16 int16
type Int32 int32
type Int64 int64
type Uint uint
type Uint8 uint8
type Uint16 uint16
type Uint32 uint32
type Uint64 uint64
type String string

type Float32Slice []float32
type Float64Slice []float64
type IntSlice []int
type Int8Slice []int8
type Int16Slice []int16
type Int32Slice []int32
type Int64Slice []int64
type UintSlice []uint
type Uint8Slice []uint8
type Uint16Slice []uint16
type Uint32Slice []uint32
type Uint64Slice []uint64
type StringSlice []string

func (f *factory) registerNumber() {
	f.Set(`Int64`, f.Int64)
}
func (f *factory) Int64(v Int64) Int64 {
	return v
}

type emptyInterface struct {
	typ  *emptyInterface
	word unsafe.Pointer
}

func wrapArray(value reflect.Value) reflect.Value {
	typ := reflect.SliceOf(value.Type().Elem())
	slice := reflect.New(typ)

	p0 := (*reflect.SliceHeader)(unsafe.Pointer(slice.Pointer()))
	p0.Len = value.Len()
	p0.Cap = p0.Len

	p1 := (*emptyInterface)(unsafe.Pointer(&value))
	p0.Data = uintptr(p1.word)

	return reflect.Indirect(slice)
}
func Wrap(value reflect.Value) interface{} {
	if value.Kind() == reflect.Array {
		return wrapArray(value).Interface()
	}
	i := value.Interface()
	switch i := i.(type) {
	case int:
		return Int(i)
	case int8:
		return Int8(i)
	case int16:
		return Int16(i)
	case int32:
		return Int32(i)
	case int64:
		return Int64(i)
	case uint:
		return Uint(i)
	case uint8:
		return Uint8(i)
	case uint16:
		return Uint16(i)
	case uint32:
		return Uint32(i)
	case uint64:
		return Uint64(i)
	case float32:
		return Float32(i)
	case float64:
		return Float64(i)

	case []int:
		return IntSlice(i)
	case []int8:
		return Int8Slice(i)
	case []int16:
		return Int16Slice(i)
	case []int32:
		return Int32Slice(i)
	case []int64:
		return Int64Slice(i)
	case []uint:
		return UintSlice(i)
	case []uint8:
		return Uint8Slice(i)
	case []uint16:
		return Uint16Slice(i)
	case []uint32:
		return Uint32Slice(i)
	case []uint64:
		return Uint64Slice(i)
	case []float32:
		return Float32Slice(i)
	case []float64:
		return Float64Slice(i)
	default:
	}
	return i
}
func Unwrap(i interface{}) reflect.Value {
	return reflect.ValueOf(unwrap(i))
}
func unwrap(i interface{}) interface{} {
	switch i := i.(type) {
	case Int:
		return int(i)
	case Int8:
		return int8(i)
	case Int16:
		return int16(i)
	case Int32:
		return int32(i)
	case Int64:
		return int64(i)
	case Uint:
		return uint(i)
	case Uint8:
		return uint8(i)
	case Uint16:
		return uint16(i)
	case Uint32:
		return uint32(i)
	case Uint64:
		return uint64(i)
	case Float32:
		return float32(i)
	case Float64:
		return float64(i)
	case IntSlice:
		return []int(i)
	case Int8Slice:
		return []int8(i)
	case Int16Slice:
		return []int16(i)
	case Int32Slice:
		return []int32(i)
	case Int64Slice:
		return []int64(i)
	case UintSlice:
		return []uint(i)
	case Uint8Slice:
		return []uint8(i)
	case Uint16Slice:
		return []uint16(i)
	case Uint32Slice:
		return []uint32(i)
	case Uint64Slice:
		return []uint64(i)
	case Float32Slice:
		return []float32(i)
	case Float64Slice:
		return []float64(i)
	default:
	}
	return i
}
