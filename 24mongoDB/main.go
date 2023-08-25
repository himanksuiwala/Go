package main

import (
	"fmt"
	"log"
	"mongoDB/router"
	"net/http"
)

func main() {
	fmt.Println("HEYS")
	r := router.Router()
	fmt.Println("SEVER is getting started...4000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
