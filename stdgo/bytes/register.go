package bytes

import (
	"bytes"

	"github.com/powerpuffpenguin/goja"
)

func (f *factory) register() {
	f.Set(`Bytes`, Bytes)
	f.Set(`String`, String)

	f.Accessor(`MinRead`, f.getMinRead, nil)
	f.Accessor(`ErrTooLarge`, f.getErrTooLarge, nil)

	f.Set(`Compare`, bytes.Compare)
	f.Set(`Contains`, bytes.Contains)
	f.Set(`ContainsAny`, bytes.ContainsAny)
	f.Set(`ContainsRune`, bytes.ContainsRune)
	f.Set(`Count`, bytes.Count)
	f.Set(`Equal`, bytes.Equal)
	f.Set(`EqualFold`, bytes.EqualFold)
	f.Set(`Fields`, bytes.Fields)
	f.Set(`FieldsFunc`, bytes.FieldsFunc)
	f.Set(`HasPrefix`, bytes.HasPrefix)
	f.Set(`HasSuffix`, bytes.HasSuffix)
	f.Set(`Index`, bytes.Index)
	f.Set(`IndexAny`, bytes.IndexAny)
	f.Set(`IndexByte`, bytes.IndexByte)
	f.Set(`IndexFunc`, bytes.IndexFunc)
	f.Set(`IndexRune`, bytes.IndexRune)
	f.Set(`Join`, bytes.Join)
	f.Set(`LastIndex`, bytes.LastIndex)
	f.Set(`LastIndexAny`, bytes.LastIndexAny)
	f.Set(`LastIndexByte`, bytes.LastIndexByte)
	f.Set(`LastIndexFunc`, bytes.LastIndexFunc)
	f.Set(`Map`, bytes.Map)
	f.Set(`Repeat`, bytes.Repeat)
	f.Set(`Replace`, bytes.Replace)
	f.Set(`ReplaceAll`, bytes.ReplaceAll)
	f.Set(`Runes`, bytes.Runes)
	f.Set(`Split`, bytes.Split)
	f.Set(`SplitAfter`, bytes.SplitAfter)
	f.Set(`SplitAfterN`, bytes.SplitAfterN)
	f.Set(`SplitN`, bytes.SplitN)
	f.Set(`Title`, bytes.Title)
	f.Set(`ToLower`, bytes.ToLower)
	f.Set(`ToLowerSpecial`, bytes.ToLowerSpecial)
	f.Set(`ToTitle`, bytes.ToTitle)
	f.Set(`ToTitleSpecial`, bytes.ToTitleSpecial)
	f.Set(`ToUpper`, bytes.ToUpper)
	f.Set(`ToUpperSpecial`, bytes.ToUpperSpecial)
	f.Set(`ToValidUTF8`, bytes.ToValidUTF8)
	f.Set(`Trim`, bytes.Trim)
	f.Set(`TrimFunc`, bytes.TrimFunc)
	f.Set(`TrimLeft`, bytes.TrimLeft)
	f.Set(`TrimLeftFunc`, bytes.TrimLeftFunc)
	f.Set(`TrimPrefix`, bytes.TrimPrefix)
	f.Set(`TrimRight`, bytes.TrimRight)
	f.Set(`TrimRightFunc`, bytes.TrimRightFunc)
	f.Set(`TrimSpace`, bytes.TrimSpace)
	f.Set(`TrimSuffix`, bytes.TrimSuffix)

	f.Set(`NewBuffer`, bytes.NewBuffer)
	f.Set(`NewBufferString`, bytes.NewBufferString)

	f.Set(`NewReader`, bytes.NewReader)

	f.Set(`isBufferPointer`, isBufferPointer)
	f.Set(`isReaderPointer`, isReaderPointer)
}
func Bytes(str string) []byte {
	return []byte(str)
}
func String(b []byte) string {
	return string(b)
}

func isBufferPointer(i interface{}) bool {
	_, result := i.(*bytes.Buffer)
	return result
}
func isReaderPointer(i interface{}) bool {
	_, result := i.(*bytes.Reader)
	return result
}
func (f *factory) getMinRead(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(bytes.MinRead)
}
func (f *factory) getErrTooLarge(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(bytes.ErrTooLarge)
}
