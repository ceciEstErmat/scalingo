package handlers

import (
	"github.com/Scalingo/go-handlers"
	"github.com/Scalingo/sclng-backend-test-v1/internal/services"
)

func InitHttpHandler(router *handlers.Router, g *services.Github) {
	health(router)
	githubRoute(router, g)
}
