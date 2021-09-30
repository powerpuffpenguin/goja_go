package builtin_test

import (
	"math"
	"reflect"
	"testing"

	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
	"github.com/powerpuffpenguin/goja_go/stdgo/builtin"
)

func TestSlice(t *testing.T) {
	r := goja.New()
	new(require.Registry).Enable(r)
	r.Set(`print`, print)
	r.Set(`printType`, printType)
	r.Set(`make`, func() interface{} {
		var result = [3]int64{0, 1, 2}
		return builtin.Wrap(reflect.ValueOf(result))
	})
	r.Set(`MaxInt64`, func() builtin.Int64 {
		return builtin.Int64(math.MaxInt64)
	})
	r.Set(`checkMax`, func(v builtin.Int64) {
		if v != math.MaxInt64 {
			t.Fatal(`checkMax err`, v, math.MaxInt64)
		}
	})
	_, e := r.RunString(`
function check(ok,msg){
	if(!ok){
		if(msg){
			throw new Error("not pass -> "+msg)
		}else{
			throw new Error("not pass")
		}
	}
}
const builtin = require("stdgo/builtin")
var a0=make()
check(builtin.len(a0)==3)
check(builtin.cap(a0)==3)
var a1=builtin.append(a0,builtin.Int64(3),builtin.Int64(4))
check(builtin.len(a0)==3)
check(builtin.cap(a0)==3)
check(builtin.len(a1)==3+2)
check(builtin.cap(a1)==3*2)
var a2 = builtin.slice(a1,0,3,4)
check(builtin.len(a2)==3)
check(builtin.cap(a2)==4)
a2 = builtin.slice(a1,0,3)
check(builtin.len(a2)==3)
check(builtin.cap(a2)==3*2)
for(var i=0;i<builtin.len(a0);i++)
{
	check(builtin.getIndex(a0,i)==i)
}
for(var i=0;i<builtin.len(a1);i++)
{
	check(builtin.getIndex(a1,i)==i)
}
builtin.setIndex(a1,2,MaxInt64())

checkMax(builtin.getIndex(a1,2))
checkMax(builtin.getIndex(a2,2))
`)
	if e != nil {
		t.Fatal(e)
	}
}
