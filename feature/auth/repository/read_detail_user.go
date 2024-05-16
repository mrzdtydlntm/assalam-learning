package repository

import (
	"assalam-learn/feature/auth/models"
	"context"
	"log"

	"github.com/jmoiron/sqlx"
)

type ReadDetailUserParam struct {
	ID    int64  `db:"id"`
	Email string `db:"email"`
}

func ReadDetailUserQuery(ctx context.Context, db *sqlx.DB, user ReadDetailUserParam) (data models.User, err error) {
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
			id = ?;
	`

	row := db.QueryRowxContext(ctx, sqlStatement, user.ID)

	err = row.Scan(
		&data.ID,
		&data.Fullname,
		&data.Email,
		&data.Password,
		&data.CreatedAt,
		&data.UpdatedAt,
	)
	if err != nil {
		log.Printf("Error read detail user with error: %s", err)
		return
	}

	return
}

func ReadDetailUserByEmailQuery(ctx context.Context, db *sqlx.DB, user ReadDetailUserParam) (data models.User, err error) {
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
			email = ?;
	`

	row := db.QueryRowxContext(ctx, sqlStatement, user.Email)

	err = row.Scan(
		&data.ID,
		&data.Fullname,
		&data.Email,
		&data.Password,
		&data.CreatedAt,
		&data.UpdatedAt,
	)
	if err != nil {
		log.Printf("Error read detail user with error: %s", err)
		return
	}

	return
}
