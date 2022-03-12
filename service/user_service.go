package service

import (
	"context"
	"fmt"

	"gitlab.com/golang-team-template/monolith/model"
	userrepo "gitlab.com/golang-team-template/monolith/storage/user"
)

//UserService interface
type UserService interface {
	SignUp(context.Context, *model.User) (*model.User, error)

}

type userServiceImpl struct {
	userRepo userrepo.Repository
}

//NewUserService returns a new UserService
func NewUserService(userRepo userrepo.Repository) UserService {
	return &userServiceImpl{
		userRepo: userRepo,
	}
}

//Verify returns the user verified


//SingUp return the user datas
func (s *userServiceImpl) SignUp(ctx context.Context, request *model.User) (*model.User, error) {
	response, err := s.userRepo.SaveUser(ctx, request)
	if err != nil {
		fmt.Println(err)
		return &model.User{}, err
	}

	return response, nil
}