package builtin

import (
	"errors"
	"reflect"

	"github.com/powerpuffpenguin/goja"
	"github.com/powerpuffpenguin/goja_go/core/utils"
)

func (f *factory) register() {
	f.Accessor(`native`, f.getNative, nil)
	f.Set(`append`, f.append)
	f.Set(`cap`, f.cap)
	f.Set(`len`, f.len)
	f.Set(`copy`, f.copy)
	f.Set(`getIndex`, f.getIndex)
	f.Set(`setIndex`, f.setIndex)
	f.Set(`slice`, f.slice)
	f.Set(`getKey`, f.getKey)
	f.Set(`setKey`, f.setKey)
	f.Set(`deleteKey`, f.deleteKey)
	f.Set(`typeOf`, f.typeOf)
	f.Set(`newType`, f.newType)
	f.Set(`makeSlice`, f.makeSlice)
	f.Set(`makeMap`, f.makeMap)
	f.Set(`makeChan`, f.makeChan)
	f.Set(`recv`, f.recv)
	f.Set(`tryRecv`, f.tryRecv)
	f.Set(`send`, f.send)
	f.Set(`trySend`, f.trySend)
	f.Set(`selectDefault`, f.selectDefault)
	f.Set(`selectRecv`, f.selectRecv)
	f.Set(`selectSend`, f.selectSend)
	f.Set(`select`, f._select)
}
func (f *factory) getNative(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(f)
}
func (f *factory) append(slice interface{}, elems ...interface{}) interface{} {
	defer utils.Recover(f.runtime)
	count := len(elems)
	if count != 0 {
		typ := reflect.TypeOf(slice)
		s := Unwrap(slice)
		x := make([]reflect.Value, count)
		for i, elem := range elems {
			x[i] = Unwrap(elem)
		}
		return reflect.Append(s, x...).Convert(typ).Interface()
	}
	return slice
}
func (f *factory) cap(i interface{}) Int {
	defer utils.Recover(f.runtime)
	return Int(Unwrap(i).Cap())
}
func (f *factory) len(i interface{}) Int {
	defer utils.Recover(f.runtime)
	return Int(Unwrap(i).Len())
}
func (f *factory) copy(dst, src interface{}) Int {
	defer utils.Recover(f.runtime)
	return Int(reflect.Copy(Unwrap(dst), Unwrap(src)))
}
func (f *factory) getIndex(i interface{}, index int) interface{} {
	defer utils.Recover(f.runtime)
	result := reflect.ValueOf(i).Index(index)
	return Wrap(result)
}
func (f *factory) setIndex(i interface{}, index int, x interface{}) {
	defer utils.Recover(f.runtime)
	Unwrap(i).Index(index).Set(
		Unwrap(x),
	)
}

