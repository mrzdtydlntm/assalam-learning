package payload

import (
	"assalam-learn/feature/auth/models"
	"time"
)

type ResponseUser struct {
	ID        int64     `json:"id"`
	Fullname  string    `json:"fullname"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ResponseLogin struct {
	User  ResponseUser `json:"user"`
	Token string       `json:"token"`
}

func ToResponseUser(user models.User) (res ResponseUser) {
	res = ResponseUser{
		ID:        user.ID,
		Fullname:  user.Fullname,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
	}

	if user.UpdatedAt.Valid {
		res.UpdatedAt = user.UpdatedAt.Time
	}

	return
}

func ToResponseListUser(user []models.User) (res []*ResponseUser) {
	res = make([]*ResponseUser, len(user))

	for i := range user {
		res[i] = new(ResponseUser)
		data := ToResponseUser(user[i])
		res[i] = &data
	}

	return
}

func ToResponseLogin(user models.User, token string) (res ResponseLogin) {
	res = ResponseLogin{
		User: ResponseUser{
			ID:        user.ID,
			Fullname:  user.Fullname,
			Email:     user.Email,
			Password:  user.Password,
			CreatedAt: user.CreatedAt,
		},
		Token: token,
	}

	if user.UpdatedAt.Valid {
		res.User.UpdatedAt = user.UpdatedAt.Time
	}

	return
}
