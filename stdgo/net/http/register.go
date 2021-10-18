package http

import (
	"errors"
	"net/http"

	"github.com/powerpuffpenguin/goja"
	"github.com/powerpuffpenguin/goja_go/stdgo/builtin"
)

func (f *factory) register() {
	f.Accessor(`MethodGet`, f.getMethodGet, nil)
	f.Accessor(`MethodHead`, f.getMethodHead, nil)
	f.Accessor(`MethodPost`, f.getMethodPost, nil)
	f.Accessor(`MethodPut`, f.getMethodPut, nil)
	f.Accessor(`MethodPatch`, f.getMethodPatch, nil)
	f.Accessor(`MethodDelete`, f.getMethodDelete, nil)
	f.Accessor(`MethodConnect`, f.getMethodConnect, nil)
	f.Accessor(`MethodOptions`, f.getMethodOptions, nil)
	f.Accessor(`MethodTrace`, f.getMethodTrace, nil)

	f.Accessor(`StatusContinue`, f.getStatusContinue, nil)
	f.Accessor(`StatusSwitchingProtocols`, f.getStatusSwitchingProtocols, nil)
	f.Accessor(`StatusProcessing`, f.getStatusProcessing, nil)
	f.Accessor(`StatusEarlyHints`, f.getStatusEarlyHints, nil)
	f.Accessor(`StatusOK`, f.getStatusOK, nil)
	f.Accessor(`StatusCreated`, f.getStatusCreated, nil)
	f.Accessor(`StatusAccepted`, f.getStatusAccepted, nil)
	f.Accessor(`StatusNonAuthoritativeInfo`, f.getStatusNonAuthoritativeInfo, nil)
	f.Accessor(`StatusNoContent`, f.getStatusNoContent, nil)
	f.Accessor(`StatusResetContent`, f.getStatusResetContent, nil)
	f.Accessor(`StatusPartialContent`, f.getStatusPartialContent, nil)
	f.Accessor(`StatusMultiStatus`, f.getStatusMultiStatus, nil)
	f.Accessor(`StatusAlreadyReported`, f.getStatusAlreadyReported, nil)
	f.Accessor(`StatusIMUsed`, f.getStatusIMUsed, nil)
	f.Accessor(`StatusMultipleChoices`, f.getStatusMultipleChoices, nil)
	f.Accessor(`StatusMovedPermanently`, f.getStatusMovedPermanently, nil)
	f.Accessor(`StatusFound`, f.getStatusFound, nil)
	f.Accessor(`StatusSeeOther`, f.getStatusSeeOther, nil)
	f.Accessor(`StatusNotModified`, f.getStatusNotModified, nil)
	f.Accessor(`StatusUseProxy`, f.getStatusUseProxy, nil)
	f.Accessor(`StatusTemporaryRedirect`, f.getStatusTemporaryRedirect, nil)
	f.Accessor(`StatusPermanentRedirect`, f.getStatusPermanentRedirect, nil)
	f.Accessor(`StatusBadRequest`, f.getStatusBadRequest, nil)
	f.Accessor(`StatusUnauthorized`, f.getStatusUnauthorized, nil)
	f.Accessor(`StatusPaymentRequired`, f.getStatusPaymentRequired, nil)
	f.Accessor(`StatusForbidden`, f.getStatusForbidden, nil)
	f.Accessor(`StatusNotFound`, f.getStatusNotFound, nil)
	f.Accessor(`StatusMethodNotAllowed`, f.getStatusMethodNotAllowed, nil)
	f.Accessor(`StatusNotAcceptable`, f.getStatusNotAcceptable, nil)
	f.Accessor(`StatusProxyAuthRequired`, f.getStatusProxyAuthRequired, nil)
	f.Accessor(`StatusRequestTimeout`, f.getStatusRequestTimeout, nil)
	f.Accessor(`StatusConflict`, f.getStatusConflict, nil)
	f.Accessor(`StatusGone`, f.getStatusGone, nil)
	f.Accessor(`StatusLengthRequired`, f.getStatusLengthRequired, nil)
	f.Accessor(`StatusPreconditionFailed`, f.getStatusPreconditionFailed, nil)
	f.Accessor(`StatusRequestEntityTooLarge`, f.getStatusRequestEntityTooLarge, nil)
	f.Accessor(`StatusRequestURITooLong`, f.getStatusRequestURITooLong, nil)
	f.Accessor(`StatusUnsupportedMediaType`, f.getStatusUnsupportedMediaType, nil)
	f.Accessor(`StatusRequestedRangeNotSatisfiable`, f.getStatusRequestedRangeNotSatisfiable, nil)
	f.Accessor(`StatusExpectationFailed`, f.getStatusExpectationFailed, nil)
	f.Accessor(`StatusTeapot`, f.getStatusTeapot, nil)
	f.Accessor(`StatusMisdirectedRequest`, f.getStatusMisdirectedRequest, nil)
	f.Accessor(`StatusUnprocessableEntity`, f.getStatusUnprocessableEntity, nil)
	f.Accessor(`StatusLocked`, f.getStatusLocked, nil)
	f.Accessor(`StatusFailedDependency`, f.getStatusFailedDependency, nil)
	f.Accessor(`StatusTooEarly`, f.getStatusTooEarly, nil)
	f.Accessor(`StatusUpgradeRequired`, f.getStatusUpgradeRequired, nil)
	f.Accessor(`StatusPreconditionRequired`, f.getStatusPreconditionRequired, nil)
	f.Accessor(`StatusTooManyRequests`, f.getStatusTooManyRequests, nil)
	f.Accessor(`StatusRequestHeaderFieldsTooLarge`, f.getStatusRequestHeaderFieldsTooLarge, nil)
	f.Accessor(`StatusUnavailableForLegalReasons`, f.getStatusUnavailableForLegalReasons, nil)
	f.Accessor(`StatusInternalServerError`, f.getStatusInternalServerError, nil)
	f.Accessor(`StatusNotImplemented`, f.getStatusNotImplemented, nil)
	f.Accessor(`StatusBadGateway`, f.getStatusBadGateway, nil)
	f.Accessor(`StatusServiceUnavailable`, f.getStatusServiceUnavailable, nil)
	f.Accessor(`StatusGatewayTimeout`, f.getStatusGatewayTimeout, nil)
	f.Accessor(`StatusHTTPVersionNotSupported`, f.getStatusHTTPVersionNotSupported, nil)
	f.Accessor(`StatusVariantAlsoNegotiates`, f.getStatusVariantAlsoNegotiates, nil)
	f.Accessor(`StatusInsufficientStorage`, f.getStatusInsufficientStorage, nil)
	f.Accessor(`StatusLoopDetected`, f.getStatusLoopDetected, nil)
	f.Accessor(`StatusNotExtended`, f.getStatusNotExtended, nil)
	f.Accessor(`StatusNetworkAuthenticationRequired`, f.getStatusNetworkAuthenticationRequired, nil)

	f.Accessor(`DefaultMaxHeaderBytes`, f.getDefaultMaxHeaderBytes, nil)
	f.Accessor(`DefaultMaxIdleConnsPerHost`, f.getDefaultMaxIdleConnsPerHost, nil)
	f.Accessor(`TimeFormat`, f.getTimeFormat, nil)
	f.Accessor(`TrailerPrefix`, f.getTrailerPrefix, nil)

	f.Accessor(`ErrNotSupported`, f.getErrNotSupported, nil)
	f.Accessor(`ErrMissingBoundary`, f.getErrMissingBoundary, nil)
	f.Accessor(`ErrNotMultipart`, f.getErrNotMultipart, nil)
	f.Accessor(`ErrHeaderTooLong`, f.getErrHeaderTooLong, nil)
	f.Accessor(`ErrShortBody`, f.getErrShortBody, nil)
	f.Accessor(`ErrMissingContentLength`, f.getErrMissingContentLength, nil)
	f.Accessor(`ErrBodyNotAllowed`, f.getErrBodyNotAllowed, nil)
	f.Accessor(`ErrHijacked`, f.getErrHijacked, nil)
	f.Accessor(`ErrContentLength`, f.getErrContentLength, nil)
	f.Accessor(`ErrWriteAfterFlush`, f.getErrWriteAfterFlush, nil)

	f.Accessor(`ServerContextKey`, f.getServerContextKey, nil)
	f.Accessor(`LocalAddrContextKey`, f.getLocalAddrContextKey, nil)

	f.Accessor(`DefaultClient`, f.getDefaultClient, nil)
	f.Accessor(`DefaultServeMux`, f.getDefaultServeMux, nil)
	f.Accessor(`ErrAbortHandler`, f.getErrAbortHandler, nil)
	f.Accessor(`ErrBodyReadAfterClose`, f.getErrBodyReadAfterClose, nil)
	f.Accessor(`ErrHandlerTimeout`, f.getErrHandlerTimeout, nil)
	f.Accessor(`ErrLineTooLong`, f.getErrLineTooLong, nil)
	f.Accessor(`ErrMissingFile`, f.getErrMissingFile, nil)
	f.Accessor(`ErrNoCookie`, f.getErrNoCookie, nil)
	f.Accessor(`ErrNoLocation`, f.getErrNoLocation, nil)
	f.Accessor(`ErrServerClosed`, f.getErrServerClosed, nil)
	f.Accessor(`ErrSkipAltProtocol`, f.getErrSkipAltProtocol, nil)
	f.Accessor(`ErrUseLastResponse`, f.getErrUseLastResponse, nil)
	f.Accessor(`NoBody`, f.getNoBody, nil)

	f.Set(`CanonicalHeaderKey`, http.CanonicalHeaderKey)
	f.Set(`DetectContentType`, http.DetectContentType)
	f.Set(`Error`, http.Error)
	f.Set(`Handle`, http.Handle)
	f.Set(`HandleFunc`, http.HandleFunc)
	f.Set(`ListenAndServe`, http.ListenAndServe)
	f.Set(`ListenAndServeTLS`, http.ListenAndServeTLS)
	f.Set(`MaxBytesReader`, http.MaxBytesReader)
	f.Set(`NotFound`, http.NotFound)
	f.Set(`ParseHTTPVersion`, http.ParseHTTPVersion)
	f.Set(`ParseTime`, http.ParseTime)
	f.Set(`ProxyFromEnvironment`, http.ProxyFromEnvironment)
	f.Set(`ProxyURL`, http.ProxyURL)
	f.Set(`Redirect`, http.Redirect)
	f.Set(`Serve`, http.Serve)
	f.Set(`ServeContent`, http.ServeContent)
	f.Set(`ServeFile`, http.ServeFile)
	f.Set(`ServeTLS`, http.ServeTLS)
	f.Set(`SetCookie`, http.SetCookie)
	f.Set(`StatusText`, http.StatusText)

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
	f.Set(`Server`, Server)
	f.Set(`Transport`, Transport)

	f.Set(`isClientPointer`, isClientPointer)
	f.Set(`isConnState`, isConnState)
	f.Set(`isCookiePointer`, isCookiePointer)
	f.Set(`isDir`, isDir)
	f.Set(`isFile`, isFile)
	f.Set(`isFileSystem`, isFileSystem)
	f.Set(`isFlusher`, isFlusher)
	f.Set(`isHandler`, isHandler)
	f.Set(`isHandlerFunc`, isHandlerFunc)
	f.Set(`isHeader`, isHeader)
	f.Set(`isHijacker`, isHijacker)
	f.Set(`isPushOptionsPointer`, isPushOptionsPointer)
	f.Set(`isPusher`, isPusher)
	f.Set(`isRequestPointer`, isRequestPointer)
	f.Set(`isResponsePointer`, isResponsePointer)
	f.Set(`isResponseWriter`, isResponseWriter)
	f.Set(`isRoundTripper`, isRoundTripper)
	f.Set(`isSameSite`, isSameSite)
	f.Set(`isServeMuxPointer`, isServeMuxPointer)
	f.Set(`isServerPointer`, isServerPointer)

	f.Set(`Handler`, f.Handler)
}
func (f *factory) Handler(call goja.FunctionCall) goja.Value {
	callable, ok := goja.AssertFunction(call.Argument(0))
	if !ok {
		f.runtime.NewGoError(errors.New(`Handler arg0 not callable`))
	}
	return f.runtime.ToValue(newHandler(f.runtime, callable))
}
func (f *factory) getDefaultClient(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(http.DefaultClient)
}
func (f *factory) getDefaultServeMux(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(http.DefaultServeMux)
}
func (f *factory) getErrAbortHandler(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(http.ErrAbortHandler)
}
func (f *factory) getErrBodyReadAfterClose(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(http.ErrBodyReadAfterClose)
}
func (f *factory) getErrHandlerTimeout(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(http.ErrHandlerTimeout)
}
func (f *factory) getErrLineTooLong(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(http.ErrLineTooLong)
}
func (f *factory) getErrMissingFile(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(http.ErrMissingFile)
}
func (f *factory) getErrNoCookie(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(http.ErrNoCookie)
}
func (f *factory) getErrNoLocation(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(http.ErrNoLocation)
}
func (f *factory) getErrServerClosed(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(http.ErrServerClosed)
}
func (f *factory) getErrSkipAltProtocol(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(http.ErrSkipAltProtocol)
}
func (f *factory) getErrUseLastResponse(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(http.ErrUseLastResponse)
}
func (f *factory) getNoBody(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(http.NoBody)
}
func (f *factory) getServerContextKey(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(http.ServerContextKey)
}
func (f *factory) getLocalAddrContextKey(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(http.LocalAddrContextKey)
}
func (f *factory) getErrNotSupported(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(http.ErrNotSupported)
}
func (f *factory) getErrMissingBoundary(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(http.ErrMissingBoundary)
}
func (f *factory) getErrNotMultipart(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(http.ErrNotMultipart)
}
func (f *factory) getErrHeaderTooLong(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(http.ErrHeaderTooLong)
}
func (f *factory) getErrShortBody(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(http.ErrShortBody)
}
func (f *factory) getErrMissingContentLength(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(http.ErrMissingContentLength)
}
func (f *factory) getErrBodyNotAllowed(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(http.ErrBodyNotAllowed)
}
func (f *factory) getErrHijacked(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(http.ErrHijacked)
}
func (f *factory) getErrContentLength(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(http.ErrContentLength)
}
func (f *factory) getErrWriteAfterFlush(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(http.ErrWriteAfterFlush)
}
func (f *factory) getDefaultMaxHeaderBytes(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(http.DefaultMaxHeaderBytes))
}
func (f *factory) getDefaultMaxIdleConnsPerHost(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(http.DefaultMaxIdleConnsPerHost))
}
func (f *factory) getTimeFormat(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(http.TimeFormat)
}
func (f *factory) getTrailerPrefix(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(http.TrailerPrefix)
}
func (f *factory) getStatusContinue(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(http.StatusContinue))
}
func (f *factory) getStatusSwitchingProtocols(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(http.StatusSwitchingProtocols))
}
func (f *factory) getStatusProcessing(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(http.StatusProcessing))
}
func (f *factory) getStatusEarlyHints(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(http.StatusEarlyHints))
}
func (f *factory) getStatusOK(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(http.StatusOK))
}
func (f *factory) getStatusCreated(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(http.StatusCreated))
}
func (f *factory) getStatusAccepted(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(http.StatusAccepted))
}
func (f *factory) getStatusNonAuthoritativeInfo(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(http.StatusNonAuthoritativeInfo))
}
func (f *factory) getStatusNoContent(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(http.StatusNoContent))
}
func (f *factory) getStatusResetContent(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(http.StatusResetContent))
}
func (f *factory) getStatusPartialContent(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(http.StatusPartialContent))
}
func (f *factory) getStatusMultiStatus(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(http.StatusMultiStatus))
}
func (f *factory) getStatusAlreadyReported(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(http.StatusAlreadyReported))
}
func (f *factory) getStatusIMUsed(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(http.StatusIMUsed))
}
func (f *factory) getStatusMultipleChoices(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(http.StatusMultipleChoices))
}
func (f *factory) getStatusMovedPermanently(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(http.StatusMovedPermanently))
}
func (f *factory) getStatusFound(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(http.StatusFound))
}
func (f *factory) getStatusSeeOther(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(http.StatusSeeOther))
}
func (f *factory) getStatusNotModified(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(http.StatusNotModified))
}
func (f *factory) getStatusUseProxy(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(http.StatusUseProxy))
}
func (f *factory) getStatusTemporaryRedirect(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(http.StatusTemporaryRedirect))
}
func (f *factory) getStatusPermanentRedirect(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(http.StatusPermanentRedirect))
}
func (f *factory) getStatusBadRequest(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(http.StatusBadRequest))
}
func (f *factory) getStatusUnauthorized(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(http.StatusUnauthorized))
}
func (f *factory) getStatusPaymentRequired(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(http.StatusPaymentRequired))
}
func (f *factory) getStatusForbidden(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(http.StatusForbidden))
}
func (f *factory) getStatusNotFound(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(http.StatusNotFound))
}
func (f *factory) getStatusMethodNotAllowed(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(http.StatusMethodNotAllowed))
}
func (f *factory) getStatusNotAcceptable(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(http.StatusNotAcceptable))
}
func (f *factory) getStatusProxyAuthRequired(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(http.StatusProxyAuthRequired))
}
func (f *factory) getStatusRequestTimeout(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(http.StatusRequestTimeout))
}
func (f *factory) getStatusConflict(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(http.StatusConflict))
}
func (f *factory) getStatusGone(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(http.StatusGone))
}
func (f *factory) getStatusLengthRequired(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(http.StatusLengthRequired))
}
func (f *factory) getStatusPreconditionFailed(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(http.StatusPreconditionFailed))
}
func (f *factory) getStatusRequestEntityTooLarge(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(http.StatusRequestEntityTooLarge))
}
func (f *factory) getStatusRequestURITooLong(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(http.StatusRequestURITooLong))
}
func (f *factory) getStatusUnsupportedMediaType(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(http.StatusUnsupportedMediaType))
}
func (f *factory) getStatusRequestedRangeNotSatisfiable(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(http.StatusRequestedRangeNotSatisfiable))
}
func (f *factory) getStatusExpectationFailed(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(http.StatusExpectationFailed))
}
func (f *factory) getStatusTeapot(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(http.StatusTeapot))
}
func (f *factory) getStatusMisdirectedRequest(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(http.StatusMisdirectedRequest))
}
func (f *factory) getStatusUnprocessableEntity(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(http.StatusUnprocessableEntity))
}
func (f *factory) getStatusLocked(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(http.StatusLocked))
}
func (f *factory) getStatusFailedDependency(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(http.StatusFailedDependency))
}
func (f *factory) getStatusTooEarly(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(http.StatusTooEarly))
}
func (f *factory) getStatusUpgradeRequired(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(http.StatusUpgradeRequired))
}
func (f *factory) getStatusPreconditionRequired(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(http.StatusPreconditionRequired))
}
func (f *factory) getStatusTooManyRequests(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(http.StatusTooManyRequests))
}
func (f *factory) getStatusRequestHeaderFieldsTooLarge(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(http.StatusRequestHeaderFieldsTooLarge))
}
func (f *factory) getStatusUnavailableForLegalReasons(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(http.StatusUnavailableForLegalReasons))
}
func (f *factory) getStatusInternalServerError(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(http.StatusInternalServerError))
}
func (f *factory) getStatusNotImplemented(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(http.StatusNotImplemented))
}
func (f *factory) getStatusBadGateway(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(http.StatusBadGateway))
}
func (f *factory) getStatusServiceUnavailable(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(http.StatusServiceUnavailable))
}
func (f *factory) getStatusGatewayTimeout(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(http.StatusGatewayTimeout))
}
func (f *factory) getStatusHTTPVersionNotSupported(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(http.StatusHTTPVersionNotSupported))
}
func (f *factory) getStatusVariantAlsoNegotiates(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(http.StatusVariantAlsoNegotiates))
}
func (f *factory) getStatusInsufficientStorage(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(http.StatusInsufficientStorage))
}
func (f *factory) getStatusLoopDetected(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(http.StatusLoopDetected))
}
func (f *factory) getStatusNotExtended(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(http.StatusNotExtended))
}
func (f *factory) getStatusNetworkAuthenticationRequired(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(builtin.Int(http.StatusNetworkAuthenticationRequired))
}
func (f *factory) getMethodGet(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(http.MethodGet)
}
func (f *factory) getMethodHead(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(http.MethodHead)
}
func (f *factory) getMethodPost(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(http.MethodPost)
}
func (f *factory) getMethodPut(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(http.MethodPut)
}
func (f *factory) getMethodPatch(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(http.MethodPatch)
}
func (f *factory) getMethodDelete(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(http.MethodDelete)
}
func (f *factory) getMethodConnect(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(http.MethodConnect)
}
func (f *factory) getMethodOptions(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(http.MethodOptions)
}
func (f *factory) getMethodTrace(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(http.MethodTrace)
}
func isClientPointer(i interface{}) bool {
	_, result := i.(*http.Client)
	return result
}
func isConnState(i interface{}) bool {
	_, result := i.(http.ConnState)
	return result
}
func isCookiePointer(i interface{}) bool {
	_, result := i.(*http.Cookie)
	return result
}
func isDir(i interface{}) bool {
	_, result := i.(http.Dir)
	return result
}
func isFile(i interface{}) bool {
	_, result := i.(http.File)
	return result
}
func isFileSystem(i interface{}) bool {
	_, result := i.(http.FileSystem)
	return result
}
func isFlusher(i interface{}) bool {
	_, result := i.(http.Flusher)
	return result
}
func isHandler(i interface{}) bool {
	_, result := i.(http.Handler)
	return result
}
func isHandlerFunc(i interface{}) bool {
	_, result := i.(http.HandlerFunc)
	return result
}
func isHeader(i interface{}) bool {
	_, result := i.(http.Header)
	return result
}
func isHijacker(i interface{}) bool {
	_, result := i.(http.Hijacker)
	return result
}
func isPushOptionsPointer(i interface{}) bool {
	_, result := i.(*http.PushOptions)
	return result
}
func isPusher(i interface{}) bool {
	_, result := i.(http.Pusher)
	return result
}
func isRequestPointer(i interface{}) bool {
	_, result := i.(*http.Request)
	return result
}
func isResponsePointer(i interface{}) bool {
	_, result := i.(*http.Response)
	return result
}
func isResponseWriter(i interface{}) bool {
	_, result := i.(http.ResponseWriter)
	return result
}
func isRoundTripper(i interface{}) bool {
	_, result := i.(http.RoundTripper)
	return result
}
func isSameSite(i interface{}) bool {
	_, result := i.(http.SameSite)
	return result
}
func isServeMuxPointer(i interface{}) bool {
	_, result := i.(*http.ServeMux)
	return result
}
func isServerPointer(i interface{}) bool {
	_, result := i.(*http.Server)
	return result
}
func Transport() *http.Transport {
	return &http.Transport{}
}
func Server() *http.Server {
	return &http.Server{}
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
