package handler

import (
	"iteration-backend/database"
	"iteration-backend/dto"
	"iteration-backend/service"
	"iteration-backend/tools"
	"net/http"

	"github.com/labstack/echo/v4"
)

func SignUp(c echo.Context) error {
	su := &dto.AccountSignUp{}
	if err := c.Bind(su); err != nil {
		return err
	}
	if tools.EmailValid(su.Email) {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid_email"})
	}
	if !tools.PhoneValid(su.Phone) {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid_phone"})
	}
	if len(su.Password) < 8 {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid_password"})
	}
	if su.Password != su.PasswordConfirmation {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid_password_confirmation"})
	}
	uuid, _ := database.CreateAccount(su)
	return c.JSON(http.StatusCreated, echo.Map{"uuid": uuid})
}

func SignIn(c echo.Context) error {
	si := &dto.AccountSignIn{}
	if err := c.Bind(si); err != nil {
		return err
	}
	if !tools.PhoneValid(si.Phone) {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid_phone"})
	}
	if len(si.Password) < 8 {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid_password"})
	}
	uuid, err := database.SignIn(si)
	if err != nil || uuid == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err})
	}
	token, err := service.GenerateToken(uuid)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusAccepted, echo.Map{"token": token})
}
