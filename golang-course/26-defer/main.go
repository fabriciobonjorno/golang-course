package main

import (
	"io"
	"net/http"
)

func main() {
	// Make a GET request to https://google.com and print the response body
	req, err := http.Get("https://google.com")
	if err != nil {
		panic(err)
	}
	// Ensure the response body is closed after we're done with it
	defer req.Body.Close()

	// Read and print the response body
	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	println(string(res))

}
