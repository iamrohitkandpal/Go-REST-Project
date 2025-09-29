package main

import (
	"context"
	// "fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/iamrohitkandpal/Go-REST-Project/internal/config"
	"github.com/iamrohitkandpal/Go-REST-Project/internal/http/handlers/student"
	"github.com/iamrohitkandpal/Go-REST-Project/internal/storage/sqlite"
)

func main() {
	// Loading Config
	cfg := config.MustLoad()

	// Database Setup
	storage, err := sqlite.New(*cfg)
	if err != nil {
		log.Fatal(err)
	}
	slog.Info("Storage Initialized", slog.String("env", cfg.Env), slog.String("version", "1.0.0"))

	// Router Setup
	router := http.NewServeMux()

	router.HandleFunc("POST /api/students", student.New(storage))
	router.HandleFunc("GET /api/students/{id}", student.GetById(storage))
	router.HandleFunc("GET /api/students/", student.GetAll(storage))
	router.HandleFunc("PUT /api/students/{id}", student.UpdateOne(storage))
	router.HandleFunc("DELETE /api/students/{id}", student.DeleteOne(storage))


	// Server Setup
	server := http.Server{
		Addr:    cfg.Address,
		Handler: router,
	}

	slog.Info("Server Started", slog.String("At Address", cfg.Address))

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func (){
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("Failed to Run Server")
		}
	}()

	<-done 

	slog.Info("Shutting down the server")

	ctx ,cancel := context.WithTimeout(context.Background(), 5 * time.Second)

	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Failed to Shutdown Server", slog.String("Error", err.Error()))
	}

	slog.Info("Server SHutdown Successfully")
}
