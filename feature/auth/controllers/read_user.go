package controllers

import (
	"assalam-learn/feature/auth/payload"
	"assalam-learn/feature/auth/repository"
	"fmt"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
)

func readUser(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		var request payload.ReadUserPayload
		if err = c.Bind(&request); err != nil {
			log.Printf("Error bind read user json request with error: %s", err)
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "Error bind request struct",
				"error":   err.Error(),
			})
		}

		// Checking middleware context
		fmt.Println(c.Get("user_id"))

		data, err := repository.ReadUserQuery(c.Request().Context(), db, request.ToRepositoryParam())
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "Error read user",
				"error":   err.Error(),
			})
		}

		return c.JSON(http.StatusOK, echo.Map{
			"message": "Successful read user",
			"data":    payload.ToResponseListUser(data),
		})
	}
}
