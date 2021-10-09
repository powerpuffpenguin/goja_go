package errors

import (
	"errors"

	"github.com/powerpuffpenguin/goja"
)

func (f *factory) register() {
	f.Set(`GoError`, f.GoError)
	f.Set(`Wrap`, Wrap)
	f.Set(`New`, errors.New)
	f.Set(`Unwrap`, errors.Unwrap)
	f.Set(`Is`, errors.Is)
	f.Set(`As`, errors.As)

	f.Set(`isError`, isError)
}
func isError(i interface{}) bool {
	_, result := i.(error)
	return result
}
func (f *factory) GoError(err error) *goja.Object {
	return f.runtime.NewGoError(err)
}
func Wrap(text string, err error) error {
	return &wrapError{
		msg: text,
		err: err,
	}
}

type wrapError struct {
	msg string
	err error
}

func (e *wrapError) Error() string {
	return e.msg
}

func (e *wrapError) Unwrap() error {
	return e.err
}
