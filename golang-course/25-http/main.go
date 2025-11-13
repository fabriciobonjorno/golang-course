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
	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	println(string(res))
	req.Body.Close()

}
