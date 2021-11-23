package usecase

import "golang-lab/dto"

// NewUsecase ..
func NewUsecase(repo Repository) *usecase {
	return &usecase{
		repo: repo,
	}
}

// Repository ..
type Repository interface {
	ListUsers() (dto.TableUsers, error)
	CreateUser(dto.User) error
}

type usecase struct {
	repo Repository
}
