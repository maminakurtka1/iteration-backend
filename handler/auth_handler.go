package handler

import (
	"iteration-backend/database"
	"iteration-backend/dto"
	"iteration-backend/service"
	"iteration-backend/tools"
	"net/http"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/echo/v4"
)

// SignUp godoc
// @Summary Sign up new user.
// @Description Create new account.
// @Tags authorize
// @Accept */sign-up*
// @Produce json
// @Success 201
// @Router /sign-up [post]
func SignUp(db *pgxpool.Pool) echo.HandlerFunc {
	return func(c echo.Context) error {
		su := &dto.AccountSignUp{}
		if err := c.Bind(su); err != nil {
			return c.NoContent(http.StatusBadRequest)
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
		_, err := database.CreateAccount(db, su)
		if err != nil {
			return c.NoContent(http.StatusInternalServerError)
		}
		return c.NoContent(http.StatusCreated)
	}
}

func SignIn(db *pgxpool.Pool) echo.HandlerFunc {
	return func(c echo.Context) error {
		si := &dto.AccountSignIn{}
		if err := c.Bind(si); err != nil {
			return c.NoContent(http.StatusBadRequest)
		}
		if !tools.PhoneValid(si.Phone) {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "invalid_phone"})
		}
		if len(si.Password) < 8 {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "invalid_password"})
		}
		uuid, err := database.SignIn(db, si)
		if err != nil || uuid == "" {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "invalid_creds"})
		}
		token, err := service.GenerateToken(uuid)
		if err != nil || token == "" {
			return c.NoContent(http.StatusInternalServerError)
		}
		return c.JSON(http.StatusOK, echo.Map{"token": token})
	}

}
