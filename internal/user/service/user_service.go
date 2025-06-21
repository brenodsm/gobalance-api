package service

import (
	"errors"
	"strings"

	"github.com/badoux/checkmail"
	"github.com/brenodsm/gobalance-api/internal/user/model"
	"github.com/brenodsm/gobalance-api/internal/user/security"
)

// userService contém regras de negócio relacionadas a usuários.
type userService struct{}

// NewUserService cria e retorna uma instância de userService.
func NewUserService() *userService {
	return &userService{}
}

// Register aplica validações e sanitiza os dados do usuário antes do cadastro.
func (u *userService) Register(user *model.User) error {
	u.sanitizeUserInput(user)
	if err := u.validate(user); err != nil {
		return err
	}

	if err := u.validatePassword(user.Password); err != nil {
		return err
	}

	if err := u.hashPassword(user); err != nil {
		return err
	}
	return nil
}

// validate verifica se os campos obrigatórios do usuário estão preenchidos e se o e-mail é válido.
func (u *userService) validate(user *model.User) error {
	if user.Name == "" {
		return errors.New("o nome não pode estar em branco")
	}
	if user.Email == "" {
		return errors.New("o email não pode estar em branco")
	}
	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("email inválido")
	}
	if user.Password == "" {
		return errors.New("a senha não pode estar em branco")
	}
	return nil
}

// validatePassword verifica se a senha atende aos requisitos mínimos de segurança.
func (u *userService) validatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("a senha deve ter no mínimo 8 caracteres")
	}
	for _, l := range password {
		if l == ' ' {
			return errors.New("a senha não pode conter espaços")
		}
	}
	return nil
}

// sanitizeUserInput padroniza os campos de entrada do usuário, removendo espaços em branco.
func (u *userService) sanitizeUserInput(user *model.User) {
	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(user.Email)
	user.Password = strings.TrimSpace(user.Password)
}

// hashPassword gera o hash da senha e substitui o valor original no modelo do usuário.
func (u *userService) hashPassword(user *model.User) error {
	hash, err := security.Hash(user.Password)
	if err != nil {
		return err
	}
	user.Password = hash
	return nil
}
