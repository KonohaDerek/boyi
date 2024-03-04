package errors

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/pkg/errors"
)

// 自定义的 errors
var (
	ErrBadRequest               = &_error{Code: "400000", Message: http.StatusText(http.StatusBadRequest), Status: http.StatusBadRequest}
	ErrInvalidInput             = &_error{Code: "400001", Message: "One of the request inputs is not valid.", Status: http.StatusBadRequest}
	ErrResultBeenBound          = &_error{Code: "400002", Message: "The result already been bound, Need to check which resource has been used", Status: http.StatusBadRequest}
	ErrAccountAlreadyRegistered = &_error{Code: "400003", Message: "The username Already Registered.", Status: http.StatusBadRequest}
	ErrRoomAlreadyDeactivated   = &_error{Code: "400004", Message: "The consulting room already deactivated.", Status: http.StatusBadRequest}
	ErrSpinachAllowed           = &_error{Code: "400005", Message: "The request from Spinach has been refused or access is not allowed.", Status: http.StatusBadRequest}

	ErrUnauthorized                  = &_error{Code: "401001", Message: http.StatusText(http.StatusUnauthorized), Status: http.StatusUnauthorized}
	ErrTokenUnavailable              = &_error{Code: "401002", Message: "token not found", Status: http.StatusUnauthorized}
	ErrUsernameOrPasswordUnavailable = &_error{Code: "401003", Message: "username or password is unavailable", Status: http.StatusUnauthorized}
	ErrInvalidAuthenticationInfo     = &_error{Code: "401004", Message: "The authentication information was not provided in the correct format. Verify the value of Authorization header.", Status: http.StatusUnauthorized}
	ErrHostsDeny                     = &_error{Code: "401005", Message: "IP非法.", Status: http.StatusUnauthorized}

	// Spinach Payment
	ErrConnectPaymentServer = &_error{Code: "400201", Message: "暂时无法连接支付服务器!", Status: http.StatusServiceUnavailable}

	ErrTooManyRequests = &_error{Code: "402001", Message: "Too many requests.", Status: http.StatusTooManyRequests}

	ErrForbidden       = &_error{Code: "403000", Message: http.StatusText(http.StatusForbidden), Status: http.StatusForbidden}
	ErrNotAllowed      = &_error{Code: "403001", Message: "The request is understood, but it has been refused or access is not allowed.", Status: http.StatusForbidden}
	ErrAccountDisabled = &_error{Code: "403002", Message: "The account is been disabled.", Status: http.StatusForbidden}

	// ErrNotFound         =  &_error{Code: "404000", Message: http.StatusText(http.StatusNotFound), Status: http.StatusNotFound}
	ErrResourceNotFound       = &_error{Code: "404001", Message: "The specified resource does not exist.", Status: http.StatusNotFound}
	ErrResourceHasBeenDeleted = &_error{Code: "404002", Message: "The specified resource has been deleted.", Status: http.StatusNotFound}
	ErrMerchantOriginNotFound = &_error{Code: "404101", Message: "该域名尚未绑定商户", Status: http.StatusNotFound}

	ErrMethodNotAllowed = &_error{Code: "405001", Message: "Server has received and recognized the request, but has rejected the specific HTTP method it’s using.", Status: http.StatusMethodNotAllowed}

	ErrRequestTime = &_error{Code: "408001", Message: "request time out", Status: http.StatusRequestTimeout}

	ErrConflict              = &_error{Code: "409000", Message: http.StatusText(http.StatusConflict), Status: http.StatusConflict}
	ErrResourceAlreadyExists = &_error{Code: "409001", Message: "The specified resource already exists.", Status: http.StatusConflict}

	ErrContextCancel = &_error{Code: "499000", Message: "Client Closed Request", Status: 499}

	ErrInternalServerError = &_error{Code: "500000", Message: http.StatusText(http.StatusInternalServerError), Status: http.StatusInternalServerError}
	ErrInternalError       = &_error{Code: "500001", Message: "The server encountered an internal error. Please retry the request.", Status: http.StatusInternalServerError}
	Err3PartyInternalError = &_error{Code: "500002", Message: "The server 3 party encountered an internal error. Please retry the request.", Status: http.StatusInternalServerError}
)

