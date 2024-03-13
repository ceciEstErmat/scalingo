package handlers

import (
	"net/http"

	"github.com/Scalingo/go-handlers"
)

func health(router *handlers.Router) {
	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		return
	})
}
