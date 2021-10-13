package fnv

import "hash/fnv"

func (f *factory) register() {
	f.Set(`New128`, fnv.New128)
	f.Set(`New128a`, fnv.New128a)
	f.Set(`New32`, fnv.New32)
	f.Set(`New32a`, fnv.New32a)
	f.Set(`New64`, fnv.New64)
	f.Set(`New64a`, fnv.New64a)
}
