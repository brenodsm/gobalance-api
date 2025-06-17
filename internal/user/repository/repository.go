package repository

import (
	"database/sql"
	"errors"

	"github.com/brenodsm/gobalance-api/internal/user/model"
	"github.com/jackc/pgx/v5/pgconn"
)

// ErrUserAlreadyRegister indica que o usuário já está cadastrado no sistema.
var ErrUserAlreadyRegister error = errors.New("usuário já cadastrado")

// userRepository implementa operações de acesso a dados para usuários.
type userRepository struct {
	db *sql.DB
}

// NewUserRepository cria e retorna um novo repositório de usuários.
func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{db}
}

// Create insere um novo usuário no banco de dados e retorna seu ID.
func (u *userRepository) Create(user model.User) (uint64, error) {
	var id uint64
	query := "INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id"
	err := u.db.QueryRow(query, user.Name, user.Email, user.Password).Scan(&id)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return 0, ErrUserAlreadyRegister
		}
		return 0, err
	}
	return id, nil
}
