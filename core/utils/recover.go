package utils

import (
	"errors"
	"fmt"

	"github.com/powerpuffpenguin/goja"
)

func Recover(r *goja.Runtime) {
	if e := recover(); e != nil {
		if v, ok := e.(goja.Value); ok {
			panic(v)
		} else if err, ok := e.(error); ok {
			panic(r.NewGoError(err))
		} else {
			panic(r.NewGoError(errors.New(fmt.Sprint(e))))
		}
	}
}
