package services

import "context"

type UserService struct {
	Name     string
	Email    string
	Password string
}

func NewUserService() *UserService {
	return &UserService{}
}

func (u *UserService) Createuser(ctx context.Context) (*UserService, error) {

}
