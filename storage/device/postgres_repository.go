package device

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"gitlab.com/golang-team-template/monolith/model"
)

type pgRepoImpl struct {
	pool *pgxpool.Pool
}

func newPGRepo(pool *pgxpool.Pool) *pgRepoImpl {
	return &pgRepoImpl{
		pool: pool,
	}
}

func (r *pgRepoImpl) SaveDevice(ctx context.Context, device *model.Device) error {
	return nil
}
func (r *pgRepoImpl) UpdateDevice(ctx context.Context, device *model.Device) error {
	return nil
}
func (r *pgRepoImpl) DeleteDeviceByGUID(ctx context.Context, deviceGUID string) error {
	return nil
}
