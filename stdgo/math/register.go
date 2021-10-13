package math

import (
	"math"

	"github.com/powerpuffpenguin/goja"
	"github.com/powerpuffpenguin/goja_go/stdgo/builtin"
)

func (f *factory) register() {
	f.Accessor(`E`, f.getE, nil)
	f.Accessor(`Pi`, f.getPi, nil)
	f.Accessor(`Phi`, f.getPhi, nil)
	f.Accessor(`Sqrt2`, f.getSqrt2, nil)
	f.Accessor(`SqrtE`, f.getSqrtE, nil)
	f.Accessor(`SqrtPi`, f.getSqrtPi, nil)
	f.Accessor(`SqrtPhi`, f.getSqrtPhi, nil)
	f.Accessor(`Ln2`, f.getLn2, nil)
	f.Accessor(`Log2E`, f.getLog2E, nil)
	f.Accessor(`Ln10`, f.getLn10, nil)
	f.Accessor(`Log10E`, f.getLog10E, nil)

	f.Accessor(`MaxFloat32`, f.getMaxFloat32, nil)
	f.Accessor(`SmallestNonzeroFloat32`, f.getSmallestNonzeroFloat32, nil)

	f.Accessor(`MaxFloat64`, f.getMaxFloat64, nil)
	f.Accessor(`SmallestNonzeroFloat64`, f.getSmallestNonzeroFloat64, nil)

	f.Accessor(`MaxInt64`, f.getMaxInt64, nil)
	f.Accessor(`MaxInt32`, f.getMaxInt32, nil)
	f.Accessor(`MaxInt16`, f.getMaxInt16, nil)
	f.Accessor(`MaxInt8`, f.getMaxInt8, nil)
	f.Accessor(`MaxUint64`, f.getMaxUint64, nil)
	f.Accessor(`MaxUint32`, f.getMaxUint32, nil)
	f.Accessor(`MaxUint16`, f.getMaxUint16, nil)
	f.Accessor(`MaxUint8`, f.getMaxUint8, nil)
	f.Accessor(`MinInt64`, f.getMinInt64, nil)
	f.Accessor(`MinInt32`, f.getMinInt32, nil)
	f.Accessor(`MinInt16`, f.getMinInt16, nil)
	f.Accessor(`MinInt8`, f.getMinInt8, nil)

	f.Set(`Abs`, math.Abs)
	f.Set(`Acos`, math.Acos)
	f.Set(`Acosh`, math.Acosh)
	f.Set(`Asin`, math.Asin)
	f.Set(`Asinh`, math.Asinh)
	f.Set(`Atan`, math.Atan)
	f.Set(`Atan2`, math.Atan2)
	f.Set(`Atanh`, math.Atanh)
	f.Set(`Cbrt`, math.Cbrt)
	f.Set(`Ceil`, math.Ceil)
	f.Set(`Copysign`, math.Copysign)
	f.Set(`Cos`, math.Cos)
	f.Set(`Cosh`, math.Cosh)
	f.Set(`Dim`, math.Dim)
	f.Set(`Erf`, math.Erf)
	f.Set(`Erfc`, math.Erfc)
	f.Set(`Erfcinv`, math.Erfcinv)
	f.Set(`Erfinv`, math.Erfinv)
	f.Set(`Exp`, math.Exp)
	f.Set(`Exp2`, math.Exp2)
	f.Set(`Expm1`, math.Expm1)
	f.Set(`FMA`, math.FMA)
	f.Set(`Float32bits`, math.Float32bits)
	f.Set(`Float32frombits`, math.Float32frombits)
	f.Set(`Float64bits`, math.Float64bits)
	f.Set(`Float64frombits`, math.Float64frombits)
	f.Set(`Floor`, math.Floor)
	f.Set(`Frexp`, math.Frexp)
	f.Set(`Gamma`, math.Gamma)
	f.Set(`Hypot`, math.Hypot)
	f.Set(`Ilogb`, math.Ilogb)
	f.Set(`Inf`, math.Inf)
	f.Set(`IsInf`, math.IsInf)
	f.Set(`IsNaN`, math.IsNaN)
	f.Set(`J0`, math.J0)
	f.Set(`J1`, math.J1)
	f.Set(`Jn`, math.Jn)
	f.Set(`Ldexp`, math.Ldexp)
	f.Set(`Lgamma`, math.Lgamma)
	f.Set(`Log`, math.Log)
	f.Set(`Log10`, math.Log10)
	f.Set(`Log1p`, math.Log1p)
	f.Set(`Log2`, math.Log2)
	f.Set(`Logb`, math.Logb)
	f.Set(`Max`, math.Max)
	f.Set(`Min`, math.Min)
	f.Set(`Mod`, math.Mod)
	f.Set(`Modf`, math.Modf)
	f.Set(`NaN`, math.NaN)
	f.Set(`Nextafter`, math.Nextafter)
	f.Set(`Nextafter32`, math.Nextafter32)
	f.Set(`Pow`, math.Pow)
	f.Set(`Pow10`, math.Pow10)
	f.Set(`Remainder`, math.Remainder)
	f.Set(`Round`, math.Round)
	f.Set(`RoundToEven`, math.RoundToEven)
	f.Set(`Signbit`, math.Signbit)
	f.Set(`Sin`, math.Sin)
	f.Set(`Sincos`, math.Sincos)
	f.Set(`Sinh`, math.Sinh)
	f.Set(`Sqrt`, math.Sqrt)
	f.Set(`Tan`, math.Tan)
	f.Set(`Tanh`, math.Tanh)
	f.Set(`Trunc`, math.Trunc)
	f.Set(`Y0`, math.Y0)
	f.Set(`Y1`, math.Y1)
	f.Set(`Yn`, math.Yn)
}

