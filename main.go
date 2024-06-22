package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello World")

	greeting := func (w http.ResponseWriter, r *http.Request) {
		methodType := r.Method
		io.WriteString(w,"Hello World");
		io.WriteString(w,methodType)
	}

	http.HandleFunc("/",greeting);
	log.Fatal(http.ListenAndServe(":8808", nil))
}
