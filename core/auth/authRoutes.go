package auth

import (

	"github.com/gorilla/mux"
)

func Route(router *mux.Router) *mux.Router{
	authRouter := router.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/register",RegisterController).Methods("POST")
	authRouter.HandleFunc("/login",LoginController).Methods("POST")
	authRouter.HandleFunc("/get-user",GetUserController).Methods("GET")

	return authRouter
}