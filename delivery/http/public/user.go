package public

import (
	"fmt"
	"golang-lab/dto"
	"net/http"

	"github.com/labstack/echo/v4"
)

// ListUsers..
func (h *handler) ListUsers(c echo.Context) error {
	users, err := h.uc.ListUsers()
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, dto.ResponseMessage{
			Success: false,
		})
	}
	return c.JSON(http.StatusOK, dto.ResponseMessage{
		Data:    users,
		Success: true,
	})
}

// CreateUser..
func (h *handler) CreateUser(c echo.Context) error {
	user := dto.User{}
	if err := c.Bind(&user); err != nil {
		errorMessage := err.Error()
		return c.JSON(http.StatusBadRequest, dto.ResponseMessage{
			Error:   &errorMessage,
			Success: false,
		})
	}
	if user.Name == "" {
		errorMessage := "ลืมใส่ชื่อมาอะคร๊าาาบ"
		return c.JSON(http.StatusBadRequest, dto.ResponseMessage{
			Error:   &errorMessage,
			Success: false,
		})
	}

	if err := h.uc.CreateUser(user); err != nil {
		errorMessage := err.Error()
		return c.JSON(http.StatusInternalServerError, dto.ResponseMessage{
			Error:   &errorMessage,
			Success: false,
		})
	}
	return c.JSON(http.StatusOK, dto.ResponseMessage{
		Success: true,
	})
}
