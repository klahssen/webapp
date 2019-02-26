package accounts

import (
	"context"
	"fmt"

	pb "github.com/klahssen/webapp/domain"
	"github.com/klahssen/webapp/internal/validators"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	//"github.com/klahssen/webapp/internal/errors"
)

//Repository represents the storage layer
type Repository interface {
	GetByEmail(ctx context.Context, email string) error
	Get(ctx context.Context, uid string) (*pb.AccountEntity, error)
	CountByEmail(ctx context.Context, email string) (int, error)
	Delete(ctx context.Context, key string) error
	PutNew(ctx context.Context, entity *pb.AccountEntity) (string, error) //return the uid
	Put(ctx context.Context, uid string, entity *pb.AccountEntity) error
}

type service struct {
	repo   Repository
	emails pb.EmailsServer
}

//NewInstance returns an accounts server with the attached repo
func NewInstance(repo Repository, emailsSrv pb.EmailsServer) (pb.AccountsServer, error) {
	srv := &service{repo: repo, emails: emailsSrv}
	if err := srv.validate(); err != nil {
		return nil, err
	}
	return srv, fmt.Errorf("not implemented")
}

func (srv *service) validate() error {
	if srv.repo == nil {
		return status.Error(codes.Internal, "invalid dependency: repository: nil")
	}
	if srv.emails == nil {
		return status.Error(codes.Internal, "invalid dependency: email service: nil")
	}
	return nil
}

func (srv *service) isAvailableEmail(ctx context.Context, email string) error {
	//check for conflicts
	n, err := srv.repo.CountByEmail(ctx, email)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}
	if n > 0 {
		return status.Error(codes.InvalidArgument, pb.ErrAccountEmailConflict.Error())
	}
	return nil
}

func (srv *service) Create(ctx context.Context, params *pb.AccountParams) (*pb.AccountResp, error) {
	if params == nil {
		return nil, status.Error(codes.InvalidArgument, pb.ErrNothingToProcess.Error())
	}
	//validate content
	params.Format()
	a := &pb.AccountEntity{Em: params.Em, Pw: params.Pw, Type: pb.AccountType_USER}
	var err error
	if err = a.ValidateNew(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	//check if email is available
	if err = srv.isAvailableEmail(ctx, params.Em); err != nil {
		return nil, err
	}
	//hash password
	if err = a.Hash(); err != nil {
		return nil, err
	}
	uid, err := srv.repo.PutNew(ctx, a)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.AccountResp{Uid: uid}, nil
}
func (srv *service) UpdateEmail(ctx context.Context, params *pb.AccountParams) (*pb.AccountResp, error) {
	if params == nil {
		return nil, status.Error(codes.InvalidArgument, pb.ErrNothingToProcess.Error())
	}
	params.Format()
	//get account by UID
	a, err := srv.repo.Get(ctx, params.Uid)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}
	if err = validators.EmailAddress(params.Em); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	a.Em = params.Em
	if err = srv.isAvailableEmail(ctx, params.Em); err != nil {
		return nil, status.Error(codes.AlreadyExists, err.Error())
	}
	a.RecordUpdate() //update timestamp
	err = srv.repo.Put(ctx, params.Uid, a)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.AccountResp{Uid: params.Uid}, nil
}
func (srv *service) UpdatePassword(ctx context.Context, params *pb.AccountParams) (*pb.AccountResp, error) {
	return nil, fmt.Errorf("not implemented")
}
func (srv *service) UpdatePrivileges(ctx context.Context, params *pb.AccountPrivileges) (*pb.AccountResp, error) {
	return nil, fmt.Errorf("not implemented")
}
func (srv *service) Get(ctx context.Context, params *pb.AccountID) (*pb.AccountJwtTokens, error) {
	return nil, fmt.Errorf("not implemented")
}
func (srv *service) Authenticate(ctx context.Context, params *pb.AccountCredentials) (*pb.AccountJwtTokens, error) {
	return nil, fmt.Errorf("not implemented")
}
