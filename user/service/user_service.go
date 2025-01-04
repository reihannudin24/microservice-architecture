package service

import (
	"Blast/internal/user/model"
	"Blast/internal/user/repository"
)

type UserService struct {
	Repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) RegisterEmail(user model.User) (int64, error) {
	return s.Repo.RegisterEmail(user)
}

func (s *UserService) GetUserByID(id int64) (*model.User, error) {
	return s.Repo.GetByID(id)
}
