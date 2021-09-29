package builtin

import (
	"reflect"

	"github.com/powerpuffpenguin/goja_go/core/utils"
)

func (f *factory) register() {
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
}
func (f *factory) append(slice interface{}, elems ...interface{}) interface{} {
	defer utils.Recover(f.runtime)
	count := len(elems)
	if count != 0 {
		typ := reflect.TypeOf(slice)
		s := UnwrapValueOf(slice)
		x := make([]reflect.Value, count)
		for i, elem := range elems {
			x[i] = UnwrapValueOf(elem)
		}
		return reflect.Append(s, x...).Convert(typ).Interface()
	}
	return slice
}
func (f *factory) cap(i interface{}) Int {
	defer utils.Recover(f.runtime)
	return Int(UnwrapValueOf(i).Cap())
}
func (f *factory) len(i interface{}) Int {
	defer utils.Recover(f.runtime)
	return Int(UnwrapValueOf(i).Len())
}
func (f *factory) copy(dst, src interface{}) Int {
	defer utils.Recover(f.runtime)
	return Int(reflect.Copy(UnwrapValueOf(dst), UnwrapValueOf(src)))
}
func (f *factory) getIndex(i interface{}, index int) interface{} {
	defer utils.Recover(f.runtime)
	result := reflect.ValueOf(i).Index(index).Interface()
	return Wrap(result)
}
func (f *factory) setIndex(i interface{}, index int, x interface{}) {
	defer utils.Recover(f.runtime)
	UnwrapValueOf(i).Index(index).Set(
		UnwrapValueOf(x),
	)
}
func (f *factory) slice(i interface{}, index int, x interface{}) {
	defer utils.Recover(f.runtime)

}
func (f *factory) getKey(i, key interface{}) (interface{}, bool) {
	defer utils.Recover(f.runtime)
	result := reflect.ValueOf(i).MapIndex(UnwrapValueOf(key))
	if !result.IsValid() {
		return nil, false
	}
	return Wrap(result), true
}
func (f *factory) setKey(i, key, elem interface{}) {
	defer utils.Recover(f.runtime)
	reflect.ValueOf(i).SetMapIndex(UnwrapValueOf(key), UnwrapValueOf(elem))
}
func (f *factory) deleteKey(i, key interface{}) {
	defer utils.Recover(f.runtime)
	reflect.ValueOf(i).SetMapIndex(UnwrapValueOf(key), reflect.Value{})
}
