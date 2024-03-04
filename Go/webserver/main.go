package main

import (
	"html/template"
	"log"
	"net/http"
)

type Family struct {
	Father string
	Son    string
	Mother string
}

func main() {
	families := map[string][]Family{
		"Fam": {
			{Father: "Mitchell", Son: "Jack", Mother: "Rose"},
			{Father: "Roy", Son: "Sonu", Mother: "Natasha"},
			{Father: "Prabhu", Son: "Narendra", Mother: "Sayali"},
		},
	}
	root := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.Execute(w, families)
	}
	http.HandleFunc("/", root)
	log.Fatal(http.ListenAndServe("127.0.0.1:80", nil))
}
