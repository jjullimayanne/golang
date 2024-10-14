package database

import (
    "log"
    "api/internal/infra/db"
    "api/configs/db"
)


func InitDatabase() (*db.PostgresDB, error) {
    databaseConfig, err := configs.LoadDBConfig()
    if err != nil {
        log.Printf("Falha ao carregar a configuração do banco de dados: %v", err)
        return nil, err
    }

    postgresDatabase := db.NewPostgresDB(databaseConfig)
    _, err = postgresDatabase.Connect()
    if err != nil {
        log.Printf("Não foi possível conectar ao banco de dados: %v", err)
        return nil, err
    }

    return postgresDatabase, nil
}
