package public

import "golang-lab/dto"

// HTTPUsecase ..
type HTTPUsecase interface {
	ListUsers() ([]dto.User, error)
	CreateUser(dto.User) error
}

type handler struct {
	uc HTTPUsecase
}

func newHandler(uc HTTPUsecase) *handler {
	return &handler{
		uc: uc,
	}
}
