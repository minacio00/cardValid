package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"

	cardvalidation "github.com/minacio00/cardValid/cardValidation"
)

type card struct {
	Number string `json:"number"`
}
type response struct {
	IsValid bool `json:"valid"`
}

func str2intArr(s string) []int {
	intArr := make([]int, len(s))
	for i, v := range s {
		digit, err := strconv.Atoi(string(v))
		if err != nil {
			fmt.Println("Error:", err)
			return nil
		}
		intArr[i] = digit
	}
	return intArr
}
func handleCardCheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	card := &card{}
	err := decoder.Decode(&card)
	if err != nil {
		s := fmt.Sprintf("Invalid input, %s", err.Error())
		http.Error(w, s, http.StatusBadRequest)
		return
	}

	intArr := str2intArr(card.Number)
	result := &response{}
	result.IsValid = cardvalidation.IsValid(intArr)

	w.Header().Set("Content-type", "application/json")
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	println(result.IsValid)
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
	log.Print("Server started at :8080")

	<-c
	log.Print("Server stopped")
}
