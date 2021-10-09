package ioutil

import (
	"io/ioutil"

	"github.com/powerpuffpenguin/goja"
)

func (f *factory) register() {
	f.Accessor(`Discard`, f.getDiscard, nil)
	f.Set(`NopCloser`, ioutil.NopCloser)
	f.Set(`ReadAll`, ioutil.ReadAll)
	f.Set(`ReadDir`, ioutil.ReadDir)
	f.Set(`ReadFile`, ioutil.ReadFile)
	f.Set(`TempDir`, ioutil.TempDir)
	f.Set(`TempFile`, ioutil.TempFile)
	f.Set(`WriteFile`, ioutil.WriteFile)
}

func (f *factory) getDiscard(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(ioutil.Discard)
}
