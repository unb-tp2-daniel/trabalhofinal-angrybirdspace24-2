package main

import (
	"net/http"
)

type User struct {
	Nome string `json:"nome"`
}

var users User

func main() {

	http.ListenAndServe(":8080", nil)
}
