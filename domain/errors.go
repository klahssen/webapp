package domain

import (
	"fmt"
	"net/http"

	"github.com/klahssen/webapp/internal/errors"
	//"google.golang.org/grpc/codes"
	//"google.golang.org/grpc/status"
)

//errors
var (
	ErrNotImplemented           = errors.Error{Code: 000, Msg: "not implemented"}
	ErrAccountInvalidEmail      = errors.Error{Code: 001, Msg: "invalid email address"}
	ErrAccountEmailConflict     = errors.Error{Code: 002, Msg: "another account exists with this email"}
	ErrAccountInvalidPassword   = errors.Error{Code: 003, Msg: fmt.Sprintf("invalid password, must be between %d and %d characters", minPasswordLength, maxPasswordLength)}
	ErrNothingToProcess         = errors.Error{Code: http.StatusBadRequest, Msg: "nothing to process"}
	ErrAuthenticationFailed     = errors.Error{Code: http.StatusUnauthorized, Msg: "invalid password or identifier"}
	ErrMissingDependency        = errors.Error{Code: http.StatusInternalServerError, Msg: "missing/nil dependency injection"}
	ErrNotFound                 = errors.Error{Code: http.StatusNotFound, Msg: "not found"}
	ErrNotAuthorized            = errors.Error{Code: http.StatusUnauthorized, Msg: "not authorized"}
	ErrEmptyContext             = errors.Error{Code: http.StatusBadRequest, Msg: "empty context"}
	ErrMissingAccessToken       = errors.Error{Code: http.StatusBadRequest, Msg: "missing access token"}
	ErrInvalidAccessToken       = errors.Error{Code: http.StatusBadRequest, Msg: "invalid access token"}
	ErrAccessTokenExpired       = errors.Error{Code: http.StatusBadRequest, Msg: "access token expired"}
	ErrInternal                 = errors.Error{Code: http.StatusInternalServerError, Msg: "internal server error"}
	ErrInvalidKeyID             = errors.Error{Code: http.StatusInternalServerError, Msg: "invalid access token key ID"}
	ErrInvalidPermissions       = errors.Error{Code: http.StatusInternalServerError, Msg: "invalid permissions, check all fields"}
	ErrInvalidClaims            = errors.Error{Code: http.StatusInternalServerError, Msg: "invalid claims, check all fields"}
	ErrFailedToGenerateJwtToken = errors.Error{Code: http.StatusInternalServerError, Msg: "failed to generate jwt token"}
)
