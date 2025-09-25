package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/iamrohitkandpal/Go-REST-Project/internal/config"
)

func main() {
	cfg := config.MustLoad()

	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcom to project api"))
	})

	server := http.Server{
		Addr:    cfg.Address,
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Failed to Run Server")
	}

	fmt.Println("Server Started")
}
