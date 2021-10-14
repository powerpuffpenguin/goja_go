package big

import (
	"math/big"

	"github.com/powerpuffpenguin/goja"
	"github.com/powerpuffpenguin/goja_go/stdgo/builtin"
)

func (f *factory) register() {
	f.Accessor(`MaxExp`, f.getMaxExp, nil)
	f.Accessor(`MinExp`, f.getMinExp, nil)
	f.Accessor(`MaxPrec`, f.getMaxPrec, nil)
	f.Accessor(`MaxBase`, f.getMaxBase, nil)

	f.Set(`Jacobi`, big.Jacobi)

	f.Set(`Accuracy`, Accuracy)
	f.Accessor(`Below`, f.getBelow, nil)
	f.Accessor(`Exact`, f.getExact, nil)
	f.Accessor(`Above`, f.getAbove, nil)

	f.Set(`NewFloat`, big.NewFloat)
	f.Set(`ParseFloat`, big.ParseFloat)

	f.Set(`NewInt`, big.NewInt)

	f.Set(`NewRat`, big.NewRat)

	f.Set(`RoundingMode`, RoundingMode)
	f.Accessor(`ToNearestEven`, f.getToNearestEven, nil)
	f.Accessor(`ToNearestAway`, f.getToNearestAway, nil)
	f.Accessor(`ToZero`, f.getToZero, nil)
	f.Accessor(`AwayFromZero`, f.getAwayFromZero, nil)
	f.Accessor(`ToNegativeInf`, f.getToNegativeInf, nil)
	f.Accessor(`ToPositiveInf`, f.getToPositiveInf, nil)

	f.Set(`Word`, Word)

	f.Set(`isAccuracy`, isAccuracy)
	f.Set(`isErrNaN`, isErrNaN)
	f.Set(`isFloatPointer`, isFloatPointer)
	f.Set(`isIntPointer`, isIntPointer)
	f.Set(`isRatPointer`, isRatPointer)
	f.Set(`isRoundingMode`, isRoundingMode)
	f.Set(`isWord`, isWord)
}
func isAccuracy(i interface{}) bool {
	_, result := i.(big.Accuracy)
	return result
}
func isErrNaN(i interface{}) bool {
	_, result := i.(big.ErrNaN)
	return result
}
func isFloatPointer(i interface{}) bool {
	_, result := i.(*big.Float)
	return result
}
func isIntPointer(i interface{}) bool {
	_, result := i.(*big.Int)
	return result
}
func isRatPointer(i interface{}) bool {
	_, result := i.(*big.Rat)
	return result
}
func isRoundingMode(i interface{}) bool {
	_, result := i.(big.RoundingMode)
	return result
}
func isWord(i interface{}) bool {
	_, result := i.(big.Word)
	return result
}
func Word(v uint) big.Word {
	return big.Word(v)
}
func (f *factory) getToNearestEven(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(big.ToNearestEven)
}
func (f *factory) getToNearestAway(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(big.ToNearestAway)
}
func (f *factory) getToZero(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(big.ToZero)
}
func (f *factory) getAwayFromZero(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(big.AwayFromZero)
}
func (f *factory) getToNegativeInf(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(big.ToNegativeInf)
}
func (f *factory) getToPositiveInf(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(big.ToPositiveInf)
}
func RoundingMode(v byte) big.RoundingMode {
	return big.RoundingMode(v)
}
func (f *factory) getBelow(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(big.Below)
}
func (f *factory) getExact(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(big.Exact)
}
func (f *factory) getAbove(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(big.Above)
}
func Accuracy(v int8) big.Accuracy {
	return big.Accuracy(v)
}
func (f *factory) getMaxExp(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int32(big.MaxExp))
}
func (f *factory) getMinExp(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int32(big.MinExp))
}
func (f *factory) getMaxPrec(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Uint32(big.MaxPrec))
}
func (f *factory) getMaxBase(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Uint8(big.MaxBase))
}
