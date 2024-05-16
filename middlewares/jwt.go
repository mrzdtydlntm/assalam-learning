package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/pkg/errors"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func ParseJWT(g *echo.Group) {
	g.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get(echo.HeaderAuthorization)
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, echo.Map{
					"message": "Error get header token",
					"error":   errors.New("header token not found").Error(),
				})
			}

			token := strings.Fields(authHeader)
			if len(token) < 2 && token[1] == "" {
				return c.JSON(http.StatusUnauthorized, echo.Map{
					"message": "Error get header token",
					"error":   errors.New("header token not found").Error(),
				})
			}

			// Parse jwt decode
			claims := jwt.MapClaims{}
			res, err := jwt.ParseWithClaims(token[1], claims, func(t *jwt.Token) (interface{}, error) {
				if jwt.GetSigningMethod("HS512") != t.Method {
					return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
				}

				return []byte(os.Getenv("JWT_SECRET")), nil
			})
			if err != nil {
				log.Printf("Error parse jwt token with err: %s", err)
				return c.JSON(http.StatusUnauthorized, echo.Map{
					"message": "Error parse jwt token",
					"error":   err.Error(),
				})
			}

			// Check decoded data if it's valid or not
			resMap, ok := res.Claims.(jwt.MapClaims)
			if !ok && !res.Valid {
				log.Printf("Invalid token jwt")
				return c.JSON(http.StatusUnauthorized, echo.Map{
					"message": "Error parse jwt token",
					"error":   errors.New("error validate jwt parse").Error(),
				})
			}

			// Sync decoded data with database
			issuer := resMap["iss"]
			if issuer == nil {
				return c.JSON(http.StatusUnauthorized, echo.Map{
					"message": "Error get issuer",
					"error":   errors.New("issuer unknown").Error(),
				})
			}

			c.Set("user_id", issuer)

			return next(c)
		}
	})
}
