package app

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Run() {
	fmt.Println(http.ListenAndServe(":8080", GetRouter()))
}

func GetRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/transactions", addTransaction).Methods("POST")
	r.HandleFunc("/transactions", getTransactions).Methods("GET")
	r.HandleFunc("/transactions/{hashid}", getTransaction).Methods("GET")
	r.HandleFunc("/search/{str}", search).Methods("GET")
	r.HandleFunc("/publish", publish).Methods("GET")
	r.HandleFunc("/subscribe", subscribe)
	return r

}
