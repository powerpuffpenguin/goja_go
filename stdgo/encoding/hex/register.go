package hex

import (
	"encoding/hex"

	"github.com/powerpuffpenguin/goja"
)

func (f *factory) register() {
	f.Accessor(`ErrLength`, f.getErrLength, nil)

	f.Set(`Decode`, hex.Decode)
	f.Set(`DecodeString`, hex.DecodeString)
	f.Set(`DecodedLen`, hex.DecodedLen)
	f.Set(`Dump`, hex.Dump)
	f.Set(`Dumper`, hex.Dumper)
	f.Set(`Encode`, hex.Encode)
	f.Set(`EncodeToString`, hex.EncodeToString)
	f.Set(`EncodedLen`, hex.EncodedLen)
	f.Set(`NewDecoder`, hex.NewDecoder)
	f.Set(`NewEncoder`, hex.NewEncoder)

	f.Set(`isInvalidByteError`, isInvalidByteError)
}
func (f *factory) getErrLength(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(hex.ErrLength)
}
func isInvalidByteError(i interface{}) bool {
	_, result := i.(hex.InvalidByteError)
	return result
}
