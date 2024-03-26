package types

import (
	"net/mail"
	"stratplus/errors"
	"strconv"
	"unicode"

	"gorm.io/gorm"

	"github.com/dgrijalva/jwt-go"
)

// User represents a user.
type User struct {
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	Phone    string `json:"phone" gorm:"unique"`
	User     string `json:"user"`
	gorm.Model
}

// ValidateFields validates the fields of a user.
func (u *User) ValidateFields() error {
	if u.User == "" {
		return errors.ErrUserRequired
	}

	if u.Email == "" {
		return errors.ErrEmailRequired
	}

	_, err := mail.ParseAddress(u.Email)
	if err != nil {
		return errors.ErrInvalidEmail
	}

	if u.Password == "" {
		return errors.ErrPasswordRequired
	}

	if err := u.verifyPassword(); err != nil {
		return err
	}

	if u.Phone == "" {
		return errors.ErrPhoneRequired
	}

	if err := u.validatePhone(); err != nil {
		return err
	}

	return nil
}

// ValidateLoginFields validates the fields of a user for login.
func (u *User) ValidateLoginFields() error {
	if u.Email == "" {
		return errors.ErrEmailRequired
	}

	_, err := mail.ParseAddress(u.Email)
	if err != nil {
		return errors.ErrInvalidEmail
	}

	if u.Password == "" {
		return errors.ErrPasswordRequired
	}

	return nil
}

// verifyPassword verifies the password with rules.
func (u *User) verifyPassword() error {
	var (
		upp, low, num, sym bool
		tot                uint8
	)

	for _, char := range u.Password {
		switch {
		case unicode.IsUpper(char):
			upp = true
			tot++
		case unicode.IsLower(char):
			low = true
			tot++
		case unicode.IsNumber(char):
			num = true
			tot++
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			sym = true
			tot++
		default:
			return errors.ErrInvalidPassword
		}
	}

	if !upp || !low || !num || !sym || tot < 6 || tot > 12 {
		return errors.ErrInvalidPassword
	}

	return nil
}

func (u *User) validatePhone() error {
	if len(u.Phone) != 10 {
		return errors.ErrPhoneNumberInvalid
	}

	if _, err := strconv.Atoi(u.Phone); err != nil {
		return errors.ErrPhoneNumberInvalid
	}

	return nil
}

// GenerateToken generates a token for a user.
func (u *User) GenerateToken() string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = u.Email
	claims["exp"] = 15000

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return errors.ErrTokenGeneration.Error()
	}

	return tokenString
}
