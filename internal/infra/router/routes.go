package router

import (
	core "api/internal/core/routes"
	healthController "api/internal/modules/health/controllers"
	"api/internal/modules/signup/controllers"
	"log"
	"net/http"
	"github.com/gorilla/mux"
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
	log.Println("Endpoint /signup registrado")

	r.muxRouter.HandleFunc("/health", healthController.HealthCheckHandler).Methods("GET")
	log.Println("Endpoint /health registrado")

	r.PrintRoutes()
}

func (r *MuxRouter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.muxRouter.ServeHTTP(w, req)
}

func (r *MuxRouter) PrintRoutes() {
	r.muxRouter.Walk(func(route *mux.Route, _ *mux.Router, _ []*mux.Route) error {
		path, err := route.GetPathTemplate()
		if err != nil {
			return err
		}
		methods, err := route.GetMethods()
		if err != nil {
			return err
		}
		log.Printf("Rota registrada: %s, Métodos: %v\n", path, methods)
		return nil
	})
}
