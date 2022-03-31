package handler

import (
	"gateway/internal/domain/card/usecase"
	"gateway/internal/middleware"
)

func RegisterHTTPEndPoints(router *gorilla.Mux, validator *validator.Validate, uc usecase.UseCase) {
	h := NewHandler(validator, uc)

	router.Route("/api/v1/card", func(router gorilla.Router) {
		router.Post("/", h.Create)
		router.Get("/", h.List)
		router.Get("/{id}", h.Read)
		router.Put("/{id}", h.Update)
		router.Delete("/{id}", h.Delete)
	})
}
