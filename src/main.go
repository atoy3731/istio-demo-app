package main

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Status string
	Code   int
}

type Vars struct {
	IntraUrl    string
	CrossUrl    string
	InternetUrl string
	Port        string
}

//go:embed templates
var indexHTML embed.FS

//go:embed static
var staticFiles embed.FS

func main() {
	vars := Vars{
		IntraUrl:    os.Getenv("INTRA_URL"),
		CrossUrl:    os.Getenv("CROSS_URL"),
		InternetUrl: os.Getenv("INTERNET_URL"),
		Port:        os.Getenv("PORT"),
	}

	log.Println("===============")
	log.Println("Configuration:")
	log.Println(fmt.Sprintf("  - Port = %s", vars.Port))
	log.Println(fmt.Sprintf("  - IntraUrl = %s", vars.IntraUrl))
	log.Println(fmt.Sprintf("  - CrossUrl = %s", vars.CrossUrl))
	log.Println(fmt.Sprintf("  - InternetUrl = %s", vars.InternetUrl))
	log.Println("===============")

	tmpl, err := template.ParseFS(indexHTML, "templates/index.html.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	var staticFS = http.FS(staticFiles)
	fs := http.FileServer(staticFS)

	http.Handle("/static/", fs)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := vars
		tmpl.Execute(w, data)
	})

	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		status(w, r)
	})

	http.HandleFunc("/api/intra", func(w http.ResponseWriter, r *http.Request) {
		intraStatus(w, r, vars)
	})

	http.HandleFunc("/api/cross", func(w http.ResponseWriter, r *http.Request) {
		crossStatus(w, r, vars)
	})

	http.HandleFunc("/api/internet", func(w http.ResponseWriter, r *http.Request) {
		internetStatus(w, r, vars)
	})

	if vars.Port == "" {
		log.Fatal("ERROR: Missing required variable 'PORT'")
	}

	log.Println(fmt.Sprintf("Listening for requests at http://0.0.0.0:%s", vars.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", vars.Port), nil))
}
