package accounts

import (
	"context"

	pb "github.com/klahssen/webapp/domain"
	//"github.com/klahssen/webapp/internal/errors"
)

//Create an Account
func Create(ctx context.Context, params *pb.AccountParams, repo pb.AccountsRepoServer) (*pb.AccountEntity, error) {
	if params == nil {
		return nil, &pb.ErrNothingToProcess
	}
	if repo == nil {
		return nil, &pb.ErrMissingDependency
	}
	//validate content
	params.Format()
	a := &pb.AccountEntity{Em: params.Em, Pw: params.Pw, Type: pb.AccountType_USER}
	var err error
	if err = a.ValidateNew(); err != nil {
		return nil, err
	}
	//check for conflicts
	_, err = repo.GetByEmail(ctx, &pb.AccountEmail{Email: params.Em})
	if err == nil {
		return nil, &pb.ErrAccountEmailConflict
	}
	//hash password
	if err = a.Hash(); err != nil {
		return nil, err
	}
	resp, err := repo.PutNew(ctx, a)
	if err == nil && resp != nil {
		a.Uid = resp.Uid
	}
	return &pb.AccountEntity{}, nil
}

//ChangeEmail address of an Account
func ChangeEmail(ctx context.Context, params *pb.AccountParams, repo pb.AccountsRepoServer, emails pb.EmailsServer) (*pb.AccountResp, error) {
	if params == nil {
		return nil, &pb.ErrNothingToProcess
	}
	if repo == nil {
		return nil, &pb.ErrNothingToProcess
	}
	params.Format()
	//params.IsEmpty()
	a := &pb.AccountEntity{Em: params.Em, Pw: params.Pw}
	var err error
	if err = a.ValidateNew(); err != nil {
		return nil, err
	}
	if err = a.Hash(); err != nil {
		return nil, err
	}
	resp, err := repo.Put(ctx, a)
	if err == nil && resp != nil {
		a.Uid = resp.Uid
	}
	return &pb.AccountEntity{}, nil
}
