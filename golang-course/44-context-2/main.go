package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Request Started")
	defer log.Println("Request Ended")

	select {
	case <-time.After(5 * time.Second):
		log.Println("Processing complete")
		w.Write([]byte("Request processed successfully"))
	case <-ctx.Done():
		err := ctx.Err()
		log.Println("Request cancelled:", err)
		http.Error(w, "Request cancelled", http.StatusRequestTimeout)
	}
}
