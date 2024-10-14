package main

import (
    "log"
    "api/internal/setup"
    "net/http"
)

func main() {
    _, database, err := setup.InitAll()
    if err != nil {
        log.Fatalf("Erro ao iniciar a aplicação: %v", err)
    }
    defer database.Close()

    log.Println("Servidor iniciado na porta 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
