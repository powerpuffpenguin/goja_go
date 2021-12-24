package builtin

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
	"unsafe"

	"github.com/powerpuffpenguin/goja"
	"github.com/powerpuffpenguin/goja/loop"
	"github.com/powerpuffpenguin/goja_go/core/utils"
)

type CallStyle int

const (
	callAsync CallStyle = 1 << iota
	callErr
	callNoWrap
)

func (f *factory) registerBox() {
	f.Set(`print`, print)
	f.Set(`printType`, printType)
	f.Set(`async`, f.async)
	f.Set(`error`, f.error)
}
func (f *factory) error(call goja.FunctionCall) goja.Value {
	runtime := f.runtime
	callable, ok := goja.AssertFunction(call.Argument(0))
	if !ok {
		panic(runtime.NewTypeError(`arg0 not callable`))
	}
	defer utils.Recover(runtime)
	call.Arguments[0] = runtime.ToValue(callErr)
	val, e := callable(goja.Undefined())
	if e != nil {
		panic(runtime.NewGoError(e))
	}
	return val
}
func (f *factory) async(call goja.FunctionCall) goja.Value {
	runtime := f.runtime
	callable, ok := goja.AssertFunction(call.Argument(0))
	if !ok {
		panic(runtime.NewTypeError(`arg0 not callable`))
	}
	defer utils.Recover(runtime)
	call.Arguments[0] = runtime.ToValue(callAsync)
	val, e := callable(call.This, call.Arguments...)
	if e != nil {
		panic(runtime.NewGoError(e))
	}
	return val
}
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

type emptyInterface struct {
	typ  *emptyInterface
	word unsafe.Pointer
}

func wrapArray(value reflect.Value) reflect.Value {
	typ := reflect.SliceOf(value.Type().Elem())
	slice := reflect.New(typ)

	p0 := (*reflect.SliceHeader)(unsafe.Pointer(slice.Pointer()))
	p0.Len = value.Len()
	p0.Cap = p0.Len

	p1 := (*emptyInterface)(unsafe.Pointer(&value))
	p0.Data = uintptr(p1.word)

	return reflect.Indirect(slice)
}
func Wrap(value reflect.Value) interface{} {
	if value.Kind() == reflect.Array {
		return wrapArray(value).Interface()
	}
	i := value.Interface()
	switch i := i.(type) {
	case int:
		return Int(i)
	case int8:
		return Int8(i)
	case int16:
		return Int16(i)
	case int32:
		return Int32(i)
	case int64:
		return Int64(i)
	case uint:
		return Uint(i)
	case uint8:
		return Uint8(i)
	case uint16:
		return Uint16(i)
	case uint32:
		return Uint32(i)
	case uint64:
		return Uint64(i)
	case float32:
		return Float32(i)
	case float64:
		return Float64(i)

	case []int:
		return IntSlice(i)
	case []int8:
		return Int8Slice(i)
	case []int16:
		return Int16Slice(i)
	case []int32:
		return Int32Slice(i)
	case []int64:
		return Int64Slice(i)
	case []uint:
		return UintSlice(i)
	case []uint8:
		return Uint8Slice(i)
	case []uint16:
		return Uint16Slice(i)
	case []uint32:
		return Uint32Slice(i)
	case []uint64:
		return Uint64Slice(i)
	case []float32:
		return Float32Slice(i)
	case []float64:
		return Float64Slice(i)
	default:
	}
	return i
}

func Unwrap(i interface{}) reflect.Value {
	return reflect.ValueOf(unwrap(i))
}
func unwrap(i interface{}) interface{} {
	switch i := i.(type) {
	case Int:
		return int(i)
	case Int8:
		return int8(i)
	case Int16:
		return int16(i)
	case Int32:
		return int32(i)
	case Int64:
		return int64(i)
	case Uint:
		return uint(i)
	case Uint8:
		return uint8(i)
	case Uint16:
		return uint16(i)
	case Uint32:
		return uint32(i)
	case Uint64:
		return uint64(i)
	case Float32:
		return float32(i)
	case Float64:
		return float64(i)
	case IntSlice:
		return []int(i)
	case Int8Slice:
		return []int8(i)
	case Int16Slice:
		return []int16(i)
	case Int32Slice:
		return []int32(i)
	case Int64Slice:
		return []int64(i)
	case UintSlice:
		return []uint(i)
	case Uint8Slice:
		return []uint8(i)
	case Uint16Slice:
		return []uint16(i)
	case Uint32Slice:
		return []uint32(i)
	case Uint64Slice:
		return []uint64(i)
	case Float32Slice:
		return []float32(i)
	case Float64Slice:
		return []float64(i)
	default:
	}
	return i
}

func FieldGetter(value reflect.Value) reflect.Value {
	return reflect.ValueOf(Wrap(value))
}

type caller struct {
	scheduler loop.Scheduler
	err       bool
	noWrap    bool
}

func (c *caller) Reset() {
	c.scheduler = nil
	c.err = false
	c.noWrap = false
}
func (c *caller) Before(runtime *goja.Runtime, call *goja.FunctionCall) (err error) {
	var (
		style CallStyle
	)
	if len(call.Arguments) != 0 {
		if tmp, ok := call.Arguments[0].Export().(CallStyle); ok {
			call.Arguments = call.Arguments[1:]

			style = tmp
		}

		if len(call.Arguments) != 0 {
			offset := len(call.Arguments) - 1
			arg := call.Arguments[offset]
			if scheduler, ok := arg.Export().(loop.Scheduler); ok {
				call.Arguments = call.Arguments[:offset]
				c.scheduler = scheduler
			}
		}
	}
	if c.scheduler == nil {
		if (style & callAsync) != 0 {
			c.scheduler = runtime.Loop().GetScheduler()
		}
	}
	if (style & callErr) != 0 {
		c.err = true
	}
	if (style & callNoWrap) != 0 {
		c.noWrap = true
	}
	return nil
}

