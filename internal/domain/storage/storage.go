package storage

import "orders/internal/domain/entity"

type Storage interface{
	SaveUser(user entity.User) error
	//GetUserById(id int64) entity.User
	//DeleteUserById(id int64) error
	//UpdateUserById(id int64) error
} 
