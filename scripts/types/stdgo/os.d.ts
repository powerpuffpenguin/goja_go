declare module "stdgo/os" {
    import {
        Float32, Float64,
        Int64, Int32, Int16, Int8, Int,
        Uint64, Uint32, Uint16, Uint8, Uint,
        Number, NumberLike,
        Byte, Bytes, Rune, Runes,
        Float32Slice, Float64Slice,
        Int64Slice, Int32Slice, Int16Slice, Int8Slice, IntSlice,
        Uint64Slice, Uint32Slice, Uint16Slice, Uint8Slice, UintSlice,
        Error,
        ReadChannel, WriteChannel, Channel,
        Slice, Map,
        Uintptr, Native,
    } from "stdgo/builtin";
    import * as io from 'stdgo/io';
    import * as fs from 'stdgo/io/fs';
    import * as syscall from 'stdgo/syscall';
    import * as time from 'stdgo/time';

    // Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified.
    const O_RDONLY: Int // open the file read-only.
    const O_WRONLY: Int // open the file write-only.
    const O_RDWR: Int   // open the file read-write.
    // The remaining values may be or'ed in to control behavior.
    const O_APPEND: Int // append data to the file when writing.
    const O_CREATE: Int  // create a new file if none exists.
    const O_EXCL: Int   // used with O_CREATE, file must not exist.
    const O_SYNC: Int   // open for synchronous I/O.
    const O_TRUNC: Int  // truncate regular writable file when opened.

    const SEEK_SET: Int = 0 // seek relative to the origin of the file
    const SEEK_CUR: Int = 1 // seek relative to the current offset
    const SEEK_END: Int = 2 // seek relative to the end

    const PathSeparator: Rune// OS-specific path separator
    const PathListSeparator: Rune // OS-specific path list separator

    const ModeDir: FileMode        // d: is a directory
    const ModeAppend: FileMode                                     // a: append-only
    const ModeExclusive: FileMode                                  // l: exclusive use
    const ModeTemporary: FileMode                                  // T: temporary file; Plan 9 only
    const ModeSymlink: FileMode                                    // L: symbolic link
    const ModeDevice: FileMode                                     // D: device file
    const ModeNamedPipe: FileMode                                  // p: named pipe (FIFO)
    const ModeSocket: FileMode                                     // S: Unix domain socket
    const ModeSetuid: FileMode                                     // u: setuid
    const ModeSetgid: FileMode                                     // g: setgid
    const ModeCharDevice: FileMode                                 // c: Unix character device, when ModeDevice is set
    const ModeSticky: FileMode                                     // t: sticky
    const ModeIrregular: FileMode                                  // ?: non-regular file; nothing else is known about this file
    // Mask for the type bits. For regular files, none will be set.
    const ModeType: FileMode = ModeDir | ModeSymlink | ModeNamedPipe | ModeSocket | ModeDevice | ModeCharDevice | ModeIrregular
    const ModePerm: FileMode = 0777 // Unix permission bits

    const DevNull: string

    // ErrInvalid indicates an invalid argument.
    // Methods on File will return this error when the receiver is nil.
    const ErrInvalid = fs.ErrInvalid // "invalid argument"

    const ErrPermission = fs.ErrPermission // "permission denied"
    const ErrExist = fs.ErrExist      // "file already exists"
    const ErrNotExist = fs.ErrNotExist   // "file does not exist"
    const ErrClosed = fs.ErrClosed     // "file already closed"

    const ErrNoDeadline: Error       // "file type does not support deadline"
    const ErrDeadlineExceeded: Error // "i/o timeout"

    const Stdin: FilePointer
    const Stdout: FilePointer
    const Stderr: FilePointer

    const Args: Array<string>
    const ErrProcessDone: Error

    function Chdir(dir: string): void
    function Chmod(name: string, mode: FileMode | NumberLike): void
    function Chown(name: string, uid: Int | NumberLike, gid: Int | NumberLike): void
    function Chtimes(name: string, atime: time.Time, mtime: time.Time): void
    function Clearenv(): void
    /** added in go1.16 */
    function DirFS(dir: string): fs.FS
    function Environ(): Array<string>
    function Exit(code: Int | NumberLike)
    function Expand(s: string, mapping: (varname: string) => string): string
    function ExpandEnv(s: string): string
    function Getegid(): Int
    function Getenv(key: string): string
    function Geteuid(): Int
    function Getgid(): Int
    function Getgroups(): Array<Int>
    function Getpagesize(): Int
    function Getpid(): Int
    function Getppid(): Int
    function Getuid(): Int
    function Getwd(): string
    function Hostname(): string
    function IsExist(err: Error): boolean
    function IsNotExist(err: Error): boolean
    function IsPathSeparator(c: Byte | NumberLike): Boolean
    function IsPermission(err: Error): boolean
    /** added in go1.10 */
    function IsTimeout(err: Error): boolean
    function Lchown(name: string, uid: Int | NumberLike, gid: Int | NumberLike): void
    function Link(oldname: string, newname: string): void
    function LookupEnv(key: string): [string, boolean]
    function Mkdir(name: string, perm: FileMode | NumberLike): void
    function MkdirAll(path: string, perm: FileMode | NumberLike): void
    function MkdirTemp(dir: string, pattern: string): string
    /** return (r *File, w *File, err error) */
    function Pipe(): [File, File]
    function NewSyscallError(syscall: string, err: Error): void
    function ReadFile(name: string): Bytes
    function Readlink(name: string): string
    function Remove(name: string): void
    function RemoveAll(path: string): void
    function Rename(oldpath: string, newpath: string): void
    function SameFile(fi1: FileInfo, fi2: FileInfo): boolean
    function Setenv(key: string, value: string): void
    function Symlink(oldname: string, newname: string): void
    function TempDir(): string
    function Truncate(name: string, size: Int64 | NumberLike): void
    function Unsetenv(key: string): void
    function UserCacheDir(): string
    function UserConfigDir(): string
    function UserHomeDir(): string
    function WriteFile(name: string, data: Bytes, perm: FileMode | NumberLike)


    function ReadDir(name: string): Array<DirEntry>
    type DirEntry = fs.DirEntry

    function Create(name: string): FilePointer
    function CreateTemp(dir: string, pattern: string): FilePointer
    function NewFile(fd: Uintptr, name: string): FiFilePointerle
    function Open(name: string): FilePointer
    function OpenFile(name: string, flag: Int | NumberLike, perm: FileMode | NumberLike): FilePointer
    interface FilePointer extends Native {
        readonly __FilePointer: FilePointer

        Chdir(): void
        Chmod(mode: FileMode): void
        Chown(uid: Int | NumberLike, gid: Int | NumberLike): void
        Close(): void
        Fd(): Uintptr
        Name(): string
        Read(b: Bytes): Int
        ReadAt(b: Bytes, off: Int64 | NumberLike): Int
        ReadDir(n: Int | NumberLike): Array<DirEntry>
        ReadFrom(r: io.Reader): Int64
        Readdir(n: Int | NumberLike): Array<FileInfo>
        Readdirnames(n: Int | NumberLike): Array<string>
        Seek(offset: Int64 | NumberLike, whence: Int | NumberLike): Int64
        SetDeadline(t: time.Time): void
        SetReadDeadline(t: time.Time): void
        SetWriteDeadline(t: time.Time): void
        Stat(): FileInfo
        Sync(): void
        SyscallConn(): syscall.RawConn
        Truncate(size: Int64 | NumberLike): void
        Write(b: Bytes): Int
        WriteAt(b: Bytes, off: Int64 | NumberLike): Int
        WriteString(s: string): Int
    }

    function Lstat(name: string): FileInfo
    function Stat(name: string): FileInfo
    type FileInfo = fs.FileInfo

    type FileMode = fs.FileMode

    interface LinkErrorPointer extends Error {
        readonly __LinkErrorPointer: LinkErrorPointer
        Op: string
        Old: string
        New: string
        Err: Error

        Unwrap(): Error
    }
    type PathErrorPointer = fs.PathErrorPointer

    function NewProcAttr(): ProcAttrPointer
    interface ProcAttrPointer extends Native {
        readonly __ProcAttrPointer: ProcAttrPointer

        Dir: string
        Env: Array<string>
        Files: Array<FilePointer>
        Sys: syscall.SysProcAttr
    }

    function FindProcess(pid: Int | NumberLike): ProcessPointer
    function StartProcess(name: string, argv: Array<string>, attr: ProcAttrPointer): ProcessPointer
    interface ProcessPointer extends Native {
        readonly __ProcessPointer: ProcessPointer
        Pid: Int

        Kill()
        Release()
        Signal(sig: Signal)
        Wait(): ProcessStatePointer
    }

    interface ProcessStatePointer extends Native {
        readonly __ProcessStatePointer: ProcessStatePointer

        ExitCode(): Int
        Exited(): boolean
        Pid(): Int
        String(): string
        Success(): boolean
        Sys(): any
        SysUsage(): any
        SystemTime(): time.Duration
        UserTime(): time.Duration
    }
    interface Signal extends Native {
        String(): string
        Signal(): void // to distinguish from other Stringers
    }

    const Interrupt: Signal
    const Kill: Signal

    interface SyscallErrorPointer extends Error {
        readonly __SyscallErrorPointer: SyscallErrorPointer

        Syscall: string
        Err: Error

        Timeout(): boolean
        Unwrap(): Error
    }

    function isDirEntry(v: any): v is DirEntry
    function isFilePointer(v: any): v is FilePointer
    function isFileInfo(v: any): v is FileInfo
    function isFileMode(v: any): v is FileMode
    function isLinkErrorPointer(v: any): v is LinkErrorPointer
    function isPathErrorPointer(v: any): v is PathErrorPointer
    function isProcAttrPointer(v: any): v is ProcAttrPointer
    function isProcessPointer(v: any): v is ProcessPointer
    function isProcessStatePointer(v: any): v is ProcessStatePointer
    function isSignal(v: any): v is Signal
    function isSyscallErrorPointer(v: any): v is SyscallErrorPointer
}