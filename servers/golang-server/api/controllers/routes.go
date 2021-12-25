package controllers

import "github.com/tabinnorway/dive-arch/golang-server/api/middlewares"

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Login Route
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	//Users routes
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")

	//Dives routes
	s.Router.HandleFunc("/dives", middlewares.SetMiddlewareJSON(s.CreateDive)).Methods("POST")
	s.Router.HandleFunc("/dives", middlewares.SetMiddlewareJSON(s.GetDives)).Methods("GET")
	s.Router.HandleFunc("/dives/{id}", middlewares.SetMiddlewareJSON(s.GetDive)).Methods("GET")
	s.Router.HandleFunc("/dives/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateDive))).Methods("PUT")
	s.Router.HandleFunc("/dives/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteDive)).Methods("DELETE")
}
