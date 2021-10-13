package bzip2

import (
	"compress/bzip2"
)

func (f *factory) register() {
	f.Set(`NewReader`, bzip2.NewReader)

	f.Set(`isStructuralError`, isStructuralError)
}
func isStructuralError(i interface{}) bool {
	_, result := i.(bzip2.StructuralError)
	return result
}
