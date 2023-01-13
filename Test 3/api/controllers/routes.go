package controllers

import "github.com/zhafranammar/rest-api/api/middlewares"

func (s *Server) initializeRoutes() {

	// Login Route
	s.Router.HandleFunc("/api/v1/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	//Users routes
	s.Router.HandleFunc("/api/v1/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/api/v1/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/api/v1/users/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/api/v1/users/{id}", middlewares.SetMiddlewareJSON(s.DeleteUser)).Methods("DELETE")

	//Author routes
	s.Router.HandleFunc("/api/v1/authors", middlewares.SetMiddlewareJSON(s.CreateAuthor)).Methods("POST")
	s.Router.HandleFunc("/api/v1/authors", middlewares.SetMiddlewareJSON(s.GetAuthors)).Methods("GET")
	s.Router.HandleFunc("/api/v1/authors/{id}", middlewares.SetMiddlewareJSON(s.GetAuthor)).Methods("GET")
	s.Router.HandleFunc("/api/v1/authors/{id}", middlewares.SetMiddlewareJSON(s.UpdateAuthor)).Methods("PUT")
	s.Router.HandleFunc("/api/v1/authors/{id}", middlewares.SetMiddlewareJSON(s.DeleteAuthor)).Methods("DELETE")

	//Publisher routes
	s.Router.HandleFunc("/api/v1/publishers", middlewares.SetMiddlewareJSON(s.CreatePublisher)).Methods("POST")
	s.Router.HandleFunc("/api/v1/publishers", middlewares.SetMiddlewareJSON(s.GetPublishers)).Methods("GET")
	s.Router.HandleFunc("/api/v1/publishers/{id}", middlewares.SetMiddlewareJSON(s.GetPublisher)).Methods("GET")
	s.Router.HandleFunc("/api/v1/publishers/{id}", middlewares.SetMiddlewareJSON(s.UpdatePublisher)).Methods("PUT")
	s.Router.HandleFunc("/api/v1/publishers/{id}", middlewares.SetMiddlewareJSON(s.DeletePublisher)).Methods("DELETE")

	//Book routes
	s.Router.HandleFunc("/api/v1/books", middlewares.SetMiddlewareJSON(s.CreateBook)).Methods("POST")
	s.Router.HandleFunc("/api/v1/books", middlewares.SetMiddlewareJSON(s.GetBooks)).Methods("GET")
	s.Router.HandleFunc("/api/v1/books/{id}", middlewares.SetMiddlewareJSON(s.GetBook)).Methods("GET")
	s.Router.HandleFunc("/api/v1/books/publisher/{id}", middlewares.SetMiddlewareJSON(s.GetBookPublisher)).Methods("GET")
	s.Router.HandleFunc("/api/v1/books/author/{id}", middlewares.SetMiddlewareJSON(s.GetBookAuthor)).Methods("GET")
	s.Router.HandleFunc("/api/v1/books/{id}", middlewares.SetMiddlewareJSON(s.UpdateBook)).Methods("PUT")
	s.Router.HandleFunc("/api/v1/books/{id}", middlewares.SetMiddlewareJSON(s.DeleteBook)).Methods("DELETE")

}