package handlers

import (
	"orders/internal/domain/storage"
	createuseruc "orders/internal/domain/usecases/create_user_uc"
)

type Handlers struct {
	CreateUserUc createuseruc.UseCase	
}

func New(storage storage.Storage) Handlers {
	return Handlers{
		CreateUserUc: createuseruc.UseCase{
			Storage: storage,
		},
	}
}
