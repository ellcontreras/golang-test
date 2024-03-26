package service

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"net/http"
	"stratplus/types"
)

// InitCreateUser handles the creation of a new user.
func InitCreateUser(c echo.Context, db *gorm.DB) error {
	u := new(types.User)
	err := c.Bind(u)
	if err != nil {
		return err
	}

	err = u.ValidateFields()
	if err != nil {
		return c.JSON(http.StatusBadRequest, types.JSON{
			"error": err.Error(),
		})
	}

	err = db.Create(u).Error
	if err != nil {
		// Im not sending the error that credentials are already in use for security reasons, avoid leaking information.
		return c.JSON(http.StatusInternalServerError, types.JSON{
			"error": "error creating user",
		})
	}

	return c.JSON(http.StatusOK, types.JSON{
		"message": fmt.Sprintf("user created with id: %d", u.ID),
	})
}

// InitAuthenticateUser handles the authentication of a user.
func InitAuthenticateUser(c echo.Context, db *gorm.DB) error {
	u := new(types.User)
	err := c.Bind(u)
	if err != nil {
		return err
	}

	err = u.ValidateLoginFields()
	if err != nil {
		return c.JSON(http.StatusBadRequest, types.JSON{
			"error": err.Error(),
		})
	}

	err = db.Where("email = ? AND password = ?", u.Email, u.Password).First(u).Error
	if err != nil {
		return c.JSON(http.StatusUnauthorized, types.JSON{
			"error": "invalid credentials",
		})
	}

	token := u.GenerateToken()

	return c.JSON(http.StatusOK, types.JSON{
		"jwt": token,
	})
}
