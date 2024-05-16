package payload

import "assalam-learn/feature/auth/repository"

type LoginUserPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (request *LoginUserPayload) ToRepositoryParam() (param repository.ReadDetailUserParam) {
	param = repository.ReadDetailUserParam{
		Email: request.Email,
	}

	return
}
