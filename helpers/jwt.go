package helpers

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateJWT(id int64) (token string, err error) {
	var (
		secret []byte             = []byte(os.Getenv("JWT_SECRET"))
		claims jwt.StandardClaims = jwt.StandardClaims{
			ExpiresAt: time.Now().Add(15 * time.Minute).Unix(),
			Issuer:    fmt.Sprintf("%d", id),
		}
	)

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	token, err = rawToken.SignedString(secret)
	if err != nil {
		log.Printf("Error signed string jwt with err: %s\n", err)
		return
	}

	return
}
