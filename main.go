package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)
type Film struct {
	Title string
	Director string
}
func main() {

	greeting := func(w http.ResponseWriter, r *http.Request) {
		// methodType := r.Method
		films := map[string][]Film{
			"Films" : {
				{Title : "Inside out",Director: "Kalsey Mann"},
				{Title : "Chandu Champion",Director: "Kabir Khan"},
			},
		}

		temp := template.Must(template.ParseFiles("index.html"))
		temp.Execute(w,films); // renders the html page
	}

	requestHandler:= func(w http.ResponseWriter, r *http.Request){
		time.Sleep(2 *time.Second)
		title := r.PostFormValue("title")
		director := r.PostFormValue("director")

		htmlStr := fmt.Sprintf("<li class='list-group-item bg-primary text-white'> %s - %s </li>",title,director);

		temp,_ := template.New("list-item").Parse(htmlStr);
		temp.Execute(w,nil);

	}
	http.HandleFunc("/", greeting)
	http.HandleFunc("/add-film/", requestHandler)
	log.Fatal(http.ListenAndServe(":8808", nil)) // ' : ' this is impotent other wise show error missing port in address

}
