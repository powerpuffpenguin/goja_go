package utils

import (
	"errors"

	"github.com/powerpuffpenguin/goja"
)

type Definer struct {
	runtime *goja.Runtime
	obj     *goja.Object
}

func MakeDefiner(runtime *goja.Runtime, obj *goja.Object) Definer {
	return Definer{
		runtime: runtime,
		obj:     obj,
	}
}

func (d *Definer) Set(name string, value interface{}) {
	e := d.obj.Set(name, value)
	if e != nil {
		panic(d.runtime.ToValue(e))
	}
}
func (d *Definer) GetObject(name string) *goja.Object {
	v := d.obj.Get(name)
	if result, ok := v.(*goja.Object); ok {
		return result
	}
	panic(d.runtime.ToValue(errors.New(" not a Object property: " + name)))
}
func (d *Definer) ChildDefiner(name string) Definer {
	obj := d.GetObject(name)
	return MakeDefiner(d.runtime, obj)
}
func (d *Definer) Accessor(name string,
	getter, setter func(goja.FunctionCall) goja.Value,
) {
	var g, s goja.Value
	if getter == nil {
		g = goja.Undefined()
	} else {
		g = d.runtime.ToValue(getter)
	}
	if setter == nil {
		s = goja.Undefined()
	} else {
		s = d.runtime.ToValue(setter)
	}
	e := d.obj.DefineAccessorProperty(name,
		g,
		s,
		goja.FLAG_TRUE, goja.FLAG_TRUE,
	)
	if e != nil {
		panic(d.runtime.ToValue(e))
	}
}
