package elliptic

import "crypto/elliptic"

func (f *factory) register() {
	f.Set(`GenerateKey`, elliptic.GenerateKey)
	f.Set(`Marshal`, elliptic.Marshal)
	f.Set(`MarshalCompressed`, elliptic.MarshalCompressed)
	f.Set(`Unmarshal`, elliptic.Unmarshal)
	f.Set(`UnmarshalCompressed`, elliptic.UnmarshalCompressed)

	f.Set(`P224`, elliptic.P224)
	f.Set(`P256`, elliptic.P256)
	f.Set(`P384`, elliptic.P384)
	f.Set(`P521`, elliptic.P521)

	f.Set(`isCurve`, isCurve)
	f.Set(`isCurveParamsPointer`, isCurveParamsPointer)
}
func isCurve(i interface{}) bool {
	_, result := i.(elliptic.Curve)
	return result
}
func isCurveParamsPointer(i interface{}) bool {
	_, result := i.(*elliptic.CurveParams)
	return result
}
