declare module "stdgo/syscall" {
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

    interface SysProcAttr extends Native {
        readonly __SysProcAttr: SysProcAttr
    }
    interface RawConn extends Native {
        // Control invokes f on the underlying connection's file
        // descriptor or handle.
        // The file descriptor fd is guaranteed to remain valid while
        // f executes but not after f returns.
        Control(f: (fd: Uintptr) => void)

        // Read invokes f on the underlying connection's file
        // descriptor or handle; f is expected to try to read from the
        // file descriptor.
        // If f returns true, Read returns. Otherwise Read blocks
        // waiting for the connection to be ready for reading and
        // tries again repeatedly.
        // The file descriptor is guaranteed to remain valid while f
        // executes but not after f returns.
        Read(f: (fd: Uintptr) => boolean)

        // Write is like Read but for writing.
        Write(f: (fd: Uintptr) => boolean)
    }
}