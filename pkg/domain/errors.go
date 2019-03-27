package domain

import (
	"fmt"
	"net/http"

	"github.com/klahssen/webapp/pkg/internal/errors"
	"google.golang.org/grpc/codes"
	//"google.golang.org/grpc/codes"
	//"google.golang.org/grpc/status"
)

//errors
var (
	ErrNotImplemented           = errors.New("not implemented", codes.Internal)
	ErrAccountInvalidEmail      = errors.New("invalid email address", codes.InvalidArgument)
	ErrAccountEmailConflict     = errors.New("another account exists with this email", codes.AlreadyExists)
	ErrAccountInvalidPassword   = errors.New(fmt.Sprintf("invalid password, must be between %d and %d characters", minPasswordLength, maxPasswordLength), codes.InvalidArgument)
	ErrNothingToProcess         = errors.New("nothing to process", codes.InvalidArgument)
	ErrAuthenticationFailed     = errors.New("invalid password or identifier", codes.Unauthenticated)
	ErrMissingDependency        = errors.New("missing/nil dependency injection", codes.Internal)
	ErrNotFound                 = errors.New("not found", codes.NotFound)
	ErrNotAuthorized            = errors.New("not authorized", codes.Unauthenticated)
	ErrEmptyContext             = errors.New("empty context", codes.InvalidArgument)
	ErrMissingAccessToken       = errors.New("missing access token", codes.Unauthenticated)
	ErrInvalidAccessToken       = errors.New("invalid access token", codes.Unauthenticated)
	ErrAccessTokenExpired       = errors.New("access token expired", codes.Unauthenticated)
	ErrInternal                 = errors.New("internal server error", http.StatusInternalServerError)
	ErrInvalidKeyID             = errors.New("invalid access token key ID", codes.InvalidArgument)
	ErrInvalidPermissions       = errors.New("invalid permissions, check all fields", codes.PermissionDenied)
	ErrInvalidClaims            = errors.New("invalid claims, check all fields", codes.PermissionDenied)
	ErrFailedToGenerateJwtToken = errors.New("failed to generate jwt token", http.StatusInternalServerError)
)
