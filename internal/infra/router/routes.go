package router

import (
    "github.com/gorilla/mux"
    "api/internal/core/routes"
    "net/http"
    "api/internal/modules/signup/controllers"
    "api/internal/modules/health/controllers"  

)

type MuxRouter struct {
    muxRouter *mux.Router
}

func (r *MuxRouter) NewRouter() core.Router {
    r.muxRouter = mux.NewRouter()
    return r
}

func (r *MuxRouter) SetupRoutes(authController *controllers.AuthController) {
    r.muxRouter.HandleFunc("/signup", authController.RegisterUser).Methods("POST")

    r.muxRouter.HandleFunc("/health", healthController.HealthCheckHandler).Methods("GET")
}

func (r *MuxRouter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    r.muxRouter.ServeHTTP(w, req)
}
