package gob

import "encoding/gob"

func (f *factory) register() {
	f.Set(`Register`, gob.Register)
	f.Set(`RegisterName`, gob.RegisterName)

	f.Set(`NewDecoder`, gob.NewDecoder)
	f.Set(`NewEncoder`, gob.NewEncoder)

	f.Set(`isCommonType`, isCommonType)
	f.Set(`isDecoderPointer`, isDecoderPointer)
	f.Set(`isEncoderPointer`, isEncoderPointer)
	f.Set(`isGobDecoder`, isGobDecoder)
	f.Set(`isGobEncoder`, isGobEncoder)
}
func isCommonType(i interface{}) bool {
	_, result := i.(gob.CommonType)
	return result
}
func isDecoderPointer(i interface{}) bool {
	_, result := i.(*gob.Decoder)
	return result
}
func isEncoderPointer(i interface{}) bool {
	_, result := i.(*gob.Encoder)
	return result
}
func isGobDecoder(i interface{}) bool {
	_, result := i.(gob.GobDecoder)
	return result
}
func isGobEncoder(i interface{}) bool {
	_, result := i.(gob.GobEncoder)
	return result
}
