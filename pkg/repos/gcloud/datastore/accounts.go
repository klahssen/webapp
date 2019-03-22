package datastore

import (
	"context"
	"fmt"

	"cloud.google.com/go/datastore"
	pb "github.com/klahssen/webapp/pkg/domain"
	"github.com/klahssen/webapp/pkg/repos"
)

const (
	kind = "accounts"
)

type accountsRepo struct {
	projectID string
	namespace string
}

var (
	errInternalServerError = fmt.Errorf("internal server error")
	errNotFound            = fmt.Errorf("not found")
	errInvalidUID          = fmt.Errorf("invalid uid")
)

//NewAccountsRepo returns an instance of the accounts repo
func NewAccountsRepo(projectID, namespace string) (repos.Accounts, error) {
	if projectID == "" {
		return nil, fmt.Errorf("invalid projectID '%s'", projectID)
	}
	return &accountsRepo{
		projectID: projectID, namespace: namespace,
	}, nil
}

func (r *accountsRepo) getClient(ctx context.Context) (*datastore.Client, error) {
	client, err := datastore.NewClient(ctx, r.projectID)
	if err != nil {
		logger.Errorf("failed to get new datastore client: %v", err)
		return nil, errInternalServerError
	}
	return client, nil
}

func (r *accountsRepo) CountByStatus(ctx context.Context, status pb.AccountStatus) (int, error) {
	client, err := r.getClient(ctx)
	if err != nil {
		return 0, err
	}
	q := datastore.NewQuery(kind).Filter("status=", status).KeysOnly()
	n, err := client.Count(ctx, q)
	if err != nil {
		logger.Errorf("query failed: %v", err)
		return 0, errInternalServerError
	}
	return n, nil
}
func (r *accountsRepo) CountByType(ctx context.Context, typ pb.AccountType) (int, error) {
	client, err := r.getClient(ctx)
	if err != nil {
		return 0, err
	}
	q := datastore.NewQuery(kind).Filter("type=", typ).KeysOnly()
	n, err := client.Count(ctx, q)
	if err != nil {
		logger.Errorf("query failed: %v", err)
		return 0, errInternalServerError
	}
	return n, nil
}
func (r *accountsRepo) CountByEmail(ctx context.Context, email string) (int, error) {
	client, err := r.getClient(ctx)
	if err != nil {
		return 0, err
	}
	q := datastore.NewQuery(kind).Filter("email=", email).KeysOnly()
	n, err := client.Count(ctx, q)
	if err != nil {
		logger.Errorf("query failed: %v", err)
		return 0, errInternalServerError
	}
	return n, nil
}
func (r *accountsRepo) GetByEmail(ctx context.Context, email string) ([]*pb.AccountEntity, error) {
	client, err := r.getClient(ctx)
	if err != nil {
		return nil, err
	}
	res := []*pb.AccountEntity{}
	q := datastore.NewQuery(kind).Filter("email=", email)
	keys, err := client.GetAll(ctx, q, &res)
	if err != nil {
		logger.Errorf("query failed: %v", err)
		return nil, errInternalServerError
	}
	l1 := len(keys)
	l2 := len(res)
	if l1 != l2 {
		logger.Errorf("mismatch: %d keys and %d res", l1, l2)
		return nil, errInternalServerError
	}
	for i := range res {
		res[i].Uid = keys[i].Encode()
	}
	return res, nil
}
func (r *accountsRepo) Get(ctx context.Context, uid string) (*pb.AccountEntity, error) {
	client, err := r.getClient(ctx)
	if err != nil {
		return nil, err
	}
	res := &pb.AccountEntity{}
	key, err := decodeKey(uid)
	if err != nil {
		return nil, err
	}
	err = client.Get(ctx, key, res)
	if err != nil {
		logger.Errorf("query failed: %v", err)
		if err.Error() == datastore.ErrNoSuchEntity.Error() {
			return nil, errInternalServerError
		} else if err.Error() == datastore.ErrInvalidKey.Error() {
			return nil, errInvalidUID
		}
		return nil, errInternalServerError
	}
	return res, nil
}
func (r *accountsRepo) Delete(ctx context.Context, uid string) error {
	client, err := r.getClient(ctx)
	if err != nil {
		return err
	}
	key, err := decodeKey(uid)
	if err != nil {
		return err
	}
	err = client.Delete(ctx, key)
	if err != nil {
		logger.Errorf("query failed: %v", err)
		if err.Error() == datastore.ErrNoSuchEntity.Error() {
			return errInternalServerError
		} else if err.Error() == datastore.ErrInvalidKey.Error() {
			return errInvalidUID
		}
		return errInternalServerError
	}
	return nil
}
func (r *accountsRepo) PutNew(ctx context.Context, entity *pb.AccountEntity) (string, error) {
	client, err := r.getClient(ctx)
	if err != nil {
		return "", err
	}
	if entity == nil {
		return "", &pb.ErrNothingToProcess
	}
	k, err := client.Put(ctx, newKey(r.namespace), entity)
	if err != nil {
		logger.Errorf("query failed: %v", err)
		return "", errInternalServerError
	}
	entity.Uid = k.Encode()
	return k.Encode(), nil
}
func (r *accountsRepo) Put(ctx context.Context, uid string, entity *pb.AccountEntity) error {
	client, err := r.getClient(ctx)
	if err != nil {
		return err
	}
	if entity == nil {
		return &pb.ErrNothingToProcess
	}
	k, err := decodeKey(uid)
	if err != nil {
		return err
	}
	_, err = client.Put(ctx, k, entity)
	if err != nil {
		logger.Errorf("query failed: %v", err)
		return errInternalServerError
	}
	return nil
}

func newKey(namespace string) *datastore.Key {
	k := datastore.IncompleteKey(kind, nil)
	k.Namespace = namespace
	return k
}

func decodeKey(uid string) (*datastore.Key, error) {
	k, err := datastore.DecodeKey(uid)
	if err != nil {
		logger.Errorf("failed to decode key: %v", err)
		return nil, errInvalidUID
	}
	return k, nil
}
