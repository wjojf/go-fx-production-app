package models

import "fmt"

var (
	ValidFields = map[string]interface{}{
		"username":   nil,
		"email":      nil,
		"password":   nil,
		"isVerified": nil,
	}
)

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

type UserValueObjectPartial struct {
	UserValueObject
	Fields []string
}

func NewUserValueObjectPartial(
	Username string,
	Email string,
	Password string,
	isVerified bool,
	fields []string,
) (UserValueObjectPartial, error) {

	vo, err := NewUserValueObject(Username, Email, Password, isVerified)
	if err != nil {
		return UserValueObjectPartial{}, err
	}

	for _, field := range fields {
		if _, ok := ValidFields[field]; !ok {
			return UserValueObjectPartial{}, fmt.Errorf("invalid field %s", field)
		}
	}

	return UserValueObjectPartial{UserValueObject: vo, Fields: fields}, nil
}
