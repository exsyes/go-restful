package simple

import "errors"

type SimpleRepository struct {
	Error bool
}

func NewSimpleRepository(isError bool) *SimpleRepository {
	return &SimpleRepository{
		Error: isError,
	}
}

type SimpleService struct {
	*SimpleRepository
}

func NewSimpleService(Repository *SimpleRepository) (*SimpleService, error) {
	if Repository.Error {
		return nil, errors.New("failed create service")
	} else {
		return &SimpleService{SimpleRepository: Repository}, nil
	}
}
