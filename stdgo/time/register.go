package time

import (
	"time"

	"github.com/powerpuffpenguin/goja"
)

func (f *factory) register() {
	f.Accessor(`Layout`, f.getLayout, nil)
	f.Accessor(`ANSIC`, f.getANSIC, nil)
	f.Accessor(`UnixDate`, f.getUnixDate, nil)
	f.Accessor(`RubyDate`, f.getRubyDate, nil)
	f.Accessor(`RFC822`, f.getRFC822, nil)
	f.Accessor(`RFC822Z`, f.getRFC822Z, nil)
	f.Accessor(`RFC850`, f.getRFC850, nil)
	f.Accessor(`RFC1123`, f.getRFC1123, nil)
	f.Accessor(`RFC1123Z`, f.getRFC1123Z, nil)
	f.Accessor(`RFC3339`, f.getRFC3339, nil)
	f.Accessor(`RFC3339Nano`, f.getRFC3339Nano, nil)
	f.Accessor(`Kitchen`, f.getKitchen, nil)
	f.Accessor(`Stamp`, f.getStamp, nil)
	f.Accessor(`StampMilli`, f.getStampMilli, nil)
	f.Accessor(`StampMicro`, f.getStampMicro, nil)
	f.Accessor(`StampNano`, f.getStampNano, nil)

	f.Accessor(`Nanosecond`, f.getNanosecond, nil)
	f.Accessor(`Microsecond`, f.getMicrosecond, nil)
	f.Accessor(`Millisecond`, f.getMillisecond, nil)
	f.Accessor(`Second`, f.getSecond, nil)
	f.Accessor(`Minute`, f.getMinute, nil)
	f.Accessor(`Hour`, f.getHour, nil)

	f.Set(`After`, time.After)
	f.Set(`Sleep`, time.Sleep)
	f.Set(`Tick`, time.Tick)

	f.Set(`Month`, Month)
	f.Accessor(`January`, f.getJanuary, nil)
	f.Accessor(`February`, f.getFebruary, nil)
	f.Accessor(`March`, f.getMarch, nil)
	f.Accessor(`April`, f.getApril, nil)
	f.Accessor(`May`, f.getMay, nil)
	f.Accessor(`June`, f.getJune, nil)
	f.Accessor(`July`, f.getJuly, nil)
	f.Accessor(`August`, f.getAugust, nil)
	f.Accessor(`September`, f.getSeptember, nil)
	f.Accessor(`October`, f.getOctober, nil)
	f.Accessor(`November`, f.getNovember, nil)
	f.Accessor(`December`, f.getDecember, nil)

	f.Set(`FixedZone`, time.FixedZone)
	f.Set(`LoadLocation`, time.LoadLocation)
	f.Set(`LoadLocationFromTZData`, time.LoadLocationFromTZData)

	f.Set(`NewTicker`, time.NewTicker)

	f.Set(`AfterFunc`, time.AfterFunc)
	f.Set(`NewTimer`, time.NewTimer)

	f.Set(`Weekday`, Weekday)
	f.Accessor(`Sunday`, f.getSunday, nil)
	f.Accessor(`Monday`, f.getMonday, nil)
	f.Accessor(`Tuesday`, f.getTuesday, nil)
	f.Accessor(`Wednesday`, f.getWednesday, nil)
	f.Accessor(`Thursday`, f.getThursday, nil)
	f.Accessor(`Friday`, f.getFriday, nil)
	f.Accessor(`Saturday`, f.getSaturday, nil)

	f.Set(`isDuration`, isDuration)
	f.Set(`isLocationPointer`, isLocationPointer)
	f.Set(`isMonth`, isMonth)
	f.Set(`isParseErrorPointer`, isParseErrorPointer)
	f.Set(`isTicker`, isTickerPointer)
	f.Set(`isTime`, isTime)
	f.Set(`isTimerPointer`, isTimerPointer)
	f.Set(`isWeekday`, isWeekday)
}
func Month(val int) time.Month {
	return time.Month(val)
}
func Weekday(val int) time.Weekday {
	return time.Weekday(val)
}

func isDuration(i interface{}) bool {
	_, result := i.(time.Duration)
	return result
}
func isLocationPointer(i interface{}) bool {
	_, result := i.(*time.Location)
	return result
}
func isMonth(i interface{}) bool {
	_, result := i.(time.Month)
	return result
}
func isParseErrorPointer(i interface{}) bool {
	_, result := i.(*time.ParseError)
	return result
}
func isTickerPointer(i interface{}) bool {
	_, result := i.(*time.Ticker)
	return result
}
func isTime(i interface{}) bool {
	_, result := i.(time.Time)
	return result
}
func isTimerPointer(i interface{}) bool {
	_, result := i.(*time.Timer)
	return result
}
func isWeekday(i interface{}) bool {
	_, result := i.(time.Weekday)
	return result
}

func (f *factory) getSunday(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(time.Sunday)
}
func (f *factory) getMonday(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(time.Monday)
}
func (f *factory) getTuesday(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(time.Tuesday)
}
func (f *factory) getWednesday(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(time.Wednesday)
}
func (f *factory) getThursday(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(time.Thursday)
}
func (f *factory) getFriday(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(time.Friday)
}
func (f *factory) getSaturday(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(time.Saturday)
}
func (f *factory) getLayout(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue("01/02 03:04:05PM '06 -0700")
}
func (f *factory) getANSIC(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(time.ANSIC)
}
func (f *factory) getUnixDate(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(time.UnixDate)
}
func (f *factory) getRubyDate(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(time.RubyDate)
}
func (f *factory) getRFC822(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(time.RFC822)
}
func (f *factory) getRFC822Z(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(time.RFC822Z)
}
func (f *factory) getRFC850(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(time.RFC850)
}
func (f *factory) getRFC1123(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(time.RFC1123)
}
func (f *factory) getRFC1123Z(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(time.RFC1123Z)
}
func (f *factory) getRFC3339(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(time.RFC3339)
}
func (f *factory) getRFC3339Nano(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(time.RFC3339Nano)
}
func (f *factory) getKitchen(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(time.Kitchen)
}
func (f *factory) getStamp(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(time.Stamp)
}
func (f *factory) getStampMilli(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(time.StampMilli)
}
func (f *factory) getStampMicro(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(time.StampMicro)
}
func (f *factory) getStampNano(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(time.StampNano)
}
func (f *factory) getNanosecond(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(time.Nanosecond)
}
func (f *factory) getMicrosecond(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(time.Microsecond)
}
func (f *factory) getMillisecond(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(time.Millisecond)
}
func (f *factory) getSecond(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(time.Second)
}
func (f *factory) getMinute(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(time.Minute)
}
func (f *factory) getHour(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(time.Hour)
}
func (f *factory) getJanuary(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(time.January)
}
func (f *factory) getFebruary(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(time.February)
}
func (f *factory) getMarch(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(time.March)
}
func (f *factory) getApril(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(time.April)
}
func (f *factory) getMay(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(time.May)
}
func (f *factory) getJune(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(time.June)
}
func (f *factory) getJuly(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(time.July)
}
func (f *factory) getAugust(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(time.August)
}
func (f *factory) getSeptember(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(time.September)
}
func (f *factory) getOctober(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(time.October)
}
func (f *factory) getNovember(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(time.November)
}
func (f *factory) getDecember(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(time.December)
}
