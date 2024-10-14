package router

import (
    "github.com/gorilla/mux"
    "api/internal/core/routes"
    "net/http"
    "api/internal/modules/signup/controllers"
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
}

func (r *MuxRouter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    r.muxRouter.ServeHTTP(w, req)
}
