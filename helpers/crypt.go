package helpers

import (
	"os"
	"strconv"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

func GenerateHashPassword(password string) (hashed string, err error) {
	costStr := os.Getenv("BCRYPT_COST")

	cost, err := strconv.Atoi(costStr)
	if err != nil {
		err = errors.Wrapf(err, "error parse int on bcrypt cost env : %s", costStr)
		return
	}

	if cost < bcrypt.MinCost {
		cost = bcrypt.DefaultCost
	}

	crypt, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		err = errors.Wrap(err, "error generate bcrypt hash password")
		return
	}

	hashed = string(crypt)

	return
}

// Compare bcrypt hashed password.
func CompareHashPassword(passwordInput, passwordDB string) (err error) {
	err = bcrypt.CompareHashAndPassword([]byte(passwordDB), []byte(passwordInput))
	return
}
