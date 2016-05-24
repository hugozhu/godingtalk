package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path"
)

func GetTemplate(tpl string) *template.Template {
	t, _ := template.ParseFiles(tpl)
	return t
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	lp := path.Join("templates", "layout.html")
	fp := path.Join("templates", "index.html")

	tmpl, err := template.ParseFiles(lp, fp)
	if err == nil {
		err = tmpl.ExecuteTemplate(w, "layout", nil)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	fs := http.FileServer(path.Join(os.Getenv("root"), "public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))
	http.HandleFunc("/", serveTemplate)
	http.ListenAndServe(":8000", nil)
}
