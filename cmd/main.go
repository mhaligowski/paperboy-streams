package main

import (
	_ "github.com/mhaligowski/streams"
	"net/http"
)


func main() {
	http.ListenAndServe(":8080", nil)
}
