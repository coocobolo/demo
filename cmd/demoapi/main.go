package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	hostName := os.Getenv("HOSTNAME")
	if hostName == "" {
		hostName = "unknown"
	}

	http.HandleFunc("/readyz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		msg := fmt.Sprintf("hello-world from %s", hostName)
		w.Write([]byte(msg))
	})

	server := http.Server{
		Addr: ":8080",
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
