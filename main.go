package main

import (
	"fmt"
	"net/http"
	"os"

	httpHandler "github.com/Scalingo/go-handlers"
	"github.com/Scalingo/go-utils/logger"
	"github.com/google/go-github/v60/github"

	"github.com/Scalingo/sclng-backend-test-v1/internal/handlers"
	"github.com/Scalingo/sclng-backend-test-v1/internal/services"
)

func main() {
	log := logger.Default()
	log.Info("Initializing app")
	cfg, err := newConfig()
	if err != nil {
		log.WithError(err).Error("Fail to initialize configuration")
		os.Exit(1)
	}

	githubSdk := github.NewClient(nil)
	if cfg.GithubAuthToken != "" {
		githubSdk = githubSdk.WithAuthToken(cfg.GithubAuthToken)
	}
	githubService := services.InitGithub(githubSdk)

	router := httpHandler.NewRouter(log)
	handlers.InitHttpHandler(router, &githubService)

	log = log.WithField("port", cfg.Port)
	log.Info("Listening...")
	err = http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), router)
	if err != nil {
		log.WithError(err).Error("Fail to listen to the given port")
		os.Exit(2)
	}
}
