package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hugozhu/godingtalk/demo/github"
)

func init() {
	http.HandleFunc("/github", github.Handle)
	http.HandleFunc("/_ah/health", healthCheckHandler)
	http.HandleFunc("/", handler)
}

func main() {
	log.Print("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "Hello Github Alert!")
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}
