package main

import (
	"awesomeProject4/api"
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func main() {
	var (
		listen = flag.String("listen", ":8080", "Listen specified address.")
	)
	flag.Parse()
	ctx, cancel := context.WithCancel(context.Background())

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	server := api.NewServer(logger)
	r := chi.NewRouter()
	r.Route("/api/v1", func(r chi.Router) {
		r.Use(middleware.Timeout(time.Duration(20) * time.Second))
		api.HandlerFromMux(server, r)
	})
	r.Get("/api/v1/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	srv := &http.Server{
		Handler: r,
		Addr:    *listen,
	}
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// ES is a http.Handler, so you can pass it directly to your mux

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("Error starting server", zap.Error(err))
		}
	}()
	logger.Info(fmt.Sprintf("server started on port %v", *listen))

	<-done
	logger.Info("server terminating...")
	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("server terminating failed", zap.Error(err))
	}

	// cleanup: closing db...etc
	cancel()
	logger.Info("exited")
}
