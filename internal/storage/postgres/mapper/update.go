package mapper

import (
	"fmt"

	"github.com/wjojf/go-uber-fx/internal/domain/users/models"
)

func GetUpdateUserQuery(userID string, vo models.UserValueObjectPartial) string {

	if len(vo.Fields) == 0 {
		return ""
	}

	query := "UPDATE users SET "

	for i, field := range vo.Fields {

		if field == "username" {
			query += "username = " + fmt.Sprintf("'%s'", vo.Username)
		}

		if field == "email" {
			query += "email = " + fmt.Sprintf("'%s'", vo.Email)
		}

		if field == "password" {
			query += "password = " + fmt.Sprintf("'%s'", vo.Password)
		}

		if field == "isVerified" {
			value := "false"
			if vo.IsVerified {
				value = "true"
			}
			query += "is_verified = " + value
		}

		if i != len(vo.Fields)-1 {
			query += ", "
		}
	}

	return fmt.Sprintf("%s WHERE id = '%s';", query, userID)
}
