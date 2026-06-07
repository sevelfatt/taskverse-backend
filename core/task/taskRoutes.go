package task

import (
	"github.com/gorilla/mux"
	"github.com/sevelfatt/taskverse-backend/middlewares"
)

func Route(router *mux.Router) *mux.Router{
	taskRouter := router.PathPrefix("/task").Subrouter()
	taskRouter.HandleFunc("/", middlewares.AuthMiddelware(GetAllTasksByUserUUIDController)).Methods("GET")
	taskRouter.HandleFunc("/", middlewares.AuthMiddelware(CreateTaskController)).Methods("POST")

	return taskRouter
}