func (f *factory) slice(call goja.FunctionCall) goja.Value {
	var (
		runtime = f.runtime
		arg0    = call.Argument(0)
	)
	export := arg0.Export()
	kind := arg0.ExportType().Kind()
	if kind != reflect.Slice {
		panic(runtime.NewGoError(&reflect.ValueError{
			Method: `reflect.Value.Slice`,
			Kind:   kind,
		}))
	}
	args := call.Arguments[1:]
	count := len(args)
	if count == 0 {
		return arg0
	} else if count == 1 {
		var i int
		e := runtime.ExportTo(args[0], &i)
		if e != nil {
			panic(runtime.NewGoError(e))
		}
		defer utils.Recover(f.runtime)
		value := reflect.ValueOf(export)
		return runtime.ToValue(value.Slice(i, value.Len()).Interface())
	} else if count == 2 {
		var i int
		e := runtime.ExportTo(args[0], &i)
		if e != nil {
			panic(runtime.NewGoError(e))
		}
		var j int
		e = runtime.ExportTo(args[1], &j)
		if e != nil {
			panic(runtime.NewGoError(e))
		}
		defer utils.Recover(f.runtime)
		value := reflect.ValueOf(export)
		return runtime.ToValue(value.Slice(i, j).Interface())
	}
	var i int
	e := runtime.ExportTo(args[0], &i)
	if e != nil {
		panic(runtime.NewGoError(e))
	}
	var j int
	e = runtime.ExportTo(args[1], &j)
	if e != nil {
		panic(runtime.NewGoError(e))
	}
	var k int
	e = runtime.ExportTo(args[2], &k)
	if e != nil {
		panic(runtime.NewGoError(e))
	}
	defer utils.Recover(f.runtime)
	value := reflect.ValueOf(export)
	return runtime.ToValue(value.Slice3(i, j, k).Interface())
}
func (f *factory) getKey(i, key interface{}) (interface{}, bool) {
	defer utils.Recover(f.runtime)
	result := reflect.ValueOf(i).MapIndex(Unwrap(key))
	if !result.IsValid() {
		return nil, false
	}
	return Wrap(result), true
}
func (f *factory) setKey(i, key, elem interface{}) {
	defer utils.Recover(f.runtime)
	reflect.ValueOf(i).SetMapIndex(Unwrap(key), Unwrap(elem))
}
func (f *factory) deleteKey(i, key interface{}) {
	defer utils.Recover(f.runtime)
	reflect.ValueOf(i).SetMapIndex(Unwrap(key), reflect.Value{})
}
func (f *factory) typeOf(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(call.Argument(0).ExportType())
}
func (f *factory) newType(typ reflect.Type) interface{} {
	defer utils.Recover(f.runtime)
	return reflect.New(typ).Interface()
}
func (f *factory) makeSlice(call goja.FunctionCall) goja.Value {
	var (
		runtime = f.runtime
		arg0    = call.Argument(0)
		typ     reflect.Type
	)
	e := runtime.ExportTo(arg0, &typ)
	if e != nil {
		panic(runtime.NewGoError(e))
	}
	args := call.Arguments[1:]
	count := len(args)
	if count == 0 {
		defer utils.Recover(f.runtime)
		return runtime.ToValue(reflect.MakeSlice(typ, 0, 0).Interface())
	} else if count == 1 {
		var l int
		e := runtime.ExportTo(args[0], &l)
		if e != nil {
			panic(runtime.NewGoError(e))
		}
		defer utils.Recover(f.runtime)
		return runtime.ToValue(reflect.MakeSlice(typ, l, l).Interface())
	}
	var l int
	e = runtime.ExportTo(args[0], &l)
	if e != nil {
		panic(runtime.NewGoError(e))
	}
	var c int
	e = runtime.ExportTo(args[1], &c)
	if e != nil {
		panic(runtime.NewGoError(e))
	}
	defer utils.Recover(f.runtime)
	return runtime.ToValue(reflect.MakeSlice(typ, l, c).Interface())
}
func (f *factory) makeMap(typ reflect.Type) interface{} {
	defer utils.Recover(f.runtime)
	return reflect.MakeMap(typ).Interface()
}
func (f *factory) makeChan(call goja.FunctionCall) goja.Value {
	var (
		runtime = f.runtime
		arg0    = call.Argument(0)
		typ     reflect.Type
	)
	e := runtime.ExportTo(arg0, &typ)
	if e != nil {
		panic(runtime.NewGoError(e))
	}
	args := call.Arguments[1:]
	count := len(args)
	if count == 0 {
		defer utils.Recover(f.runtime)
		return runtime.ToValue(reflect.MakeChan(typ, 0).Interface())
	}
	var buffer int
	e = runtime.ExportTo(args[0], &buffer)
	if e != nil {
		panic(runtime.NewGoError(e))
	}
	defer utils.Recover(f.runtime)
	return runtime.ToValue(reflect.MakeChan(typ, buffer).Interface())
}
func (f *factory) recv(i interface{}) (result interface{}, ok bool) {
	defer utils.Recover(f.runtime)
	x, ok := reflect.ValueOf(i).Recv()
	if ok {
		result = Wrap(x)
	}
	return
}
func (f *factory) tryRecv(i interface{}) (result interface{}, ok bool) {
	defer utils.Recover(f.runtime)
	x, ok := reflect.ValueOf(i).TryRecv()
	if ok {
		result = Wrap(x)
	}
	return
}
func (f *factory) send(i, x interface{}) {
	defer utils.Recover(f.runtime)
	reflect.ValueOf(i).Send(Unwrap(x))
}
func (f *factory) trySend(i, x interface{}) bool {
	defer utils.Recover(f.runtime)
	return reflect.ValueOf(i).TrySend(Unwrap(x))
}
func (f *factory) selectDefault() reflect.SelectCase {
	return reflect.SelectCase{
		Dir: reflect.SelectDefault,
	}
}
func (f *factory) selectRecv(call goja.FunctionCall) goja.Value {
	var (
		runtime = f.runtime
		arg0    = call.Argument(0)
	)
	kind := arg0.ExportType().Kind()
	if kind != reflect.Chan {
		panic(runtime.NewGoError(&reflect.ValueError{
			Method: `reflect.SelectRecv`,
			Kind:   kind,
		}))
	}
	return runtime.ToValue(reflect.SelectCase{
		Dir:  reflect.SelectRecv,
		Chan: reflect.ValueOf(arg0.Export()),
	})
}
func (f *factory) selectSend(call goja.FunctionCall) goja.Value {
	var (
		runtime = f.runtime
		arg0    = call.Argument(0)
	)
	kind := arg0.ExportType().Kind()
	if kind != reflect.Chan {
		panic(runtime.NewGoError(&reflect.ValueError{
			Method: `reflect.SelectSend`,
			Kind:   kind,
		}))
	} else if len(call.Arguments) < 2 {
		panic(runtime.NewGoError(errors.New(`send value is not set`)))
	}
	return runtime.ToValue(reflect.SelectCase{
		Dir:  reflect.SelectSend,
		Chan: reflect.ValueOf(arg0.Export()),
		Send: Unwrap(call.Arguments[1].Export()),
	})
}
func (f *factory) _select(cases ...reflect.SelectCase) (chosen int, recv interface{}, recvOK bool) {
	chosen, v, recvOK := reflect.Select(cases)
	if recvOK {
		recv = Wrap(v)
	}
	return
}
