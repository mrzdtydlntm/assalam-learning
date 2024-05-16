package repository

import (
	"assalam-learn/feature/auth/models"
	"context"
	"log"

	"github.com/jmoiron/sqlx"
)

type CreateUserParam struct {
	Fullname string `db:"fullname"`
	Email    string `db:"email"`
	Password string `db:"password"`
}

func CreateUserQuery(ctx context.Context, db *sqlx.DB, user CreateUserParam) (data models.User, err error) {
	sqlStatement := `
		INSERT INTO users (
			fullname,
			email,
			password,
			created_at
		)
		VALUES (
			?,
			?,
			?,
			now()
		);
	`

	_, err = db.ExecContext(ctx, sqlStatement,
		user.Fullname,
		user.Email,
		user.Password,
	)
	if err != nil {
		log.Printf("Error create user with error: %s", err)
		return
	}

	sqlStatement = `
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
			id = LAST_INSERT_ID();
	`
	row := db.QueryRowxContext(ctx, sqlStatement)

	err = row.Scan(
		&data.ID,
		&data.Fullname,
		&data.Email,
		&data.Password,
		&data.CreatedAt,
		&data.UpdatedAt,
	)
	if err != nil {
		log.Printf("Error scan new user with error: %s", err)
		return
	}

	return
}
