package time

import (
	"time"
)

func (f *factory) registerDuration() {
	f.Set(`Duration`, Duration)
	f.Set(`AddDuration`, AddDuration)
	f.Set(`SubDuration`, SubDuration)
	f.Set(`ParseDuration`, time.ParseDuration)
	f.Set(`Since`, time.Since)
	f.Set(`Until`, time.Until)
}
func Duration(val int64) time.Duration {
	return time.Duration(val)
}
func AddDuration(result time.Duration, vals ...time.Duration) time.Duration {
	for _, val := range vals {
		result += val
	}
	return result
}

func SubDuration(result time.Duration, vals ...time.Duration) time.Duration {
	for _, val := range vals {
		result -= val
	}
	return result
}
