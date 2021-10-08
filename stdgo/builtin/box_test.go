package builtin_test

import (
	"log"
	"testing"

	"github.com/powerpuffpenguin/goja"
	"github.com/powerpuffpenguin/goja/loop"
	"github.com/powerpuffpenguin/goja/require"
	"github.com/powerpuffpenguin/goja_go/stdgo"
	"github.com/powerpuffpenguin/goja_go/stdgo/builtin"
)

func New() *goja.Runtime {
	runtime := goja.New(
		goja.WithScheduler(loop.NewScheduler(32)),
		goja.WithFieldGetter(builtin.FieldGetter),
		goja.WithCallerFactory(builtin.NewCallerFactory()),
	)

	// enable require
	registry := require.NewRegistry(
		require.WithGlobalFolders(`node_modules`),
		require.WithLoader(require.DefaultSourceLoader),
	)
	registry.Enable(runtime)

	// enable stdgo
	stdgo.Enable(runtime)
	stdgo.RegisterNativeModuleToRegistry(registry)
	return runtime
}
func TestBox(t *testing.T) {
	r := New()
	r.Set(`print`, print)
	r.Set(`err`, func(args ...interface{}) {
		t.Fatal(args...)
	})
	r.Set(`max`, func(x, y int) int {
		if x > y {
			return x
		}
		return y
	})
	_, e := r.RunStringAndServe(`
const builtin =  require('stdgo/builtin');
if(2!=max(1,2)){
	err("not equal")
}
let promise=builtin.async(()=>max(1,2));
if(!(promise instanceof Promise)){
	err("not Promise")
}
promise.then((v)=>{
	if(v!=2){
		err("not equal")
	}
},(e)=>{
	err(e)
});
`)
	if e != nil {
		log.Fatalln(e)
	}
}
