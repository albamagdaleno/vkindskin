package main

import (
	"log"
	"net/http"
	"os"
)

var port = os.Getenv("PORT")

func main() {
	log.Print("Hola Mundo, me llamo Alba")

	if port == "" {
		port = "3000"
	}

	server := http.NewServeMux()

	server.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
		log.Println("Request acepted")
	})

	if err := http.ListenAndServe(":"+port, server); err != nil {
		log.Fatal("failed to start server: %v", err)
	}

}
