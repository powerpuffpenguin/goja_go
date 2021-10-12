package url

import (
	"net/url"
)

func (f *factory) register() {
	f.Set(`PathEscape`, url.PathEscape)
	f.Set(`PathUnescape`, url.PathUnescape)
	f.Set(`QueryEscape`, url.QueryEscape)
	f.Set(`QueryUnescape`, url.QueryUnescape)

	f.Set(`Parse`, url.Parse)
	f.Set(`ParseRequestURI`, url.ParseRequestURI)

	f.Set(`User`, url.User)
	f.Set(`UserPassword`, url.UserPassword)

	f.Set(`ParseQuery`, url.ParseQuery)
	f.Set(`Values`, Values)

	f.Set(`isErrorPointer`, isErrorPointer)
	f.Set(`isEscapeError`, isEscapeError)
	f.Set(`isInvalidHostError`, isInvalidHostError)
	f.Set(`isURLPointer`, isURLPointer)
	f.Set(`isUserinfoPointer`, isUserinfoPointer)
	f.Set(`isValues`, isValues)
}
func isErrorPointer(i interface{}) bool {
	_, result := i.(*url.Error)
	return result
}
func isEscapeError(i interface{}) bool {
	_, result := i.(url.EscapeError)
	return result
}
func isInvalidHostError(i interface{}) bool {
	_, result := i.(url.InvalidHostError)
	return result
}
func isURLPointer(i interface{}) bool {
	_, result := i.(*url.URL)
	return result
}
func isUserinfoPointer(i interface{}) bool {
	_, result := i.(*url.Userinfo)
	return result
}
func isValues(i interface{}) bool {
	_, result := i.(url.Values)
	return result
}
func Values() url.Values {
	return make(url.Values)
}
