package main


import (
	"fmt"
	"net/http"
)

// function that returns an http.Handler
func exampleHandler() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		_, err := fmt.Fprintln(writer, "Hello World")
		if err != nil {
			return
		}
	})
}
