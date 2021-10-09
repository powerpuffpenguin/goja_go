package fs

import (
	"io/fs"

	"github.com/powerpuffpenguin/goja"
)

func (f *factory) registerFileMode() {
	f.Accessor(`ModeDir`, f.getModeDir, nil)
	f.Accessor(`ModeAppend`, f.getModeAppend, nil)
	f.Accessor(`ModeExclusive`, f.getModeExclusive, nil)
	f.Accessor(`ModeTemporary`, f.getModeTemporary, nil)
	f.Accessor(`ModeSymlink`, f.getModeSymlink, nil)
	f.Accessor(`ModeDevice`, f.getModeDevice, nil)
	f.Accessor(`ModeNamedPipe`, f.getModeNamedPipe, nil)
	f.Accessor(`ModeSocket`, f.getModeSocket, nil)
	f.Accessor(`ModeSetuid`, f.getModeSetuid, nil)
	f.Accessor(`ModeSetgid`, f.getModeSetgid, nil)
	f.Accessor(`ModeCharDevice`, f.getModeCharDevice, nil)
	f.Accessor(`ModeSticky`, f.getModeSticky, nil)
	f.Accessor(`ModeIrregular`, f.getModeIrregular, nil)
	f.Accessor(`ModeType`, f.getModeType, nil)
	f.Accessor(`ModePerm`, f.getModePerm, nil)

	f.Set(`FileMode`, FileMode)
}
func FileMode(v uint32) fs.FileMode {
	return fs.FileMode(v)
}
func (f *factory) getModeDir(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(fs.ModeDir)
}
func (f *factory) getModeAppend(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(fs.ModeAppend)
}
func (f *factory) getModeExclusive(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(fs.ModeExclusive)
}
func (f *factory) getModeTemporary(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(fs.ModeTemporary)
}
func (f *factory) getModeSymlink(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(fs.ModeSymlink)
}
func (f *factory) getModeDevice(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(fs.ModeDevice)
}
func (f *factory) getModeNamedPipe(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(fs.ModeNamedPipe)
}
func (f *factory) getModeSocket(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(fs.ModeSocket)
}
func (f *factory) getModeSetuid(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(fs.ModeSetuid)
}
func (f *factory) getModeSetgid(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(fs.ModeSetgid)
}
func (f *factory) getModeCharDevice(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(fs.ModeCharDevice)
}
func (f *factory) getModeSticky(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(fs.ModeSticky)
}
func (f *factory) getModeIrregular(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(fs.ModeIrregular)
}
func (f *factory) getModeType(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(fs.ModeType)
}
func (f *factory) getModePerm(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(fs.ModePerm)
}
