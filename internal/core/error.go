package core

type ErrorType string
type ErrorCode string

const (
	InvalidRequestError     ErrorType = "invalid_request" // 请求错误
	InvalidRequestErrorCode ErrorCode = "1000"
	InvalidArgument         ErrorType = "invalid_argument" // 参数错误
	InvalidArgumentCode     ErrorCode = "1001"
	InternalServerError     ErrorType = "internal_server_error" // 内部服务器错误
	InternalServerErrorCode ErrorCode = "1002"
)

type AppError struct {
	Type    ErrorType
	Code    ErrorCode
	Message string
	Err     error
}

func NewInvalidRequestError(err error) *AppError {
	return &AppError{
		Type:    InvalidRequestError,
		Code:    InvalidRequestErrorCode,
		Message: "invalid request",
		Err:     err,
	}
}

func (e *AppError) Error() string {
	return e.Message
}

func (e *AppError) Unwrap() error {
	return e.Err
}
