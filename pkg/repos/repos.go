package repos

import (
	"context"

	pb "github.com/klahssen/webapp/pkg/domain"
)

//Accounts represents the storage layer
type Accounts interface {
	CountByStatus(ctx context.Context, status pb.AccountStatus) (int, error)
	CountByType(ctx context.Context, typ pb.AccountType) (int, error)
	CountByEmail(ctx context.Context, email string) (int, error)
	GetByEmail(ctx context.Context, email string) ([]*pb.AccountEntity, error)
	Get(ctx context.Context, uid string) (*pb.AccountEntity, error)
	Delete(ctx context.Context, uid string) error
	PutNew(ctx context.Context, entity *pb.AccountEntity) (string, error) //return the uid
	Put(ctx context.Context, uid string, entity *pb.AccountEntity) error
}
