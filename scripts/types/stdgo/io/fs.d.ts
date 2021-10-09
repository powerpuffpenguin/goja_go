declare module "stdgo/io/fs" {
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
    } from "stdgo/builtin";
    import *as time from "stdgo/time";

    const ErrInvalid: Error
    const ErrPermission: Error
    const ErrExist: Error
    const ErrNotExist: Error
    const ErrClosed: Error

    const SkipDir: Error

    function Glob(fsys: FS, pattern: string): Array<string>
    function ReadFile(fsys: FS, name: string): Bytes
    function ValidPath(name: string): boolean
    function WalkDir(fsys: FS, root: string, fn: WalkDirFunc)

    function ReadDir(fsys: FS, name: string): Array<DirEntry>
    interface DirEntry extends Native {
        Name(): string
        IsDir(): boolean
        Type(): FileMode
        Info(): FileInfo
    }

    function Sub(fsys: FS, dir: string): FS
    interface FS extends Native {
        Open(name: string): File
    }

    interface File extends Native {
        Stat(): FileInfo
        Read(b: Bytes): Int
        Close(): void
    }

    function Stat(fsys: FS, name: string): FileInfo
    interface FileInfo extends Native {
        Name(): string       // base name of the file
        Size(): Int64        // length in bytes for regular files; system-dependent for others
        Mode(): FileMode     // file mode bits
        ModTime(): time.Time // modification time
        IsDir(): boolean        // abbreviation for Mode().IsDir()
        Sys(): any  // underlying data source (can return nil)
    }

    function FileMode(v: Uint32 | NumberLike): FileMode
    interface FileMode extends Native {
        readonly __FileMode: FileMode

        IsDir(): boolean
        IsRegular(): boolean
        Perm(): FileMode
        String(): string
        Type(): FileMode
    }
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

    interface GlobFS extends FS {
        Glob(pattern: string): Array<string>
        Glob(pattern: string, scheduler: Scheduler): Promise<Array<string>>
    }
    interface PathErrorPointer extends Error {
        readonly __PathErrorPointer: PathErrorPointer

        Op: string
        Path: string
        Err: Error

        Timeout(): boolean
        Unwrap(): Error
    }
    interface ReadDirFS extends FS {
        ReadDir(name: string): Array<DirEntry>
    }
    interface ReadDirFile extends File {
        ReadDir(n: Int | NumberLike): Array<DirEntry>
    }
    interface ReadFileFS extends FS {
        ReadFile(name: string): Bytes
    }
    interface StatFS extends FS {
        Stat(name: string): FileInfo
    }
    interface SubFS extends FS {
        Sub(dir: string): FS
    }
    type WalkDirFunc = (path: string, d: DirEntry, err: Error) => Error

    function isDirEntry(v: any): v is DirEntry
    function isFS(v: any): v is FS
    function isFile(v: any): v is File
    function isFileInfo(v: any): v is FileInfo
    function isFileMode(v: any): v is FileMode
    function isGlobFS(v: any): v is GlobFS
    function isPathErrorPointer(v: any): v is PathErrorPointer
    function isReadDirFS(v: any): v is ReadDirFS
    function isReadDirFile(v: any): v is ReadDirFile
    function isReadFileFS(v: any): v is ReadFileFS
    function isStatFS(v: any): v is StatFS
    function isSubFS(v: any): v is SubFS
    function isWalkDirFunc(v: any): v is WalkDirFunc
}