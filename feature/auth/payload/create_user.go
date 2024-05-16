package payload

import (
	"assalam-learn/feature/auth/repository"
	"assalam-learn/helpers"
	"log"
)

type CreateUserPayload struct {
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (request *CreateUserPayload) ToRepositoryParam() (param repository.CreateUserParam, err error) {
	password, err := helpers.GenerateHashPassword(request.Password)
	if err != nil {
		log.Printf("Error hash password with err: %s", err)
		return
	}

	param = repository.CreateUserParam{
		Fullname: request.Fullname,
		Email:    request.Email,
		Password: password,
	}

	return
}
