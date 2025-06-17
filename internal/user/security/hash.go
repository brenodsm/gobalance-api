package security

import "golang.org/x/crypto/bcrypt"

// Hash gera um hash bcrypt para a senha fornecida.
func Hash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}
