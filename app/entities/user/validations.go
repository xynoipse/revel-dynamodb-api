package user

import (
	validate "revel-dynamodb-api/app/utils/validation"

	"github.com/revel/revel"
)

// Validate User
func (user *Model) Validate(v *revel.Validation, service int) {
	switch service {

	case CreateService:
		validate.Email(v, user.Email).Key("email")
		validate.Password(v, user.Password).Key("password")
		validate.Name(v, user.FirstName).Key("first_name")
		validate.Name(v, user.LastName).Key("last_name")

	}
}
