package models

type User struct {
	ID           int64
	Name         string
	Phone        string
	Email        string
	PasswordHash []byte
}
