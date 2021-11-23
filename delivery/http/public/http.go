package public

import (
	"github.com/labstack/echo/v4"
)

// NewHTTP ..
func NewHTTP(e *echo.Group, uc HTTPUsecase) *handler {
	h := newHandler(uc)

	e.GET("/users", h.ListUsers)
	e.POST("/user", h.CreateUser)

	return h
}
