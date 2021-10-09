package path

import (
	"path"

	"github.com/powerpuffpenguin/goja"
)

func (f *factory) register() {
	f.Accessor(`ErrBadPattern`, f.getErrBadPattern, nil)

	f.Set(`Base`, path.Base)
	f.Set(`Clean`, path.Clean)
	f.Set(`Dir`, path.Dir)
	f.Set(`Ext`, path.Ext)
	f.Set(`IsAbs`, path.IsAbs)
	f.Set(`Join`, path.Join)
	f.Set(`Match`, path.Match)
	f.Set(`Split`, path.Split)
}

func (f *factory) getErrBadPattern(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(path.ErrBadPattern)
}
