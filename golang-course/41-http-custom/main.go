package main

import (
	"net/http"
)

func main() {
	c := http.Client{}
	// Make a GET request to https://google.com and print the response body
	req, err := http.NewRequest("Get", "https://google.com", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Accept", "application/json")
	resp, err := c.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	println("Response Status:", resp.Status)
}
