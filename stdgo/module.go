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

	crypto "github.com/powerpuffpenguin/goja_go/stdgo/crypto"
	crypto_aes "github.com/powerpuffpenguin/goja_go/stdgo/crypto/aes"
	crypto_cipher "github.com/powerpuffpenguin/goja_go/stdgo/crypto/cipher"
	crypto_des "github.com/powerpuffpenguin/goja_go/stdgo/crypto/des"
	crypto_dsa "github.com/powerpuffpenguin/goja_go/stdgo/crypto/dsa"
	crypto_ecdsa "github.com/powerpuffpenguin/goja_go/stdgo/crypto/ecdsa"
	crypto_ed25519 "github.com/powerpuffpenguin/goja_go/stdgo/crypto/ed25519"
	crypto_elliptic "github.com/powerpuffpenguin/goja_go/stdgo/crypto/elliptic"
	crypto_hmac "github.com/powerpuffpenguin/goja_go/stdgo/crypto/hmac"
	crypto_md5 "github.com/powerpuffpenguin/goja_go/stdgo/crypto/md5"
	crypto_rand "github.com/powerpuffpenguin/goja_go/stdgo/crypto/rand"
	crypto_rc4 "github.com/powerpuffpenguin/goja_go/stdgo/crypto/rc4"
	crypto_rsa "github.com/powerpuffpenguin/goja_go/stdgo/crypto/rsa"
	crypto_sha1 "github.com/powerpuffpenguin/goja_go/stdgo/crypto/sha1"
	crypto_sha256 "github.com/powerpuffpenguin/goja_go/stdgo/crypto/sha256"
	crypto_sha512 "github.com/powerpuffpenguin/goja_go/stdgo/crypto/sha512"
	crypto_subtle "github.com/powerpuffpenguin/goja_go/stdgo/crypto/subtle"
	crypto_tls "github.com/powerpuffpenguin/goja_go/stdgo/crypto/tls"
	crypto_x509 "github.com/powerpuffpenguin/goja_go/stdgo/crypto/x509"
	crypto_x509_pkix "github.com/powerpuffpenguin/goja_go/stdgo/crypto/x509/pkix"

	database_sql "github.com/powerpuffpenguin/goja_go/stdgo/database/sql"
	database_sql_driver "github.com/powerpuffpenguin/goja_go/stdgo/database/sql/driver"

	encoding "github.com/powerpuffpenguin/goja_go/stdgo/encoding"
	encoding_ascii85 "github.com/powerpuffpenguin/goja_go/stdgo/encoding/ascii85"
	encoding_asn1 "github.com/powerpuffpenguin/goja_go/stdgo/encoding/asn1"
	encoding_base32 "github.com/powerpuffpenguin/goja_go/stdgo/encoding/base32"
	encoding_base64 "github.com/powerpuffpenguin/goja_go/stdgo/encoding/base64"
	encoding_binary "github.com/powerpuffpenguin/goja_go/stdgo/encoding/binary"
	encoding_csv "github.com/powerpuffpenguin/goja_go/stdgo/encoding/csv"
	encoding_gob "github.com/powerpuffpenguin/goja_go/stdgo/encoding/gob"
	encoding_hex "github.com/powerpuffpenguin/goja_go/stdgo/encoding/hex"
	encoding_json "github.com/powerpuffpenguin/goja_go/stdgo/encoding/json"
	encoding_pem "github.com/powerpuffpenguin/goja_go/stdgo/encoding/pem"
	encoding_xml "github.com/powerpuffpenguin/goja_go/stdgo/encoding/xml"

	errors "github.com/powerpuffpenguin/goja_go/stdgo/errors"

	fmt "github.com/powerpuffpenguin/goja_go/stdgo/fmt"

	hash "github.com/powerpuffpenguin/goja_go/stdgo/hash"
	hash_adler32 "github.com/powerpuffpenguin/goja_go/stdgo/hash/adler32"
	hash_crc32 "github.com/powerpuffpenguin/goja_go/stdgo/hash/crc32"
	hash_crc64 "github.com/powerpuffpenguin/goja_go/stdgo/hash/crc64"
	hash_fnv "github.com/powerpuffpenguin/goja_go/stdgo/hash/fnv"
	hash_maphash "github.com/powerpuffpenguin/goja_go/stdgo/hash/maphash"

	image "github.com/powerpuffpenguin/goja_go/stdgo/image"
	image_color "github.com/powerpuffpenguin/goja_go/stdgo/image/color"
	image_color_palette "github.com/powerpuffpenguin/goja_go/stdgo/image/color/palette"
	image_draw "github.com/powerpuffpenguin/goja_go/stdgo/image/draw"
	image_gif "github.com/powerpuffpenguin/goja_go/stdgo/image/gif"
	image_jpeg "github.com/powerpuffpenguin/goja_go/stdgo/image/jpeg"
	image_png "github.com/powerpuffpenguin/goja_go/stdgo/image/png"

	index_suffixarray "github.com/powerpuffpenguin/goja_go/stdgo/index/suffixarray"

	io "github.com/powerpuffpenguin/goja_go/stdgo/io"
	io_fs "github.com/powerpuffpenguin/goja_go/stdgo/io/fs"
	io_ioutil "github.com/powerpuffpenguin/goja_go/stdgo/io/ioutil"

	math "github.com/powerpuffpenguin/goja_go/stdgo/math"
	math_big "github.com/powerpuffpenguin/goja_go/stdgo/math/big"
	math_bits "github.com/powerpuffpenguin/goja_go/stdgo/math/bits"
	math_cmplx "github.com/powerpuffpenguin/goja_go/stdgo/math/cmplx"

	net "github.com/powerpuffpenguin/goja_go/stdgo/net"
	net_http "github.com/powerpuffpenguin/goja_go/stdgo/net/http"
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

	{crypto.ModuleID, crypto.Require},
	{crypto_aes.ModuleID, crypto_aes.Require},
	{crypto_cipher.ModuleID, crypto_cipher.Require},
	{crypto_des.ModuleID, crypto_des.Require},
	{crypto_dsa.ModuleID, crypto_dsa.Require},
	{crypto_ecdsa.ModuleID, crypto_ecdsa.Require},
	{crypto_ed25519.ModuleID, crypto_ed25519.Require},
	{crypto_elliptic.ModuleID, crypto_elliptic.Require},
	{crypto_hmac.ModuleID, crypto_hmac.Require},
	{crypto_md5.ModuleID, crypto_md5.Require},
	{crypto_rand.ModuleID, crypto_rand.Require},
	{crypto_rc4.ModuleID, crypto_rc4.Require},
	{crypto_rsa.ModuleID, crypto_rsa.Require},
	{crypto_sha1.ModuleID, crypto_sha1.Require},
	{crypto_sha256.ModuleID, crypto_sha256.Require},
	{crypto_sha512.ModuleID, crypto_sha512.Require},
	{crypto_subtle.ModuleID, crypto_subtle.Require},
	{crypto_tls.ModuleID, crypto_tls.Require},
	{crypto_x509.ModuleID, crypto_x509.Require},
	{crypto_x509_pkix.ModuleID, crypto_x509_pkix.Require},

	{database_sql_driver.ModuleID, database_sql_driver.Require},
	{database_sql.ModuleID, database_sql.Require},

	{encoding.ModuleID, encoding.Require},
	{encoding_ascii85.ModuleID, encoding_ascii85.Require},
	{encoding_asn1.ModuleID, encoding_asn1.Require},
	{encoding_base32.ModuleID, encoding_base32.Require},
	{encoding_base64.ModuleID, encoding_base64.Require},
	{encoding_binary.ModuleID, encoding_binary.Require},
	{encoding_csv.ModuleID, encoding_csv.Require},
	{encoding_gob.ModuleID, encoding_gob.Require},
	{encoding_hex.ModuleID, encoding_hex.Require},
	{encoding_json.ModuleID, encoding_json.Require},
	{encoding_pem.ModuleID, encoding_pem.Require},
	{encoding_xml.ModuleID, encoding_xml.Require},

	{errors.ModuleID, errors.Require},

	{fmt.ModuleID, fmt.Require},

	{hash.ModuleID, hash.Require},
	{hash_adler32.ModuleID, hash_adler32.Require},
	{hash_crc32.ModuleID, hash_crc32.Require},
	{hash_crc64.ModuleID, hash_crc64.Require},
	{hash_fnv.ModuleID, hash_fnv.Require},
	{hash_maphash.ModuleID, hash_maphash.Require},

	{image.ModuleID, image.Require},
	{image_color.ModuleID, image_color.Require},
	{image_color_palette.ModuleID, image_color_palette.Require},
	{image_draw.ModuleID, image_draw.Require},
	{image_gif.ModuleID, image_gif.Require},
	{image_jpeg.ModuleID, image_jpeg.Require},
	{image_png.ModuleID, image_png.Require},

	{index_suffixarray.ModuleID, index_suffixarray.Require},

	{io.ModuleID, io.Require},
	{io_fs.ModuleID, io_fs.Require},
	{io_ioutil.ModuleID, io_ioutil.Require},

	{math.ModuleID, math.Require},
	{math_big.ModuleID, math_big.Require},
	{math_bits.ModuleID, math_bits.Require},
	{math_cmplx.ModuleID, math_cmplx.Require},

	{net.ModuleID, net.Require},
	{net_http.ModuleID, net_http.Require},
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
