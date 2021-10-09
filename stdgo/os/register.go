package os

import (
	"io"
	"io/fs"
	"os"

	"github.com/powerpuffpenguin/goja"
	"github.com/powerpuffpenguin/goja_go/stdgo/builtin"
)

func (f *factory) register() {
	f.Accessor(`O_RDONLY`, f.getO_RDONLY, nil)
	f.Accessor(`O_WRONLY`, f.getO_WRONLY, nil)
	f.Accessor(`O_RDWR`, f.getO_RDWR, nil)
	f.Accessor(`O_APPEND`, f.getO_APPEND, nil)
	f.Accessor(`O_CREATE`, f.getO_CREATE, nil)
	f.Accessor(`O_EXCL`, f.getO_EXCL, nil)
	f.Accessor(`O_SYNC`, f.getO_SYNC, nil)
	f.Accessor(`O_TRUNC`, f.getO_TRUNC, nil)

	f.Accessor(`SEEK_SET`, f.getSEEK_SET, nil)
	f.Accessor(`SEEK_CUR`, f.getSEEK_CUR, nil)
	f.Accessor(`SEEK_END`, f.getSEEK_END, nil)

	f.Accessor(`PathSeparator`, f.getPathSeparator, nil)
	f.Accessor(`PathListSeparator`, f.getPathListSeparator, nil)

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

	f.Accessor(`DevNull`, f.getDevNull, nil)

	f.Accessor(`ErrInvalid`, f.getErrInvalid, nil)
	f.Accessor(`ErrPermission`, f.getErrPermission, nil)
	f.Accessor(`ErrExist`, f.getErrExist, nil)
	f.Accessor(`ErrNotExist`, f.getErrNotExist, nil)
	f.Accessor(`ErrClosed`, f.getErrClosed, nil)

	f.Accessor(`ErrNoDeadline`, f.getErrNoDeadline, nil)
	f.Accessor(`ErrDeadlineExceeded`, f.getErrDeadlineExceeded, nil)

	f.Accessor(`Stdin`, f.getStdin, nil)
	f.Accessor(`Stdout`, f.getStdout, nil)
	f.Accessor(`Stderr`, f.getStderr, nil)

	f.Accessor(`Args`, f.getArgs, nil)
	f.Accessor(`ErrProcessDone`, f.getErrProcessDone, nil)

	f.Set(`Chdir`, os.Chdir)
	f.Set(`Chmod`, os.Chmod)
	f.Set(`Chown`, os.Chown)
	f.Set(`Chtimes`, os.Chtimes)
	f.Set(`Clearenv`, os.Clearenv)
	f.Set(`DirFS`, os.DirFS)
	f.Set(`Environ`, os.Environ)
	f.Set(`Exit`, os.Exit)
	f.Set(`Expand`, os.Expand)
	f.Set(`ExpandEnv`, os.ExpandEnv)
	f.Set(`Getegid`, os.Getegid)
	f.Set(`Getenv`, os.Getenv)
	f.Set(`Geteuid`, os.Geteuid)
	f.Set(`Getgid`, os.Getgid)
	f.Set(`Getgroups`, os.Getgroups)
	f.Set(`Getpagesize`, os.Getpagesize)
	f.Set(`Getpid`, os.Getpid)
	f.Set(`Getppid`, os.Getppid)
	f.Set(`Getuid`, os.Getuid)
	f.Set(`Getwd`, os.Getwd)
	f.Set(`Hostname`, os.Hostname)
	f.Set(`IsExist`, os.IsExist)
	f.Set(`IsNotExist`, os.IsNotExist)
	f.Set(`IsPathSeparator`, os.IsPathSeparator)
	f.Set(`IsPermission`, os.IsPermission)
	f.Set(`IsTimeout`, os.IsTimeout)
	f.Set(`Lchown`, os.Lchown)
	f.Set(`Link`, os.Link)
	f.Set(`LookupEnv`, os.LookupEnv)
	f.Set(`Mkdir`, os.Mkdir)
	f.Set(`MkdirAll`, os.MkdirAll)
	f.Set(`MkdirTemp`, os.MkdirTemp)
	f.Set(`Pipe`, os.Pipe)
	f.Set(`NewSyscallError`, os.NewSyscallError)
	f.Set(`ReadFile`, os.ReadFile)
	f.Set(`Readlink`, os.Readlink)
	f.Set(`Remove`, os.Remove)
	f.Set(`RemoveAll`, os.RemoveAll)
	f.Set(`Rename`, os.Rename)
	f.Set(`SameFile`, os.SameFile)
	f.Set(`Setenv`, os.Setenv)
	f.Set(`Symlink`, os.Symlink)
	f.Set(`TempDir`, os.TempDir)
	f.Set(`Truncate`, os.Truncate)
	f.Set(`Unsetenv`, os.Unsetenv)
	f.Set(`UserCacheDir`, os.UserCacheDir)
	f.Set(`UserConfigDir`, os.UserConfigDir)
	f.Set(`UserHomeDir`, os.UserHomeDir)
	f.Set(`WriteFile`, os.WriteFile)

	f.Set(`ReadDir`, os.ReadDir)

	f.Set(`Create`, os.Create)
	f.Set(`CreateTemp`, os.CreateTemp)
	f.Set(`NewFile`, os.NewFile)
	f.Set(`Open`, os.Open)
	f.Set(`OpenFile`, os.OpenFile)

	f.Set(`Lstat`, os.Lstat)
	f.Set(`Stat`, os.Stat)

	f.Set(`NewProcAttr`, NewProcAttr)
	f.Set(`FindProcess`, os.FindProcess)
	f.Set(`StartProcess`, os.StartProcess)

	f.Accessor(`Interrupt`, f.getInterrupt, nil)
	f.Accessor(`Kill`, f.getKill, nil)

	f.Set(`isDirEntry`, isDirEntry)
	f.Set(`isFilePointer`, isFilePointer)
	f.Set(`isFileInfo`, isFileInfo)
	f.Set(`isFileMode`, isFileMode)
	f.Set(`isLinkErrorPointer`, isLinkErrorPointer)
	f.Set(`isPathErrorPointer`, isPathErrorPointer)
	f.Set(`isProcAttrPointer`, isProcAttrPointer)
	f.Set(`isProcessPointer`, isProcessPointer)
	f.Set(`isProcessStatePointer`, isProcessStatePointer)
	f.Set(`isSignal`, isSignal)
	f.Set(`isSyscallErrorPointer`, isSyscallErrorPointer)
}
func (f *factory) getInterrupt(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(os.Interrupt)
}
func (f *factory) getKill(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(os.Kill)
}
func NewProcAttr() *os.ProcAttr {
	return &os.ProcAttr{}
}
func isDirEntry(i interface{}) bool {
	_, result := i.(os.DirEntry)
	return result
}
func isFilePointer(i interface{}) bool {
	_, result := i.(*os.File)
	return result
}
func isFileInfo(i interface{}) bool {
	_, result := i.(os.FileInfo)
	return result
}
func isFileMode(i interface{}) bool {
	_, result := i.(os.FileMode)
	return result
}
func isLinkErrorPointer(i interface{}) bool {
	_, result := i.(*os.LinkError)
	return result
}
func isPathErrorPointer(i interface{}) bool {
	_, result := i.(*os.PathError)
	return result
}
func isProcAttrPointer(i interface{}) bool {
	_, result := i.(*os.ProcAttr)
	return result
}
func isProcessPointer(i interface{}) bool {
	_, result := i.(*os.Process)
	return result
}
func isProcessStatePointer(i interface{}) bool {
	_, result := i.(*os.ProcessState)
	return result
}
func isSignal(i interface{}) bool {
	_, result := i.(os.Signal)
	return result
}
func isSyscallErrorPointer(i interface{}) bool {
	_, result := i.(*os.SyscallError)
	return result
}

