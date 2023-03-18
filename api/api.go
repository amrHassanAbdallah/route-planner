package api

import (
	"awesomeProject4/service"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
)

type RouteRequest struct {
	Route [][]string `json:"route" validate:"required,min=1,max=100,dive,required,min=2,max=2"`
}

func validateRouteRequest(req RouteRequest) error {
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return err
	}
	return nil
}

type ServerInterfaceGlobal interface {
	ServerInterface
}
type server struct {
	logger *zap.Logger
}

func (s server) PostRoute(w http.ResponseWriter, r *http.Request) {
	logger := s.logger
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

	point := service.GetRouteStartAndEnd(req.Route)
	if point.Source == "" || point.Destination == "" {
		logger.Error("Invalid route")
		http.Error(w, "Invalid route", http.StatusBadRequest)
		return
	}

	resp := service.Point{Source: point.Source, Destination: point.Destination}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		logger.Error("Error encoding response", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logger.Info("Request successful", zap.Any("request", req), zap.Any("response", resp))
}

func NewServer(log *zap.Logger) ServerInterfaceGlobal {
	return &server{
		logger: log,
	}
}
