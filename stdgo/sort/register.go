package sort

import "sort"

func (f *factory) register() {
	f.Set(`Float64s`, sort.Float64s)
	f.Set(`Float64sAreSorted`, sort.Float64sAreSorted)
	f.Set(`Ints`, sort.Ints)
	f.Set(`IntsAreSorted`, sort.IntsAreSorted)
	f.Set(`IsSorted`, sort.IsSorted)
	f.Set(`Search`, sort.Search)
	f.Set(`SearchFloat64s`, sort.SearchFloat64s)
	f.Set(`SearchInts`, sort.SearchInts)
	f.Set(`SearchStrings`, sort.SearchStrings)
	f.Set(`Slice`, sort.Slice)
	f.Set(`SliceIsSorted`, sort.SliceIsSorted)
	f.Set(`SliceStable`, sort.SliceStable)
	f.Set(`Sort`, sort.Sort)
	f.Set(`Stable`, sort.Stable)
	f.Set(`Strings`, sort.Strings)
	f.Set(`StringsAreSorted`, sort.StringsAreSorted)

	f.Set(`Float64Slice`, Float64Slice)
	f.Set(`IntSlice`, IntSlice)
	f.Set(`Reverse`, sort.Reverse)
	f.Set(`StringSlice`, StringSlice)

	f.Set(`isFloat64Slice`, isFloat64Slice)
	f.Set(`isInterface`, isInterface)
	f.Set(`isIntSlice`, isIntSlice)
	f.Set(`isStringSlice`, isStringSlice)
}

func isFloat64Slice(i interface{}) bool {
	_, result := i.(sort.Float64Slice)
	return result
}
func isInterface(i interface{}) bool {
	_, result := i.(sort.Interface)
	return result
}
func isIntSlice(i interface{}) bool {
	_, result := i.(sort.IntSlice)
	return result
}
func isStringSlice(i interface{}) bool {
	_, result := i.(sort.StringSlice)
	return result
}
func Float64Slice(v []float64) sort.Float64Slice {
	return sort.Float64Slice(v)
}
func IntSlice(v []int) sort.IntSlice {
	return sort.IntSlice(v)
}
func StringSlice(v []string) sort.StringSlice {
	return sort.StringSlice(v)
}
