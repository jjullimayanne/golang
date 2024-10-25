package main

import (
	"api/internal/setup"
	"log"
	"net/http"
)

func main() {
	authController, database, muxRouter, err := setup.InitAll()
	if err != nil {
		log.Fatalf("Erro ao iniciar a aplicação: %v", err)
	}
	defer database.Close()

	log.Println("Servidor iniciado na porta 8080")

	if authController != nil {
		log.Println("AuthController inicializado com sucesso.")
	}

	log.Fatal(http.ListenAndServe(":8080", muxRouter))
}
