package auth

import (
	"github.com/gorilla/mux"
	"github.com/sevelfatt/taskverse-backend/middlewares"
)

func Route(router *mux.Router) *mux.Router{
	authRouter := router.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/register",RegisterController).Methods("POST")
	authRouter.HandleFunc("/login",LoginController).Methods("POST")
	authRouter.HandleFunc("/get-user",middlewares.AuthMiddelware(GetUserController)).Methods("GET")

	return authRouter
}