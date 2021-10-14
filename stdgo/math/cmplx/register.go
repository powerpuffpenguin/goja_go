package cmplx

import "math/cmplx"

func (f *factory) register() {
	f.Set(`Abs`, cmplx.Abs)
	f.Set(`Acos`, cmplx.Acos)
	f.Set(`Acosh`, cmplx.Acosh)
	f.Set(`Asin`, cmplx.Asin)
	f.Set(`Asinh`, cmplx.Asinh)
	f.Set(`Atan`, cmplx.Atan)
	f.Set(`Atanh`, cmplx.Atanh)
	f.Set(`Conj`, cmplx.Conj)
	f.Set(`Cos`, cmplx.Cos)
	f.Set(`Cosh`, cmplx.Cosh)
	f.Set(`Cot`, cmplx.Cot)
	f.Set(`Exp`, cmplx.Exp)
	f.Set(`Inf`, cmplx.Inf)
	f.Set(`IsInf`, cmplx.IsInf)
	f.Set(`IsNaN`, cmplx.IsNaN)
	f.Set(`Log`, cmplx.Log)
	f.Set(`Log10`, cmplx.Log10)
	f.Set(`NaN`, cmplx.NaN)
	f.Set(`Phase`, cmplx.Phase)
	f.Set(`Polar`, cmplx.Polar)
	f.Set(`Pow`, cmplx.Pow)
	f.Set(`Rect`, cmplx.Rect)
	f.Set(`Sin`, cmplx.Sin)
	f.Set(`Sinh`, cmplx.Sinh)
	f.Set(`Sqrt`, cmplx.Sqrt)
	f.Set(`Tan`, cmplx.Tan)
	f.Set(`Tanh`, cmplx.Tanh)
}