func (f *factory) getArgs(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(os.Args)
}
func (f *factory) getErrProcessDone(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(os.ErrProcessDone)
}
func (f *factory) getStdin(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(os.Stdin)
}
func (f *factory) getStdout(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(os.Stdout)
}
func (f *factory) getStderr(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(os.Stderr)
}
func (f *factory) getErrDeadlineExceeded(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(os.ErrDeadlineExceeded)
}
func (f *factory) getErrNoDeadline(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(os.ErrNoDeadline)
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
func (f *factory) getDevNull(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(os.DevNull)
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

func (f *factory) getPathSeparator(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int32(os.PathSeparator))
}
func (f *factory) getPathListSeparator(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int32(os.PathListSeparator))
}
func (f *factory) getSEEK_SET(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(io.SeekStart))
}
func (f *factory) getSEEK_CUR(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(io.SeekCurrent))
}
func (f *factory) getSEEK_END(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(io.SeekEnd))
}
func (f *factory) getO_RDONLY(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(os.O_RDONLY))
}
func (f *factory) getO_WRONLY(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(os.O_WRONLY))
}
func (f *factory) getO_RDWR(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(os.O_RDWR))
}
func (f *factory) getO_APPEND(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(os.O_APPEND))
}
func (f *factory) getO_CREATE(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(os.O_CREATE))
}
func (f *factory) getO_EXCL(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(os.O_EXCL))
}
func (f *factory) getO_SYNC(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(os.O_SYNC))
}
func (f *factory) getO_TRUNC(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(os.O_TRUNC))
}
