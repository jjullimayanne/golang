package di

import (
	dbInterface "api/internal/core/database/interface"
	"api/internal/infra/auth/keycloak"
	"api/internal/modules/signup/controllers"
	"api/internal/modules/signup/data/repositories"
	"api/internal/modules/signup/usecases"
)

func InjectDependencies(dbConnection dbInterface.Database, keycloakAuth *keycloak.KeycloakAuthenticator) (*controllers.AuthController, error) {
	userRepository := repositories.NewUserRepository(keycloakAuth)
	registerUserUseCase := usecases.NewRegisterUserUseCase(userRepository)
	authController := controllers.NewAuthController(registerUserUseCase)

	return authController, nil
}