func (c *caller) Call(runtime *goja.Runtime, callSlice bool, callable reflect.Value, in []reflect.Value) (out []reflect.Value, err error) {
	defer func() {
		if e := recover(); e != nil {
			if v, ok := e.(goja.Value); ok {
				panic(v)
			} else if err, ok := e.(error); ok {
				panic(runtime.NewGoError(err))
			} else {
				panic(runtime.NewGoError(errors.New(fmt.Sprint(e))))
			}
		}
	}()

	if c.scheduler == nil {
		if callSlice {
			out = callable.CallSlice(in)
		} else {
			out = callable.Call(in)
		}
	} else {
		promise, resolve, reject := runtime.NewPromise()
		loop := runtime.Loop()
		err = loop.Async(nil)
		if err != nil {
			return
		}
		c.scheduler.Go(newAsyncImpl(c, runtime, promise, resolve, reject, callSlice, callable, in))
		out = append(out, reflect.ValueOf(promise))
	}
	return
}

func (c *caller) After(runtime *goja.Runtime, out []reflect.Value) (goja.Value, error) {
	if c.err {
		return c.after(runtime, out)
	} else {
		if len(out) != 0 {
			if last := out[len(out)-1]; last.Type().Name() == "error" {
				if !last.IsNil() {
					err := last.Interface()
					if e, ok := err.(*goja.Exception); ok {
						var v goja.Value = e.Value()
						panic(v)
					}
					return nil, err.(error)
				}
				out = out[:len(out)-1]
			}
		}
		return c.after(runtime, out)
	}
}
func (c *caller) after(runtime *goja.Runtime, out []reflect.Value) (result goja.Value, err error) {
	switch len(out) {
	case 0:
		result = goja.Undefined()
	case 1:
		if c.noWrap {
			result = runtime.ToValue(out[0].Interface())
		} else {
			result = runtime.ToValue(Wrap(out[0]))
		}
	default:
		s := make([]interface{}, len(out))
		if c.noWrap {
			for i, v := range out {
				s[i] = v.Interface()
			}
		} else {
			for i, v := range out {
				s[i] = Wrap(v)
			}
		}
		result = runtime.ToValue(s)
	}
	return
}

type callerFactory sync.Pool

func NewCallerFactory() *callerFactory {
	pool := &sync.Pool{
		New: func() interface{} {
			return &caller{}
		},
	}
	return (*callerFactory)(pool)
}
func (f *callerFactory) Get() goja.Caller {
	caller := ((*sync.Pool)(f)).Get().(*caller)
	caller.Reset()
	return caller
}
func (f *callerFactory) Put(caller goja.Caller) {
	((*sync.Pool)(f)).Put(caller)
}

type asyncImpl struct {
	caller          *caller
	runtime         *goja.Runtime
	promise         *goja.Promise
	resolve, reject func(interface{})

	callSlice bool
	value     reflect.Value
	in        []reflect.Value

	result []reflect.Value
	err    goja.Value
}

func newAsyncImpl(caller *caller,
	runtime *goja.Runtime,
	promise *goja.Promise,
	resolve, reject func(interface{}),
	callSlice bool, value reflect.Value, in []reflect.Value,
) *asyncImpl {
	return &asyncImpl{
		caller:    caller,
		runtime:   runtime,
		promise:   promise,
		resolve:   resolve,
		reject:    reject,
		callSlice: callSlice,
		value:     value,
		in:        in,
	}
}
func (a *asyncImpl) Serve() {
	a.serve()
	a.runtime.Loop().Result(a)
}
func (a *asyncImpl) serve() {
	defer func() {
		if e := recover(); e != nil {
			if v, ok := e.(goja.Value); ok {
				a.err = v
			} else if err, ok := e.(error); ok {
				a.err = a.runtime.NewGoError(err)
			} else {
				a.err = a.runtime.NewGoError(errors.New(fmt.Sprint(e)))
			}
		}
	}()
	if a.callSlice {
		a.result = a.value.CallSlice(a.in)
	} else {
		a.result = a.value.Call(a.in)
	}
}
func (a *asyncImpl) OnResult(closed bool) (completed bool) {
	completed = true
	if a.err != nil {
		a.reject(a.err)
		return
	}
	var (
		runtime = a.runtime
		caller  = a.caller
		out     = a.result
		result  goja.Value
		e       error
	)
	if caller.err {
		result, e = caller.after(runtime, out)
	} else {
		if len(out) != 0 {
			if last := out[len(out)-1]; last.Type().Name() == "error" {
				if !last.IsNil() {
					err := last.Interface()
					if e, ok := err.(*goja.Exception); ok {
						var v goja.Value = e.Value()
						a.reject(v)
						return
					}
					a.reject(runtime.NewGoError(err.(error)))
					return
				}
				out = out[:len(out)-1]
			}
		}
		result, e = caller.after(runtime, out)
	}
	if e == nil {
		a.resolve(result)
	} else {
		a.reject(runtime.NewGoError(e))
	}
	return
}
