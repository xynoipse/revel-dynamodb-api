package user

import (
	pass "revel-dynamodb-api/app/utils/password"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/revel/revel"
)

// GenerateAuthToken returns User jwt access token string
func (user *Model) GenerateAuthToken() (accessToken string, err error) {
	jwtSecret, found := revel.Config.String("jwt.secret")
	if !found {
		revel.RevelLog.Fatal("jwt.secret not configured")
	}

	// Create token and set claims
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   user.UserID,
		"email": user.Email,
		"iat":   time.Now().Unix(),                                     // Issued at
		"exp":   time.Now().Add(time.Minute * 15).Unix(),               // Expiration time
		"nbf":   time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(), // Not valid before
	})

	// Generate signed access token using the jwt secret
	accessToken, err = jwtToken.SignedString([]byte(jwtSecret))
	if err != nil {
		return
	}

	return
}

// EncryptPassword encrypt password using argon2 hashing algorithm
func (user *Model) EncryptPassword() (err error) {
	user.Password, err = pass.Hash(user.Password)
	if err != nil {
		return
	}
	return
}

// ValidatePassword compare password to the hashed password if valid
func (user *Model) ValidatePassword(password string) (match bool, err error) {
	match, err = pass.Compare(password, user.Password)
	if err != nil {
		return
	}
	return
}

// IsEmailNew returns if email is new and doesn't already exist
func (user *Model) IsEmailNew() bool {
	if _, err := FindByEmail(user.Email); err != nil {
		return true
	}
	return false
}
