package rand

import "math/rand"

func (f *factory) register() {
	f.Set(`ExpFloat64`, rand.ExpFloat64)
	f.Set(`Float32`, rand.Float32)
	f.Set(`Float64`, rand.Float64)
	f.Set(`Int`, rand.Int)
	f.Set(`Int31`, rand.Int31)
	f.Set(`Int31n`, rand.Int31n)
	f.Set(`Int63`, rand.Int63)
	f.Set(`Int63n`, rand.Int63n)
	f.Set(`Intn`, rand.Intn)
	f.Set(`NormFloat64`, rand.NormFloat64)
	f.Set(`Perm`, rand.Perm)
	f.Set(`Read`, rand.Read)
	f.Set(`Seed`, rand.Seed)
	f.Set(`Shuffle`, rand.Shuffle)
	f.Set(`Uint32`, rand.Uint32)
	f.Set(`Uint64`, rand.Uint64)

	f.Set(`New`, rand.New)
	f.Set(`NewSource`, rand.NewSource)
	f.Set(`NewZipf`, rand.NewZipf)

	f.Set(`isRandPointer`, isRandPointer)
	f.Set(`isSource`, isSource)
	f.Set(`isSource64`, isSource64)
	f.Set(`isZipfPointer`, isZipfPointer)
}
func isRandPointer(i interface{}) bool {
	_, result := i.(*rand.Rand)
	return result
}
func isSource(i interface{}) bool {
	_, result := i.(rand.Source)
	return result
}
func isSource64(i interface{}) bool {
	_, result := i.(rand.Source64)
	return result
}
func isZipfPointer(i interface{}) bool {
	_, result := i.(*rand.Zipf)
	return result
}
