package core

import (
    "net/http"
    "api/internal/modules/signup/controllers"
)

type Router interface {
    NewRouter() Router                              
    SetupRoutes(authController *controllers.AuthController)  
    ServeHTTP(w http.ResponseWriter, r *http.Request) 
}
