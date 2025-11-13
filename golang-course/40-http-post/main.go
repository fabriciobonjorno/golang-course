package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main() {
	c := http.Client{}
	jsonData := bytes.NewBuffer([]byte(`{"name":"fabricio"}`))
	// Make a GET request to https://google.com and print the response body
	resp, err := c.Post("https://google.com", "application/json", jsonData)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	io.CopyBuffer(os.Stdout, resp.Body, nil)

}
