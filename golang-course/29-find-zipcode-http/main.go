package main

import "net/http"

func main() {
	http.HandleFunc("/", findZipCode)
	http.ListenAndServe(":8080", nil)
}

func findZipCode(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Find Zip Code"))
}