func (f *factory) getMaxInt64(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int64(math.MaxInt64))
}
func (f *factory) getMaxInt32(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int32(math.MaxInt32))
}
func (f *factory) getMaxInt16(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int16(math.MaxInt16))
}
func (f *factory) getMaxInt8(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int8(math.MaxInt8))
}
func (f *factory) getMaxUint64(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Uint64(math.MaxUint64))
}
func (f *factory) getMaxUint32(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Uint32(math.MaxUint32))
}
func (f *factory) getMaxUint16(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Uint16(math.MaxUint16))
}
func (f *factory) getMaxUint8(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Uint8(math.MaxUint8))
}
func (f *factory) getMaxFloat64(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Float64(math.MaxFloat64))
}

func (f *factory) getMinInt64(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int64(math.MinInt64))
}
func (f *factory) getMinInt32(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int32(math.MinInt32))
}
func (f *factory) getMinInt16(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int16(math.MinInt16))
}
func (f *factory) getMinInt8(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int8(math.MinInt8))
}

func (f *factory) getMaxFloat32(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Float32(math.MaxFloat32))
}
func (f *factory) getSmallestNonzeroFloat32(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Float32(math.SmallestNonzeroFloat32))
}
func (f *factory) getSmallestNonzeroFloat64(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Float64(math.SmallestNonzeroFloat64))
}
func (f *factory) getE(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Float64(math.E))
}
func (f *factory) getPi(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Float64(math.Pi))
}
func (f *factory) getPhi(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Float64(math.Phi))
}
func (f *factory) getSqrt2(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Float64(math.Sqrt2))
}
func (f *factory) getSqrtE(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Float64(math.SqrtE))
}
func (f *factory) getSqrtPi(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Float64(math.SqrtPi))
}
func (f *factory) getSqrtPhi(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Float64(math.SqrtPhi))
}
func (f *factory) getLn2(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Float64(math.Ln2))
}
func (f *factory) getLog2E(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Float64(math.Log2E))
}
func (f *factory) getLn10(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Float64(math.Ln10))
}
func (f *factory) getLog10E(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Float64(math.Log10E))
}
