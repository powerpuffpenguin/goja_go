package time

import (
	"time"
)

func (f *factory) registerTime() {
	f.Set(`Date`, time.Date)
	f.Set(`Now`, time.Now)
	f.Set(`Parse`, time.Parse)
	f.Set(`ParseInLocation`, time.ParseInLocation)
	f.Set(`Unix`, time.Unix)
	f.Set(`UnixMilli`, UnixMilli)
	f.Set(`UnixMicro`, UnixMicro)
}

// UnixMilli returns the local Time corresponding to the given Unix time,
// msec milliseconds since January 1, 1970 UTC.
func UnixMilli(msec int64) time.Time {
	return time.Unix(msec/1e3, (msec%1e3)*1e6)
}

// UnixMicro returns the local Time corresponding to the given Unix time,
// usec microseconds since January 1, 1970 UTC.
func UnixMicro(usec int64) time.Time {
	return time.Unix(usec/1e6, (usec%1e6)*1e3)
}
