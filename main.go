package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/gorilla/mux"
)

type Quote struct {
	ID     int    `json:"id"`
	Author string `json:"author"`
	Quote  string `json:"quote"`
}

var (
	quotes    = make([]Quote, 0)
	idCounter = 1
	mu        sync.Mutex
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/quotes", createQuote).Methods("POST")
	router.HandleFunc("/quotes", getQuotes).Methods("GET")
	router.HandleFunc("/quotes/random", getRandomQuote).Methods("GET")
	router.HandleFunc("/quotes/{id}", deleteQuote).Methods("DELETE")

	log.Println("Сервер запущен на порту: 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func createQuote(w http.ResponseWriter, r *http.Request) {
	var q Quote
	if err := json.NewDecoder(r.Body).Decode(&q); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mu.Lock()
	q.ID = idCounter
	idCounter++
	quotes = append(quotes, q)
	mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(q)
}

func getQuotes(w http.ResponseWriter, r *http.Request) {
	author := r.URL.Query().Get("author")
	mu.Lock()
	defer mu.Unlock()

	var result []Quote
	if author == "" {
		result = quotes
	} else {
		for _, q := range quotes {
			if strings.EqualFold(q.Author, author) {
				result = append(result, q)
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func getRandomQuote(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	if len(quotes) == 0 {
		http.Error(w, "No quotes available", http.StatusNotFound)
		return
	}
	q := quotes[rand.Intn(len(quotes))]

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(q)
}

func deleteQuote(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()
	for i, q := range quotes {
		if q.ID == id {
			quotes = append(quotes[:i], quotes[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Quote not found", http.StatusNotFound)
}
