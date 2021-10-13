package maphash

import (
	"hash/maphash"
)

func (f *factory) register() {
	f.Set(`MakeSeed`, maphash.MakeSeed)

	f.Set(`isHashPointer`, isHashPointer)
	f.Set(`isSeed`, isSeed)
}
func isHashPointer(i interface{}) bool {
	_, result := i.(*maphash.Hash)
	return result
}
func isSeed(i interface{}) bool {
	_, result := i.(maphash.Seed)
	return result
}
