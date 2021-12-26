package controllers

import (
	"fmt"
	"net/http"

	"github.com/tabinnorway/dive-arch/golang-server/api/middlewares"
)

func (s *Server) initializeRoutes() {
	fmt.Println("Initializing routes")
	// Home Route
	// s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Login Route
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	//Users routes
	s.Router.HandleFunc("/api/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/api/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/api/users/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/api/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")

	//Dives routes
	s.Router.HandleFunc("/api/dives", middlewares.SetMiddlewareJSON(s.CreateDive)).Methods("POST")
	s.Router.HandleFunc("/api/dives", middlewares.SetMiddlewareJSON(s.GetDives)).Methods("GET")
	s.Router.HandleFunc("/api/dives/{id}", middlewares.SetMiddlewareJSON(s.GetDive)).Methods("GET")
	s.Router.HandleFunc("/api/dives/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateDive))).Methods("PUT")
	s.Router.HandleFunc("/api/dives/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteDive)).Methods("DELETE")

	s.Router.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))
}
