package routes

import (
	"log"
	"net/http"

	"github.com/geo-albin/bookStore/pkg/controllers"

	"github.com/gorilla/mux"
)

var (
	r *mux.Router
)

func init() {
	r = mux.NewRouter()
}

func AddRoutes() error {
	log.Println("Start adding the route")
	r.HandleFunc("/books", controllers.GetBooks).Methods("GET")
	r.HandleFunc("/books/{id:[0-9]+}", controllers.GetBook).Methods("GET")
	r.HandleFunc("/books/{id:[0-9]+}", controllers.UpdateBook).Methods("PUT")
	r.HandleFunc("/books", controllers.CreateBook).Methods("POST")
	r.HandleFunc("/books/{id:[0-9]+}", controllers.DeleteBook).Methods("DELETE")

	http.Handle("/", r)
	return nil
}

func StartAndServe() error {
	err := http.ListenAndServe(":8080", r)
	return err
}
