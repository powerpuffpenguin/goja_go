package subtle

import "crypto/subtle"

func (f *factory) register() {
	f.Set(`ConstantTimeByteEq`, subtle.ConstantTimeByteEq)
	f.Set(`ConstantTimeCompare`, subtle.ConstantTimeCompare)
	f.Set(`ConstantTimeCopy`, subtle.ConstantTimeCopy)
	f.Set(`ConstantTimeEq`, subtle.ConstantTimeEq)
	f.Set(`ConstantTimeLessOrEq`, subtle.ConstantTimeLessOrEq)
	f.Set(`ConstantTimeSelect`, subtle.ConstantTimeSelect)
}
