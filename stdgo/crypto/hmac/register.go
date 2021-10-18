package hmac

import "crypto/hmac"

func (f *factory) register() {
	f.Set(`Equal`, hmac.Equal)
	f.Set(`New`, hmac.New)
}
