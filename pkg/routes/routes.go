package routes

import (
	"main/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterRoutes = func(r *mux.Router) {
	r.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	r.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	r.HandleFunc("/book/{bookId}", controllers.GetBookByID).Methods("GET")
	r.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	r.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
