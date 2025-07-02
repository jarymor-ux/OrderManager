package createuseruc

import (
	"orders/internal/domain/entity"

	"braces.dev/errtrace"
)

type Storage interface {
	SaveUser(user entity.User) error
}

type UseCase struct {
	Storage Storage
}

func New(storage Storage) UseCase {
	return UseCase{
		Storage: storage,
	}
}

func (uc UseCase) Run(user entity.User) error {
	if err := uc.Storage.SaveUser(user); err != nil{
		return errtrace.Wrap(err)
	}

	return nil
}
