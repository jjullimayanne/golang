package database

import (
	dbInterface "api/internal/core/database/interface"
	dbConfig "api/internal/core/database/struct"
	coreError "api/internal/core/error"
	"api/internal/infra/db"
	"time"
)

const (
	maxRetries    = 5               
	retryInterval = 2 * time.Second 
	initialDelay  = 5 * time.Second
)

func InitDatabase() (dbInterface.Database, error) {
	databaseConfig, err := dbConfig.LoadDBConfig()
	if err != nil {
		coreError.LogError(err)
		return nil, coreError.WrapError(err, "falha ao carregar a configuração do banco de dados")
	}

	time.Sleep(initialDelay)

	database := db.NewDatabase(databaseConfig)

	for i := 0; i < maxRetries; i++ {
		_, err = database.Connect()
		if err == nil {
			coreError.LogError(coreError.WrapError(err, "DONEDONEDONEDONEDONEDOEEEEEEEEEE"))
			return database, nil
		}

		coreError.LogError(coreError.WrapError(err, "tentativa de conexão falhou"))

		time.Sleep(retryInterval)
	}

	wrappedErr := coreError.WrapError(err, "não foi possível conectar ao banco de dados após várias tentativas")
	coreError.LogError(wrappedErr)
	return nil, wrappedErr
}

//TODO: change to health check later 
