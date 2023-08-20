package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
)

func handleCardCheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
	w.WriteHeader(200)

	// algoritmo de validacao
}
func main() {
	http.HandleFunc("/validateCard", handleCardCheck)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Print("Server started")

	<-c
	log.Print("Server stopped")
}
