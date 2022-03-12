package device

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"gitlab.com/golang-team-template/monolith/model"
)

// Repository interface
type Repository interface {
	SaveDevice(ctx context.Context, device *model.Device) error
	UpdateDevice(ctx context.Context, device *model.Device) error
	DeleteDeviceByGUID(ctx context.Context, deviceGUID string) error
}

type repoImpl struct {
	pgRepo *pgRepoImpl
}

// NewRepository returns a new repository
func NewRepository(pool *pgxpool.Pool) Repository {
	return &repoImpl{
		pgRepo: newPGRepo(pool),
	}
}

func (r *repoImpl) SaveDevice(ctx context.Context, device *model.Device) error {
	return r.pgRepo.SaveDevice(ctx, device)
}

func (r *repoImpl) UpdateDevice(ctx context.Context, device *model.Device) error {
	return r.pgRepo.UpdateDevice(ctx, device)
}

func (r *repoImpl) DeleteDeviceByGUID(ctx context.Context, deviceGUID string) error {
	return r.pgRepo.DeleteDeviceByGUID(ctx, deviceGUID)
}
