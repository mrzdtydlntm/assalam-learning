package controllers

import (
	"assalam-learn/middlewares"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
)

func Controller(e *echo.Group, db *sqlx.DB) {
	auth := e.Group("/auth")

	auth.POST("/create", createUser(db))
	auth.POST("/login", loginUser(db))

	user := auth.Group("/user")
	middlewares.ParseJWT(user)
	user.GET("", readUser(db))
}
