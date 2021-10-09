package core

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/powerpuffpenguin/goja/require"
	core_strings "github.com/powerpuffpenguin/goja_go/core/strings"
)

func Loader(filename string) ([]byte, error) {
	source, e := Load(filename)
	if e != nil {
		return nil, e
	}
	return core_strings.StringToBytes(source), nil
}
func Load(filename string) (source string, e error) {
	var b []byte
	f, e := os.Open(filename)
	if e != nil {
		if os.IsNotExist(e) {
			e = require.ModuleFileDoesNotExistError
		}
		return
	}
	stat, e := f.Stat()
	if e != nil {
		f.Close()
		return
	}
	if stat.IsDir() {
		f.Close()
		e = require.ModuleFileDoesNotExistError
		return
	}
	b, e = ioutil.ReadAll(f)
	f.Close()
	if e != nil {
		return
	}
	ext := strings.ToLower(filepath.Ext(filename))
	if ext == `.js` {
		source = fmt.Sprintf(`var require = require || {};var module = module || { exports: {} };(function(__dirname,__filename,exports,require,module){%s})(%q,%q,module.exports,require,module);`,
			b,
			filepath.Dir(filename), filename,
		)
	} else {
		source = core_strings.BytesToString(b)
	}
	return
}
