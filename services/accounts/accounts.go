package accounts

import (
	"context"
	"fmt"

	pb "github.com/klahssen/webapp/domain"
	repos "github.com/klahssen/webapp/repos"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	//"github.com/klahssen/webapp/internal/errors"
)

type accountsService struct {
	repo   repos.Accounts
	emails pb.EmailsServer
}

//New returns an accounts server with the attached repo
func New(repo repos.Accounts, emailsSrv pb.EmailsServer) (pb.AccountsServer, error) {
	srv := &accountsService{repo: repo, emails: emailsSrv}
	if err := srv.validate(); err != nil {
		return nil, err
	}
	return srv, fmt.Errorf("not implemented")
}

func (srv *accountsService) validate() error {
	if srv.repo == nil {
		return status.Error(codes.Internal, "invalid dependency: repository: nil")
	}
	if srv.emails == nil {
		return status.Error(codes.Internal, "invalid dependency: email service: nil")
	}
	return nil
}

func (srv *accountsService) isAvailableEmail(ctx context.Context, email string) error {
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

func (srv *accountsService) Create(ctx context.Context, params *pb.AccountParams) (*pb.AccountResp, error) {
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
func (srv *accountsService) UpdateEmail(ctx context.Context, params *pb.AccountParams) (*pb.AccountResp, error) {
	if params == nil {
		return nil, status.Error(codes.InvalidArgument, pb.ErrNothingToProcess.Error())
	}
	params.Format()
	if err := pb.ValidateEmail(params.Em); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	//get account by UID
	a, err := srv.repo.Get(ctx, params.Uid)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
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
func (srv *accountsService) UpdatePassword(ctx context.Context, params *pb.AccountParams) (*pb.AccountResp, error) {
	if params == nil {
		return nil, status.Error(codes.InvalidArgument, pb.ErrNothingToProcess.Error())
	}
	params.Format()
	if err := pb.ValidatePassword(params.Pw); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	//get account by UID
	a, err := srv.repo.Get(ctx, params.Uid)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}
	a.Pw = params.Pw
	if err := a.Hash(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	a.RecordUpdate()
	err = srv.repo.Put(ctx, params.Uid, a)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.AccountResp{Uid: params.Uid}, nil
}
func (srv *accountsService) UpdatePrivileges(ctx context.Context, params *pb.AccountPrivileges) (*pb.AccountResp, error) {
	return nil, fmt.Errorf("not implemented")
}
func (srv *accountsService) Get(ctx context.Context, params *pb.AccountID) (*pb.AccountJwtTokens, error) {
	return nil, fmt.Errorf("not implemented")
}
func (srv *accountsService) Authenticate(ctx context.Context, params *pb.AccountCredentials) (*pb.AccountJwtTokens, error) {
	return nil, fmt.Errorf("not implemented")
}
