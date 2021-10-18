package http

import (
	"errors"
	"net/http"

	"github.com/powerpuffpenguin/goja"
)

type handler struct {
	runtime  *goja.Runtime
	callable goja.Callable
}

func newHandler(runtime *goja.Runtime, callable goja.Callable) http.Handler {
	return &handler{
		runtime:  runtime,
		callable: callable,
	}
}
func (h *handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	result := &handlerResult{
		runtime:  h.runtime,
		callable: h.callable,
		ch:       make(chan struct{}),
		w:        w,
		req:      req,
	}
	if !h.runtime.Loop().Result(result) {
		http.Error(w, `loop closed`, http.StatusInternalServerError)
		return
	}
	<-result.ch
	err := result.err
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

type handlerResult struct {
	runtime  *goja.Runtime
	callable goja.Callable
	ch       chan struct{}

	w   http.ResponseWriter
	req *http.Request
	err error
}

func (h *handlerResult) OnResult(closed bool) (completed bool) {
	runtime := h.runtime
	undefined := goja.Undefined()
	v, e := h.callable(goja.Undefined(), runtime.ToValue(h.w), runtime.ToValue(h.req))
	if e != nil {
		h.err = e
		close(h.ch)
		return
	}
	if obj, ok := v.(*goja.Object); ok {
		f, ok := goja.AssertFunction(obj.Get(`then`))
		if ok {
			f(undefined, runtime.ToValue(h.onfulfilled), runtime.ToValue(h.onrejected))
		}
	}
	return
}
func (h *handlerResult) onfulfilled(call goja.FunctionCall) goja.Value {
	defer close(h.ch)
	return goja.Undefined()
}
func (h *handlerResult) onrejected(call goja.FunctionCall) goja.Value {
	defer close(h.ch)
	h.err = errors.New(call.Argument(0).String())
	return goja.Undefined()
}
