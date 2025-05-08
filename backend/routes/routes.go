package routes

import (
    "github.com/gorilla/mux"
    "net/http"
)

func RegisterRoutes(router *mux.Router) {
    router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Welcome to the Social Network!"))
    }).Methods("GET")
}
