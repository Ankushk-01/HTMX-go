package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

type Film struct {
	Title    string
	Director string
}

func main() {
	temp := template.Must(template.ParseFiles("index.html"))

	greeting := func(w http.ResponseWriter, r *http.Request) {
		// methodType := r.Method
		films := map[string][]Film{
			"Films": {
				{Title: "Inside out", Director: "Kalsey Mann"},
				{Title: "Chandu Champion", Director: "Kabir Khan"},
			},
		}

		temp.Execute(w, films) // renders the html page
	}

	requestHandler := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second)
		title := r.PostFormValue("title")
		director := r.PostFormValue("director")
		temp.ExecuteTemplate(w, "file-director-list", Film{Title: title, Director: director})
	}
	http.HandleFunc("/", greeting)
	http.HandleFunc("/add-film/", requestHandler)
	log.Fatal(http.ListenAndServe(":8808", nil)) // ' : ' this is impotent other wise show error missing port in address

}
