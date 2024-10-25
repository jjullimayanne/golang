package database

import (
    "api/configs/db"
    "api/internal/infra/db"
    "api/internal/core/error"
)

func InitDatabase() (*db.PostgresDB, error) {
    databaseConfig, err := configs.LoadDBConfig()
    if err != nil {
        coreError.LogError(err) 
        return nil, coreError.WrapError(err, "falha ao carregar a configuração do banco de dados")
    }

    postgresDatabase := db.NewPostgresDB(databaseConfig)
    _, err = postgresDatabase.Connect()
    if err != nil {
        wrappedErr := coreError.WrapError(err, "não foi possível conectar ao banco de dados")
        coreError.LogError(wrappedErr) 
        return nil, wrappedErr
    }

    return postgresDatabase, nil
}
