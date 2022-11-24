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

	// Admin
	admin := mux.PathPrefix("/admin").Subrouter()

	adminAuth := admin.PathPrefix("/auth").Subrouter()
	adminAuth.HandleFunc("/sign-in", h.adminSignIn).Methods("POST", "OPTIONS")

	adminApi := admin.PathPrefix("/api").Subrouter()
	adminApi.Use(h.adminIdentity)

	adminUser := adminApi.PathPrefix("/users").Subrouter()
	adminUser.HandleFunc("", h.adminGetUsers).Methods("GET", "OPTIONS")
	adminUser.HandleFunc("", h.adminCreateUser).Methods("POST", "OPTIONS")
	adminUser.HandleFunc("/{id}", h.adminDeleteUser).Methods("DELETE", "OPTIONS")
	adminUser.HandleFunc("/{id}", h.adminUpdateUser).Methods("PUT", "OPTIONS")

	adminProducts := adminApi.PathPrefix("/products").Subrouter()
	adminProducts.HandleFunc("", h.adminGetProducts).Methods("GET", "OPTIONS")
	adminProducts.HandleFunc("", h.adminCreateProducts).Methods("POST", "OPTIONS")
	adminProducts.HandleFunc("/{id}", h.adminDeleteProducts).Methods("DELETE", "OPTIONS")
	adminProducts.HandleFunc("/{id}", h.adminUpdateProducts).Methods("PUT", "OPTIONS")

	// User
	auth := mux.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/sign-up", h.signUp).Methods("POST", "OPTIONS")
	auth.HandleFunc("/sign-in", h.signIn).Methods("POST", "OPTIONS")

	mux.HandleFunc("/product/all", h.productGetAll)

	api := mux.PathPrefix("/api").Subrouter()
	api.Use(h.userIdentity)

	user := api.PathPrefix("/user").Subrouter()
	user.HandleFunc("", h.profilePage).Methods("GET", "OPTIONS")

	product := api.PathPrefix("/product").Subrouter()
	product.HandleFunc("", h.productCreate).Methods("POST", "OPTIONS")
	product.HandleFunc("", h.productFind).Methods("GET", "OPTIONS")
	product.HandleFunc("/{id}", h.productUpdate).Methods("PUT", "OPTIONS")
	product.HandleFunc("/{id}", h.productDelete).Methods("DELETE", "OPTIONS")

	return mux
}
