package ascii85

import "encoding/ascii85"

func (f *factory) register() {
	f.Set(`Decode`, ascii85.Decode)
	f.Set(`Encode`, ascii85.Encode)
	f.Set(`MaxEncodedLen`, ascii85.MaxEncodedLen)
	f.Set(`NewDecoder`, ascii85.NewDecoder)
	f.Set(`NewEncoder`, ascii85.NewEncoder)

	f.Set(`isCorruptInputError`, isCorruptInputError)
}
func isCorruptInputError(i interface{}) bool {
	_, result := i.(ascii85.CorruptInputError)
	return result
}
