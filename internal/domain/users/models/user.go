package models

type User struct {
	id       string
	username string
}

func NewUser(Id string, Username string) (User, error) {
	return User{id: Id, username: Username}, nil
}

func (u User) ID() string {
	return u.id
}

func (u User) Username() string {
	return u.username
}

type UserValueObject struct {
	Username string
	Email    string
	Password string
}

func NewUserValueObject(Username string, Email string, Password string) (UserValueObject, error) {
	return UserValueObject{
		Username: Username,
		Email:    Email,
		Password: Password,
	}, nil
}
