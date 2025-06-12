package database

import (
	"database/sql"

	"github.com/brenodsm/gobalance-api/config"
	_ "github.com/jackc/pgx/v5/stdlib"
)

// OpenConnection estabelece e retorna uma conexão com o banco de dados PostgreSQL.
func OpenConnection() (*sql.DB, error) {
	cfg, err := config.NewConfig()
	if err != nil {
		return nil, err
	}
	db, err := sql.Open("pgx", cfg.ConnectionString)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}
