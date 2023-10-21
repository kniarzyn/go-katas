package main

import (
	"log"
	"net/http"
	"text/template"
)

func main() {
	http.Handle("/static/", http.FileServer(http.Dir("static")))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("./templates/index.html"))
		tmpl.Execute(w, nil)
	})

	log.Println("App running on 8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
