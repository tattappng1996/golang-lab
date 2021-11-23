package usecase

import "golang-lab/dto"

// ListUsers ..
func (u *usecase) ListUsers() ([]dto.User, error) {
	users := []dto.User{}

	tableUser, err := u.repo.ListUsers()
	if err != nil {
		return users, err
	}

	for _, user := range tableUser.Users {
		users = append(users, user)
	}
	return users, nil
}

// CreateUser ..
func (u *usecase) CreateUser(user dto.User) error {
	return u.repo.CreateUser(user)
}
