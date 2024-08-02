package users

import "github.com/wjojf/go-uber-fx/internal/domain/users/models"

type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r CreateUserRequest) ToValueObject() (models.UserValueObject, error) {
	return models.NewUserValueObject(r.Username, r.Email, r.Password, false)
}

type UpdateFullRequest struct {
	Username   string `json:"username"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	IsVerified bool   `json:"isVerified"`
}

func (r UpdateFullRequest) ToValueObject() (models.UserValueObjectPartial, error) {

	var fields []string = make([]string, 0, len(models.ValidFields))
	for k := range models.ValidFields {
		fields = append(fields, k)
	}

	return models.NewUserValueObjectPartial(r.Username, r.Email, r.Password, r.IsVerified, fields)
}

type UpdatePartialRequest struct {
	Username   string `json:"username"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	IsVerified bool   `json:"isVerified"`

	Fields []string `json:"fields"`
}

func (r UpdatePartialRequest) ToValueObject() (models.UserValueObjectPartial, error) {
	return models.NewUserValueObjectPartial(r.Username, r.Email, r.Password, r.IsVerified, r.Fields)
}
