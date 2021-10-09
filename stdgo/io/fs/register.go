package fs

import (
	"io/fs"

	"github.com/powerpuffpenguin/goja"
)

func (f *factory) register() {
	f.Accessor(`ErrInvalid`, f.getErrInvalid, nil)
	f.Accessor(`ErrPermission`, f.getErrPermission, nil)
	f.Accessor(`ErrExist`, f.getErrExist, nil)
	f.Accessor(`ErrNotExist`, f.getErrNotExist, nil)
	f.Accessor(`ErrClosed`, f.getErrClosed, nil)

	f.Accessor(`SkipDir`, f.getSkipDir, nil)

	f.Set(`Glob`, fs.Glob)
	f.Set(`ReadFile`, fs.ReadFile)
	f.Set(`ValidPath`, fs.ValidPath)
	f.Set(`WalkDir`, fs.WalkDir)

	f.Set(`ReadDir`, fs.ReadDir)
	f.Set(`Sub`, fs.Sub)
	f.Set(`Stat`, fs.Stat)

	f.Set(`isDirEntry`, isDirEntry)
	f.Set(`isFS`, isFS)
	f.Set(`isFile`, isFile)
	f.Set(`isFileInfo`, isFileInfo)
	f.Set(`isFileMode`, isFileMode)
	f.Set(`isGlobFS`, isGlobFS)
	f.Set(`isPathErrorPointer`, isPathErrorPointer)
	f.Set(`isReadDirFS`, isReadDirFS)
	f.Set(`isReadDirFile`, isReadDirFile)
	f.Set(`isReadFileFS`, isReadFileFS)
	f.Set(`isStatFS`, isStatFS)
	f.Set(`isSubFS`, isSubFS)
	f.Set(`isWalkDirFunc`, isWalkDirFunc)
}

func isDirEntry(i interface{}) bool {
	_, result := i.(fs.DirEntry)
	return result
}
func isFS(i interface{}) bool {
	_, result := i.(fs.FS)
	return result
}
func isFile(i interface{}) bool {
	_, result := i.(fs.File)
	return result
}
func isFileInfo(i interface{}) bool {
	_, result := i.(fs.FileInfo)
	return result
}
func isFileMode(i interface{}) bool {
	_, result := i.(fs.FileMode)
	return result
}
func isGlobFS(i interface{}) bool {
	_, result := i.(fs.GlobFS)
	return result
}
func isPathErrorPointer(i interface{}) bool {
	_, result := i.(*fs.PathError)
	return result
}
func isReadDirFS(i interface{}) bool {
	_, result := i.(fs.ReadDirFS)
	return result
}
func isReadDirFile(i interface{}) bool {
	_, result := i.(fs.ReadDirFile)
	return result
}
func isReadFileFS(i interface{}) bool {
	_, result := i.(fs.ReadFileFS)
	return result
}
func isStatFS(i interface{}) bool {
	_, result := i.(fs.StatFS)
	return result
}
func isSubFS(i interface{}) bool {
	_, result := i.(fs.SubFS)
	return result
}
func isWalkDirFunc(i interface{}) bool {
	_, result := i.(fs.WalkDirFunc)
	return result
}

func (f *factory) getSkipDir(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(fs.SkipDir)
}
func (f *factory) getErrInvalid(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(fs.ErrInvalid)
}
func (f *factory) getErrPermission(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(fs.ErrPermission)
}
func (f *factory) getErrExist(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(fs.ErrExist)
}
func (f *factory) getErrNotExist(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(fs.ErrNotExist)
}
func (f *factory) getErrClosed(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(fs.ErrClosed)
}
