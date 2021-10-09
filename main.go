package main

import (
	"flag"
	"log"
	"path/filepath"

	"github.com/powerpuffpenguin/goja"
	"github.com/powerpuffpenguin/goja/loop"
	"github.com/powerpuffpenguin/goja/require"
	core "github.com/powerpuffpenguin/goja_go/core"
	"github.com/powerpuffpenguin/goja_go/core/console"
	"github.com/powerpuffpenguin/goja_go/stdgo"
	"github.com/powerpuffpenguin/goja_go/stdgo/builtin"
)

func main() {
	var (
		help     bool
		filename string
	)
	flag.BoolVar(&help, `help`, false, `display help`)
	flag.Parse()
	if help {
		flag.PrintDefaults()
	} else {
		filename = flag.Arg(0)
		if filename == "" {
			log.Fatalln(`no script specified`)
		}
		var e error
		if filepath.IsAbs(filename) {
			filename = filepath.Clean(filename)
		} else {
			filename, e = filepath.Abs(filename)
			if e != nil {
				log.Fatalln((e))
			}
		}
		// new runtime
		runtime := goja.New(
			goja.WithScheduler(loop.NewScheduler(32)),
			goja.WithFieldGetter(builtin.FieldGetter),
			goja.WithCallerFactory(builtin.NewCallerFactory()),
		)

		// enable require
		registry := require.NewRegistry(
			require.WithGlobalFolders(`node_modules`),
			require.WithLoader(core.Loader),
		)
		registry.Enable(runtime)
		// enable console
		console.Enable(runtime)

		// enable stdgo
		stdgo.Enable(runtime)
		stdgo.RegisterNativeModuleToRegistry(registry)

		source, e := core.Load(filename)
		if e != nil {
			log.Fatalln(e)
		}
		_, e = runtime.RunScriptAndServe(filename, string(source)) // run and serve
		if e != nil {
			log.Fatalln(e)
		}
	}
}
