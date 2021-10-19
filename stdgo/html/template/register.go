package template

import (
	"html/template"

	"github.com/powerpuffpenguin/goja"
)

func (f *factory) register() {
	f.Set(`HTMLEscape`, template.HTMLEscape)
	f.Set(`HTMLEscapeString`, template.HTMLEscapeString)
	f.Set(`HTMLEscaper`, template.HTMLEscaper)
	f.Set(`IsTrue`, template.IsTrue)
	f.Set(`JSEscape`, template.JSEscape)
	f.Set(`JSEscapeString`, template.JSEscapeString)
	f.Set(`JSEscaper`, template.JSEscaper)
	f.Set(`URLQueryEscaper`, template.URLQueryEscaper)

	f.Set(`ErrorCode`, ErrorCode)
	f.Accessor(`OK`, f.getOK, nil)
	f.Accessor(`ErrAmbigContext`, f.getErrAmbigContext, nil)
	f.Accessor(`ErrBadHTML`, f.getErrBadHTML, nil)
	f.Accessor(`ErrBranchEnd`, f.getErrBranchEnd, nil)
	f.Accessor(`ErrEndContext`, f.getErrEndContext, nil)
	f.Accessor(`ErrNoSuchTemplate`, f.getErrNoSuchTemplate, nil)
	f.Accessor(`ErrOutputContext`, f.getErrOutputContext, nil)
	f.Accessor(`ErrPartialCharset`, f.getErrPartialCharset, nil)
	f.Accessor(`ErrPartialEscape`, f.getErrPartialEscape, nil)
	f.Accessor(`ErrRangeLoopReentry`, f.getErrRangeLoopReentry, nil)
	f.Accessor(`ErrSlashAmbig`, f.getErrSlashAmbig, nil)
	f.Accessor(`ErrPredefinedEscaper`, f.getErrPredefinedEscaper, nil)
	f.Set(`FuncMap`, FuncMap)
	f.Set(`Must`, template.Must)
	f.Set(`New`, template.New)
	f.Set(`ParseFS`, template.ParseFS)
	f.Set(`ParseFiles`, template.ParseFiles)
	f.Set(`ParseGlob`, template.ParseGlob)
	f.Set(`isCSS`, isCSS)
	f.Set(`isErrorPointer`, isErrorPointer)
	f.Set(`isErrorCode`, isErrorCode)
	f.Set(`isFuncMap`, isFuncMap)
	f.Set(`isHTML`, isHTML)
	f.Set(`isHTMLAttr`, isHTMLAttr)
	f.Set(`isJS`, isJS)
	f.Set(`isJSStr`, isJSStr)
	f.Set(`isSrcset`, isSrcset)
	f.Set(`isTemplatePointer`, isTemplatePointer)
	f.Set(`isURL`, isURL)
}
func isCSS(i interface{}) bool {
	_, result := i.(template.CSS)
	return result
}
func isErrorPointer(i interface{}) bool {
	_, result := i.(*template.Error)
	return result
}
func isErrorCode(i interface{}) bool {
	_, result := i.(template.ErrorCode)
	return result
}
func isFuncMap(i interface{}) bool {
	_, result := i.(template.FuncMap)
	return result
}
func isHTML(i interface{}) bool {
	_, result := i.(template.HTML)
	return result
}
func isHTMLAttr(i interface{}) bool {
	_, result := i.(template.HTMLAttr)
	return result
}
func isJS(i interface{}) bool {
	_, result := i.(template.JS)
	return result
}
func isJSStr(i interface{}) bool {
	_, result := i.(template.JSStr)
	return result
}
func isSrcset(i interface{}) bool {
	_, result := i.(template.Srcset)
	return result
}
func isTemplatePointer(i interface{}) bool {
	_, result := i.(*template.Template)
	return result
}
func isURL(i interface{}) bool {
	_, result := i.(template.URL)
	return result
}
func FuncMap() template.FuncMap {
	return make(template.FuncMap)
}
func (f *factory) getOK(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(template.OK)
}
func (f *factory) getErrAmbigContext(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(template.ErrAmbigContext)
}
func (f *factory) getErrBadHTML(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(template.ErrBadHTML)
}
func (f *factory) getErrBranchEnd(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(template.ErrBranchEnd)
}
func (f *factory) getErrEndContext(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(template.ErrEndContext)
}
func (f *factory) getErrNoSuchTemplate(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(template.ErrNoSuchTemplate)
}
func (f *factory) getErrOutputContext(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(template.ErrOutputContext)
}
func (f *factory) getErrPartialCharset(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(template.ErrPartialCharset)
}
func (f *factory) getErrPartialEscape(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(template.ErrPartialEscape)
}
func (f *factory) getErrRangeLoopReentry(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(template.ErrRangeLoopReentry)
}
func (f *factory) getErrSlashAmbig(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(template.ErrSlashAmbig)
}
func (f *factory) getErrPredefinedEscaper(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(template.ErrPredefinedEscaper)
}
func ErrorCode(v int) template.ErrorCode {
	return template.ErrorCode(v)
}
