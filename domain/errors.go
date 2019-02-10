package domain

import (
	"fmt"
	"net/http"

	"github.com/klahssen/webapp/internal/errors"
)

//errors
var (
	ErrAccountInvalidEmail    = errors.Error{Code: 001, Msg: "invalid email address"}
	ErrAccountEmailConflict   = errors.Error{Code: 002, Msg: "another account exists with this email"}
	ErrAccountInvalidPassword = errors.Error{Code: 003, Msg: fmt.Sprintf("invalid password, must be between %d and %d characters", minPasswordLength, maxPasswordLength)}
	ErrNothingToProcess       = errors.Error{Code: http.StatusBadRequest, Msg: "nothing to process"}
	ErrAuthenticationFailed   = errors.Error{Code: http.StatusUnauthorized, Msg: "invalid password or identifier"}
	ErrMissingDependency      = errors.Error{Code: http.StatusInternalServerError, Msg: "missing/nil dependency injection"}
)
