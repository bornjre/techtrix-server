package app

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Run() {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodOptions},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Api-key", "Content-Type", "Accept"},
	})
	fmt.Println(http.ListenAndServe(":8080", c.Handler(GetRouter())))
}

func GetRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/transactions", addTransaction).Methods("POST")
	r.HandleFunc("/transactions", updateTransaction).Methods("PUT")
	r.HandleFunc("/transactions", getTransactions).Methods("GET")
	r.HandleFunc("/transactions/{hashid}", getTransaction).Methods("GET")
	r.HandleFunc("/verify", verify).Methods("GET")
	r.HandleFunc("/search/{str}", search).Methods("GET")
	r.HandleFunc("/publish", publish).Methods("GET")
	r.HandleFunc("/subscribe", subscribersService.subscribe)
	r.HandleFunc("/login", SignIn).Methods("POST")

	return r

}
