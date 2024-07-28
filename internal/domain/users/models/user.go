package models

type User struct {
	id         string
	username   string
	isVerified bool
}

func NewUser(id string, username string, isVerified bool) (User, error) {
	return User{id: id, username: username}, nil
}

func (u User) ID() string {
	return u.id
}

func (u User) Username() string {
	return u.username
}

func (u User) IsVerified() bool {
	return u.isVerified
}

type UserValueObject struct {
	Username   string
	Email      string
	Password   string
	IsVerified bool
}

func NewUserValueObject(Username string, Email string, Password string, isVerified bool) (UserValueObject, error) {
	return UserValueObject{
		Username:   Username,
		Email:      Email,
		Password:   Password,
		IsVerified: isVerified,
	}, nil
}
