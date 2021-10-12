# goja_go

# stdgo

The following js environment supports the go standard library:

* archive
   * tar
   * zip
* bufio
* bytes
* ~~compress~~
   * ~~bzip2~~
   * ~~flate~~
   * ~~gzip~~
   * ~~lzw~~
   * ~~zlib~~
* ~~container~~
* context
* ~~crypto~~
   * ~~aes~~
   * ~~cipher~~
   * ~~des~~
   * ~~dsa~~
   * ~~ecdsa~~
   * ~~ed25519~~
   * ~~elliptic~~
   * ~~hmac~~
   * ~~md5~~
   * ~~rand~~
   * ~~rc4~~
   * ~~rsa~~
   * ~~sha1~~
   * ~~sha256~~
   * ~~sha512~~
   * ~~subtle~~
   * ~~tls~~
   * ~~x509~~
      * ~~pkix~~
* ~~database~~
* ~~debug~~
* ~~embed~~
* ~~encoding~~
* errors
* ~~expvar~~
* ~~flag~~
* fmt
* ~~go~~
* ~~hash~~
   * ~~adler32~~
   * ~~crc32~~
   * ~~crc64~~
   * ~~fnv~~
   * ~~maphash~~
* ~~html~~
* ~~image~~
* ~~index~~
* io
   * fs
   * ioutil
* ~~log~~
* ~~math~~
* ~~mime~~
* ~~net~~
   * url
* os
* path
   * filepath
* ~~plugin~~
* ~~reflect~~
* ~~regexp~~
* ~~runtime~~
* ~~sort~~
* ~~strconv~~
* ~~strings~~
* ~~sync~~
* ~~syscall~~
* ~~testing~~
* ~~text~~
* time
* ~~unicode~~
* ~~unsafe~~

1. The typescript declaration of the js function is located in **scripts/types/stdgo**
2. The module name in js is the prefix of stdgo + ${go std module name}