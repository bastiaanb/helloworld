package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
)

func main() {
	log.Printf("starting helloworld")

	message := os.Args[1]
	var httpAddress = os.Getenv("LISTEN_ADDRESS")
	if httpAddress == "" {
		httpAddress = ":8081"
	}

	http.Handle("/fs/", http.StripPrefix("/fs/", http.FileServer(http.Dir("."))))
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		handleHello(w, r, message)
	})
	http.HandleFunc("/env", handleEnv)

	http.HandleFunc("/health", handleHealth)
	http.HandleFunc("/fail", func(w http.ResponseWriter, r *http.Request) {
		handleSetHealth(w, r, http.StatusServiceUnavailable)
	})
	http.HandleFunc("/ok", func(w http.ResponseWriter, r *http.Request) {
		handleSetHealth(w, r, http.StatusOK)
	})

	log.Printf("starting http service at %s", httpAddress)
	log.Fatal(http.ListenAndServe(httpAddress, nil))
}

var mutex sync.Mutex
var counter int
var healthStatus = http.StatusOK

func handleHealth(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	w.WriteHeader(healthStatus)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(healthStatus)
	mutex.Unlock()
}

func handleSetHealth(w http.ResponseWriter, r *http.Request, status int) {
	mutex.Lock()
	healthStatus = status
	w.WriteHeader(http.StatusOK)
	mutex.Unlock()
}

func handleEnv(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	var environment = make(map[string]string)
	for _, env := range os.Environ() {
		pair := strings.Split(env, "=")
		if strings.HasPrefix(pair[0], "NOMAD") {
			environment[pair[0]] = pair[1]
		}
	}
	json.NewEncoder(w).Encode(environment)
}

func handleHello(w http.ResponseWriter, r *http.Request, message string) {
	mutex.Lock()
	counter++
	fmt.Fprintf(w, "Hello %d: %s\n", counter, message)
	mutex.Unlock()
}
