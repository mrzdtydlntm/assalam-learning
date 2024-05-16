package controllers

import (
	"assalam-learn/feature/auth/payload"
	"assalam-learn/feature/auth/repository"
	"assalam-learn/helpers"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
)

func loginUser(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		var request payload.LoginUserPayload
		if err = c.Bind(&request); err != nil {
			log.Printf("Error bind login user json request with error: %s", err)
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "Error bind request struct",
				"error":   err.Error(),
			})
		}

		data, err := repository.ReadDetailUserByEmailQuery(c.Request().Context(), db, request.ToRepositoryParam())
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "Error login user",
				"error":   err.Error(),
			})
		}

		err = helpers.CompareHashPassword(request.Password, data.Password)
		if err != nil {
			log.Printf("Error compare password with err: %s", err)
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "Password not match",
				"error":   err.Error(),
			})
		}

		token, err := helpers.CreateJWT(data.ID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "Error create jwt token",
				"error":   err.Error(),
			})
		}

		return c.JSON(http.StatusOK, echo.Map{
			"message": "Successful login user",
			"data":    payload.ToResponseLogin(data, token),
		})
	}
}
