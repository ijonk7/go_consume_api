package main

import (
	"log"
	"net/http"

	post "go_consume_api/controllers"
)

func main() {

	http.HandleFunc("/", post.Index)
	http.HandleFunc("/posts", post.Index)
	http.HandleFunc("/post/create", post.Create)
	http.HandleFunc("/post/store", post.Store)
	http.HandleFunc("/post/delete", post.Delete)

	log.Print("Server started on: http://localhost:8181")
	log.Fatal(http.ListenAndServe(":8181", nil))

}
