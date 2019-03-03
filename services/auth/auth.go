package auth

import (
	"context"

	pb "github.com/klahssen/webapp/domain"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type authService struct{}

//New returns a new instance of the Service
func New() pb.AuthorizationServer {
	return &authService{}
}

//IsAuthorized method
func (a *authService) IsAuthorized(ctx context.Context, params *pb.AuthParams) (*pb.AuthResp, error) {
	if params == nil {
		return nil, status.Error(codes.InvalidArgument, pb.ErrNothingToProcess.Error())
	}
	perms, err := getPermissionsFromCtx(ctx)
	switch params.Action {
	case pb.Action_ACCOUNTS_CREATE:
		return notAuthorized(), nil
	}
	return authorized(), nil
}

func getPermissionsFromCtx(ctx context.Context) (*pb.Premissions, error) {
	if ctx == nil {
		return nil, &pb.ErrEmptyContext
	}
	return &pb.Premissions{}, nil
}

func notAuthorized() *pb.AuthResp {
	return &pb.AuthResp{
		Authorized: false,
	}
}
func authorized() *pb.AuthResp {
	return &pb.AuthResp{
		Authorized: true,
	}
}
