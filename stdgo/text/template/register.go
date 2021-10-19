package template

import "text/template"

func (f *factory) register() {
	f.Set(`HTMLEscape`, template.HTMLEscape)
	f.Set(`HTMLEscapeString`, template.HTMLEscapeString)
	f.Set(`HTMLEscaper`, template.HTMLEscaper)
	f.Set(`IsTrue`, template.IsTrue)
	f.Set(`JSEscape`, template.JSEscape)
	f.Set(`JSEscapeString`, template.JSEscapeString)
	f.Set(`JSEscaper`, template.JSEscaper)
	f.Set(`URLQueryEscaper`, template.URLQueryEscaper)

	f.Set(`FuncMap`, FuncMap)
	f.Set(`Must`, template.Must)
	f.Set(`New`, template.New)
	f.Set(`ParseFS`, template.ParseFS)
	f.Set(`ParseFiles`, template.ParseFiles)
	f.Set(`ParseGlob`, template.ParseGlob)
	f.Set(`isExecError`, isExecError)
	f.Set(`isFuncMap`, isFuncMap)
	f.Set(`isTemplatePointer`, isTemplatePointer)
}
func isExecError(i interface{}) bool {
	_, result := i.(template.ExecError)
	return result
}
func isFuncMap(i interface{}) bool {
	_, result := i.(template.FuncMap)
	return result
}
func isTemplatePointer(i interface{}) bool {
	_, result := i.(*template.Template)
	return result
}
func FuncMap() template.FuncMap {
	return make(template.FuncMap)
}
