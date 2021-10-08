package builtin

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
