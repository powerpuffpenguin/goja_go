package stdgo

import (
	"github.com/powerpuffpenguin/goja"
	"github.com/powerpuffpenguin/goja/require"

	archive_tar "github.com/powerpuffpenguin/goja_go/stdgo/archive/tar"
	archive_zip "github.com/powerpuffpenguin/goja_go/stdgo/archive/zip"

	compress_bzip2 "github.com/powerpuffpenguin/goja_go/stdgo/compress/bzip2"
	compress_flate "github.com/powerpuffpenguin/goja_go/stdgo/compress/flate"
	compress_gzip "github.com/powerpuffpenguin/goja_go/stdgo/compress/gzip"
	compress_lzw "github.com/powerpuffpenguin/goja_go/stdgo/compress/lzw"
	compress_zlib "github.com/powerpuffpenguin/goja_go/stdgo/compress/zlib"

	bufio "github.com/powerpuffpenguin/goja_go/stdgo/bufio"
	bytes "github.com/powerpuffpenguin/goja_go/stdgo/bytes"
	context "github.com/powerpuffpenguin/goja_go/stdgo/context"

	errors "github.com/powerpuffpenguin/goja_go/stdgo/errors"

	fmt "github.com/powerpuffpenguin/goja_go/stdgo/fmt"

	hash "github.com/powerpuffpenguin/goja_go/stdgo/hash"
	hash_adler32 "github.com/powerpuffpenguin/goja_go/stdgo/hash/adler32"
	hash_crc32 "github.com/powerpuffpenguin/goja_go/stdgo/hash/crc32"
	hash_crc64 "github.com/powerpuffpenguin/goja_go/stdgo/hash/crc64"
	hash_fnv "github.com/powerpuffpenguin/goja_go/stdgo/hash/fnv"
	hash_maphash "github.com/powerpuffpenguin/goja_go/stdgo/hash/maphash"

	io "github.com/powerpuffpenguin/goja_go/stdgo/io"
	io_fs "github.com/powerpuffpenguin/goja_go/stdgo/io/fs"
	io_ioutil "github.com/powerpuffpenguin/goja_go/stdgo/io/ioutil"

	math "github.com/powerpuffpenguin/goja_go/stdgo/math"
	math_big "github.com/powerpuffpenguin/goja_go/stdgo/math/big"
	math_bits "github.com/powerpuffpenguin/goja_go/stdgo/math/bits"
	math_cmplx "github.com/powerpuffpenguin/goja_go/stdgo/math/cmplx"
	
	net "github.com/powerpuffpenguin/goja_go/stdgo/net"
	net_url "github.com/powerpuffpenguin/goja_go/stdgo/net/url"

	os "github.com/powerpuffpenguin/goja_go/stdgo/os"
	os_exec "github.com/powerpuffpenguin/goja_go/stdgo/os/exec"
	os_signal "github.com/powerpuffpenguin/goja_go/stdgo/os/signal"

	path "github.com/powerpuffpenguin/goja_go/stdgo/path"
	path_filepath "github.com/powerpuffpenguin/goja_go/stdgo/path/filepath"

	regexp "github.com/powerpuffpenguin/goja_go/stdgo/regexp"
	sort "github.com/powerpuffpenguin/goja_go/stdgo/sort"
	strconv "github.com/powerpuffpenguin/goja_go/stdgo/strconv"
	strings "github.com/powerpuffpenguin/goja_go/stdgo/strings"
	time "github.com/powerpuffpenguin/goja_go/stdgo/time"
)

var modules = []module{
	{archive_tar.ModuleID, archive_tar.Require},
	{archive_zip.ModuleID, archive_zip.Require},

	{compress_bzip2.ModuleID, compress_bzip2.Require},
	{compress_flate.ModuleID, compress_flate.Require},
	{compress_gzip.ModuleID, compress_gzip.Require},
	{compress_lzw.ModuleID, compress_lzw.Require},
	{compress_zlib.ModuleID, compress_zlib.Require},

	{bufio.ModuleID, bufio.Require},
	{bytes.ModuleID, bytes.Require},
	{context.ModuleID, context.Require},

	{errors.ModuleID, errors.Require},

	{fmt.ModuleID, fmt.Require},

	{hash.ModuleID, hash.Require},
	{hash_adler32.ModuleID, hash_adler32.Require},
	{hash_crc32.ModuleID, hash_crc32.Require},
	{hash_crc64.ModuleID, hash_crc64.Require},
	{hash_fnv.ModuleID, hash_fnv.Require},
	{hash_maphash.ModuleID, hash_maphash.Require},

	{io.ModuleID, io.Require},
	{io_fs.ModuleID, io_fs.Require},
	{io_ioutil.ModuleID, io_ioutil.Require},

	{math.ModuleID, math.Require},
	{math_big.ModuleID, math_big.Require},
	{math_bits.ModuleID, math_bits.Require},
	{math_cmplx.ModuleID, math_cmplx.Require},

	{net.ModuleID, net.Require},
	{net_url.ModuleID, net_url.Require},

	{os.ModuleID, os.Require},
	{os_exec.ModuleID, os_exec.Require},
	{os_signal.ModuleID, os_signal.Require},

	{path.ModuleID, path.Require},
	{path_filepath.ModuleID, path_filepath.Require},

	{regexp.ModuleID, regexp.Require},
	{sort.ModuleID, sort.Require},
	{strconv.ModuleID, strconv.Require},
	{strings.ModuleID, strings.Require},
	{time.ModuleID, time.Require},
}

type module struct {
	ID      string
	Require func(runtime *goja.Runtime, module *goja.Object)
}

func RegisterNativeModule() {
	for _, m := range modules {
		require.RegisterNativeModule(m.ID, m.Require)
	}
}
func RegisterNativeModuleToRegistry(registry *require.Registry) {
	for _, m := range modules {
		registry.RegisterNativeModule(m.ID, m.Require)
	}
}
