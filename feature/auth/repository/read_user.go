package repository

import (
	"assalam-learn/feature/auth/models"
	"context"
	"log"

	"github.com/jmoiron/sqlx"
)

type ReadUserParam struct {
	SetSearch bool
	Search    string
}

func ReadUserQuery(ctx context.Context, db *sqlx.DB, user ReadUserParam) (data []models.User, err error) {
	sqlStatement := `
		SELECT
			id,
			fullname,
			email,
			password,
			created_at,
			updated_at
		FROM
			users
		WHERE
			(
				LOWER(?) = 1 AND (LOWER(fullname) LIKE LOWER(?) OR LOWER(email) LIKE LOWER(?))
			) OR
			(
				LOWER(?) = 0 AND TRUE
			)
	`

	err = db.SelectContext(ctx, &data, sqlStatement,
		user.SetSearch,
		user.Search,
		user.Search,
		user.SetSearch,
	)
	if err != nil {
		log.Printf("Error read user with error: %s", err)
		return
	}

	return
}
