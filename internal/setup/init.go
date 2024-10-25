package setup

import (
    "log"
    "api/internal/di"
    "api/internal/core/database"
    "api/internal/infra/router"
    "net/http"
    "api/internal/modules/signup/controllers"
    "api/internal/setup/database"
)

func InitAll() (*controllers.AuthController, core.Database, error) {
    dbConnection, err := database.InitDatabase()
    if err != nil {
        log.Fatalf("Erro ao inicializar o banco de dados: %v", err)
        return nil, nil, err
    }

    authController, err := di.InjectDependencies(dbConnection)
    if err != nil {
        log.Fatalf("Erro ao injetar dependências: %v", err)
        return nil, nil, err
    }

    muxRouter := &router.MuxRouter{}
    muxRouter.NewRouter()

    muxRouter.SetupRoutes(authController)

    http.Handle("/", muxRouter)

    log.Println("Serviços e dependências inicializados com sucesso.")
    
    return authController, dbConnection, nil
}
