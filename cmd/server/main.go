package main

import (
	"context"
	"fmt"
	"github.com/sergeysergeevru/go-blog-service/internal/config"
	"github.com/sergeysergeevru/go-blog-service/internal/controllers"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -package oapi -generate gin,types,strict-server -o ../../internal/controllers/oapi/api.gen.go ../../api/openapi.yaml
func main() {

	cfg, err := config.ReadConfig("configs/config.yml")
	if err != nil {
		log.Fatalf("can not read configuration: %s", err)
	}

	router := controllers.CreateHandler(cfg)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Server.Port),
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()

	stop()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}
