package auth

import (
	"context"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	pb "github.com/klahssen/webapp/domain"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type authService struct {
	keyFunc         jwt.Keyfunc
	claimsValidator pb.ClaimsValidator
	permsFunc       pb.AuthorizeFunc
}

//New returns a new instance of the Service
func New(keyFunc jwt.Keyfunc, claimsValidator pb.ClaimsValidator, permsFunc pb.AuthorizeFunc) (pb.AuthorizationServer, error) {
	if keyFunc == nil {
		return nil, fmt.Errorf("keyFunc is nil")
	}
	if claimsValidator == nil {
		return nil, fmt.Errorf("claimsValidator is nil")
	}
	if permsFunc == nil {
		return nil, fmt.Errorf("permsFunc is nil")
	}
	return &authService{keyFunc: keyFunc, claimsValidator: claimsValidator, permsFunc: permsFunc}, nil
}

//IsAuthorized method
func (a *authService) IsAuthorized(ctx context.Context, params *pb.AuthParams) (*pb.AuthResp, error) {
	if params == nil {
		return nil, status.Error(codes.InvalidArgument, pb.ErrNothingToProcess.Error())
	}
	return a.permsFunc(params)
}
