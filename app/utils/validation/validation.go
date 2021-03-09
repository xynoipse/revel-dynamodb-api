package validation

import (
	"regexp"

	"github.com/revel/revel"
)

// Required validates required field
func Required(v *revel.Validation, field interface{}) *revel.ValidationResult {
	return v.Check(field, revel.Required{}).Message(RequiredField)
}

// UUID validates a requied uuid field
func UUID(v *revel.Validation, uuid string) *revel.ValidationResult {
	return v.Check(uuid,
		revel.Required{},
		revel.Match{Regexp: regexp.MustCompile(`^[0-9a-fA-F]{8}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{12}$`)},
	).Message(InvalidUUID)
}

// Name validates a required name field
func Name(v *revel.Validation, name string) *revel.ValidationResult {
	return v.Check(name,
		revel.Required{},
		revel.MaxSize{Max: 255},
		revel.Match{Regexp: regexp.MustCompile(`^[^±!@£$%^&*_+§¡€#¢§¶•ªº«\/<>?:;|=.,]{1,20}$`)},
	).Message(InvalidName)
}

// Password validates required password field
func Password(v *revel.Validation, password string) *revel.ValidationResult {
	return v.Check(password,
		revel.Required{},
		revel.MinSize{Min: 8},
	).Message(RequiredField)
}

// Email validates required email field
func Email(v *revel.Validation, email string) *revel.ValidationResult {
	return v.Check(email,
		revel.Required{},
		revel.ValidEmail().Match,
	).Message(InvalidEmail)
}

// URL validates required url field
func URL(v *revel.Validation, url string) *revel.ValidationResult {
	return v.Check(url,
		revel.Required{},
		revel.ValidURL(),
	).Message(InvalidURL)
}
