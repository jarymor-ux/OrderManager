package entity

type User struct {
	ID           int64 `json:"id"`
	Username     string `json:"username"`
	PasswordHash string `json:"-"`
}

func (u *User) WithID(id int64) *User {
	u.ID = id
	return u
}

func (u *User) WithUsername(username string) *User{
	u.Username = username
	return  u
}

func (u *User) WithPasswordHash(hash string) *User{
	u.PasswordHash = hash
	return u
}
