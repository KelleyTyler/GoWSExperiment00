package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"sync"
)

// var (
// 	staticLink string = "./static"
// )

type Post struct {
	ID   int    `json:"ID"`
	Body string `json:"BODY"`
}

var (
	posts   = make(map[int]Post)
	nextID  = 1
	postsMu sync.Mutex
)

func main() {
	//fmt.Printf("HELLO WORLD")
	fileServer := http.FileServer(http.Dir("./static/"))
	http.Handle("/", fileServer)

	fmt.Printf("SERVER RUNNING ON LOCALHOST:8080\n")
	// http.HandleFunc("/hello", HelloHandler)
	// http.HandleFunc("/form", formHandler)
	//http.HandleFunc("/Post", postHandler)
	http.HandleFunc("/Posts", postsHandler)
	http.HandleFunc("/Post/", postHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	// id, err := strconv.Atoi(r.URL.Path[len("/Posts/"):])
	fmt.Printf("POST !!! %d, %s\n\n", len(r.URL.Path), r.URL.Path)
	id, err := strconv.Atoi(r.URL.Path[len("/Post/"):])
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	} else {
		fmt.Printf("GOOD POST !!! + %d\n", id)
	}

	switch r.Method {
	case "GET":
		getPostHandler(w, r, id)
	case "DELETE":
		deletePostHandler(w, r, id)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
func getPostHandler(w http.ResponseWriter, r *http.Request, id int) {
	fmt.Printf("GET POST !!! \n")
	postsMu.Lock()
	defer postsMu.Unlock()

	p, ok := posts[id]
	if !ok {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)

}
func deletePostHandler(w http.ResponseWriter, r *http.Request, id int) {
	fmt.Printf("DELETE POST !!! \n")
	postsMu.Lock()
	defer postsMu.Unlock()

	// If you use a two-value assignment for accessing a
	// value on a map, you get the value first then an
	// "exists" variable.
	_, ok := posts[id]
	if !ok {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	delete(posts, id)
	w.WriteHeader(http.StatusOK)
}

func postsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("POSTS !!! \n\n")
	switch r.Method {
	case "GET":
		getPostsHandler(w, r)
	case "POST":
		postPostsHandler(w, r)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
func getPostsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("GET POSTS !!! \n")
	postsMu.Lock()
	defer postsMu.Unlock()
	ps := make([]Post, 0, len(posts))
	for _, p := range posts {
		ps = append(ps, p)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ps)

}
func postPostsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("POST POSTS !!! \n")
	var p Post

	// This will read the entire body into a byte slice
	// i.e. ([]byte)
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	// Now we'll try to parse the body. This is similar
	// to JSON.parse in JavaScript.
	if err := json.Unmarshal(body, &p); err != nil {
		http.Error(w, "Error parsing request body", http.StatusBadRequest)
		return
	}

	// As we're going to mutate the posts map, we need to
	// lock the server again
	postsMu.Lock()
	defer postsMu.Unlock()

	p.ID = nextID
	nextID++
	posts[p.ID] = p

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(p)

}

// func HelloHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/hello" {
// 		http.Error(w, "404 not found.", http.StatusNotFound)
// 		return
// 	}

// 	if r.Method != "GET" {
// 		http.Error(w, "Method is not supported.", http.StatusNotFound)
// 		return
// 	}
// 	fmt.Fprint(w, "Hello World")
// }
// func formHandler(w http.ResponseWriter, r *http.Request) {
// 	if err := r.ParseForm(); err != nil {
// 		fmt.Fprintf(w, "FORM ERROR!! %v ", err)
// 		return
// 	}
// 	fmt.Fprintf(w, "POST SUCCESSFUL")
// 	name := r.FormValue("name")
// 	address := r.FormValue("address")
// 	fmt.Fprintf(w, "Name = %s\n", name)
// 	fmt.Fprintf(w, "Address = %s\n", address)
// }
