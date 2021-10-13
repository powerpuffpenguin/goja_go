package hash

import "hash"

func (f *factory) register() {
	f.Set(`isHash`, isHash)
	f.Set(`isHash32`, isHash32)
	f.Set(`isHash64`, isHash64)
}
func isHash(i interface{}) bool {
	_, result := i.(hash.Hash)
	return result
}
func isHash32(i interface{}) bool {
	_, result := i.(hash.Hash32)
	return result
}
func isHash64(i interface{}) bool {
	_, result := i.(hash.Hash64)
	return result
}
