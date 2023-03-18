package main

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type RouteRequest struct {
	Route [][]string `json:"route" validate:"required,min=1,max=1,dive,required,min=2,max=2"`
}

func validateRouteRequest(req RouteRequest) error {
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return err
	}
	return nil
}

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/route", func(w http.ResponseWriter, r *http.Request) {
		var req RouteRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			logger.Error("Error decoding request", zap.Error(err))
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := validateRouteRequest(req); err != nil {
			logger.Error("Invalid request", zap.Error(err))
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		point := GetRouteStartAndEnd(req.Route)
		if point.Source == "" || point.Destination == "" {
			logger.Error("Invalid route")
			http.Error(w, "Invalid route", http.StatusBadRequest)
			return
		}

		resp := Point{Source: point.Source, Destination: point.Destination}
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			logger.Error("Error encoding response", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		logger.Info("Request successful", zap.Any("request", req), zap.Any("response", resp))
	})

	if err := http.ListenAndServe(":8080", r); err != nil {
		logger.Error("Error starting server", zap.Error(err))
	}
}
