package models

type User struct {
	ID       int
	Username string
}

func NewUser(Id int, Username string) (User, error) {
	return User{ID: Id, Username: Username}, nil
}
