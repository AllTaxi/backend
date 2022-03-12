package user

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

func (r *pgRepoImpl) SaveUser(ctx context.Context, user *model.User) (*model.User, error) {
	_, err := r.pool.Exec(ctx, `
	INSERT INTO 
		users (
			id,
			email, 
			password,
			first_name,
			last_name,
			acces_token
		) VALUES($1, $2, $3, $4, $5, $6)`,
		user.ID,
		user.Email,
		user.Password,
		user.FirstName,
		user.LastName,
		user.AccessToken,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *pgRepoImpl) UpdateUser(ctx context.Context, user *model.User) error {
	return nil
}
func (r *pgRepoImpl) DeleteUserByGUID(ctx context.Context, userGUID string) error {
	return nil
}

// func (r *pgRepoImpl) CheckRedis(ctx context.Context, key, value string) error {

// 	conn := r.rds.Get()
// 	defer conn.Close()
// 	interfaceData, err := conn.Do("GET", key)
// 	if err != nil {
// 		return err
// 	}
// 	checkValue, err := redis.String(interfaceData, err)
// 	if checkValue != value {
// 		fmt.Println(err, checkValue)
// 		return err
// 	}
// 	return nil
// }
