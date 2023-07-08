package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func Get() {
	resp, err := http.Get("http://example.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

func Post() {
	const json = `{ "name": "Hiro", "age": 27}`
	in := strings.NewReader(json)

	// not found
	resp, err := http.Post("http://example.com/upload", "text/json", in)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	message, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf(string(message))
}

func main() {
	Post()
}
