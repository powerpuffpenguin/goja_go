package regexp

import "regexp"

func (f *factory) register() {
	f.Set(`Match`, regexp.Match)
	f.Set(`MatchReader`, regexp.MatchReader)
	f.Set(`MatchString`, regexp.MatchString)
	f.Set(`QuoteMeta`, regexp.QuoteMeta)

	f.Set(`Compile`, regexp.Compile)
	f.Set(`CompilePOSIX`, regexp.CompilePOSIX)
	f.Set(`MustCompile`, regexp.MustCompile)
	f.Set(`MustCompilePOSIX`, regexp.MustCompilePOSIX)

	f.Set(`isRegexpPointer`, isRegexpPointer)
}
func isRegexpPointer(i interface{}) bool {
	_, result := i.(*regexp.Regexp)
	return result
}
