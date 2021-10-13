package strings

import (
	"strings"

	"github.com/powerpuffpenguin/goja_go/stdgo/builtin"
)

func (f *factory) register() {
	f.Set(`At`, At)
	f.Set(`Foreach`, Foreach)
	f.Set(`String`, String)

	f.Set(`Compare`, strings.Compare)
	f.Set(`Contains`, strings.Contains)
	f.Set(`ContainsAny`, strings.ContainsAny)
	f.Set(`ContainsRune`, strings.ContainsRune)
	f.Set(`Count`, strings.Count)
	f.Set(`EqualFold`, strings.EqualFold)
	f.Set(`Fields`, strings.Fields)
	f.Set(`FieldsFunc`, strings.FieldsFunc)
	f.Set(`HasPrefix`, strings.HasPrefix)
	f.Set(`HasSuffix`, strings.HasSuffix)
	f.Set(`Index`, strings.Index)
	f.Set(`IndexAny`, strings.IndexAny)
	f.Set(`IndexByte`, strings.IndexByte)
	f.Set(`IndexFunc`, strings.IndexFunc)
	f.Set(`IndexRune`, strings.IndexRune)
	f.Set(`Join`, strings.Join)
	f.Set(`LastIndex`, strings.LastIndex)
	f.Set(`LastIndexAny`, strings.LastIndexAny)
	f.Set(`LastIndexByte`, strings.LastIndexByte)
	f.Set(`LastIndexFunc`, strings.LastIndexFunc)
	f.Set(`Map`, strings.Map)
	f.Set(`Repeat`, strings.Repeat)
	f.Set(`Replace`, strings.Replace)
	f.Set(`ReplaceAll`, strings.ReplaceAll)
	f.Set(`Split`, strings.Split)
	f.Set(`SplitAfter`, strings.SplitAfter)
	f.Set(`SplitAfterN`, strings.SplitAfterN)
	f.Set(`SplitN`, strings.SplitN)
	f.Set(`Title`, strings.Title)
	f.Set(`ToLower`, strings.ToLower)
	f.Set(`ToLowerSpecial`, strings.ToLowerSpecial)
	f.Set(`ToTitle`, strings.ToTitle)
	f.Set(`ToTitleSpecial`, strings.ToTitleSpecial)
	f.Set(`ToUpper`, strings.ToUpper)
	f.Set(`ToUpperSpecial`, strings.ToUpperSpecial)
	f.Set(`ToValidUTF8`, strings.ToValidUTF8)
	f.Set(`Trim`, strings.Trim)
	f.Set(`TrimFunc`, strings.TrimFunc)
	f.Set(`TrimLeft`, strings.TrimLeft)
	f.Set(`TrimLeftFunc`, strings.TrimLeftFunc)
	f.Set(`TrimPrefix`, strings.TrimPrefix)
	f.Set(`TrimRight`, strings.TrimRight)
	f.Set(`TrimRightFunc`, strings.TrimRightFunc)
	f.Set(`TrimSpace`, strings.TrimSpace)
	f.Set(`TrimSuffix`, strings.TrimSuffix)

	f.Set(`Builder`, Builder)
	f.Set(`NewReader`, strings.NewReader)
	f.Set(`NewReplacer`, strings.NewReplacer)

	f.Set(`isBuilderPointer`, isBuilderPointer)
	f.Set(`isReaderPointer`, isReaderPointer)
	f.Set(`isReplacerPointer`, isReplacerPointer)
}
func isBuilderPointer(i interface{}) bool {
	_, result := i.(*strings.Builder)
	return result
}
func isReaderPointer(i interface{}) bool {
	_, result := i.(*strings.Reader)
	return result
}
func isReplacerPointer(i interface{}) bool {
	_, result := i.(*strings.Replacer)
	return result
}

func Builder() *strings.Builder {
	return &strings.Builder{}
}
func At(s string, index int) rune {
	return rune(s[index])
}
func Foreach(s string, f func(rune, builtin.Int) bool) {
	for i, r := range s {
		if !f(r, builtin.Int(i)) {
			break
		}
	}
}
func String(r rune) string {
	return string(r)
}
