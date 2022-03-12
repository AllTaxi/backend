package user

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"gitlab.com/golang-team-template/monolith/model"
)

//Repository interface for user repository
type Repository interface {
	SaveUser(ctx context.Context, user *model.User) (*model.User, error)
	UpdateUser(ctx context.Context, user *model.User) error
	DeleteUserByGUID(ctx context.Context, userGUID string) error
}

type repoImpl struct {
	pgRepo *pgRepoImpl
}

//NewRepository returns a new repository
func NewRepository(pool *pgxpool.Pool) Repository {
	return &repoImpl{
		pgRepo: newPGRepo(pool),
	}
}

func (r *repoImpl) SaveUser(ctx context.Context, user *model.User) (*model.User, error) {
	return r.pgRepo.SaveUser(ctx, user)
}

func (r *repoImpl) UpdateUser(ctx context.Context, user *model.User) error {
	return r.pgRepo.UpdateUser(ctx, user)
}

func (r *repoImpl) DeleteUserByGUID(ctx context.Context, userGUID string) error {
	return r.pgRepo.DeleteUserByGUID(ctx, userGUID)
}