type _error struct {
	Status  int                    `json:"status"`
	Code    string                 `json:"code"`
	Message string                 `json:"message"`
	Details map[string]interface{} `json:"details"`
}

// HttpError ...
type HttpError struct {
	Status  int                    `json:"-"`
	Code    string                 `json:"code"`
	Message string                 `json:"message"`
	Details map[string]interface{} `json:"details"`
}

func (e *_error) Error() string {
	var b strings.Builder
	_, _ = b.WriteRune('[')
	_, _ = b.WriteString(e.Code)
	_, _ = b.WriteRune(']')
	_, _ = b.WriteRune(' ')
	_, _ = b.WriteString(e.Message)
	return b.String()
}

// Is ...
func (e *_error) Is(target error) bool {
	causeErr := errors.Cause(target)
	tErr, ok := causeErr.(*_error)
	if !ok {
		return false
	}
	return e.Code == tErr.Code
}

// GetHTTPError ,,,
func GetHTTPError(err *_error) HttpError {
	return HttpError{
		Status:  err.Status,
		Message: err.Message,
		Code:    err.Code,
		Details: err.Details,
	}
}

// ConvertToHttpError 尝试从 _error 转换成 HttpError
func ConvertToHttpError(err error) HttpError {
	causeErr := errors.Cause(err)
	_err, ok := causeErr.(*_error)
	if !ok {
		return GetHTTPError(ErrInternalError)
	}
	return GetHTTPError(_err)
}

// NewWithMessage 抽换错误讯息
// 未定义的错误会被视为 ErrInternalError 类型
func NewWithMessage(err error, message string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	causeErr := errors.Cause(err)
	_err, ok := causeErr.(*_error)
	if !ok {
		return WithStack(&_error{
			Status:  ErrInternalError.Status,
			Code:    ErrInternalError.Code,
			Message: ErrInternalError.Message,
		})
	}
	err = &_error{
		Status:  _err.Status,
		Code:    _err.Code,
		Message: message,
	}
	var msg string
	for i := 0; i < len(args); i++ {
		msg += "%+v"
	}
	return Wrapf(err, msg, args...)
}

// WithErrors 使用订好的errors code 与讯息,如果未定义message 显示对应的http status描述
func WithErrors(err error) error {
	if err == nil {
		return nil
	}
	causeErr := errors.Cause(err)
	_err, ok := causeErr.(*_error)
	if !ok {
		return WithStack(&_error{
			Status:  ErrInternalError.Status,
			Code:    ErrInternalError.Code,
			Message: http.StatusText(ErrInternalError.Status),
		})
	}
	return WithStack(&_error{
		Status:  _err.Status,
		Code:    _err.Code,
		Message: _err.Message,
	})
}

// SetDetails set details as you wish =)
func (e *_error) SetDetails(details map[string]interface{}) {
	e.Details = details
}

// CompareErrorCode 比较两个错误代码是否一致
func CompareErrorCode(errA error, errB error) bool {
	var aErr, bErr *_error
	if err, exists := errors.Cause(errA).(*_error); exists {
		aErr = err
	}
	if err, exists := errors.Cause(errB).(*_error); exists {
		bErr = err
	}
	if aErr.Code == bErr.Code {
		return true
	}
	return false
}

// NewWithMessagef 抽换错误讯息
func NewWithMessagef(err error, format string, args ...interface{}) error {
	return NewWithMessage(err, fmt.Sprintf(format, args...))
}

// GetCodeWithErrors 使用订好的errors code 与讯息,如果未定义message 显示对应的http status描述
func GetCodeWithErrors(err error) (string, string) {

	causeErr := errors.Cause(err)
	_err, ok := causeErr.(*_error)
	if !ok {
		return ErrInternalError.Code, ErrInternalError.Message
	}
	return _err.Code, _err.Message
}

// HTTPConvertToError 将 http 的 response body convert to _error
func HTTPConvertToError(b []byte) error {
	interErr := _error{}
	jErr := json.Unmarshal(b, &interErr)
	if jErr != nil {
		return ErrInternalError
	}
	return WithStack(&interErr)
}

// GetStatusWithErrors 取得 error 的 http status code
func GetStatusWithErrors(err error) int {
	causeErr := errors.Cause(err)
	_err, ok := causeErr.(*_error)
	if !ok {
		return http.StatusInternalServerError
	}
	return _err.Status
}
