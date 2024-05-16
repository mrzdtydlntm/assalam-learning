package controllers

import (
	"assalam-learn/feature/auth/payload"
	"assalam-learn/feature/auth/repository"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
)

func createUser(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		var request payload.CreateUserPayload
		if err = c.Bind(&request); err != nil {
			log.Printf("Error bind create user json request with error: %s", err)
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "Error bind request struct",
				"error":   err.Error(),
			})
		}

		param, err := request.ToRepositoryParam()
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "Error to adapter create user param",
				"error":   err.Error(),
			})
		}

		data, err := repository.CreateUserQuery(c.Request().Context(), db, param)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "Error create user",
				"error":   err.Error(),
			})
		}

		return c.JSON(http.StatusOK, echo.Map{
			"message": "Successful create user",
			"data":    payload.ToResponseUser(data),
		})
	}
}
