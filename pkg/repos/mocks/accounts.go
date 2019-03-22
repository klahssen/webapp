package mocks

import (
	"context"
	"fmt"
	"time"

	pb "github.com/klahssen/webapp/pkg/domain"
	"github.com/klahssen/webapp/pkg/repos"
)

type accountsRepo struct {
	data map[string]*pb.AccountEntity
}

//NewAccountsRepo returns an instance of the accounts repo
func NewAccountsRepo() repos.Accounts {
	return &accountsRepo{}
}

func (r *accountsRepo) CountByStatus(ctx context.Context, status pb.AccountStatus) (int, error) {
	n := 0
	for _, v := range r.data {
		if v.Status == status {
			n++
		}
	}
	return n, nil
}
func (r *accountsRepo) CountByType(ctx context.Context, typ pb.AccountType) (int, error) {
	n := 0
	for _, v := range r.data {
		if v.Type == typ {
			n++
		}
	}
	return n, nil
}
func (r *accountsRepo) CountByEmail(ctx context.Context, email string) (int, error) {
	n := 0
	for _, v := range r.data {
		if v.Em == email {
			n++
		}
	}
	return n, nil
}
func (r *accountsRepo) GetByEmail(ctx context.Context, email string) ([]*pb.AccountEntity, error) {
	res := []*pb.AccountEntity{}
	for _, v := range r.data {
		if v.Em == email {
			res = append(res, v)
		}
	}
	return res, nil
}
func (r *accountsRepo) Get(ctx context.Context, uid string) (*pb.AccountEntity, error) {
	if e, ok := r.data[uid]; ok {
		return e, nil
	}
	return nil, &pb.ErrNotFound
}
func (r *accountsRepo) Delete(ctx context.Context, uid string) error {
	delete(r.data, uid)
	return nil
}
func (r *accountsRepo) PutNew(ctx context.Context, entity *pb.AccountEntity) (string, error) {
	if entity == nil {
		return "", &pb.ErrNothingToProcess
	}
	key := r.newKey()
	r.data[key] = entity
	return key, nil
}
func (r *accountsRepo) Put(ctx context.Context, uid string, entity *pb.AccountEntity) error {
	if entity == nil {
		return &pb.ErrNothingToProcess
	}
	r.data[uid] = entity
	return nil
}

func (r *accountsRepo) newKey() string {
	return fmt.Sprintf("acct_%d", time.Now().Unix())
}
