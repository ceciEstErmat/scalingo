package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Scalingo/go-utils/logger"

	githubInterface "github.com/Scalingo/sclng-backend-test-v1/api/github"

	"github.com/Scalingo/go-handlers"
	"github.com/Scalingo/sclng-backend-test-v1/internal/services"
	"github.com/gorilla/schema"
)

func githubRoute(router *handlers.Router, g *services.Github) {
	getRouter := router.Methods(http.MethodGet).Subrouter()

	getRouter.HandleFunc("/repositories", func(writer http.ResponseWriter, request *http.Request) {
		searchRepository(writer, request, g)
	})
}

func searchRepository(w http.ResponseWriter, req *http.Request, g *services.Github) {
	log := logger.Get(req.Context())

	decoder := schema.NewDecoder()
	queryParams := githubInterface.GetGithubRepositorySearchQueryParams{}
	err := decoder.Decode(&queryParams, req.URL.Query())
	if err != nil {
		log.Println("Error in GET parameters : ", err)
	}

	filters := githubInterface.GetHttpInterfaceToService(queryParams)

	repos, err := g.GetRepository(filters)
	if err != nil {
		log.WithError(err).Error("Error fetching repositories:")
		http.Error(w, "Failed to fetch repositories", http.StatusInternalServerError)
		return
	}

	response := githubInterface.GetRepositoryStatsResponse(repos)
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.WithError(err).Error("Error encoding JSON: ")
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	// Set response headers and write JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonResponse)
	if err != nil {
		log.WithError(err).Error("Error writing response: ")
		return
	}
}
