package setup

import (
	"log"
	"net/http"
	"time"

	core "api/internal/core/database/interface"
	"api/internal/di"
	"api/internal/infra/auth/keycloak"
	"api/internal/infra/router"
	"api/internal/modules/signup/controllers"
	"api/internal/setup/database"
)

const (
	maxRetries    = 10               // Número máximo de tentativas
	retryInterval = 2 * time.Second // Intervalo entre tentativas
)

func InitAll() (*controllers.AuthController, core.Database, *router.MuxRouter, error) {
	dbConnection, err := database.InitDatabase()
	if err != nil {
		log.Fatalf("Erro ao inicializar o banco de dados: %v", err)
		return nil, nil, nil, err
	}

	var keycloakAuth *keycloak.KeycloakAuthenticator
	for i := 0; i < maxRetries; i++ {
		keycloakAuth, err = keycloak.NewKeycloakAuthenticator()
		if err == nil {
			break // Sucesso na conexão com o Keycloak
		}
		log.Printf("Tentativa %d de conexão com Keycloak falhou: %v", i+1, err)
		time.Sleep(retryInterval) // Aguarda antes de tentar novamente
	}

	if err != nil {
		log.Fatalf("Erro ao inicializar Keycloak após %d tentativas: %v", maxRetries, err)
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
