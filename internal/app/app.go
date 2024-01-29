package app

import (
	"context"
	"currency-rate/config"
	"currency-rate/internal/controller/http/v1"
	"log"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func Run(ctx context.Context, cfg *config.Config) {

	srv := &http.Server{
		Addr:           ":" + cfg.Port,
		Handler:        v1.NewRouter(),
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())

	// Shutdown

	err := srv.Shutdown(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
