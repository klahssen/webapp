package errors

import (
	"fmt"
	"net/http"

	"github.com/klahssen/webapp/pkg/log"
	"google.golang.org/grpc/codes"
)

//Error implements the error interface
type Error struct {
	code codes.Code //grpc status code
	msg  string
}

func New(msg string, code codes.Code) *Error {
	return &Error{code: code, msg: msg}
}

//GetStatusCode returns http StatusCode from grpc code
func (e *Error) GetStatusCode() int {
	return httpStatusFromGrpcCode(e.code)
}

//GetStatusCode returns http StatusCode from grpc code
func (e *Error) GetGrpcCode() codes.Code {
	return e.code
}

func (e *Error) Error() string {
	if e == nil {
		return ""
	}
	return fmt.Sprintf("%s", e.msg)
}

//httpStatusFromGrpcCode converts a grpc status code to http.StatusCode
func httpStatusFromGrpcCode(code codes.Code) int {
	switch code {
	case codes.OK:
		return http.StatusOK
	case codes.Canceled:
		return http.StatusRequestTimeout
	case codes.Unknown:
		return http.StatusInternalServerError
	case codes.InvalidArgument:
		return http.StatusBadRequest
	case codes.DeadlineExceeded:
		return http.StatusGatewayTimeout
	case codes.NotFound:
		return http.StatusNotFound
	case codes.AlreadyExists:
		return http.StatusConflict
	case codes.PermissionDenied:
		return http.StatusForbidden
	case codes.Unauthenticated:
		return http.StatusUnauthorized
	case codes.ResourceExhausted:
		return http.StatusTooManyRequests
	case codes.FailedPrecondition:
		return http.StatusPreconditionFailed
	case codes.Aborted:
		return http.StatusConflict
	case codes.OutOfRange:
		return http.StatusBadRequest
	case codes.Unimplemented:
		return http.StatusNotImplemented
	case codes.Internal:
		return http.StatusInternalServerError
	case codes.Unavailable:
		return http.StatusServiceUnavailable
	case codes.DataLoss:
		return http.StatusInternalServerError
	}

	//grpclog.Infof("Unknown gRPC error code: %v", code)
	log.Infof("Unknown gRPC error code: %v", code)
	return http.StatusInternalServerError
}
