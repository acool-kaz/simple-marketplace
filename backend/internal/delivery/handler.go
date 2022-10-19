package delivery

import (
	"context"
	"main/internal/service"

	"github.com/gorilla/mux"
)

type Handler struct {
	service *service.Service
	ctx     context.Context
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
		ctx:     context.Background(),
	}
}

func (h *Handler) InitRoutes() *mux.Router {
	mux := mux.NewRouter()
	mux.Use(h.corsMiddleware)
	mux.Use(h.loggingMiddleware)

	auth := mux.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/sign-up", h.signUp).Methods("POST", "OPTIONS")
	auth.HandleFunc("/sign-in", h.signIn).Methods("POST", "OPTIONS")

	mux.HandleFunc("/product/all", h.productGetAll)

	api := mux.PathPrefix("/api").Subrouter()
	api.Use(h.userIdentity)
	api.HandleFunc("/profile", h.profilePage).Methods("GET", "OPTIONS")
	api.HandleFunc("/find", h.findUsers).Methods("POST", "OPTIONS")

	product := api.PathPrefix("/product").Subrouter()
	product.HandleFunc("/create", h.productCreate).Methods("POST", "OPTIONS")
	product.HandleFunc("/find", h.productFind).Methods("POST", "OPTIONS")
	product.HandleFunc("/{id}", h.productUpdate).Methods("PUT", "OPTIONS")
	product.HandleFunc("/{id}", h.productDelete).Methods("DELETE", "OPTIONS")

	return mux
}
