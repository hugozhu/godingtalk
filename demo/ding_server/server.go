package main

import (
	"html/template"
	"net/http"
	"os"
	"fmt"
	"time"
	"path"

	"github.com/hugozhu/godingtalk"
)

func GetTemplate(tpl string) *template.Template {
	t, _ := template.ParseFiles(tpl)
	return t
}

var client *godingtalk.DingTalkClient
func serveTemplate(w http.ResponseWriter, r *http.Request) {
	lp := path.Join("templates", "layout.html")
	fp := path.Join("templates", "index.html")

	url := "http://" + r.Host + r.RequestURI
	timestamp := fmt.Sprintf("%d", time.Now().Unix())

	client.RefreshAccessToken()
	configString := client.GetConfig("abcdabc", timestamp, url)

	data := make(map[string]interface{})
	data["config"] = template.JS(configString)

	tmpl, err := template.ParseFiles(lp, fp)
	if err == nil {
		err = tmpl.ExecuteTemplate(w, "layout", data)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	corpId := os.Getenv("corpId")
	corpSecret := os.Getenv("corpSecret")
	client = godingtalk.NewDingTalkClient(corpId, corpSecret)
	client.AgentID = os.Getenv("agentID")

	fs := http.FileServer(http.Dir(path.Join(os.Getenv("root"), "public")))
	http.Handle("/public/", http.StripPrefix("/public/", fs))
	http.HandleFunc("/", serveTemplate)
	http.ListenAndServe(":8000", nil)
}
