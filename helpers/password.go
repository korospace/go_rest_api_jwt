package helpers

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	passwordHas, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(passwordHas), err
}

func VerifyPassword(hasPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hasPassword), []byte(password))

	return err
}
