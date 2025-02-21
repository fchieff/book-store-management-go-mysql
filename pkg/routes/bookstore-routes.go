package routes

import (
	"github.com/fchieff/book-store-management-go-mysql.git/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/books", controllers.BookController.GetBooks).Methods("GET")
	router.HandleFunc("/book", controllers.BookController.CreateBook).Methods("POST")
	router.HandleFunc("/book/{id}", controllers.BookController.GetBook).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.BookController.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{id}", controllers.BookController.DeleteBook).Methods("DELETE")
}
