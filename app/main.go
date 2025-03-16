package main

import (
	"fmt"
	"log"
	"net/http"
)

// var (
// 	staticLink string = "./static"
// )

func main() {
	//fmt.Printf("HELLO WORLD")
	fileServer := http.FileServer(http.Dir("./static/"))
	http.Handle("/", fileServer)

	fmt.Printf("SERVER RUNNING ON LOCALHOST:8080\n")
	http.HandleFunc("/hello", HelloHandler)
	http.HandleFunc("/form", postFormHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	fmt.Fprint(w, "Hello World")
}
func postFormHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "FORM ERROR!! %v ", err)
		return
	}
	fmt.Fprintf(w, "POST SUCCESSFUL")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}
