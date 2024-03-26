package handler

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"stratplus/service"
)

type UserHandler struct {
	DB *gorm.DB
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{
		DB: db,
	}
}

// CreateUserHandler creates a new user and provides dependencies to service.
func (h *UserHandler) CreateUserHandler(c echo.Context) error {
	return service.InitCreateUser(c, h.DB)
}

func (h *UserHandler) AuthenticateUserHandler(c echo.Context) error {
	return service.InitAuthenticateUser(c, h.DB)
}
