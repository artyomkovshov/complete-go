package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/artyomkovshov/femProject/internal/app"
)

func main() {
	app, err := app.NewApplication()

	if err != nil {
		panic(err)
	}

	app.Logger.Println("Application started successfully")

	http.HandleFunc("/health", HealthCheck)
	server := &http.Server{
		Addr:         ":8090",
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	err = server.ListenAndServe()

	if err != nil {
		app.Logger.Fatalf("Failed to start server: %v", err)
	}
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Status: OK\n")
}
