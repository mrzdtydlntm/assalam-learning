package payload

import "assalam-learn/feature/auth/repository"

type ReadUserPayload struct {
	Search string `query:"search"`
}

func (request *ReadUserPayload) ToRepositoryParam() (param repository.ReadUserParam) {
	if request.Search != "" {
		param = repository.ReadUserParam{
			SetSearch: true,
			Search:    "%" + request.Search + "%",
		}
	}

	return
}
