package routes

import (
	"fmt"
	"net/http"
	"os"

	"assalam-learn/config"
	authControllers "assalam-learn/feature/auth/controllers"
	"assalam-learn/middlewares"

	"github.com/labstack/echo"
)

func Routes() *echo.Echo {
	e := echo.New()
	db := config.ConnectDB()

	// Middlewares
	middlewares.LoggerMiddleware(e)
	middlewares.RecoverMiddleware(e)
	middlewares.CorsMiddleware(e)

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{
			"message": fmt.Sprintf("%s is Running", os.Getenv("APP_NAME")),
		})
	})

	api := e.Group("/api")
	authControllers.Controller(api, db)

	//Public
	api.Static("/media", "media")

	return e
}
