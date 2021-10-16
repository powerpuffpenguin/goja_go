package pem

import "encoding/pem"

func (f *factory) register() {
	f.Set(`Encode`, pem.Encode)
	f.Set(`EncodeToMemory`, pem.EncodeToMemory)

	f.Set(`Decode`, pem.Decode)

	f.Set(`isBlockPointer`, isBlockPointer)
}
func isBlockPointer(i interface{}) bool {
	_, result := i.(*pem.Block)
	return result
}
