package controller

import (
	"iteration-backend/tools"
	"net/http"

	"github.com/labstack/echo/v4"
)

type sign_up struct {
	Email                string `json:"email"`
	Phone                string `json:"phone"`
	Password             string `json:password`
	PasswordConfirmation string `json:password_confirmation`
}

func SignUp(c echo.Context) error {
	su := &sign_up{}
	if err := c.Bind(su); err != nil {
		return err
	}
	if tools.EmailValid(su.Email) {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid_email"})
	}
	if tools.PhoneValid(su.Phone) {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid_phone"})
	}
	if len(su.Password) < 8 {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid_password"})
	}
	if su.Password != su.PasswordConfirmation {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid_password_confirmation"})
	}
	return c.JSON(http.StatusCreated, su)
}

func SignIn(c echo.Context) error {
	return nil
}
