package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/brenodsm/gobalance-api/config"
	"github.com/brenodsm/gobalance-api/internal/router"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	r := router.RunServer()
	fmt.Printf("Servidor rodando na porta %d", cfg.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), r); err != nil {
		log.Fatalf("Erro ao inicializar o servidor %v", err)
	}
}
