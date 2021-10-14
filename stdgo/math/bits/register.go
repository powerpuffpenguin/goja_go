package bits

import (
	"math/bits"

	"github.com/powerpuffpenguin/goja"
)

func (f *factory) register() {
	f.Accessor(`UintSize`, f.getUintSize, nil)

	f.Set(`Add`, bits.Add)
	f.Set(`Add32`, bits.Add32)
	f.Set(`Add64`, bits.Add64)
	f.Set(`Div`, bits.Div)
	f.Set(`Div32`, bits.Div32)
	f.Set(`Div64`, bits.Div64)
	f.Set(`LeadingZeros`, bits.LeadingZeros)
	f.Set(`LeadingZeros16`, bits.LeadingZeros16)
	f.Set(`LeadingZeros32`, bits.LeadingZeros32)
	f.Set(`LeadingZeros64`, bits.LeadingZeros64)
	f.Set(`LeadingZeros8`, bits.LeadingZeros8)
	f.Set(`Len`, bits.Len)
	f.Set(`Len16`, bits.Len16)
	f.Set(`Len32`, bits.Len32)
	f.Set(`Len64`, bits.Len64)
	f.Set(`Len8`, bits.Len8)
	f.Set(`Mul`, bits.Mul)
	f.Set(`Mul32`, bits.Mul32)
	f.Set(`Mul64`, bits.Mul64)
	f.Set(`OnesCount`, bits.OnesCount)
	f.Set(`OnesCount16`, bits.OnesCount16)
	f.Set(`OnesCount32`, bits.OnesCount32)
	f.Set(`OnesCount64`, bits.OnesCount64)
	f.Set(`OnesCount8`, bits.OnesCount8)
	f.Set(`Rem`, bits.Rem)
	f.Set(`Rem32`, bits.Rem32)
	f.Set(`Rem64`, bits.Rem64)
	f.Set(`Reverse`, bits.Reverse)
	f.Set(`Reverse16`, bits.Reverse16)
	f.Set(`Reverse32`, bits.Reverse32)
	f.Set(`Reverse64`, bits.Reverse64)
	f.Set(`Reverse8`, bits.Reverse8)
	f.Set(`ReverseBytes`, bits.ReverseBytes)
	f.Set(`ReverseBytes16`, bits.ReverseBytes16)
	f.Set(`ReverseBytes32`, bits.ReverseBytes32)
	f.Set(`ReverseBytes64`, bits.ReverseBytes64)
	f.Set(`RotateLeft`, bits.RotateLeft)
	f.Set(`RotateLeft16`, bits.RotateLeft16)
	f.Set(`RotateLeft32`, bits.RotateLeft32)
	f.Set(`RotateLeft64`, bits.RotateLeft64)
	f.Set(`RotateLeft8`, bits.RotateLeft8)
	f.Set(`Sub`, bits.Sub)
	f.Set(`Sub32`, bits.Sub32)
	f.Set(`Sub64`, bits.Sub64)
	f.Set(`TrailingZeros`, bits.TrailingZeros)
	f.Set(`TrailingZeros16`, bits.TrailingZeros16)
	f.Set(`TrailingZeros32`, bits.TrailingZeros32)
	f.Set(`TrailingZeros64`, bits.TrailingZeros64)
	f.Set(`TrailingZeros8`, bits.TrailingZeros8)
}
func (f *factory) getUintSize(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(bits.UintSize)
}
