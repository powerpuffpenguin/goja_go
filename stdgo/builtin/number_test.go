package builtin_test

import (
	"fmt"
	"math"
	"reflect"
	"testing"
	"time"

	"github.com/powerpuffpenguin/goja"
	"github.com/powerpuffpenguin/goja/require"
	"github.com/powerpuffpenguin/goja_go/stdgo/builtin"
)

func print(call goja.FunctionCall) goja.Value {
	for i, v := range call.Arguments {
		if i != 0 {
			fmt.Print(` `)
		}
		fmt.Print(v.Export())
	}
	fmt.Println()
	return goja.Undefined()
}
func printType(call goja.FunctionCall) goja.Value {
	for i, v := range call.Arguments {
		if i != 0 {
			fmt.Print(` `)
		}
		fmt.Print(v.ExportType())
	}
	fmt.Println()
	return goja.Undefined()
}

func TestNumber(t *testing.T) {
	r := goja.New()
	new(require.Registry).Enable(r)
	r.Set(`print`, print)
	r.Set(`printType`, printType)

	r.Set(`make`, func() interface{} {
		return builtin.Wrap(reflect.ValueOf([]int64{math.MaxInt64, 1, 2}))
	})
	r.Set(`makeSecond`, func() time.Duration {
		return time.Second
	})
	r.Set(`Int64`, func(v int64) int64 {
		return v
	})
	_, e := r.RunString(`
// const builtin = require("stdgo/builtin")
// var x=make()
// x=builtin.append(x,1,Int64(makeSecond()))
// print(builtin.len(x),builtin.cap(x))
// printType(x)
`)
	if e != nil {
		t.Fatal(e)
	}
}
