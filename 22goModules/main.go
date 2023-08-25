package main

// https://go.dev/blog/get-familiar-with-workspaces
import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("HELLO")
	greeter()
}

func greeter() {
	fmt.Println("Hey there mode users")
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":3001", r))
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>This is server side rendered page</h1>"))
}
