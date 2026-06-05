package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/sevelfatt/taskverse-backend/core/auth"
	"github.com/sevelfatt/taskverse-backend/initilization"
	"github.com/sevelfatt/taskverse-backend/lib"
	_ "github.com/sevelfatt/taskverse-backend/middlewares"
)

func init() {
	initilization.DotEnvInit()
	lib.ConnectMongoDB()
}

func main(){
	port := "8000"
	if port == "" {
		log.Fatal("PORT environment variable is not set!")
	}

	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()

	// Register routes
	auth.Route(api)

	// Print all registered routes
	printRegisteredRoutes(r)

	server := &http.Server{
		Addr: ":" + port,
		Handler: r,
	}

	log.Printf("Server is running on port %s", port)
	err := server.ListenAndServe()
	log.Fatal(err)
}

func printRegisteredRoutes(r *mux.Router) {
	log.Println("Registered Routes:")
	err := r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		pathTemplate, err := route.GetPathTemplate()
		if err == nil {
			methods, _ := route.GetMethods()
			methodStr := "ANY"
			if len(methods) > 0 {
				methodStr = strings.Join(methods, ",")
			}
			log.Printf("  %-6s %s", methodStr, pathTemplate)
		}
		return nil
	})
	if err != nil {
		log.Println("Error printing routes:", err)
	}
}