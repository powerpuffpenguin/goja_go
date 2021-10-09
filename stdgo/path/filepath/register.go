package filepath

import (
	"path/filepath"

	"github.com/powerpuffpenguin/goja"
	"github.com/powerpuffpenguin/goja_go/stdgo/builtin"
)

func (f *factory) register() {
	f.Accessor(`Separator`, f.getSeparator, nil)
	f.Accessor(`ListSeparator`, f.getListSeparator, nil)

	f.Accessor(`ErrBadPattern`, f.getErrBadPattern, nil)
	f.Accessor(`SkipDir`, f.getSkipDir, nil)

	f.Set(`Abs`, filepath.Abs)
	f.Set(`Base`, filepath.Base)
	f.Set(`Clean`, filepath.Clean)
	f.Set(`Dir`, filepath.Dir)
	f.Set(`EvalSymlinks`, filepath.EvalSymlinks)
	f.Set(`Ext`, filepath.Ext)
	f.Set(`FromSlash`, filepath.FromSlash)
	f.Set(`Glob`, filepath.Glob)

	f.Set(`IsAbs`, filepath.IsAbs)
	f.Set(`Join`, filepath.Join)
	f.Set(`Match`, filepath.Match)
	f.Set(`Rel`, filepath.Rel)
	f.Set(`Split`, filepath.Split)
	f.Set(`SplitList`, filepath.SplitList)
	f.Set(`ToSlash`, filepath.ToSlash)
	f.Set(`VolumeName`, filepath.VolumeName)
	f.Set(`Walk`, filepath.Walk)
	f.Set(`WalkDir`, filepath.WalkDir)

	f.Set(`isWalkFunc`, isWalkFunc)
}
func isWalkFunc(i interface{}) bool {
	_, result := i.(filepath.WalkFunc)
	return result
}
func (f *factory) getErrBadPattern(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(filepath.ErrBadPattern)
}
func (f *factory) getSkipDir(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(filepath.SkipDir)
}
func (f *factory) getSeparator(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int32(filepath.Separator))
}
func (f *factory) getListSeparator(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int32(filepath.ListSeparator))
}
