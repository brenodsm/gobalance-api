package config

import (
	"fmt"
	"net/url"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config armazena as configurações essenciais para a aplicação,
// incluindo dados de conexão e porta do servidor.
type Config struct {
	ConnectionString string
	Port             int
}

// NewConfig carrega as configurações da aplicação a partir do arquivo .env
// e variáveis de ambiente, aplicando valores padrão quando necessário.
func NewConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("erro ao carregar .env: %v", err)
	}

	port, err := strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		port = 9000
	}

	password := url.QueryEscape(os.Getenv("DATABASE_PASSWORD"))

	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv("DATABASE_USER"),
		password,
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_NAME"),
		os.Getenv("DATABASE_SSLMODE"),
	)

	return &Config{
		ConnectionString: connectionString,
		Port:             port,
	}, nil
}
