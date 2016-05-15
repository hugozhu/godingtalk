package main

import (
	"log"
	"net/http"

	"github.com/hugozhu/godingtalk/demo/github"
)

func init() {
	http.HandleFunc("/github", github.Handle)
}

func main() {
	log.Print("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
