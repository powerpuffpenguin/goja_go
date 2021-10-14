package suffixarray

import "index/suffixarray"

func (f *factory) register() {
	f.Set(`New`, suffixarray.New)

	f.Set(`isIndexPointer`, isIndexPointer)
}
func isIndexPointer(i interface{}) bool {
	_, result := i.(*suffixarray.Index)
	return result
}
