package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"log"
	"time"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)
const (
	hosts      = "my-mongo-container:27017"
	database   = "test_mongo"
	username   = "root"
	password   = "root"
	collection = "post"
)

type Post struct {
	Tile    string
	Content string
}

func main() {
	info := &mgo.DialInfo{
		Addrs:    []string{hosts},
		Timeout:  5 * time.Second,
		Database: database,
		Username: username,
		Password: password,
	}
	
	session, err := mgo.DialWithInfo(info)
	if err != nil {
		panic(err)
	}

	col := session.DB(database).C(collection)
	session.DB(database).C(collection).RemoveAll(nil)
	err = col.Insert(&Post{"Hello World", "Đây là bài		 viết hướng dẫn thao tác với MongoDB trong Golang"},
	&Post{"Thời tiết", "Hôm nay nắng đẹp quá!!!"})
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()

	// // request handler to specific HTTP methods
    // r.HandleFunc("/books/{title}", CreateBook).Methods("POST")
    // r.HandleFunc("/books/{title}", ReadBook).Methods("GET")
    // r.HandleFunc("/books/{title}", UpdateBook).Methods("PUT")
	// r.HandleFunc("/books/{title}", DeleteBook).Methods("DELETE")
	
	// // hostnames or subdomanis
    // r.HandleFunc("/books/{title}", BookHandler).Host("www.mybookstore.com")

    // // restrict the request handler to http/https.
    // r.HandleFunc("/secure", SecureHandler).Schemes("https")
    // r.HandleFunc("/insecure", InsecureHandler).Schemes("http")

	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r) // map[string]string

		title := vars["title"]
		page := vars["page"]

		
		var posts []Post
		err = col.Find(bson.M{}).All(&posts)
		if err != nil {
			log.Fatal(err)
		}
		for _, post := range posts {
			log.Println(post)
		}
		fmt.Fprintf(w, "You've requested the books: %s on page %s\n", title, page)
	})
	
	// Order
	r.HandleFunc("/v1/api", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "V1 Api \n")
	})
	r.HandleFunc("/v1/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r) // map[string]string
		fmt.Fprintf(w, "V1 ID %s \n", vars["id"])
	})
	r.HandleFunc("/v2/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		fmt.Fprintf(w, "V2 ID %s \n", vars["id"])
	})
	r.HandleFunc("/v2/api", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "V2 Api \n")
	})

	// dynamic segments
	r.HandleFunc("/books/{title}/page/{page}", bookHandler)
	// Group Route
	// Example products
	s := r.PathPrefix("/products").Subrouter()
	// Subrouter nest Subrouter
	x := s.PathPrefix("/test").Subrouter()
	// path http://localhost:8081/products/test
	x.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Subrouter nest Subrouter \n")
	})
	// Subrouter
	// path "http://localhost:8081/products/"
	s.HandleFunc("/", productsHandlerfunc)
	// path "http://localhost:8081/products/{key}"
	s.HandleFunc("/{key}", productHandler)
	// path "http://localhost:8081/products/{key}/details"
	s.HandleFunc("/{key}/details", productDetailsHandler)

	http.ListenAndServe(":8080", r)
}

func bookHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	page := vars["page"]
	fmt.Fprintf(w, "You've requested the books: %s on page %s\n", title, page)
}
// productsHandlerfunc func
func productsHandlerfunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Product List OK \n")
}
// productHandler func
func productHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "Products %s OK \n", vars["key"])
}
// productDetailsHandler func
func productDetailsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "Product Detail %s OK \n", vars["key"])
}
