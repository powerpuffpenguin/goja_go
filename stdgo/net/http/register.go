package http

import (
	"net/http"

	"github.com/powerpuffpenguin/goja"
)

func (f *factory) register() {
	f.Set(`Client`, Client)
	f.Accessor(`StateNew`, f.getStateNew, nil)
	f.Accessor(`StateActive`, f.getStateActive, nil)
	f.Accessor(`StateIdle`, f.getStateIdle, nil)
	f.Accessor(`StateHijacked`, f.getStateHijacked, nil)
	f.Accessor(`StateClosed`, f.getStateClosed, nil)
	f.Set(`FS`, http.FS)
	f.Set(`FileServer`, http.FileServer)
	f.Set(`NotFoundHandler`, http.NotFoundHandler)
	f.Set(`RedirectHandler`, http.RedirectHandler)
	f.Set(`StripPrefix`, http.StripPrefix)
	f.Set(`TimeoutHandler`, http.TimeoutHandler)
	f.Set(`Header`, Header)
	f.Set(`PushOptions`, PushOptions)
	f.Set(`NewRequest`, http.NewRequest)
	f.Set(`NewRequestWithContext`, http.NewRequestWithContext)
	f.Set(`ReadRequest`, http.ReadRequest)
	f.Set(`Get`, http.Get)
	f.Set(`Head`, http.Head)
	f.Set(`Post`, http.Post)
	f.Set(`PostForm`, http.PostForm)
	f.Set(`ReadResponse`, http.ReadResponse)
	f.Set(`NewFileTransport`, http.NewFileTransport)
	f.Accessor(`SameSiteDefaultMode`, f.getSameSiteDefaultMode, nil)
	f.Accessor(`SameSiteLaxMode`, f.getSameSiteLaxMode, nil)
	f.Accessor(`SameSiteStrictMode`, f.getSameSiteStrictMode, nil)
	f.Accessor(`SameSiteNoneMode`, f.getSameSiteNoneMode, nil)
	f.Set(`NewServeMux`, http.NewServeMux)
}
func (f *factory) getSameSiteDefaultMode(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(http.SameSiteDefaultMode)
}
func (f *factory) getSameSiteLaxMode(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(http.SameSiteLaxMode)
}
func (f *factory) getSameSiteStrictMode(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(http.SameSiteStrictMode)
}
func (f *factory) getSameSiteNoneMode(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(http.SameSiteNoneMode)
}
func PushOptions(method string, header http.Header) *http.PushOptions {
	return &http.PushOptions{
		Method: method,
		Header: header,
	}
}
func Header() http.Header {
	return make(http.Header)
}
func Cookie(name, value string) *http.Cookie {
	return &http.Cookie{
		Name:  name,
		Value: value,
	}
}
func (f *factory) getStateNew(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(http.StateNew)
}
func (f *factory) getStateActive(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(http.StateActive)
}
func (f *factory) getStateIdle(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(http.StateIdle)
}
func (f *factory) getStateHijacked(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(http.StateHijacked)
}
func (f *factory) getStateClosed(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(http.StateClosed)
}
func Client() *http.Client {
	return &http.Client{}
}
