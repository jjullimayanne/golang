package setup

import (
	"log"
	"net/http"

	core "api/internal/core/database/interface" // Importa a interface do banco de dados
	"api/internal/di"
	"api/internal/infra/auth/keycloak"
	"api/internal/infra/router"
	"api/internal/modules/signup/controllers"
	"api/internal/setup/database"
)

func InitAll() (*controllers.AuthController, core.Database, *router.MuxRouter, error) {
	dbConnection, err := database.InitDatabase()
	if err != nil {
		log.Fatalf("Erro ao inicializar o banco de dados: %v", err)
		return nil, nil, nil, err
	}

	keycloakAuth, err := keycloak.NewKeycloakAuthenticator()
	if err != nil {
		log.Fatalf("Erro ao inicializar Keycloak: %v", err)
		return nil, nil, nil, err
	}

	authController, err := di.InjectDependencies(dbConnection, keycloakAuth)
	if err != nil {
		log.Fatalf("Erro ao injetar dependências: %v", err)
		return nil, nil, nil, err
	}

	muxRouter := &router.MuxRouter{}
	muxRouter.NewRouter()
	muxRouter.SetupRoutes(authController)

	http.Handle("/", muxRouter)

	log.Println("Serviços e dependências inicializados com sucesso.")

	return authController, dbConnection, muxRouter, nil
}
