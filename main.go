package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"go-simple-api/configs"
	"go-simple-api/controllers"
	"go-simple-api/helpers"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.WithFields(log.Fields{
		"status": "Running Well",
		"number": 1,
	}).Info("Run API")

	// Initialize Channel for OS Signal
	signalOS := make(chan os.Signal, 1)

	// Initialize Configuration
	configs.Initialize()

	// Initialize CORS Configurataion
	corsAllowedHeaders := handlers.AllowedHeaders(configs.CORSAllowedHeaders)
	corsAllowedOrigins := handlers.AllowedHeaders(configs.CORSAllowedOrigins)
	corsAllowedMethods := handlers.AllowedHeaders(configs.CORSAllowedMethods)

	// Initialize Router
	router := mux.NewRouter()

	// Initialize Router Endpoint
	router.HandleFunc("/", controllers.GetIndex).Methods("GET")

	// Initialize Router Endpoint Secured With Basic Auth
	router.Handle("/basic-auth", helpers.AuthBasic(controllers.GetAuthentication)).Methods("GET", "POST")

	// Initialize Router Endpoint Secured With Authorization
	router.Handle("/users", helpers.AuthJWT(controllers.GetUser)).Methods("GET")
	router.Handle("/users", helpers.AuthJWT(controllers.AddUser)).Methods("POST")
	router.Handle("/users/{id}", helpers.AuthJWT(controllers.GetUserById)).Methods("GET")
	router.Handle("/users/{id}", helpers.AuthJWT(controllers.PutUserById)).Methods("PUT", "PATCH")
	router.Handle("/users/{id}", helpers.AuthJWT(controllers.DelUserById)).Methods("DELETE")

	// Set Router Handler with Logging & CORS Support
	routerHandler := handlers.LoggingHandler(os.Stdout, handlers.CORS(corsAllowedHeaders, corsAllowedOrigins, corsAllowedMethods)(router))

	// Initialize Server With Initialized Router
	server := helpers.NewServer(routerHandler)

	// Starting Server
	server.Start()
	defer server.Stop()

	// Catch OS Signal from Channel
	signal.Notify(signalOS, os.Interrupt, syscall.SIGTERM)

	// Return OS Signal as Exit Code
	<-signalOS

	// Add Some Spaces When Done
	fmt.Println("")
}
